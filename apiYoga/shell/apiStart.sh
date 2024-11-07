#!/bin/bash
echo "         _                       _ "
echo "   ___  | |   ___    _   _    __| |"
echo "  / __| | |  / _ \  | | | |  / _  |"
echo " | (__  | | | (_) | | |_| | | (_| |"
echo "  \___| |_|  \___/   \__,_|  \__,_|"
echo "                                   "

docker start postgreYoga
docker start redisYoga
sleep 1
echo "start"
cd ~/workspace/yoga/apiYoga
pwd
rm ./output/*
cd ./src
go run main.go

docker stop postgreYoga
docker stop redisYoga