package kucoin

import (
	"encoding/json"
	"fmt"
)

// SpecificDealtOrder struct represents kucoin data model.
type SpecificDealtOrder struct {
	Datas []struct {
		Oid       string  `json:"oid"`
		DealPrice float64 `json:"dealPrice"`
		OrderOid  string  `json:"orderOid"`
		Direction string  `json:"direction"`
		Amount    float64 `json:"amount"`
		DealValue float64 `json:"dealValue"`
		CreatedAt int64   `json:"createdAt"`
	} `json:"datas"`
	Total           int         `json:"total"`
	Limit           int         `json:"limit"`
	PageNos         int         `json:"pageNos"`
	CurrPageNo      int         `json:"currPageNo"`
	NavigatePageNos []int       `json:"navigatePageNos"`
	UserOid         string      `json:"userOid"`
	Direction       interface{} `json:"direction"`
	StartRow        int         `json:"startRow"`
	FirstPage       bool        `json:"firstPage"`
	LastPage        bool        `json:"lastPage"`
}

type rawSpecificDealtOrder struct {
	Success bool               `json:"success"`
	Code    string             `json:"code"`
	Msg     string             `json:"msg"`
	Data    SpecificDealtOrder `json:"data"`
}

// DealtOrder struct represents kucoin data model.
type DealtOrder struct {
	Timestamp int64
	Side      string
	DealPrice float64
	DealQty   float64
	DealValue float64
}

//DealtOrders ...
type DealtOrders []DealtOrder

//UnmarshalJSON method for DealtOrder format
func (tp *DealtOrder) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Printf("Error whilde decoding %v\n", err)
		return err
	}
	tp.Timestamp = int64(v[0].(float64))
	tp.Side = string(v[1].(string))
	tp.DealPrice = float64(v[2].(float64))
	tp.DealQty = float64(v[3].(float64))
	tp.DealValue = float64(v[4].(float64))
	return nil
}

type rawDealtOrder struct {
	Success   bool         `json:"success"`
	Code      string       `json:"code"`
	Msg       string       `json:"msg"`
	Timestamp int64        `json:"timestamp"`
	Data      []DealtOrder `json:"data"`
}

// MergedDealtOrder struct represents kucoin data model.
type MergedDealtOrder struct {
	Total int `json:"total"`
	Datas []struct {
		CreatedAt     int64   `json:"createdAt"`
		Amount        float64 `json:"amount"`
		DealValue     float64 `json:"dealValue"`
		DealPrice     float64 `json:"dealPrice"`
		Fee           float64 `json:"fee"`
		FeeRate       float64 `json:"feeRate"`
		Oid           string  `json:"oid"`
		OrderOid      string  `json:"orderOid"`
		CoinType      string  `json:"coinType"`
		CoinTypePair  string  `json:"coinTypePair"`
		Direction     string  `json:"direction"`
		DealDirection string  `json:"dealDirection"`
	} `json:"datas"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type rawMergedDealtOrder struct {
	Success   bool             `json:"success"`
	Code      string           `json:"code"`
	Msg       string           `json:"msg"`
	Timestamp int64            `json:"timestamp"`
	Data      MergedDealtOrder `json:"data"`
}
