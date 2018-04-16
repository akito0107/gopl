#!/bin/bash

go build -o out
./out -n 10000000
rm out
