#!/bin/bash

go build -o out
./out -n 1000000
rm out
