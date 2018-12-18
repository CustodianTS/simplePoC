#!/bin/bash
#
# Start script for the Asset Custody usecase. There are 6 nodes and each node is stopped / started in this script.
#
# Exit on first error, print all commands.
#set -ev

SLEEPY=5
CHNAME=tradingchannel
CCNAME=simple01

# STOP all the containers
if [ `sudo docker ps | grep "regulator" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-regulator.yml down
    sleep $SLEEPY 
fi

if [ `sudo docker ps | grep "depository" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-depository.yml down
    sleep $SLEEPY 
fi

if [ `sudo docker ps | grep "bank" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-bank.yml down
    sleep $SLEEPY 
fi

if [ `sudo docker ps | grep "exchange" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-exchange.yml down
    sleep $SLEEPY 
fi

if [ `sudo docker ps | grep "custodian" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-custodian.yml down
    sleep $SLEEPY 
fi

if [ `sudo docker ps | grep "investor" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-investor.yml down
    sleep $SLEEPY
fi


# START all the containers
sudo docker-compose -f docker-compose-investor.yml up -d
sleep $SLEEPY
sudo docker-compose -f docker-compose-custodian.yml up -d
sleep $SLEEPY
sudo docker-compose -f docker-compose-exchange.yml up -d
sleep $SLEEPY
sudo docker-compose -f docker-compose-bank.yml up -d
sleep $SLEEPY
#sudo docker-compose -f docker-compose-depository.yml up -d
#sleep $SLEEPY
#sudo docker-compose -f docker-compose-regulator.yml up -d
#sleep $SLEEPY

# Create the channel
sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer channel create -o orderer.example.com:7050 -c $CHNAME -f /etc/hyperledger/configtx/channel.tx
sleep $SLEEPY

# Join peer0.investor.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.custodian.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=CustodianMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/custodian.example.com/users/Admin@custodian.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.custodian.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.exchange.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=ExchangeMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/exchange.example.com/users/Admin@exchange.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.exchange.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.bank.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=BankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/bank.example.com/users/Admin@bank.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.bank.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.depository.example.com to the channel.
#sudo docker exec -e "CORE_PEER_LOCALMSPID=DepositoryMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/depository.example.com/users/Admin@depository.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.depository.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.regulator.example.com to the channel.
#sudo docker exec -e "CORE_PEER_LOCALMSPID=RegulatorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/regulator.example.com/users/Admin@regulator.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.regulator.example.com:7051" cli peer channel join -b tradingchannel.block

# Install chaincode
# Install code on investor peer
sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode install -n $CCNAME -v 1.0 -p github.com -l golang

# Install code on custodian peer
sudo docker exec -e "CORE_PEER_LOCALMSPID=CustodianMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/custodian.example.com/users/Admin@custodian.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.custodian.example.com:7051" cli peer chaincode install -n $CCNAME -v 1.0 -p github.com -l golang

# Install code on exchange peer
sudo docker exec -e "CORE_PEER_LOCALMSPID=ExchangeMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/exchange.example.com/users/Admin@exchange.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.exchange.example.com:7051" cli peer chaincode install -n $CCNAME -v 1.0 -p github.com -l golang

# Install code on bank peer
sudo docker exec -e "CORE_PEER_LOCALMSPID=BankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/bank.example.com/users/Admin@bank.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.bank.example.com:7051" cli peer chaincode install -n $CCNAME -v 1.0 -p github.com -l golang

# Install code on depository peer
#sudo docker exec -e "CORE_PEER_LOCALMSPID=DepositoryMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/depository.example.com/users/Admin@depository.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.depository.example.com:7051" cli peer chaincode install -n $CCNAME -v 1.0 -p github.com -l golang

# Install code on depository peer
#sudo docker exec -e "CORE_PEER_LOCALMSPID=RegulatorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/regulator.example.com/users/Admin@regulator.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.regulator.example.com:7051" cli peer chaincode install -n $CCNAME -v 1.0 -p github.com -l golang

# Instantiate chain code
sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode instantiate -o orderer.example.com:7050 -C $CHNAME -n $CCNAME -l golang -v 1.0 -c '{"Args":[""]}' -P "OR ('InvestorMSP.member','CustodianMSP.member','ExchangeMSP.member','BankMSP.member')"
