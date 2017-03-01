#!/bin/bash
for i in `oc get all | grep beego | awk '{print $1}'`;do oc delete $i;done
