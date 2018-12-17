#!/bin/bash
#
# Start script for the Asset Custody usecase. There are 6 nodes and each node is stopped / started in this script.
#
# Exit on first error, print all commands.
#set -ev

CHNAME=tradingchannel
CCNAME=simple1

# Test harness - onboard_investor
sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"onboard_investor","Args":["johndoe01","John","Doe","DEPO00001","BANK00001"]}'

# Test harness - get_bank_master
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"get_bank_master","Args":["johndoe01"]}'
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"get_bank_master","Args":["johndoe02"]}'
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"get_bank_master","Args":["johndoe03"]}'

# Test harness - execute_transaction - Debit
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"execute_transaction","Args":["johndoe01, "BANK00001","DEBIT","10000"]}'
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"execute_transaction","Args":["johndoe02, "BANK00002","DEBIT","20000"]}'
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"execute_transaction","Args":["johndoe03, "BANK00003","DEBIT","30000"]}'

# Test harness - execute_transaction - Credit
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"execute_transaction","Args":["johndoe01, "BANK00001","CREDIT","10000"]}'
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"execute_transaction","Args":["johndoe02, "BANK00002","CREDIT","20000"]}'
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"execute_transaction","Args":["johndoe03, "BANK00003","CREDIT","30000"]}'

# Test harness - get_bank_master
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"get_bank_master","Args":["johndoe01"]}'
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"get_bank_master","Args":["johndoe02"]}'
#sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -c '{"function":"get_bank_master","Args":["johndoe03"]}'