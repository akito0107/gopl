#!/bin/bash

go build -o echo
e=`./echo $@ | awk '{print $1}'`
if [ $e = "./echo" ]; then
    echo "passed"
else
    echo "fail ${e} must be \"./echo\""
fi
rm echo
