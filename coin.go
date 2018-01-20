package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
)

func return_subslice(li []string, low int, high int) []string {
	if low > len(li)-1 {
		fmt.Println("\033[H\033[2J")
		show(li)
	}

	subslice := li[low:]

	if high < len(li)-1 {
		subslice = li[low:high]
	}

	return subslice
}

func show(coins []string) {
	low      := 0
	high     := 10
	iterator := 10

	for {
		subslice := return_subslice(coins, low, high)

		var sos [][]string

		var wg sync.WaitGroup
		wg.Add(len(subslice))

		for _, coin := range subslice {
			go func(coin string, wg *sync.WaitGroup) {
				slice := cmc(coin, wg)
				sos    = append(sos, slice)
			}(coin, &wg)
		}

		wg.Wait()

		render_table(sos)

		time.Sleep(time.Second * 65)

		low  += iterator
		high += iterator
	}
}

func return_response(coin string) []byte {
	url     := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s/", coin)
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return body
}

type Response struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Volume24         string `json:"24h_volume_usd"`
	MarketCap        string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	Lastcsvd         string `json:"last_csvd"`
}

type Responses []Response

func unmarshal_response(body []byte) Responses {
	var responses Responses
	json.Unmarshal(body, &responses)

	return responses
}

func return_slice(responses Responses) []string {
	response   := responses[0]
	rank       := response.Rank
	name       := response.Name
	price      := response.PriceUsd
	market_cap := response.MarketCap
	volume24h  := response.Volume24
	change1h   := response.PercentChange1H
	change24h  := response.PercentChange24H
	slice      := []string{rank, name, price, market_cap, volume24h, change1h, change24h}

	return slice
}

func cmc(coin string, wg *sync.WaitGroup) []string {
	body      := return_response(coin)
	responses := unmarshal_response(body)
	slice     := return_slice(responses)

	wg.Done()

	return slice
}

func render_table(sos [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rank", "Name", "Price", "Market Cap", "Volume (24h)", "Change (1h)", "Change (24h)"})
	table.SetBorder(true)
	table.AppendBulk(sos)
	table.Render()

	return
}

func main() {
	filename := flag.String(
		"f",
		"coins.txt",
		"time to sleep in seconds")
	flag.Parse()

	contents, _ := ioutil.ReadFile(*filename)
	coins       := strings.Split(string(contents), "\n")

	show(coins)
}

