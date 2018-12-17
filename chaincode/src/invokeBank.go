package main

import (
    "fmt"
    "encoding/json"
    //"strings"
    "strconv"
    "time"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// METHOD TO INITIATE RECORDS IN BANK MASTER
func initBank(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Println("*********************************")
    fmt.Println("---------- IN INIT BANK----------")

    // RETURN ERROR IF ARGS IS NOT 3 IN NUMBER
    if len(args) != 3 {
        fmt.Println("**************************")
        fmt.Println("Too few argments... Need 3")
        fmt.Println("**************************")
        return shim.Error("Invalid argument count. Expecting 7.")
    }

    // INITIALIZE BANK MASTER WITH USERDATA
    _balance,_ := strconv.ParseFloat(args[2], 64)
    _bankMaster := bankMaster {
                     userName: args[0],
                     bankAC:   args[1],
                     balance:  _balance,
                    }
    // PREPARE THE KEY VALUE PAIR TO PERSIST THE INVESTOR
    _bankKey, err := stub.CreateCompositeKey(prefixBank, []string{_bankMaster.userName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }

    // MARSHAL THE BANK MASTER RECORD
    _bankMasterAsBytes, err := json.Marshal(_bankMaster)
    // CHECK FOR ERROR IN MARSHALING
    if err != nil {
        return shim.Error(err.Error())
    }

    // NOW WRITE THE BANK MASTER RECORD
    err = stub.PutState(_bankKey, _bankMasterAsBytes)
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }

    fmt.Println("---------- OUT INIT BANK ----------")
    fmt.Println("***********************************")

    // RETURN SUCCESS
    return shim.Success(_bankMasterAsBytes)
}

// METHOD TO EXECUTE DEBIT OR CREDIT TRANSACTIONS ON A BANK ACCOUNT
// PARAMETERS: 1. BANK ACCOUNT, 2. USERNAME, 3. DEBIT OR CREDIT, 4. AMOUNT
func executeTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    fmt.Println("************************************************")
    fmt.Println("---------- IN EXECUTE TRANSACTION BANK----------")

    // RETURN ERROR IF ARGS IS NOT 4 IN NUMBER
    if len(args) != 4 {
        fmt.Println("**************************")
        fmt.Println("Too few argments... Need 4")
        fmt.Println("**************************")
        return shim.Error("Invalid argument count. Expecting 4.")
    }

    // SET ARGUMENTS INTO LOCAL VARIABLES
    _bankAC := args[0]
    _userName := args[1]
    _transactionType := args[2]
    _amount, _ := strconv.ParseFloat(args[3], 64)

    // PREPARE THE KEY TO GET INVESTOR BANK MASTER
    _bankKey, err := stub.CreateCompositeKey(prefixBank, []string{_userName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTransaction: set arguments and prepare key completed")

    // USE THE KEY TO RETRIEVE BANK MASTER
    _bankMasterAsBytesRead, err := stub.GetState(_bankKey)
    if err != nil {
        return shim.Error(err.Error())
    }
    _bankMaster := bankMaster{}
	  err = json.Unmarshal(_bankMasterAsBytesRead, &_bankMaster)
	  if err != nil {
		  return shim.Error(err.Error())
    }
    fmt.Println("executeTransaction: retrieve bank master completed")

    // READY TO EXECUTE TRANSACTION
    _balance := _bankMaster.balance
    if _transactionType == "DEBIT" {
        if (_balance < _amount) {
            fmt.Println("Not enought balance")
            return shim.Error(err.Error())
        }
        _balance = _balance - _amount
        fmt.Println("executeTransaction: debit completed")
    } else if _transactionType == "CREDIT" {
        _balance = _balance + _amount
        fmt.Println("executeTransaction: credit completed")
    }

    // NOW UPDATE BANK MASTER RECORD
    _bankMasterUpdate := bankMaster {
                     userName: _userName,
                     bankAC:   _bankAC,
                     balance:  _balance,
                    }

    // MARSHAL THE BANK MASTER RECORD
    _bankMasterAsBytesWrite, err := json.Marshal(_bankMasterUpdate)
    // CHECK FOR ERROR IN MARSHALING
    if err != nil {
        return shim.Error(err.Error())
    }

    // NOW WRITE THE BANK MASTER RECORD
    err = stub.PutState(_bankKey, _bankMasterAsBytesWrite)
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTransaction: update bank master completed")

    // NOW PREPARE BANK TRANSACTION RECORD TO WRITE
    _currentTime := time.Now()
    _currentTimeStr := _currentTime.String()
    _bankTransaction := bankTransaction {
        transUUID:    _currentTimeStr,
        userName:     _userName,
        bankAC:       _bankAC,
        transDate:    _currentTime,
        transAmount:  _amount,
        balance:      _bankMaster.balance,
    }
    // PREPARE THE KEY TO WRITE BANK TRANSACTION
    _bankTransactionKey, err := stub.CreateCompositeKey(prefixBank, []string{_currentTimeStr})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTransaction: prepare transaction key completed")

    // MARSHAL THE BANK TRANSACTION RECORD
    _bankTransactionAsBytes, err := json.Marshal(_bankTransaction)
    // CHECK FOR ERROR IN MARSHALING
    if err != nil {
        return shim.Error(err.Error())
    }

    // NOW WRITE THE BANK TRANSACTION RECORD
    err = stub.PutState(_bankTransactionKey, _bankTransactionAsBytes)
    // CHECK FOR ERROR
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTransaction: writing bank transaction completed")

    fmt.Println("---------- OUT EXECUTE TRANSACTION BANK----------")
    fmt.Println("*************************************************")

    // RETURN SUCCESS
    return shim.Success(_bankTransactionAsBytes)
}
