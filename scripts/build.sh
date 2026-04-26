#!/bin/sh

set -e

echo 'Preparing build directory...'
rm -rf bin/
mkdir -p bin/

echo 'Compiling the Ledger API Server...'
go build -o bin/api_server ./cmd/ledger-api/main.go

echo 'Compiling the Fraud Detection Worker...'
go build -o bin/fraud_worker ./cmd/fraud-worker/main.go

echo 'All services compiled successfully.'
echo 'Executables are located in the bin/ directory.'
