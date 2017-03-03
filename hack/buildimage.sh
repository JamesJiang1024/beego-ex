#!/bin/bash

version=`git rev-parse HEAD | cut -b 1-6`

rm -rf .output/*
bee pack
mv beego-ex.tar.gz .output/
tar -zxvf .output/beego-ex.tar.gz -C .output/
sudo docker build -t beego-ex:v1-$version .
