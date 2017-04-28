#!/usr/bin/env bash

runDirName=bjdaos_tool_running

function init_Dir() {
   if [ -d ~/${runDirName} ]; then
      rm -rf ~/${runDirName}
   fi
    mkdir ~/${runDirName}
}

function clean() {
    cd ~
    rm bjdaos_tool
}

function target::Kill_old() {
    pid=$(pgrep -u heartdata bjdaos_tool_$1)
    kill -9 $pid
}

function target::Prepare() {
    cp bjdaos_tool bjdaos_tool_$1
    mv bjdaos_tool_$1 ~/${runDirName}/
}

function target::Start() {
    cd ~/${runDirName}
    ./bjdaos_tool_$1 $1 start> ./$1.log 2>&1 &
}

function startTarget() {
    cd ~
    target::Kill_old $1
    target::Prepare $1
    target::Start $1
}

init_Dir
startTarget hd
clean


