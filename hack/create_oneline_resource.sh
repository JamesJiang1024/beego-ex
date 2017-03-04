#!/bin/bash
version=`git rev-parse HEAD | cut -b 1-6`
oc new-app --image-stream="svcrouter-test2/beego:v1-$version" --name beego-v1
oc create -f yml/beego_svc.yml
