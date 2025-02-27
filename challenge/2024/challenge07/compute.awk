#!/bin/sh
awk -F\t 'BEGIN {f=57.1/8.8} {printf "\"%s\",%.4f\n", $1, $2*f}' ill.d > data.csv
