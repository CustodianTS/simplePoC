package main

import (
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

const prefixCustodian = "CUSTDN"
const prefixBank = "BNK"
//const prefixExchange = "XCHNG"
//const prefixDepository = "DPSTRY"

var logger = shim.NewLogger("main")

type SmartContract struct {
}


// MAPPING BETWEEN FUNCTION NAMES IN APIs and GO METHODS
var bcFunctions = map[string] func(shim.ChaincodeStubInterface, []string) pb.Response {

    // CUSTODIAN PEER
    "onboard_investor":       onboardInvestor,
    //"check_kyc":              checkKYC,
    //"buy_share":              buyShare,
    //"sell_share":             sellShare,
    //"get_investor_dashboard": getInvestorDashboards,

    // BANK PEER
    "init_bank":             initBank,
    "execute_transaction": executeTransaction,

    // EXCHANGE PEER
    //"init_exchange": initExchange,
    //"execute_trade": executeTrade,

    // DEPOSITORY PEER
    //"init_depository": initDepository,
    //"record_trade":    recordTrade,
}


// INIT CALLBACK REPRESENTING INVOCATION OF CHAINCODE
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
    //_, args := stub.GetFunctionAndParameters()
    fmt.Println("**********************************")
    fmt.Println("----------IN INIT METHOD----------")
    fmt.Println("**********************************")
    return shim.Success(nil)
}

// INVOKE FUNCTION ACCEPS BLOCKCHAIN CODE INVOCATIONS
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

    fmt.Println("************************************")
    fmt.Println("----------IN INVOKE METHOD----------")
    fmt.Println("************************************")

    // GET THE FUNCION INVOKED AND ARGS FROM SHIM
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("Function From Command Line: ", function)

    // GET THE METHOD TO INVOKE FROM FUNCTION MAPPING
    bcFunc := bcFunctions[function]
    if bcFunc == nil {
        fmt.Println("ERROR: Function Mapping Not Found")
        return shim.Error("Invalid invoke function.")
    }

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
