#!/bin/bash
go build -o main
cat test.txt | xargs ./main
rm main
