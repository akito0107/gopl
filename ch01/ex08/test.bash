#!/bin/bash
go build -o main
diff <(./main http://www.gopl.io/) <(./main www.gopl.io/)
rm main
