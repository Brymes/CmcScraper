package utils

import (
	"context"
	"flag"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

var devtoolsWsURL = flag.String("devtoolsWSurl", "http://127.0.0.1:9222/", "DevTools Websocket URL")
var allocatorContext, ctxt context.Context

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateChromeContext(age int64) context.Context {
	allocatorContext, _ =
		chromedp.NewRemoteAllocator(context.Background(),
			*devtoolsWsURL)

	ctxt, _ = chromedp.NewContext(allocatorContext)

	// Parse context age to Duration as Seconds
	dur, _ := time.ParseDuration(fmt.Sprintf("%vs", age))
	ctxt, _ = context.WithTimeout(ctxt, time.Duration(dur))

	return ctxt
}
