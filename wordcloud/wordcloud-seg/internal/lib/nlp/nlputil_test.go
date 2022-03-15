package nlp_test

import (
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/nlp"
	"testing"
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
