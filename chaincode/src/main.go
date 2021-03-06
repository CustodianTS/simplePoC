package main

import (
    "fmt"
    "time"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

const PREFIX01 = "01" // For investor
const PREFIX02 = "02" // For investorPortfolio
const PREFIX03 = "03" // For investorTrades
const PREFIX04 = "04" // For bankMaster
const PREFIX05 = "05" // For bankTransactions
const PREFIX06 = "06" // For exchangeMaster
const PREFIX07 = "07" // For exchangeTrades

const PREFIX01IDX = "01IDX" // Index for investor
const PREFIX02IDX = "02IDX" // Index for investorPortfolio
const PREFIX03IDX = "03IDX" // Index for investorTrades
const PREFIX04IDX = "04IDX" // Index for bankMaster
const PREFIX05IDX = "05IDX" // Index for bankTransactions
const PREFIX06IDX = "06IDX" // Index for exchangeMaster
const PREFIX07IDX = "07IDX" // Index for exchangeTrades

const BUY = "BUY"
const SELL = "SELL"
const DEBIT = "DEBIT"
const CREDIT = "CREDIT"

var logger = shim.NewLogger("main")

func getTimeStamp() string {
    _time := time.Now()
    return _time.Format("2006-01-02 15:04:05")
}

type SmartContract struct {
}

// MAPPING BETWEEN FUNCTION NAMES IN APIs and GO METHODS
var bcFunctions = map[string] func(shim.ChaincodeStubInterface, []string) pb.Response {

    // CUSTODIAN
    "onboard_investor":       onboardInvestor,       // ONBOARD INVESTOR TO THE SYSTEM
    "trade_asset":            tradeAsset,            // TRADE - BUY OR SELL AN ASSET
    "get_investor_portfolio": getInvestorPortfolio,  // GET PORTFOLIO OF A USER
    "get_investor_trades":    getInvestorTrades,     // GET ALL TRADES OF A USER

    // BANK
    "get_bank_master":       getBankMaster,          // GET BANK ACCOUNT DETAILS OF A USER
    "execute_transaction":   executeTransaction,     // EXECUTE DEBIT OR CREDIT ON A BANK ACCOUNT
    "get_bank_transactions": getBankTransactions,    // GET ALL BANK TRANSACTIONS OF A USER

    // EXCHANGE
    "get_exchange_master_all": getExchangeMasterAll, // GET ALL STOCKTICKERS FROM EXCHANGE MASTER
    "get_exchange_master":     getExchangeMaster,    // GET ONE STOCKTICKER FROM EXCHANGE MASTER
    "execute_trade":           executeTrade,         // EXECUTE TRADE ON A STOCKTICKER
    "get_exchange_trades":     getExchangeTrades,    // GET ALL TRADES EXECUTED ON EXCHANGE
}

// INIT CALLBACK REPRESENTING INVOCATION OF CHAINCODE
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
    //_, args := stub.GetFunctionAndParameters()
    fmt.Println("**********************************")
    fmt.Println("----------IN INIT METHOD----------")
    initBank(stub)
    initExchange(stub)
    fmt.Println("----------OUT INIT METHOD----------")
    fmt.Println("***********************************")
    return shim.Success(nil)
}

// INVOKE FUNCTION ACCEPS BLOCKCHAIN CODE INVOCATIONS
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

    fmt.Println("************************************")
    fmt.Println("----------IN INVOKE METHOD----------")

    // GET THE FUNCION INVOKED AND ARGS FROM SHIM
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("Function From Command Line: ", function)

    // GET THE METHOD TO INVOKE FROM FUNCTION MAPPING
    bcFunc := bcFunctions[function]
    if bcFunc == nil {
        fmt.Println("ERROR: Function Mapping Not Found")
        return shim.Error("Invalid invoke function.")
    }

    fmt.Println("----------OUT INVOKE METHOD----------")
    fmt.Println("*************************************")

    return bcFunc(stub, args)
}

// MAIN METHOD
func main() {
    logger.SetLevel(shim.LogDebug)
    err := shim.Start(new(SmartContract))

    fmt.Println("**********************************")
    fmt.Println("----------In MAIN METHOD----------")
    fmt.Println("**********************************")

    if err != nil {
        fmt.Println("Error starting Simple chaincode: %s", err)
    }
}
