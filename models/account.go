package models

import (
	"fmt"
	"goweb/logic"
	"goweb/utils"
	"strings"
	"sync"
)

type Account struct { //创建钱包结构体
	Address         string  `gorm:"column:address" json:"address"`
	Nonce           int     `gorm:"column:Nonce" json:"nonce"`
	TotalTxValue    float64 `gorm:"column:txValue" json:"totalTxValue"`
	TotalFee        float64 `gorm:"column:TotalFee" json:"totalFee"`
	EthBalance      float64 `gorm:"column:EthBalance" json:"ethBalance"`
	UsdcBalance     float64 `gorm:"column:UsdcBalance" json:"usdcBalance"`
	ActiveDay       int     `gorm:"column:ActiveDay" json:"activeDay"`
	ActiveWeek      int     `gorm:"column:ActiveWeek" json:"activeWeek"`
	StartDay        string  `gorm:"column:StartDay" json:"startDay"`
	EndDay          string  `gorm:"column:EndDay" json:"endDay"`
	Type            string  `gorm:"column:Type" json:"type"`
	ZkMonth         int     `gorm:"column:zkMonth" json:"zkMonth"`
	Zklite_nonce    int     `gorm:"column:zklite_nonce" json:"zklite_nonce"`
	Zklite_txvalue  float64 `gorm:"column:zklite_txvalue" json:"zklite_txvalue"`
	Zklite_week     int     `gorm:"column:zklite_week" json:"zklite_week"`
	Zklite_month    int     `gorm:"column:zklite_month" json:"zklite_month"`
	Zklite_startDay string  `gorm:"column:zklite_startday" json:"zklite_startDay"`
	Zklite_eth      float64 `gorm:"column:zklite_eth" json:"zklite_eth"`
	Zklite_usdc     float64 `gorm:"column:zklite_usdc" json:"zklite_usdc"`
	Score           int32   `gorm:"column:score" json:"score"`
	Rank            int32   `gorm:"column:ranking" json:"rank"`
}

type AddressSummary struct {
	Type            string  `gorm:"column:Type" json:"type"`
	Nonce           int     `gorm:"column:Nonce" json:"nonce"`
	EthBalance      float64 `gorm:"column:EthBalance" json:"ethBalance"`
	UsdcBalance     float64 `gorm:"column:UsdcBalance" json:"usdcBalance"`
	ActiveWeek      int     `gorm:"column:ActiveWeek" json:"activeWeek"`
	TotalFee        float64 `gorm:"column:TotalFee" json:"totalFee"`
	ActiveDay       int     `gorm:"column:ActiveDay" json:"activeDay"`
	StartDay        string  `gorm:"column:StartDay" json:"startDay"`
	EndDay          string  `gorm:"column:EndDay" json:"endDay"`
	TxValue         float64 `gorm:"column:txValue" json:"totalTxValue"`
	ZkMonth         int     `gorm:"column:zkMonth" json:"zkMonth"`
	Zklite_nonce    int     `gorm:"column:zklite_nonce" json:"zklite_nonce"`
	Zklite_txvalue  float64 `gorm:"column:zklite_txvalue" json:"zklite_txvalue"`
	Zklite_week     int     `gorm:"column:zklite_week" json:"zklite_week"`
	Zklite_month    int     `gorm:"column:zklite_month" json:"zklite_month"`
	Zklite_startday string  `gorm:"column:zklite_startday" json:"zklite_startDay"`
	Zklite_eth      float64 `gorm:"column:zklite_eth" json:"zklite_eth"`
	Zklite_usdc     float64 `gorm:"column:zklite_usdc" json:"zklite_usdc"`
}

type SurveyData struct {
	Nonce       int
	EthBalance  float64
	UsdcBalance float64
	TotalFee    float64
	ActiveDay   int
	ActiveWeek  int
	StartDay    string
}

func (table *Account) TableName() string {
	return "account"
}

func GetAddressDetails(address string) *Account {
	account := &Account{}
	utils.DB.Find(&account, "address = ?", address)
	return account
}

