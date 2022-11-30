package main

import (
	"fmt"
	"github.com/influxdata/telegraf/plugins/parsers/influx"
	"github.com/influxdata/telegraf/plugins/serializers"
	"os"
	"time"
)

func main() {
	parser := influx.NewStreamParser(os.Stdin)
	serializer, _ := serializers.NewInfluxSerializer()

	for {
		metric, err := parser.Next()
		if err != nil {
			if err == influx.EOF {
				return // stream ended
			}
			if parseErr, isParseError := err.(*influx.ParseError); isParseError {
				fmt.Fprintf(os.Stderr, "parse ERR %v\n", parseErr)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "ERR %v\n", err)
			os.Exit(1)
		}

		c, found := metric.GetField("gauge")
		if !found {
			fmt.Fprintf(os.Stderr, "metric has no gauge field\n")
			os.Exit(1)
		}
		switch t := c.(type) {
		case float64:
			metric.AddField("elapsed", time.Since(time.UnixMilli(int64(t*1000))).Milliseconds())
		case int64:
			metric.AddField("elapsed", time.Since(time.UnixMilli(t*1000)).Milliseconds())
		default:
			fmt.Fprintf(os.Stderr, "gauge is not an unknown type, it's a %T\n", c)
			os.Exit(1)
		}
		b, err := serializer.Serialize(metric)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERR %v\n", err)
			os.Exit(1)
		}
		fmt.Fprint(os.Stdout, string(b))
	}
}
