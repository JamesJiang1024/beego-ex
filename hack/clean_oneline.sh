#!/bin/bash
for i in `oc get all | grep -v "is" | grep -v "po" | grep -v "rc" | grep beego | awk '{print $1}'`;do oc delete $i;done
