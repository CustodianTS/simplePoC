package main

import (
    "fmt"
    "encoding/json"
    //"strings"
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
