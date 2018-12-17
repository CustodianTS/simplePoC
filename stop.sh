#!/bin/bash
#
# Stop script for the Asset Custody usecase. There are 6 nodes and each node is stopped in this script.
#

# STOP all the containers
if [ `sudo docker ps | grep "regulator" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-regulator.yml down
fi

if [ `sudo docker ps | grep "depository" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-depository.yml down
fi

if [ `sudo docker ps | grep "bank" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-bank.yml down
fi

if [ `sudo docker ps | grep "exchange" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-exchange.yml down
fi

if [ `sudo docker ps | grep "custodian" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-custodian.yml down
fi

if [ `sudo docker ps | grep "investor" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-investor.yml down
fi

sudo docker ps
