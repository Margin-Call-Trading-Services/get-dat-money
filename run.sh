#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Please provide detach mode configuration with flag detach=(true,false)."
    exit 1
fi

detach=$1

if [ ${detach} == "true" ]; then
    echo "Spinning up environment in detach mode."
    docker compose up --build --detach
else
    docker compose up --build
fi
