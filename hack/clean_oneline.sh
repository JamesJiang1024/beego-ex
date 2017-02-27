#!/bin/bash
for i in `oc status | grep beego | awk '{print $1}'`;do oc delete $i;done
