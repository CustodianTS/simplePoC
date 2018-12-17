package main

import (
    "fmt"
    "encoding/json"
    //"strings"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// METHOD TO INITIATE RECORDS IN BANK MASTER
// USES PREFIX04, PREFIX04IDX FOR COMPOSITE KEY
func initBank(stub shim.ChaincodeStubInterface) pb.Response {

    fmt.Println("---------- IN INIT BANK----------")

    // INITIALIZE BANK MASTER WITH DATA
    _bankMaster := [] bankMaster {
                      bankMaster {UserName: "johndoe01", BankAC: "BANK00001", Balance:100000.00},
                      bankMaster {UserName: "johndoe02", BankAC: "BANK00002", Balance:200000.00},
                      bankMaster {UserName: "johndoe03", BankAC: "BANK00003", Balance:300000.00},
                    }
    
    for i := 0; i < len(_bankMaster); i++ {

        // PREPARE THE KEY VALUE PAIR TO PERSIST THE INVESTOR
        _bankKey, err := stub.CreateCompositeKey(PREFIX04IDX, []string {PREFIX04, _bankMaster[i].UserName})
        // CHECK FOR ERROR IN CREATING COMPOSITE KEY
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("Bank Key: ", _bankKey)

        // MARSHAL THE BANK MASTER RECORD 
        _bankMasterAsBytes, err := json.Marshal(_bankMaster[i])
        // CHECK FOR ERROR IN MARSHALING
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("Bank Master Record: ", string(_bankMasterAsBytes))

        // NOW WRITE THE BANK MASTER RECORD
        err = stub.PutState(_bankKey, _bankMasterAsBytes)
        // CHECK FOR ERROR
        if err != nil {
            return shim.Error(err.Error())
        }        
    }
    fmt.Println("---------- OUT INIT BANK ----------")

    // RETURN SUCCESS
    return shim.Success(nil)
}

// METHOD TO GET BANK MASTER RECORD
// PARAMETERS: 1. USERNAME
// USES PREFIX04, PREFIX04IDX FOR COMPOSITE KEY
func getBankMaster(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    fmt.Println("****************************************")
    fmt.Println("---------- IN GET BANK MASTER ----------")

    // RETURN ERROR IF ARGS IS NOT 1 IN NUMBER
    if len(args) != 1 {
        fmt.Println("**************************")
        fmt.Println("Too few argments... Need 1")
        fmt.Println("**************************")
        return shim.Error("Invalid argument count. Expecting 1.")
    }

    // SET ARGUMENTS INTO LOCAL VARIABLES
    _userName := args[0]

    // LOG THE ARGUMENTS
    fmt.Println("**** Arguments To Function ****")
    fmt.Println("User Name       : ", _userName)

    // PREPARE THE KEY TO GET RECORD FROM BANK MASTER
    _bankKey, err := stub.CreateCompositeKey(PREFIX04IDX, []string{PREFIX04, _userName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("getBankMaster: Arguments Setting and Prepare Key Completed")
    fmt.Println("Bank Key: ", _bankKey)

    // USE THE KEY TO RETRIEVE BANK MASTER
    _bankMasterAsBytes, err := stub.GetState(_bankKey)
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("getBankMaster: Record Retrieved by GetState")

    fmt.Println(string(_bankMasterAsBytes))
    fmt.Println("---------- OUT GET BANK MASTER ----------")
    fmt.Println("*****************************************")

    // RETURN SUCCESS
    return shim.Success(_bankMasterAsBytes)
}

// METHOD TO EXECUTE DEBIT OR CREDIT TRANSACTIONS ON A BANK ACCOUNT
// PARAMETERS: 1. USERNAME 2. BANK ACCOUNT 3. DEBIT OR CREDIT 4. AMOUNT
// USES PREFIX04, PREFIX04IDX FOR COMPOSITE KEY - BANK MASTER
// USES PREFIX05, PREFIX05IDX FOR COMPOSITE KEY - BANK TRANSACTIONS
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
    _userName := args[0]
    _bankAC := args[1]
    _transactionType := args[2]
    _amount, _ := strconv.ParseFloat(args[3], 64)

    // LOG THE ARGUMENTS
    fmt.Println("**** Arguments To Function ****")
    fmt.Println("User Name       : ", _userName)
    fmt.Println("Bank AC         : ", _bankAC)
    fmt.Println("Transaction Type: ", _transactionType)
    fmt.Println("Amount          : ", _amount)

    // PREPARE THE KEY TO GET RECORD FROM BANK MASTER
    _bankKey, err := stub.CreateCompositeKey(PREFIX04IDX, []string{PREFIX04, _userName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTransaction: Arguments Setting and Prepare Key Completed")
    fmt.Println("Bank Key: ", _bankKey)

    // USE THE KEY TO RETRIEVE BANK MASTER
    _bankMasterAsBytes, err := stub.GetState(_bankKey)
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTransaction: Record Retrieved by GetState")
    fmt.Println(string(_bankMasterAsBytes))

    _bankMaster := bankMaster{}
    err = json.Unmarshal(_bankMasterAsBytes, &_bankMaster)
	if err != nil {
	    return shim.Error(err.Error())
    }

    // READY TO EXECUTE TRANSACTION
    _balance := _bankMaster.Balance
    if _transactionType == DEBIT {
        if (_balance < _amount) {
            fmt.Println("executeTransaction: Not enought balance ", _balance, " < ", _amount)
            return shim.Error(err.Error())
        }
        fmt.Println("executeTransaction: Enough balance ", _balance, " > ", _amount)
        _balance = _balance - _amount
        fmt.Println("executeTransaction: Debit Completed")
        fmt.Println("executeTransaction: New balance ", _balance)
    } else if _transactionType == CREDIT {
        fmt.Println("executeTransaction: Balance ", _balance, " Amount ", _amount)
        _balance = _balance + _amount
        fmt.Println("executeTransaction: Credit Completed")
        fmt.Println("executeTransaction: New balance ", _balance)
    }

    // NOW UPDATE BANK MASTER RECORD
    _bankMaster = bankMaster {
                     UserName: _userName,
                     BankAC:   _bankAC,
                     Balance:  _balance,
                    }

    // MARSHAL THE BANK MASTER RECORD
    _bankMasterAsBytes, err = json.Marshal(_bankMaster)
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
    fmt.Println("executeTransaction: Record Updated By PutState")

    // NOW PREPARE BANK TRANSACTION RECORD TO WRITE
    _currentTimeTS := getTimeStamp()
    _bankTransaction := bankTransactions {
        TransTimestamp: _currentTimeTS,
        BankAC:         _bankAC,
        TransAmount:    _amount,
        Balance:        _balance,
    }

    // PREPARE THE KEY TO WRITE BANK TRANSACTION
    _bankTransactionKey, err := stub.CreateCompositeKey(PREFIX05IDX, []string{PREFIX05, _userName})
    // CHECK FOR ERROR IN CREATING COMPOSITE KEY
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("executeTransaction: Prepare Transaction Key Completed")

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
    fmt.Println("executeTransaction: Writing Bank Transaction Completed")

    fmt.Println("---------- OUT EXECUTE TRANSACTION BANK----------")
    fmt.Println("*************************************************")

    // RETURN SUCCESS
    return shim.Success(_bankTransactionAsBytes)
}
