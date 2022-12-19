#!/bin/bash
# simpler baldwin benchmark that excludes specific fonts
for i in $(seq 0 29)
do
	ctime -t $i pdfdeck -pagesize 1920,1080  baldwin.xml
done
