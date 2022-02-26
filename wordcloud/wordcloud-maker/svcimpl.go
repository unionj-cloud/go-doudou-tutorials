package service

import (
	"context"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/unionj-cloud/go-doudou/toolkit/pathutils"

	"github.com/chromedp/chromedp"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-resty/resty/v2"
	"github.com/minio/minio-go/v7"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/vo"
	segsvc "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg"
	segvo "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/vo"
)

type WordcloudMakerImpl struct {
	conf        *config.Config
	segClient   segsvc.WordcloudSeg
	minioClient *minio.Client
	browser     *rod.Browser
}

func NewWordcloudMaker(conf *config.Config, segClient segsvc.WordcloudSeg,
	minioClient *minio.Client, browser *rod.Browser) WordcloudMaker {
	return &WordcloudMakerImpl{
		conf,
		segClient,
		minioClient,
		browser,
	}
}

func generateWCData(data map[string]interface{}) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func wcBase(data map[string]interface{}) *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.AddSeries("wordcloud", generateWCData(data)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
				}),
		)
	return wc
}

func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

func getPublicOssUrl(endpoint, bucketName, objectName string) string {
	return fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, objectName)
}

func (receiver *WordcloudMakerImpl) Make(ctx context.Context, payload vo.MakePayload) (data string, err error) {
	srcUrl, err := url.Parse(payload.SrcUrl)
	if err != nil {
		return "", ddhttp.NewBizError(errors.Wrap(err, "srcUrl should be valid url"), ddhttp.WithStatusCode(400))
	}
	filename := srcUrl.Path[strings.LastIndex(srcUrl.Path, "/")+1:]
	span := opentracing.SpanFromContext(ctx)
	outputDir := filepath.Join(receiver.conf.BizConf.Output, fmt.Sprint(span))
	client := resty.New()
	client.SetOutputDirectory(outputDir)
	resp, err := client.R().
		SetContext(ctx).
		SetOutput(filename).
		Get(payload.SrcUrl)
	if err != nil || !resp.IsSuccess() {
		return "", ddhttp.NewBizError(errors.Wrap(err, "source file download fail"), ddhttp.WithStatusCode(400))
	}
	file := filepath.Join(outputDir, filename)
	defer os.RemoveAll(outputDir)
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return "", err
	}
	limit := 1 << (10 * 2)
	if fi.Size() > int64(limit) {
		return "", ddhttp.NewBizError(errors.New("file size is larger than 1M"), ddhttp.WithStatusCode(400))
	}
	text, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	seg, err := receiver.segClient.Seg(ctx, segvo.SegPayload{
		Text: string(text),
		Lang: payload.Lang,
	})
	if err != nil {
		return "", err
	}
	wcData := make(map[string]interface{})

	wf := seg.WordFreq
	if payload.Top > 0 {
		wf = wf[:payload.Top]
	}

	for _, item := range wf {
		wcData[item[0].(string)] = item[2]
	}
	page := components.NewPage()
	page.AddCharts(wcBase(wcData))
	var html *os.File
	outhtml := filepath.Join(outputDir, "wordcloud.html")
	html, err = os.Create(outhtml)
	if err != nil {
		return "", err
	}
	defer html.Close()
	err = page.Render(io.MultiWriter(html))
	if err != nil {
		return "", err
	}
	now := time.Now()
	rpage := receiver.browser.MustPage(fmt.Sprintf("file://%s", pathutils.Abs(outhtml))).MustWaitLoad()
	time.Sleep(1 * time.Second)
	el, err := rpage.Timeout(10 * time.Second).Element("canvas")
	if err != nil {
		return "", err
	}
	outimg := filepath.Join(outputDir, "wordcloud.png")
	el.MustScreenshot(outimg)
	logger.Info(time.Since(now))

	bucketName := receiver.conf.BizConf.OssBucket
	objectName := fmt.Sprintf("%s_wordcloud.png", fmt.Sprint(span))
	_, err = receiver.minioClient.FPutObject(ctx, bucketName, objectName, outimg, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	return getPublicOssUrl(receiver.conf.BizConf.OssEndpoint, bucketName, objectName), nil
}
