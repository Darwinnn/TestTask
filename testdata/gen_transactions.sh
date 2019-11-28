#!/bin/bash

# gen 10 random transactions
for i in {1..20}
do 
    if [ $(($RANDOM % 2)) -eq 0 ] 
    then 
        curl -X POST http://localhost:8080/api/v1/state -d "{\"amount\": \"10.0\", \"state\": \"win\", \"transactionId\": \"$(uuidgen)\"}" -H "Source-Type: game" -H "Content-Type: application/json"
    else 
        curl -X POST http://localhost:8080/api/v1/state -d "{\"amount\": \"1.0\", \"state\": \"lost\", \"transactionId\": \"$(uuidgen)\"}" -H "Source-Type: game" -H "Content-Type: application/json"
    fi
    echo "Uploaded $i/20 transactions"
done