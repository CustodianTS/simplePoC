package main

import (
    "fmt"
    "encoding/json"
    //"strings"
    //"time"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// METHOD TO CREATE INVESTOR
func onboardInvestor(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Printf("***************************************\n")
    fmt.Printf("---------- IN ONBOARDINVESTOR----------\n")

    // RETURN ERROR IF ARGS IS NOT 7 IN NUMBER
    if len(args) != 7 {
        fmt.Printf("**************************\n")
        fmt.Printf("Too few argments... Need 7\n")
        fmt.Printf("**************************\n")

        return shim.Error("Invalid argument count. Expecting 7.")
    }

    // PREPARE THE INPUT VALUES TO WRITE
    _investor := investor {
        userName:     args[0],
        userFName:    args[1],
        userLName:    args[2],
        userIdentity: args[3],
        kycStatus:    args[4],
        depositoryAC: args[5],
        bankAC:       args[6],
    }

    // PREPARE THE KEY VALUE PAIR TO PERSIST THE INVESTOR
    _investorKey, err := stub.CreateCompositeKey(prefixCustodian, []string{_investor.userName})
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

    fmt.Printf("---------- OUT ONBOARDINVESTOR----------\n")
    fmt.Printf("****************************************\n")

    // RETURN SUCCESS
    return shim.Success(nil)
}
