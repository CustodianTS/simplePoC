package main

import (
    "fmt"
    "encoding/json"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// METHOD TO INITIATE RECORDS IN EXCHANGE MASTER
// USES PREFIX06 + STOCKTICKER FOR COMPOSITE KEY, PREFIX06IDX AS INDEX - EXCHANGE MASTER
func initExchange(stub shim.ChaincodeStubInterface) pb.Response {

    fmt.Println("---------- IN INIT EXCHANGE ----------")

    // INITIALIZE EXCHANGE MASTER WITH DATA
    _exchangeMaster := [] exchangeMaster {
                          exchangeMaster {StockTicker: "ASIANPAINT", StockQty: 1000, StockPrice: 1334.85},
                          exchangeMaster {StockTicker: "AXISBANK",   StockQty: 2000, StockPrice: 617.25},
                          exchangeMaster {StockTicker: "BPCL",       StockQty: 3000, StockPrice: 350.35},
                          exchangeMaster {StockTicker: "BHEL",       StockQty: 4000, StockPrice: 67.65},
                          exchangeMaster {StockTicker: "COALINDIA",  StockQty: 5000, StockPrice: 253.00},
                          exchangeMaster {StockTicker: "CAPTRUST",   StockQty: 6000, StockPrice: 198.35},
                          exchangeMaster {StockTicker: "DABUR",      StockQty: 7000, StockPrice: 442.75},
                          exchangeMaster {StockTicker: "DLF",        StockQty: 8000, StockPrice: 178.20},
                          exchangeMaster {StockTicker: "ESCORTS",    StockQty: 9000, StockPrice: 670.90},
                          exchangeMaster {StockTicker: "EXIDEIND",   StockQty: 1000, StockPrice: 257.35},
                          }
    
    for i := 0; i < len(_exchangeMaster); i++ {

        // PREPARE THE KEY VALUE PAIR TO PERSIST THE EXCHANGE MASTER DATA
        _exchangeKey, err := stub.CreateCompositeKey(PREFIX06IDX, []string {PREFIX06, _exchangeMaster[i].StockTicker})
        // CHECK FOR ERROR IN CREATING COMPOSITE KEY
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("Exchange Key: ", _exchangeKey)

        // MARSHAL THE EXCHANGE MASTER RECORD 
        _exchangeMasterAsBytes, err := json.Marshal(_exchangeMaster[i])
        // CHECK FOR ERROR IN MARSHALING
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("Exchange Master Record: ", string(_exchangeMasterAsBytes))

        // NOW WRITE THE EXCHANGE MASTER RECORD
        err = stub.PutState(_exchangeKey, _exchangeMasterAsBytes)
        // CHECK FOR ERROR
        if err != nil {
            return shim.Error(err.Error())
        }        
    }
    fmt.Println("---------- OUT INIT EXCHANGE ----------")

    // RETURN SUCCESS
    return shim.Success(nil)
}

// METHOD TO GET EXCHANGE MASTER DATA
// PARAMETERS: 1. STOCK TICKER
// USES PREFIX06 + STOCKTICKER FOR COMPOSITE KEY, PREFIX06IDX AS INDEX - EXCHANGE MASTER
func getExchangeMaster(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Println("********************************************")
    fmt.Println("---------- IN GET EXCHANGE MASTER ----------")

    // RETURN ERROR IF ARGS IS NOT 1 IN NUMBER
    if len(args) != 1 {
        fmt.Println("**************************")
        fmt.Println("Too few argments... Need 1")
        fmt.Println("**************************")
        return shim.Error("Invalid argument count. Expecting 1.")
    }

    // SET ARGUMENTS INTO LOCAL VARIABLES
    _stockTicker := args[0]

    // LOG THE ARGUMENTS
    fmt.Println("**** Arguments To Function ****")
    fmt.Println("Stock Ticker       : ", _stockTicker)

    // PREPARE THE KEY TO GET RECORD FROM EXCHANGE MASTER
    _exchangeKey, err := stub.CreateCompositeKey(PREFIX06IDX, []string{PREFIX06, _stockTicker})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("getExchangeMaster: Prepare Key Completed")
    fmt.Println("Exchange Key: ", _exchangeKey)

    // USE THE KEY TO RETRIEVE EXCHANGE MASTER
    _exchangeMasterAsBytes, err := stub.GetState(_exchangeKey)
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("getExchangeMaster: Record Retrieved by GetState")

    fmt.Println(string(_exchangeMasterAsBytes))
    fmt.Println("---------- OUT GET EXCHANGE MASTER ----------")
    fmt.Println("*********************************************")

    // RETURN SUCCESS
    return shim.Success(_exchangeMasterAsBytes)
}

// METHOD TO GET ALL STOCK TICKERS FROM EXCHANGE MASTER
// PARAMETERS: NONE
// USES PREFIX06 + STOCKTICKER FOR COMPOSITE KEY, PREFIX06IDX AS INDEX - EXCHANGE MASTER
func getExchangeMasterAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Println("************************************************")
    fmt.Println("---------- IN GET EXCHANGE MASTER ALL ----------")

    // GET THE RESULTS ITERATOR FROM PARTIAL KEY ON EXCHANGE MASTER
    // PARTIAL KEY IS PREFIX06
    _xMasterIterator, err := stub.GetStateByPartialCompositeKey(PREFIX06IDX, []string{PREFIX06})
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }
    defer _xMasterIterator.Close()

    // ARRAY OF EXCHANGE MASTER
    _exchangeMaster := make([]exchangeMaster, 1)
    _xMaster := exchangeMaster{}
    _i := 0

    // NOW LOOP THRU THE ITERATOR TO RETRIEVE ALL THE EXCHANGE MASTER RECORDS
    for _xMasterIterator.HasNext() {
        _xResult, err := _xMasterIterator.Next()
        // CHECK FOR ERROR
        if err != nil {
            return shim.Error(err.Error())
        }

        fmt.Println("getExchangeMasterAll: Retreiving Values")
        fmt.Println(string(_xResult.Value))

        // NOW UNMARSHALL THE EXCHANGE MASTER OBJECT
        err = json.Unmarshal(_xResult.Value, &_xMaster)
        // CHECK FOR ERROR
        if err != nil {
            return shim.Error(err.Error())
        }
        if (_i == 0) {
            _exchangeMaster[_i] = _xMaster
        } else {
            _exchangeMaster = append(_exchangeMaster, _xMaster)
        }
        _i++
    }
    _exchangeMasterAsBytes, err := json.Marshal(_exchangeMaster)
    fmt.Println("getExchangeMasterAll: Get Exchange Master All Completed")
    fmt.Println(string(_exchangeMasterAsBytes))

    fmt.Println("---------- OUT GET EXCHANGE MASTER ALL ----------")
    fmt.Println("*************************************************")

    // RETURN SUCCESS
    return shim.Success(_exchangeMasterAsBytes)
}

