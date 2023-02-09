#!/bin/sh

chain_id=sim
acc_num=0
pubkey=AgT2QPS4Eu6M+cfHeba+3tumsM/hNEBGdM7nRojSZRjF

for txbytes in $(cat txs.txt)
do
	../build/sequence $txbytes $chain_id $acc_num $pubkey
done
