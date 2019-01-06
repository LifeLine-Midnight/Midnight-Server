#!/bin/bash

killall -9 midnightapisvr
sleep 1s
go run midnightapisvr.go &