// METHOD TO EXECUTE TRADE ON EXCHANGE
// PARAMETERS: 1. USERNAME 2. STOCK TICKER 3. STOCK QTY 4. STOCK PRICE
// USES PREFIX06 + STOCKTICKER FOR COMPOSITE KEY, PREFIX06IDX AS INDEX - EXCHANGE MASTER
// USES PREFIX07 + USERNAME FOR COMPOSITE KEY, PREFIX07IDX AS INDEX - EXCHANGE TRADES
func executeTrade(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    
    fmt.Println("*************************************************")
    fmt.Println("---------- IN EXECUTE TRADE - EXCHANGE ----------")

    // RETURN ERROR IF ARGS IS NOT 4 IN NUMBER
    if len(args) != 4 {
        fmt.Println("**************************")
        fmt.Println("Too few argments... Need 4")
        fmt.Println("**************************")
        return shim.Error("Invalid argument count. Expecting 4.")
    }

    // SET ARGUMENTS INTO LOCAL VARIABLES
    _userName := args[0]
    _stockTicker := args[1]
    _stockQty,_ := strconv.ParseInt(args[2], 10, 64)
    _stockPrice, _ := strconv.ParseFloat(args[3], 64)

    // LOG THE ARGUMENTS
    fmt.Println("**** Arguments To Function ****")
    fmt.Println("User Name   : ", _userName)
    fmt.Println("Stock Ticker: ", _stockTicker)
    fmt.Println("Stock Qty   : ", _stockQty)
    fmt.Println("Stock Price : ", _stockPrice)

    // PREPARE THE KEY TO GET RECORD FROM EXCHANGE MASTER
    _exchangeKey, err := stub.CreateCompositeKey(PREFIX06IDX, []string{PREFIX06, _stockTicker})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTrade: Prepare Key Completed")
    fmt.Println("Exchange Key: ", _exchangeKey)

    // USE THE KEY TO RETRIEVE EXCHANGE MASTER
    _exchangeMasterAsBytes, err := stub.GetState(_exchangeKey)
    if err != nil {
        fmt.Println("executeTrade: Stock Ticker Does Not Exist: ", _stockTicker)
        return shim.Error(err.Error())
    }
    fmt.Println("executeTrade: Record Retrieved by GetState")
    fmt.Println(string(_exchangeMasterAsBytes))

    _exchangeMaster := exchangeMaster{}
    err = json.Unmarshal(_exchangeMasterAsBytes, &_exchangeMaster)
    if err != nil {
        return shim.Error(err.Error())
    }

    // READY TO EXECUTE TRADE
    // TRADE EXECUTES IF THERE IS ENOUGH QUANTITY IN EXCHANGE AND AT PRICE >= EXCHANGE PRICE
    _xTicker := _exchangeMaster.StockTicker
    _xQty := _exchangeMaster.StockQty
    _xPrice := _exchangeMaster.StockPrice
    if (_xQty < _stockQty) {
        fmt.Println("executeTrade: Not enough stock ", _xQty, " < ", _stockQty)
        return shim.Error(err.Error())
    } else if _stockPrice < _xPrice {
        fmt.Println("executeTrade: Not enough price ", _stockPrice, " < " , _xPrice)
        return shim.Error(err.Error())
    }

    _xQty = _xQty - _stockQty
    fmt.Println("executeTrade: Trade Completed")
    fmt.Println("executeTrade: New Qty ", _xQty)

    // NOW UPDATE EXCHANGE MASTER RECORD
    _exchangeMaster = exchangeMaster {
                     StockTicker: _xTicker,
                     StockQty:    _xQty,
                     StockPrice:  _xPrice,
                    }

    // MARSHAL THE EXCHANGE MASTER RECORD
    _exchangeMasterAsBytes, err = json.Marshal(_exchangeMaster)
    // CHECK FOR ERROR IN MARSHALING
    if err != nil {
        return shim.Error(err.Error())
    }

    // NOW WRITE THE EXCHANGE MASTER RECORD
    err = stub.PutState(_exchangeKey, _exchangeMasterAsBytes)
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTrade: Exchange Master Record Updated By PutState")

    // NOW PREPARE EXCHANGE TRADES RECORD TO WRITE
    _currentTimeTS := getTimeStamp()
    _stockValue := float64(_stockQty) * _stockPrice
    _exchangeTrade := exchangeTrades {
        UserName:       _userName,
        TradeTimestamp: _currentTimeTS,
        StockTicker:    _stockTicker,
        StockQty:       _stockQty,
        StockPrice:     _stockPrice,
        StockValue:     _stockValue,
    }

    // PREPARE THE KEY TO WRITE EXCHANGE TRADE
    _exchangeTradeKey, err := stub.CreateCompositeKey(PREFIX07IDX, []string{PREFIX07, _userName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTrade: Prepare Trade Key Completed")

    // MARSHAL THE EXCHANGE TRADE RECORD
    _exchangeTradeAsBytes, err := json.Marshal(_exchangeTrade)
    // CHECK FOR ERROR IN MARSHALING
    if err != nil {
        return shim.Error(err.Error())
    }

    // NOW WRITE THE EXCHANGE TRADE RECORD
    err = stub.PutState(_exchangeTradeKey, _exchangeTradeAsBytes)
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTrade: Writing Exchange Trade Completed")

    fmt.Println("---------- OUT EXECUTE TRADE - EXCHANGE ----------")
    fmt.Println("**************************************************")

    // RETURN SUCCESS
    return shim.Success(_exchangeTradeAsBytes)
}

// METHOD TO GET ALL EXCHANGE TRADES
// PARAMETERS: NONE
// USES PREFIX07 + USERNAME AS COMPOSITE KEY, PREFIX07IDX AS INDEX - EXCHANGE TRADES
func getExchangeTrades(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Println("*******************************************************")
    fmt.Println("---------- IN GET EXCHANGE TRADES - EXCHANGE ----------")

    // GET THE RESULTS ITERATOR FROM PARTIAL KEY ON EXCHANGE TRADES
    // PARTIAL KEY IS PREFIX07
    _xTradesIterator, err := stub.GetStateByPartialCompositeKey(PREFIX07IDX, []string{PREFIX07})
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }
    defer _xTradesIterator.Close()

    // ARRAY OF EXCHANGE TRADE
    _exchangeTrades := make([]exchangeTrades, 1)
    _xTrade := exchangeTrades{}
    _i := 0

    // NOW LOOP THRU THE ITERATOR TO RETRIEVE THE EXCHANGE TRADE RECORDS
    for _xTradesIterator.HasNext() {
        _xResult, err := _xTradesIterator.Next()
        // CHECK FOR ERROR
        if err != nil {
            return shim.Error(err.Error())
        }

        fmt.Println("getExchangeTrades: Retreiving Values")
        fmt.Println(string(_xResult.Value))

        // NOW UNMARSHALL THE EXCHANGE TRADE OBJECT
        err = json.Unmarshal(_xResult.Value, &_xTrade)
        // CHECK FOR ERROR
        if err != nil {
            return shim.Error(err.Error())
        }
        if (_i == 0) {
            _exchangeTrades[_i] = _xTrade    
        } else {
            _exchangeTrades = append(_exchangeTrades, _xTrade)
        }
        _i++
    }
    _exchangeTradesAsBytes, err := json.Marshal(_exchangeTrades)
    fmt.Println("getExchangeTrades: Get Exchange Trades Completed")
    fmt.Println(string(_exchangeTradesAsBytes))

    fmt.Println("---------- OUT GET EXCHANGE TRADES - EXCHANGE ----------")
    fmt.Println("********************************************************")

    // RETURN SUCCESS
    return shim.Success(_exchangeTradesAsBytes)
}