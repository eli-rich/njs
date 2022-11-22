#! /bin/bash
cd server
go build -ldflags='-s -w' -o ../main .