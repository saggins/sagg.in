#!/bin/bash

rm main
echo "REMOVED MAIN"
export GOPATH="/home/ubuntu/saggweb/"
echo "SET GOPATH"
go build *.go
echo "FINSHED BUILDING"
sudo setcap 'cap_net_bind_service=+ep' /home/ubuntu/saggweb/src/github.com/win32prog/sagg.in/main
echo "FINISHED SETCAP"
./main
echo "RUNNING SERVER"
