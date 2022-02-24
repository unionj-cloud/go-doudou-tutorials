package nlp

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/toolkit/stringutils"
	"github.com/ztrue/tracerr"
	"golang.org/x/sync/errgroup"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"
)

var chineseSplitReg = regexp.MustCompile(`(?U)(.*[\s\x{3002}\x{ff1b}\x{ff0c}\x{ff1a}\x{201c}\x{201d}\x{ff08}\x{ff09}\x{3001}\x{ff1f}\x{300a}\x{300b}])`)

type CorpusCleaner interface {
	DoClean(text string) string
	GetSeparator() string
	DoSplit(text string) []string
}

func RemoveUrl(text string) string {
	re := regexp.MustCompile(`(https?|ftp|file)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]`)
	return re.ReplaceAllString(text, "")
}

func RemoveNonChinese(text string) string {
	re := regexp.MustCompile(`[^\p{Han}\d\s\x{3002}\x{ff1b}\x{ff0c}\x{ff1a}\x{201c}\x{201d}\x{ff08}\x{ff09}\x{3001}\x{ff1f}\x{300a}\x{300b}]`)
	return re.ReplaceAllString(text, "")
}

func CleanChineseText(text string) string {
	return strings.TrimSpace(RemoveNonChinese(RemoveUrl(text)))
}

type ChineseCorpusCleaner struct{}

func (self *ChineseCorpusCleaner) DoClean(text string) string {
	return CleanChineseText(text)
}

func (self *ChineseCorpusCleaner) GetSeparator() string {
	return "。"
}

func (self *ChineseCorpusCleaner) DoSplit(text string) []string {
	splits := chineseSplitReg.FindAllStringSubmatch(text, -1)
	var ret []string
	for _, split := range splits {
		if stringutils.IsNotEmpty(split[1]) && utf8.RuneCountInString(split[1]) > 1 {
			ret = append(ret, split[1])
		}
	}
	return ret
}

type EnglishCorpusCleaner struct{}

//TODO
func (self *EnglishCorpusCleaner) DoClean(text string) string {
	panic("implement me!")
}

func (self *EnglishCorpusCleaner) GetSeparator() string {
	return "."
}

func (self *EnglishCorpusCleaner) DoSplit(text string) []string {
	panic("implement me!")
}

type MixCorpusCleaner struct{}

//TODO
func (self *MixCorpusCleaner) DoClean(text string) string {
	panic("implement me")
}

//TODO
func (self *MixCorpusCleaner) GetSeparator() string {
	panic("implement me")
}

//TODO
func (self *MixCorpusCleaner) DoSplit(text string) []string {
	panic("implement me!")
}

func ProcessLine(r io.Reader, max int, threshold int, cleaner CorpusCleaner, callback func(s string) error) error {
	reader := bufio.NewReader(r)
	var err error

	hits := make(chan string)
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		defer close(hits)

		for {
			var line []byte
			line, err = reader.ReadBytes('\n')
			if err != nil && err != io.EOF {
				err = tracerr.Wrap(err)
				return err
			}

			sline := cleaner.DoClean(string(line))

			lineLength := utf8.RuneCount([]byte(sline))

			if max < 0 || lineLength <= max {
				select {
				case hits <- sline:
				case <-ctx.Done():
					return ctx.Err()
				}
			} else {
				splits := cleaner.DoSplit(sline)
				var index int
				for {
					if index >= len(splits) {
						break
					}

					var newline string

					for {
						if utf8.RuneCount([]byte(newline)) >= max || index >= len(splits) {
							break
						}
						newline += splits[index]
						index++
					}

					if utf8.RuneCount([]byte(newline)) > threshold {
						logrus.Println(lineLength)
						logrus.Println(newline)
						return tracerr.New(fmt.Sprintf("单行字符数超过%d，请调整后重试", threshold))
					}
					select {
					case hits <- newline:
					case <-ctx.Done():
						return ctx.Err()
					}
				}
			}

			if err == io.EOF {
				break
			}
		}

		return nil
	})

	for i := 0; i < 1; i++ {
		g.Go(func() error {
			for hit := range hits {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					if err = callback(hit); err != nil {
						err = tracerr.Wrap(err)
						return err
					}
				}
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		err = tracerr.Wrap(err)
		return err
	}
	return nil
}
