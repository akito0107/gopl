#!/bin/bash
go build -o main
echo "first time"
time ./main https://eoimages.gsfc.nasa.gov/images/imagerecords/73000/73751/world.topo.bathy.200407.3x5400x2700.jpg

echo "second time, should be loaded from cache"
time ./main https://eoimages.gsfc.nasa.gov/images/imagerecords/73000/73751/world.topo.bathy.200407.3x5400x2700.jpg
rm main
