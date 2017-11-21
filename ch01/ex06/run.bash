#!/bin/bash

echo "warnig may be mac only"

go build -o lissajous
yes > /dev/null & yes > /dev/null & yes > /dev/null & yes > /dev/null &
./lissajous > too_heavy.gif
killall yes

yes > /dev/null & yes > /dev/null &
./lissajous > heavy.gif
killall yes

./lissajous > normal.gif

rm lissajous
