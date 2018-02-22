#!/bin/bash

go build -o printer
cat test.html | ./printer
rm printer
