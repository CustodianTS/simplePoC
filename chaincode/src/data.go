package main

import (
)

// DATA STRUCTURE TO CAPTURE INVESTOR DETAILS BY CUSTODIAN
// Key consists of 01 + UserName
type investor struct {
    UserName     string  `json:"user_name"`
    UserFName    string  `json:"user_fname"`
    UserLName    string  `json:"user_lname"`
    DepositoryAC string  `json:"depository_ac"`
    BankAC       string  `json:"bank_ac"`
}

// Key consists of 02 + UserName
type investorPortfolio struct {
    StockTicker string  `json:"stock_ticker"`
    StockQty    int64   `json:"stock_qty"`
    StockPrice  float64 `json:"stock_price"`
    StockValue  float64 `json:"stock_value"`
}

// Key consists of 03 + UserName
type investorTrades struct {
    TradeTimeStamp string    `json:trade_timestamp"`
    TradeType      string    `json:"trade_type"`
    StockTicker    string    `json:"stock_ticker"`
    StockQty       int64     `json:"stock_qty"`
    StockPrice     float64   `json:"stock_price"`
    StockValue     float64   `json:"stock_value"`
}

// DATA STRUCTURE TO CAPTURE TRANSACTION DETAILS BY BANK
// Key consists of 04 + UserName
type bankMaster struct {
    UserName    string    `json:"user_name"`
    BankAC      string    `json:"bank_ac"`
    Balance     float64   `json:"balance"`
}

// Key consists of 05 + UserName
type bankTransactions struct {
    TransTimestamp string    `json:"trans_timestamp"`
    BankAC         string    `json:"bank_ac"`
    TransAmount    float64   `json:"trans_amount"`
    Balance        float64   `json:"balance"`
}

// DATA STRUCTURE TO CAPTURE TRADING DETAILS BY EXCHANGE
// Key consists of 06 + StockTicker
type exchangeMaster struct {
    StockTicker string    `json:"stock_ticker"`
    StockQty    int64     `json:"stock_qty"`
    StockPrice  float64   `json:"stock_price"`
}
// Key consists of 07 + UserName
type exchangeTrades struct {
    UserName       string    `json:"user_name"`
    TradeTimestamp string    `json:"trade_timestamp"`
    StockTicker    string    `json:"stock_ticker"`
    StockQty       int64     `json:"stock_qty"`
    StockPrice     float64   `json:"stock_price"`
    StockValue     float64   `json:"stock_value"`
}
