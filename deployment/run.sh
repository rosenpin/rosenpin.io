#!/bin/bash

binary_path=$1
config_path=$2

if [ -z $config_path ] || [ -z $config_path ]; then
    echo "Invalid paths."
    echo "Usage:"
    echo "./run.sh [[binary_path]] [[config_path]]"
    exit 1
fi

sudo ./rosenpin -c configs/production_config.yml
