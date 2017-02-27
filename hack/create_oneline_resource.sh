#!/bin/bash
version=`git rev-parse HEAD | cut -b 1-6`
oc new-app --image-stream="svcrouter-test1/beego:v1.0-$version" --name beego-v1.0
oc create -f hack/beego_svc.yml
