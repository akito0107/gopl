#!/bin/bash

go build -o dup
./dup ch14.test.2.txt ch14.test.3.txt ch14.test.txt
rm dup