func BatchQuery(address string) []Account {
	accounts := make([]Account, 0)
	addresses := strings.Split(address, " ")
	for _, addr := range addresses {
		var account Account
		utils.DB.Model(&account).Where("address = ?", addr).Find(&account)
		accounts = append(accounts, account)
	}
	fmt.Println(accounts)
	return accounts
}

func Survey(nonce, txValue, totalfee, activedays, activeweeks, usdcbalance, ethbalance, zklite_nonce, zklite_month, zklite_week, zklite_txvalue, zklite_eth, zklite_usdc string) (int64, int64) {
	var count int64
	var countTotal int64
	var account Account
	utils.DB.Model(&account).Where("txValue >= ? AND ActiveWeek >= ? AND Nonce >= ? AND EthBalance >=? AND UsdcBalance >= ? AND TotalFee >= ? AND ActiveDay >= ? AND zklite_nonce >= ? AND zklite_month >= ? AND zklite_week >= ? AND zklite_txvalue >= ? AND zklite_eth >= ? AND zklite_usdc >= ?", txValue, activeweeks, nonce, ethbalance, usdcbalance, totalfee, activedays, zklite_nonce, zklite_month, zklite_week, zklite_txvalue, zklite_eth, zklite_usdc).Count(&count)
	utils.DB.Model(&account).Count(&countTotal)
	return count, countTotal
}

func Update(addresses string) ([]Account, error) {
	addressList := strings.Split(addresses, " ")
	var updatedAccounts []Account
	updatedAccounts = make([]Account, 0)
	var wg sync.WaitGroup

	for _, address := range addressList {
		wg.Add(1)

		go func(addr string) {
			defer wg.Done()

			var account Account
			// fmt.Println(addr)
			zklite_txvalue, zklite_week, zklite_month, zklite_startDay, zklite_nonce := logic.GetValue(addr)
			zklite_eth, zklite_usdc := logic.GetBalance(addr)
			Type, SealedNonce, EthBalance, UsdcBalance, activeWeeks, fee, activeDays, firstActiveday, lastActiveday, zkMonth := logic.GetAddressDetails(addr)
			txValue := logic.GetValueDetails(addr)
			fmt.Println(UsdcBalance)

			updateData := AddressSummary{
				Type:            Type,
				Nonce:           SealedNonce,
				EthBalance:      EthBalance,
				UsdcBalance:     UsdcBalance,
				ActiveWeek:      activeWeeks,
				TotalFee:        fee,
				ActiveDay:       activeDays,
				StartDay:        firstActiveday,
				EndDay:          lastActiveday,
				TxValue:         txValue,
				ZkMonth:         zkMonth,
				Zklite_nonce:    zklite_nonce,
				Zklite_txvalue:  zklite_txvalue,
				Zklite_week:     zklite_week,
				Zklite_month:    zklite_month,
				Zklite_startday: zklite_startDay,
				Zklite_eth:      zklite_eth,
				Zklite_usdc:     zklite_usdc,
			}
			fields := []string{"Nonce", "txValue", "TotalFee", "EthBalance", "UsdcBalance", "ActiveDay", "ActiveWeek", "StartDay", "EndDay", "Type", "zklite_nonce", "zklite_txvalue", "zklite_week", "zklite_month", "zklite_startday", "zklite_eth", "zklite_usdc"}

			if err := utils.DB.Model(&account).Select(fields).Where("address = ?", addr).Updates(&updateData).Error; err != nil {
				fmt.Printf("Update for address %s failed: %v\n", addr, err)
				return
			}
			// fmt.Println(account)

			if err := utils.DB.Model(&account).Where("address = ?", addr).Find(&account).Error; err != nil {
				fmt.Printf("Find for address %s failed: %v\n", addr, err)
				return
			}
			// fmt.Println(account)
			updatedAccounts = append(updatedAccounts, account)
		}(address)
	}

	wg.Wait()

	return updatedAccounts, nil
}
