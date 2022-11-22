package service

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"google.golang.org/protobuf/types/known/wrapperspb"

	segpb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/transport/grpc"

	"github.com/go-resty/resty/v2"

	"github.com/go-rod/rod"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/config"
	pb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/transport/grpc"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-maker/vo"
)

var _ pb.WordcloudMakerServiceServer = (*WordcloudMakerImpl)(nil)

type WordcloudMakerImpl struct {
	pb.UnimplementedWordcloudMakerServiceServer

	conf        *config.Config
	segClient   segpb.WordcloudSegServiceClient
	minioClient *minio.Client
	browser     *rod.Browser
}

func NewWordcloudMaker(conf *config.Config, segClient segpb.WordcloudSegServiceClient,
	minioClient *minio.Client, browser *rod.Browser) *WordcloudMakerImpl {
	return &WordcloudMakerImpl{
		conf:        conf,
		segClient:   segClient,
		minioClient: minioClient,
		browser:     browser,
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

func getPublicOssUrl(endpoint, bucketName, objectName string) string {
	return fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, objectName)
}

func (receiver *WordcloudMakerImpl) Make(ctx context.Context, payload vo.MakePayload) (data string, err error) {
	srcUrl, err := url.Parse(payload.SrcUrl)
	if err != nil {
		return "", rest.NewBizError(errors.Wrap(err, "srcUrl should be valid url"), rest.WithStatusCode(400))
	}
	filename := srcUrl.Path[strings.LastIndex(srcUrl.Path, "/")+1:]
	imageId := uuid.NewString()
	outputDir := filepath.Join(receiver.conf.BizConf.Output, imageId)
	client := resty.New()
	client.SetOutputDirectory(outputDir)
	resp, err := client.R().
		SetContext(ctx).
		SetOutput(filename).
		Get(payload.SrcUrl)
	if err != nil || !resp.IsSuccess() {
		return "", rest.NewBizError(errors.Wrap(err, "source file download fail"), rest.WithStatusCode(400))
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
		return "", rest.NewBizError(errors.New("file size is larger than 1M"), rest.WithStatusCode(400))
	}
	text, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	in := &segpb.SegPayload{
		Text: string(text),
		Lang: payload.Lang,
	}
	segResult, err := receiver.segClient.SegRpc(ctx, in)
	if err != nil {
		return "", err
	}
	wcData := make(map[string]interface{})
	var wf [][]interface{}
	for _, item := range segResult.WordFreq {
		element := make([]interface{}, 3)

		wordAny := item.NestedAny[0]
		var wordWrapper wrapperspb.StringValue
		_ = wordAny.UnmarshalTo(&wordWrapper)
		element[0] = wordWrapper.Value

		posAny := item.NestedAny[1]
		var posWrapper wrapperspb.StringValue
		_ = posAny.UnmarshalTo(&posWrapper)
		element[1] = posWrapper.Value

		freqAny := item.NestedAny[2]
		var freqWrapper wrapperspb.DoubleValue
		_ = freqAny.UnmarshalTo(&freqWrapper)
		element[2] = freqWrapper.Value

		wf = append(wf, element)
	}

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
	zlogger.Info().Msg("path of output html: " + outhtml)
	rpage := receiver.browser.MustPage(fmt.Sprintf("file://%s", outhtml)).MustWaitLoad()
	time.Sleep(1 * time.Second)
	el, err := rpage.Timeout(10 * time.Second).Element("canvas")
	if err != nil {
		return "", err
	}
	outimg := filepath.Join(outputDir, "wordcloud.png")
	el.MustScreenshot(outimg)
	zlogger.Info().Msg(time.Since(now).String())

	bucketName := receiver.conf.BizConf.OssBucket
	objectName := fmt.Sprintf("%s_wordcloud.png", imageId)
	_, err = receiver.minioClient.FPutObject(ctx, bucketName, objectName, outimg, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}

	return getPublicOssUrl(receiver.conf.BizConf.OssEndpoint, bucketName, objectName), nil
}

func (receiver *WordcloudMakerImpl) MakeRpc(ctx context.Context, request *pb.MakePayload) (*pb.MakeRpcResponse, error) {
	payload := vo.MakePayload{
		SrcUrl: request.SrcUrl,
		Lang:   request.Lang,
		Top:    int(request.Top),
	}
	ret, err := receiver.Make(ctx, payload)
	if err != nil {
		return nil, err
	}
	return &pb.MakeRpcResponse{
		Data: ret,
	}, nil
}
