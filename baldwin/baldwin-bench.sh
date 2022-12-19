#!/bin/bash
for i in $(seq 0 29)
do
ctime -t $i pdfdeck -pagesize 1920,1080 -sans PublicSans-Medium -serif Charter-Regular -mono Inconsolata-Bold -symbol VTCWilliam-Regular baldwin.xml
done
