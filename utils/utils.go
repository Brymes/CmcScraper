package utils

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var devtoolsWsURL = flag.String("devtoolsWSurl", "http://127.0.0.1:9222/", "DevTools Websocket URL")

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateChromeContext(age int64) context.Context {
	// TODO make this function start the devTools protocol automatically
	var allocatorContext, ctxt context.Context

	allocatorContext, _ =
		chromedp.NewRemoteAllocator(context.Background(),
			*devtoolsWsURL)

	ctxt, _ = chromedp.NewContext(allocatorContext)

	// Parse context age to Duration as Seconds
	dur, _ := time.ParseDuration(fmt.Sprintf("%vs", age))
	ctxt, _ = context.WithTimeout(ctxt, time.Duration(dur))

	return ctxt
}

func ParseStringToInt(quantity string) (int, error) {
	temp := strings.Split(quantity, "$")
	chars := strings.Split(temp[1], ",")

	return strconv.Atoi(strings.Join(chars, ""))
}

func FindNumIters(numCoins int) int {
	// Divide by 100 and get number of iterations
	numIters := numCoins/100

	// Check remainder, if not 0 then add 1
	rem := numCoins % 100
	if rem != 0 {
		numIters += 1
	}

	return numIters
}

func ReadConfigFromJSON(file string) Config {
	var data Config

	//FIXME
	filePath := fmt.Sprintf("%v", file)

	dataFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func(dataFile *os.File) {
		err := dataFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(dataFile)

	jsonParser := json.NewDecoder(dataFile)
	err = jsonParser.Decode(&data)

	if err != nil {
		panic("Something Wrong with Config file")
	}
	fmt.Println(data)

	return data
}
