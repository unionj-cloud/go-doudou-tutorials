package en_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/internal/lib/en"
	"testing"
)

func TestDoSeg(t *testing.T) {
	tokenizer := en.EnglishTokenizer{}
	doc, err := tokenizer.DoSeg(context.Background(),
		"The Winter Olympics take place from 4 February to 20 February, with about 3,000 athletes competing in 109 different events. "+
			"The Winter Paralympics run from 4 March to 13 March, with about 750 competitors across 78 events.",
		en.Pos, -1, -1)
	require.NoError(t, err)
	for _, item := range doc {
		fmt.Println(item[0], item[2])
	}
}
