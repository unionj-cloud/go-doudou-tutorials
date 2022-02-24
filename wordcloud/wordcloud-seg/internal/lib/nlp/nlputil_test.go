package nlp_test

import (
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/nlp"
	"io"
	"os"
	"testing"

	"go.uber.org/atomic"
)

func TestRemoveUrl(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				text: "演讲不仅仅要求演讲者有强烈的气场和自信，而且需要一份http://www.baidu.com独特的、好的演讲稿。演讲稿越好，演讲者https://mail.qq.com才能更出色，如果你想拥有一份好的演讲稿，就不要犹豫。",
			},
			want: "演讲不仅仅要求演讲者有强烈的气场和自信，而且需要一份独特的、好的演讲稿。演讲稿越好，演讲者才能更出色，如果你想拥有一份好的演讲稿，就不要犹豫。",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nlp.RemoveUrl(tt.args.text); got != tt.want {
				t.Errorf("RemoveUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveNonChinese(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "2",
			args: args{
				text: "演讲不仅仅要求演讲者有强烈的气场和自信，而且需要一份http://www.baidu.com独特的、好的演讲稿。",
			},
			want: "演讲不仅仅要求演讲者有强烈的气场和自信，而且需要一份独特的、好的演讲稿。",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nlp.RemoveNonChinese(tt.args.text); got != tt.want {
				t.Errorf("RemoveNonChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCleanChineseText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "3",
			args: args{
				text: "演讲不仅仅要求演讲者有强烈的气场和自信，而且需要一份http://www.baidu.com独特的、好的演讲稿。",
			},
			want: "演讲不仅仅要求演讲者有强烈的气场和自信，而且需要一份独特的、好的演讲稿。",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nlp.CleanChineseText(tt.args.text); got != tt.want {
				t.Errorf("CleanChineseText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessLine(t *testing.T) {
	type args struct {
		r io.Reader
	}

	f, err := os.Open("2019.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "2",
			args: args{
				r: f,
			},
			want:    14419,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lineCount := atomic.NewInt64(0)
			err = nlp.ProcessLine(tt.args.r, 5, 60, &nlp.ChineseCorpusCleaner{}, func(s string) error {
				lineCount.Inc()
				return nil
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("CountLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if lineCount.Load() != tt.want {
				t.Errorf("CountLine() = %d, want %d", lineCount, tt.want)
			}
		})
	}
}
