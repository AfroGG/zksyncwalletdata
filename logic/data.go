package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Responsev struct {
	Itemv []Itemv `json:"items"`
	Meta  Meta    `json:"meta"`
}

type Meta struct {
	TotalPages int `json:"totalPages"`
}

type Itemv struct {
	Amount       string `json:"amount"`
	TokenAddress string `json:"tokenAddress"`
	From         string `json:"from"`
}

type Response struct {
	Item []*Item `json:"items"`
	Meta *Meta   `json:"meta"`
}

type Item struct {
	Value      string `json:"value"`
	ReceivedAt string `json:"receivedAt"`
	Fee        string `json:"fee"`
	Nonce      int    `json:"nonce"`
}

type Responses struct {
	Type        string `json:"type"`
	SealedNonce int    `json:"sealedNonce"`
	Balances    struct {
		Eth  ETH  `json:"0x000000000000000000000000000000000000800A"`
		Usdc USDC `json:"0x3355df6D4c9C3035724Fd0e3914dE96A5a83aaf4"`
	} `json:"balances"`
}

type ETH struct {
	Balance string `json:"balance"`
}

type USDC struct {
	Balance string `json:"balance"`
}

func GetValueDetails(address string) float64 {
	ethNum := 0.0
	usdcNum := 0.0
	var txvalue float64
	var valueData1 Responsev
	var valueData Responsev

	url1 := fmt.Sprintf("https://block-explorer-api.mainnet.zksync.io/address/%s/transfers?&limit=100&page=1", address)
	res, err := http.Get(url1)
	// fmt.Println(res)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0
	}

	_ = json.Unmarshal(body, &valueData1)
	// fmt.Println(valueData1)
	// if err != nil {
	// 	return 0
	// }

	for i := 1; i <= valueData1.Meta.TotalPages; i++ {
		url := fmt.Sprintf("https://block-explorer-api.mainnet.zksync.io/address/%s/transfers?&limit=100&page=%d", address, i)

		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		defer response.Body.Close()

		body, _ := ioutil.ReadAll(response.Body)

		_ = json.Unmarshal(body, &valueData)
		// fmt.Println(len(valueData.Item))

		for i := 0; i < len(valueData.Itemv); i++ {
			// fmt.Println(valueData.Item[i].Amount)
			if strings.EqualFold(valueData.Itemv[i].From, address) && valueData.Itemv[i].TokenAddress == "0x000000000000000000000000000000000000800A" {
				ethToken := valueData.Itemv[i].Amount
				// fmt.Println(valueData.Item[i].Amount)
				eth, _ := strconv.ParseFloat(ethToken, 64)
				ethNum += eth / 1e18
			} else if strings.EqualFold(valueData.Itemv[i].From, address) && valueData.Itemv[i].TokenAddress == "0x3355df6D4c9C3035724Fd0e3914dE96A5a83aaf4" {
				usdcToken := valueData.Itemv[i].Amount
				usdc, _ := strconv.ParseFloat(usdcToken, 64)
				usdcNum += usdc / 1e6
			}
		}
	}
	txvalue = ethNum*1700 + usdcNum
	fmt.Println(txvalue)
	return txvalue
}

