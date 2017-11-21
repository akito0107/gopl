#!/bin/bash
go build -o main
./main http://www.gopl.io/
rm main
