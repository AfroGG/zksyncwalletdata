package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ZkliteRes struct {
	Tx         Tx     `json:"tx"`
	Created_at string `json:"created_at"`
}

type Tx struct {
	From   string  `json:"from"`
	Type   string  `json:"type"`
	Orders []Order `json:"orders"`
	Nonce  int     `json:"nonce"`
	Token  string  `json:"token"`
	Amount string  `json:"amount"`
}

type Order struct {
	Nonce     int    `json:"nonce"`
	Amount    string `json:"amount"`
	Recipient string `json:"recipient"`
	TokenSell int    `json:"tokenSell"`
}

type ResponseLite struct {
	Result Result `json:"result"`
}

type Result struct {
	Committed Committed `json:"committed"`
}

type Committed struct {
	Balances Balance `json:"balances"`
}
type Balance struct {
	ETH  string `json:"ETH"`
	USDC string `json:"USDC"`
}

// func main() {
// 	db, err := sql.Open("mysql", "root:123456@/zkdata")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	var addressChan chan string
// 	addressChan = make(chan string, 2000)
// 	go InToAddress(db, addressChan)
// 	for i := 0; i < 20; i++ {
// 		go func() {
// 			for {
// 				address := <-addressChan
// 				// fmt.Println(address)
// 				totalValue, activeWeeks, activeMonths, startDay, nonce := GetValue(address)
// 				ETH, USDC := getBalance(address)
// 				stmt, err := db.Prepare("UPDATE addressessum SET zklite_nonce=?, zklite_txvalue=?, zklite_week=?, zklite_month=?, zklite_startDay=?,zklite_eth=?, zklite_usdc=? WHERE address=?;")
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				fmt.Println(nonce, totalValue, activeWeeks, activeMonths, startDay, ETH, USDC, address)
// 				defer stmt.Close()
// 				_, err = stmt.Exec(nonce, totalValue, activeWeeks, activeMonths, startDay, ETH, USDC, address)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 			}
// 		}()
// 	}

// 	for {

//		}
//	}
func GetNonce(address string) int {
	var valueData []ZkliteRes
	var nonce int
	url := fmt.Sprintf("https://api.zksync.io/api/v0.1/account/%s/history/0/100", address)
	res, err := http.Get(url)
	// fmt.Println(res)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	}

	if len(body) == 2 && body[0] == 91 && body[1] == 93 {
		return 0
	}
	// fmt.Println(body)
	if err != nil {
		return 0
	}
	err = json.Unmarshal(body, &valueData)
	// fmt.Println(valueData)

	for _, v := range valueData {
		if v.Tx.Type == "Transfer" && v.Tx.From == address {
			nonce = v.Tx.Nonce
			break
		} else if v.Tx.Type == "Swap" {
			nonce = v.Tx.Orders[0].Nonce
			break
		}
	}
	return nonce
}

func GetValue(address string) (float64, int, int, string, int) {
	var valueData []ZkliteRes
	var totalValue float64
	var activeWeeks int
	var activeMonths int
	var startDay string
	var allTimestamps []int64
	nonce := GetNonce(address)
	n := nonce / 100
	// fmt.Println(n)
	for i := 0; i <= n; i++ {
		url := fmt.Sprintf("https://api.zksync.io/api/v0.1/account/%s/history/%d00/100", address, i)
		// fmt.Println(url)
		res, err := http.Get(url)
		// fmt.Println(res)
		if err != nil {
			fmt.Println(err)
			return 0, 0, 0, "", 0
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
		}

		if len(body) == 2 && body[0] == 91 && body[1] == 93 {
			return 0, 0, 0, "", 0
		}

		if err != nil {
			return 0, 0, 0, "", 0
		}
		_ = json.Unmarshal(body, &valueData)
		// fmt.Println(valueData)
		ts := make([]int64, 0)
		for _, v := range valueData {
			t, err := time.Parse("2006-01-02T15:04:05.999Z", v.Created_at)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if v.Tx.Type == "ChangePubKey" {
				startDay = v.Created_at
			}
			timestamp := t.Unix()
			ts = append(ts, timestamp)
			sort.Slice(ts, func(i, j int) bool {
				return ts[i] < ts[j]
			})
			allTimestamps = append(allTimestamps, ts...)
			sort.Slice(allTimestamps, func(i, j int) bool {
				return allTimestamps[i] < allTimestamps[j]
			})
			activeMonths = getActiveMonths(allTimestamps)
			activeWeeks = getActiveWeeks(allTimestamps)

			if v.Tx.Type == "Swap" && v.Tx.Orders[0].TokenSell == 2 {
				s_val, _ := strconv.Atoi(v.Tx.Orders[0].Amount)
				val := float64(s_val) / 1e6
				totalValue = totalValue + val
			} else if v.Tx.Type == "Swap" && v.Tx.Orders[0].TokenSell == 0 {
				s_val, _ := strconv.Atoi(v.Tx.Orders[0].Amount)
				val := (float64(s_val) / 1e18) * 1700
				totalValue = totalValue + val
			} else if v.Tx.Type == "Transfer" && v.Tx.From == address && v.Tx.Token == "ETH" {
				s_val, _ := strconv.Atoi(v.Tx.Amount)
				val := (float64(s_val) / 1e18) * 1700
				totalValue = totalValue + val
			} else if v.Tx.Type == "Transfer" && v.Tx.From == address && v.Tx.Token == "USDC" {
				s_val, _ := strconv.Atoi(v.Tx.Amount)
				val := float64(s_val) / 1e6
				totalValue = totalValue + val
			} else {
				totalValue = totalValue + 0
			}
		}
	}
	return totalValue, activeWeeks, activeMonths, startDay, nonce
}

func getActiveMonths(timestamp []int64) int {
	var intervals [][]int64
	if len(timestamp) == 0 {
		return 0
	}
	intervals = generateIntervals(timestamp[0], timestamp[len(timestamp)-1], 86400*30)
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

func GetBalance(address string) (float64, float64) {
	var EthBalance float64
	var UsdcBalance float64
	var Data ResponseLite
	url := fmt.Sprintf("https://api.zksync.io/api/v0.2/accounts/%s", address)
	// fmt.Println(url)
	res, err := http.Get(url)
	// fmt.Println(res)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return 0, 0
	}
	_ = json.Unmarshal(body, &Data)
	// if Data.Result.Committed.Balances.ETH == nil {

	// }
	// fmt.Println(Data)
	ETH, _ := strconv.Atoi(Data.Result.Committed.Balances.ETH)
	EthBalance = float64(ETH) / 1e18
	USDC, _ := strconv.Atoi(Data.Result.Committed.Balances.USDC)
	UsdcBalance = float64(USDC) / 1e6
	return EthBalance, UsdcBalance
}
