// package main

// import "time"

// func main() {
// 	godur, _ := time.ParseDuration("10ms")

// 	// >
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			println("Hello")
// 			time.Sleep(godur)
// 		}
// 	}()

// 	// >
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			println("go")
// 			time.Sleep(godur)
// 		}
// 	}()

// 	dur, _ := time.ParseDuration("1s")
// 	time.Sleep(dur)

// }
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=googl
func main() {
	t := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}
	var numComplete int
	for _, symbol := range stockSymbols {
		go func(symbol string) {
			res, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)
			defer res.Body.Close()

			body, _ := ioutil.ReadAll(res.Body)
			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}

	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}
	elapsed := time.Since(t)

	fmt.Printf("Excution time %s", elapsed)
}

type QuoteResponse struct {
	Status        string
	Name          string
	LastPrice     float32
	Change        float32
	ChangePercent float32
	Timestamp     string
	MSDate        float32
	// MarketCap
	// Volume
	// ChangeYTD
	// ChangePercentYTD
	// High
	// Low
	// Open
}
