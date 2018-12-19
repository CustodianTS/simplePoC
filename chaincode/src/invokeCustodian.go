package main

import (
    "fmt"
    "encoding/json"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// METHOD TO ONBOARD AN INVESTOR - CREATE AN INVESTOR RECORD
// USES PREFIX01, PREFIX01IDX FOR COMPOSITE KEY - FOR INVESTOR
func onboardInvestor(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Println("****************************************")
    fmt.Println("---------- IN ONBOARDINVESTOR ----------")

    // RETURN ERROR IF ARGS IS NOT 5 IN NUMBER
    if len(args) != 5 {
        fmt.Println("**************************")
        fmt.Println("Too few argments... Need 5")
        fmt.Println("**************************")

        return shim.Error("Invalid argument count. Expecting 5.")
    }
    // LOG THE ARGUMENTS
    fmt.Println("**** Arguments To Function ****")
    fmt.Println("User Name     : ", args[0])
    fmt.Println("User FName    : ", args[1])
    fmt.Println("User LName    : ", args[2])
    fmt.Println("Depository AC : ", args[3])
    fmt.Println("Bank AC       : ", args[4])

    // SET ARGUMENTS INTO LOCAL STRUCTURE
    _investor := investor {
        UserName:     args[0],
        UserFName:    args[1],
        UserLName:    args[2],
        DepositoryAC: args[3],
        BankAC:       args[4],
    }

    // PREPARE THE KEY VALUE PAIR TO PERSIST THE INVESTOR
    _investorKey, err := stub.CreateCompositeKey(PREFIX01IDX, []string{PREFIX01, _investor.UserName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }

    // MARSHAL INVESTOR RECORD
    _investorBytes, err := json.Marshal(_investor)
    // CHECK FOR ERROR IN MARSHALING
    if err != nil {
        return shim.Error(err.Error())
    }

    // NOW WRITE THE INVESTOR RECORD
    err = stub.PutState(_investorKey, _investorBytes)
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }

    fmt.Println("---------- OUT ONBOARDINVESTOR----------")
    fmt.Println("****************************************")

    // RETURN SUCCESS
    return shim.Success(_investorBytes)
}

// METHOD TO TRADE AN ASSET - BUY OR SELL
// PARAMETERS: 1. USERNAME 2. TRADE TYPE 3. STOCK TICKER 4. STOCK QTY 5. STOCK PRICE
// USES PREFIX02, PREFIX02IDX FOR COMPOSITE KEY - INVESTOR PORTFOLIO
// USES PREFIX03, PREFIX03IDX FOR COMPOSITE KEY - INVESTOR TRADES
func tradeAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Println("**********************************************")
    fmt.Println("---------- IN TRADE ASSET CUSTODIAN ----------")

    // RETURN ERROR IF ARGS IS NOT 5 IN NUMBER
    if len(args) != 5 {
        fmt.Println("**************************")
        fmt.Println("Too few argments... Need 5")
        fmt.Println("**************************")
        return shim.Error("Invalid argument count. Expecting 5.")
    }

    // SET ARGUMENTS INTO LOCAL VARIABLES
    _userName := args[0]
    _tradeType := args[1]
    _stockTicker := args[2]
    _stockQty,_ := strconv.ParseInt(args[3], 10, 64)
    _stockPrice, _ := strconv.ParseFloat(args[4], 64)

    // LOG THE ARGUMENTS
    fmt.Println("**** Arguments To Function ****")
    fmt.Println("User Name   : ", _userName)
    fmt.Println("Trade Type  : ", _tradeType)
    fmt.Println("Stock Ticker: ", _stockTicker)
    fmt.Println("Stock Qty   : ", _stockQty)
    fmt.Println("Stock Price : ", _stockPrice)

    // CALL BUY IF TRADE TYPE IS BUY
    // CALL SELL IF TRADE TYPE IS SELL
    _sale := false
    if _tradeType == BUY {
        _sale = buyAsset(stub, _userName, _stockTicker, _stockQty, _stockPrice)
    } else if _tradeType == SELL {
        _sale = sellAsset(stub, _userName, _stockTicker, _stockQty, _stockPrice)
    }
    if (!_sale) {
        fmt.Println("tradeAsset: ", _tradeType, " failed")
        return shim.Error("tradeAsset failed")
    }
    // NOW PREPARE INVESTOR TRADES RECORD TO WRITE
    _currentTimeTS := getTimeStamp()
    _stockValue := float64(_stockQty) * _stockPrice
    _investorTrade := investorTrades {
        TradeTimestamp: _currentTimeTS,
        TradeType:      _tradeType,
        StockTicker:    _stockTicker,
        StockQty:       _stockQty,
        StockPrice:     _stockPrice,
        StockValue:     _stockValue,
    }

    // PREPARE THE KEY TO WRITE INVESTOR TRADE
    _investorTradeKey, err := stub.CreateCompositeKey(PREFIX03IDX, []string{PREFIX03, _userName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("tradeAsset: Prepare Investor Trade Key Completed")

    // MARSHAL THE INVESTOR TRADE RECORD
    _investorTradeAsBytes, err := json.Marshal(_investorTrade)
    // CHECK FOR ERROR IN MARSHALING
    if err != nil {
        return shim.Error(err.Error())
    }

    // NOW WRITE THE INVESTOR TRADE RECORD
    err = stub.PutState(_investorTradeKey, _investorTradeAsBytes)
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("tradeAsset: Writing Investor Trade Completed")

    fmt.Println("---------- OUT TRADE ASSET CUSTODIAN ----------")
    fmt.Println("***********************************************")

    // RETURN SUCCESS
    return shim.Success(_investorTradeAsBytes)
}

// METHOD TO BUY AN ASSET
// PARAMETERS: 1. USERNAME 2. STOCK TICKER 3. STOCK QTY 4. STOCK PRICE
// USES PREFIX02, PREFIX02IDX FOR COMPOSITE KEY - INVESTOR PORTFOLIO
func buyAsset(stub shim.ChaincodeStubInterface, _userName string, _stockTicker string, _stockQty int64, _stockPrice float64) bool {

    fmt.Println("**********************************************")
    fmt.Println("---------- IN BUY ASSET CUSTODIAN ----------")
    
    // FETCH THE RECORD FROM INVESTOR PORTFOLIO FOR UPDATE
    // PREPARE THE KEY TO READ INVESTOR PORTFOLIO
    _investorPortfolioKey, err := stub.CreateCompositeKey(PREFIX02IDX, []string{PREFIX02, _userName, _stockTicker})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return false
    }
    fmt.Println("buyAsset: Prepare Investor Trade Key Completed")

    // STRUCTURE TO GET THE INVESTOR PORTFOLIO RECORD
    _investorPortfolio := investorPortfolio {}
    _xStockQty := _stockQty

    // USE THE KEY TO RETRIEVE INVESTOR PORTOFOLIO FOR THIS STOCK TICKER
    _investorPortfolioAsBytes, err := stub.GetState(_investorPortfolioKey)

    // IF THE STOCK TICKER EXISTS IN PORTFOLIO GET THE STOCKQTY
    if _investorPortfolioAsBytes != nil {
        fmt.Println("buyAsset: Record(s) Retrieved by GetState")
        fmt.Println(string(_investorPortfolioAsBytes))

        // NOW UNMARSHALL THE INVESTOR PORTFOLIO RECORD
        err = json.Unmarshal(_investorPortfolioAsBytes, &_investorPortfolio)
        // CHECK FOR ERROR IN UNMARSHALLING
        if err != nil {
            return false
        }
        // NOW UPDATE THE STOCKQTY AS WE ARE ADDING TO EXISTING PORTFOLIO
        _xStockQty = _investorPortfolio.StockQty + _xStockQty
    }
    
    // CALCULATE THE NEW STOCK VALUE
    _stockValue := float64(_xStockQty) * _stockPrice

    // NOW PREPARE TO WRITE / UPDATE THE INVESTOR PORTFOLIO
    _investorPortfolio.StockQty = _xStockQty
    _investorPortfolio.StockValue = _stockValue
    _investorPortfolio.StockPrice = _stockPrice
    fmt.Println("buyAsset: New StockQty  : ", _xStockQty)
    fmt.Println("buyAsset: New StockValue: ", _stockValue)
    fmt.Println("buyAsset: StockPrice    : ", _stockPrice)

    // MARSHAL THE INVESTOR PORTFOLIO RECORD
    _investorPortfolioAsBytes, err = json.Marshal(_investorPortfolio)
    // CHECK FOR ERROR IN MARSHALING
    if err != nil {
        return false
    }

    // NOW WRITE THE INVESTOR PORTFOLIO RECORD
    err = stub.PutState(_investorPortfolioKey, _investorPortfolioAsBytes)
    // CHECK FOR ERROR
    if err != nil {
        return false
    }
    fmt.Println("buyAsset: Writing Investor Portfolio Completed")

    fmt.Println("---------- OUT BUY ASSET CUSTODIAN ----------")
    fmt.Println("*********************************************")

    return true
}

// METHOD TO SELL AN ASSET
// PARAMETERS: 1. USERNAME 2. STOCK TICKER 3. STOCK QTY 4. STOCK PRICE
// USES PREFIX02, PREFIX02IDX FOR COMPOSITE KEY - INVESTOR PORTFOLIO
func sellAsset(stub shim.ChaincodeStubInterface, _userName string, _stockTicker string, _stockQty int64, _stockPrice float64) bool {

    fmt.Println("**********************************************")
    fmt.Println("---------- IN SELL ASSET CUSTODIAN ----------")

    // FETCH THE RECORD FROM INVESTOR PORTFOLIO FOR UPDATE
    // PREPARE THE KEY TO READ INVESTOR PORTFOLIO
    _investorPortfolioKey, err := stub.CreateCompositeKey(PREFIX02IDX, []string{PREFIX02, _userName, _stockTicker})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return false
    }
    fmt.Println("sellAsset: Prepare Investor Trade Key Completed")

    // USE THE KEY TO RETRIEVE INVESTOR PORTOFOLIO FOR THIS STOCK TICKER
    _investorPortfolioAsBytes, err := stub.GetState(_investorPortfolioKey)
    // IF THE STOCK TICKER DOES NOT EXIST THEN RETURN ERROR
    if _investorPortfolioAsBytes == nil {
        return false
    }

    // STRUCTURE TO GET THE INVESTOR PORTFOLIO RECORD
    _investorPortfolio := investorPortfolio {}

    // IF THE STOCK TICKER EXISTS IN PORTFOLIO UPDATE STOCKQTY
    fmt.Println("sellAsset: Record Retrieved by GetState")
    fmt.Println(string(_investorPortfolioAsBytes))

    // NOW UNMARSHALL THE INVESTOR PORTFOLIO RECORD
    err = json.Unmarshal(_investorPortfolioAsBytes, &_investorPortfolio)
    // CHECK FOR ERROR IN UNMARSHALLING
    if err != nil {
        return false
    }

    // CHECK IF INVESTOR HAS ENOUGH STOCK TO SELL
    if (_investorPortfolio.StockQty < _stockQty) {
        fmt.Println("sellAsset: Not Enough Quantity To Sell ", _investorPortfolio.StockQty, " < ", _stockQty)
        return false
    }

    // NOW UPDATE THE STOCKQTY AS WE ARE ADDING TO EXISTING PORTFOLIO
    _stockQty = _investorPortfolio.StockQty - _stockQty
    
    // CALCULATE THE NEW STOCK VALUE
    _stockValue := float64(_stockQty) * _stockPrice

    // NOW PREPARE TO WRITE / UPDATE THE INVESTOR PORTFOLIO
    _investorPortfolio.StockQty = _stockQty
    _investorPortfolio.StockValue = _stockValue
    _investorPortfolio.StockPrice = _stockPrice
    fmt.Println("sellAsset: New StockQty  : ", _stockQty)
    fmt.Println("sellAsset: New StockValue: ", _stockValue)
    fmt.Println("sellAsset: StockPrice    : ", _stockPrice)

    // MARSHAL THE INVESTOR PORTFOLIO RECORD
    _investorPortfolioAsBytes, err = json.Marshal(_investorPortfolio)
    // CHECK FOR ERROR IN MARSHALING
    if err != nil {
        return false
    }

    // NOW WRITE THE INVESTOR PORTFOLIO RECORD
    err = stub.PutState(_investorPortfolioKey, _investorPortfolioAsBytes)
    // CHECK FOR ERROR
    if err != nil {
        return false
    }
    fmt.Println("sellAsset: Writing Investor Portfolio Completed")

    fmt.Println("---------- OUT SELL ASSET CUSTODIAN ----------")
    fmt.Println("*********************************************")

    return true
}