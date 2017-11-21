#!/bin/bash

go build -o echo

count=1
./echo $@ | while read line
do
    echo $line
    i=`echo "${line}" | awk '{print $1}'`
    if [ $i -ne $count ];then
        echo "fail ${i}: but ${line}"
    fi
    count=$((count += 1))
done

rm echo
