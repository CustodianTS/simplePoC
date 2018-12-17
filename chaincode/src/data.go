package main

import (
    "time"
)

// DATA STRUCTURE TO CAPTURE INVESTOR DETAILS BY CUSTODIAN
// Key consists of custodianPrefix + userName
type investor struct {
    userName     string  `json:"user_name"`
    userFName    string  `json:"user_fname"`
    userLName    string  `json:"user_lname"`
    userIdentity string  `json:"user_identity"`
    kycStatus    string  `json:"kyc_status"`
    depositoryAC string  `json:"depository_ac"`
    bankAC       string  `json:"bank_ac"`
}

// Key consists of custodianPrefix + userName
type investorPortfolio struct {
    stockTicker string  `json:"stock_ticker"`
    stockQty    int32   `json:"stock_qty"`
    stockPrice  float32 `json:"stock_price"`
    stockValue  float32 `json:"stock_value"`
}

// Key consists of custodianPrefix + userName + tradeUUID
type investorTrade struct {
    tradeUUID   string    `json:"trade_uuid"`
    tradeDate   time.Time `json:"trade_date"`
    tradeType   string    `json:"trade_type"`
    stockTicker string    `json:"stock_ticker"`
    stockQty    int32     `json:"stock_qty"`
    stockPrice  float32   `json:"stock_price"`
    stockValue  float32   `json:"stock_value"`
}

// DATA STRUCTURE TO CAPTURE TRANSACTION DETAILS BY BANK
// Key consists of bankPrefix + userName
type bankMaster struct {
    userName    string    `json:"user_name"`
    bankAC      string    `json:"bank_ac"`
    balance     float64   `json:"balance"`
}

// Key consists of bankPrefix + userName + transUUID
type bankTransaction struct {
    transUUID   string    `json:"trans_uuid"`
    userName    string    `json:"user_name"`
    bankAC      string    `json:"bank_ac"`
    transDate   time.Time `json:"trans_date"`
    transAmount float64   `json:"trans_amount"`
    balance     float64   `json:"balance"`
}

// DATA STRUCTURE TO CAPTURE TRADING DETAILS BY EXCHANGE
// Key consists of exchangePrefix only
type exchangeMaster struct {
    stockTicker string    `json:"stock_ticker"`
    stockQty    int32     `json:"stock_qty"`
    stockPrice  float32   `json:"stock_price"`
}
// Key consists of exchangePrefix + tradeUUID
type exchangeTrades struct {
    tradeUUID   string    `json:"trade_uuid"`
    tradeDate   time.Time `json:"trade_date"`
    stockTicker string    `json:"stock_ticker"`
    stockQty    int32     `json:"stock_qty"`
    stockPrice  float32   `json:"stock_price"`
    stockValue  float32   `json:"stock_value"`
}

// DATA STRUCTURE TO CAPTURE TRANSACTION DETAILS BY DEPOSITORY
// Key consists of depositoryPrefix + userName + transUUID
type depositoryTransaction struct {
    transUUID   string    `json:"trans_uuid"`
    userName    string    `json:"user_name"`
    depositoryAC string   `json:"depository_ac"`
    tradeDate   time.Time `json:"trade_date"`
    stockTicker string    `json:"stock_ticker"`
    stockQty    int32     `json:"stock_qty"`
    stockPrice  float32   `json:"stock_price"`
    stockValue  float32   `json:"stock_value"`
}
