#!/bin/bash
kill -9 `pidof map`
sleep 1
cd /data/go/src/map

go build

nohup /data/go/src/map/map -c /data/go/src/map/appserver.conf >/data/go/src/map/map.log 2>&1 &
ps -eo comm,lstart | grep map