func GetAddressDetails(address string) (string, int, float64, float64, int, float64, int, string, string, int) {
	var activeDays int = 0
	var activeWeeks int = 0
	var activeMonths int = 0
	var firstActiveday string
	var lastActiveday string
	var fee float64 = 0.0
	var EthBalance float64
	var UsdcBalance float64
	var res *Responses
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println(address)
	urlNonce := "https://block-explorer-api.mainnet.zksync.io/address/"
	res, err := getAccountDetailsNonce(urlNonce, address)
	// fmt.Println(address)
	// fmt.Println(res.SealedNonce)
	// fmt.Println(res.Type)

	if err != nil {
		fmt.Println(err)
	} else {
		f, _ := strconv.ParseFloat(res.Balances.Eth.Balance, 64)
		// fmt.Println(f)
		EthBalance = f / 1e18
		EthBalance = math.Round(EthBalance*1000) / 1000
		// fmt.Println(EthBalance)
		f2, _ := strconv.ParseFloat(res.Balances.Usdc.Balance, 64)
		UsdcBalance = f2 / 1e6
		UsdcBalance = math.Round(UsdcBalance*100) / 100

		url1 := "https://block-explorer-api.mainnet.zksync.io/transactions?page=1&limit=100&address="

		data, _ := getAccountDetails(url1, address)

		var allTimestamps []int64
		for page := 1; page <= data.Meta.TotalPages; page++ {
			url := fmt.Sprintf("https://block-explorer-api.mainnet.zksync.io/transactions?page=%v&limit=100&address=", page)

			data1, err := getAccountDetails(url, address)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if len(data1.Item) == 0 {
				fmt.Println("No data found")
				continue
			}

			fee = getFee(data1.Item, fee) + fee
			fee /= 10000000000000000
			fee = math.Round(fee*100000) / 100000
			ts := make([]int64, 0)
			for _, item := range data1.Item {

				t, err := time.Parse("2006-01-02T15:04:05.999Z", item.ReceivedAt)
				if err != nil {
					fmt.Println(err)
					continue
				}
				if page == 1 {

					lastActiveday = data1.Item[0].ReceivedAt

				}
				if page == data.Meta.TotalPages {

					firstActiveday = data1.Item[len(data1.Item)-1].ReceivedAt

				}
				timestamp := t.Unix()
				ts = append(ts, timestamp)
				sort.Slice(ts, func(i, j int) bool {
					return ts[i] < ts[j]
				})
			}
			allTimestamps = append(allTimestamps, ts...)
		}
		sort.Slice(allTimestamps, func(i, j int) bool {
			return allTimestamps[i] < allTimestamps[j]
		})
		activeDays = getActiveDays(allTimestamps)
		activeWeeks = getActiveWeeks(allTimestamps)
		activeMonths = getMonths(allTimestamps)
	}
	fmt.Println(res.Type, res.SealedNonce, EthBalance, UsdcBalance, activeWeeks, fee, activeDays, firstActiveday, lastActiveday)
	return res.Type, res.SealedNonce, EthBalance, UsdcBalance, activeWeeks, fee, activeDays, firstActiveday, lastActiveday, activeMonths

}

func getAccountDetails(url, address string) (*Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url+address, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.203")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var blockData Response
	err = json.Unmarshal(body, &blockData)
	if err != nil {
		return nil, err
	}

	return &blockData, nil
}

func removeRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func getFee(items []*Item, fee float64) float64 {
	var v int = 2
	for _, item := range items {
		n, _ := strconv.ParseInt(item.Fee[2:], 16, 64)
		res := Unwrap(n, v)
		fee = res + fee
	}
	return fee
}
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

func generateIntervals(startTimestamp, endTimestamp int64, intervalSize int64) [][]int64 {
	var intervals [][]int64
	for startTimestamp < endTimestamp {
		interval := []int64{startTimestamp, startTimestamp + intervalSize}
		intervals = append(intervals, interval)
		startTimestamp += intervalSize
	}
	return intervals
}

func getActiveWeeks(timestamp []int64) int {
	var intervals [][]int64
	if len(timestamp) == 0 {
		return 0
	}
	intervals = generateIntervals(timestamp[0], timestamp[len(timestamp)-1], 604800)
	activeWeekCount := 0
	for _, interval := range intervals {
		start := interval[0]
		end := interval[1]
		activeCount := 0
		for _, active := range timestamp {
			if active >= start && active <= end {
				activeCount++
			}
		}
		if activeCount > 0 {
			activeWeekCount++
		}
	}
	return activeWeekCount
}

func getActiveDays(timestamp []int64) int {
	var intervals [][]int64
	if len(timestamp) == 0 {
		return 0
	}
	intervals = generateIntervals(timestamp[0], timestamp[len(timestamp)-1], 86400)
	activeDayCount := 0
	for _, interval := range intervals {
		start := interval[0]
		end := interval[1]
		activeCount := 0
		for _, active := range timestamp {
			if active >= start && active <= end {
				activeCount++
			}
		}
		if activeCount > 0 {
			activeDayCount++
		}
	}
	return activeDayCount
}

func getAccountDetailsNonce(url, address string) (*Responses, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+address, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.203")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var blockData Responses
	err = json.Unmarshal(body, &blockData)
	if err != nil {
		return nil, err
	}

	return &blockData, nil
}

func getMonths(timestamp []int64) int {
	var intervals [][]int64
	if len(timestamp) == 0 {
		return 0
	}
	intervals = generateIntervals(timestamp[0], timestamp[len(timestamp)-1], 86400*30)
	activeMonthsCount := 0
	for _, interval := range intervals {
		start := interval[0]
		end := interval[1]
		activeCount := 0
		for _, active := range timestamp {
			if active >= start && active <= end {
				activeCount++
			}
		}
		if activeCount > 0 {
			activeMonthsCount++
		}
	}
	return activeMonthsCount
}
