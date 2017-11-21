#!/bin/bash

go build -o lissajous
./lissajous > out.gif
open out.gif
rm ./lissajous
