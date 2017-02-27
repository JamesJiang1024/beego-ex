#!/bin/bash
#sudo docker tag beego-ex:latest registry.travelsky.com/beego/beego-ex:$1
#sudo docker push registry.travelsky.com/beego/beego-ex:$1 
sudo docker save beego-ex:latest > .output/beego-ex.tar
scp .output/beego-ex.tar root@10.221.130.221:/root/
ssh root@10.221.130.221 "docker load < beego-ex.tar"
ssh root@10.221.130.221 "docker tag beego-ex:latest registry.travelsky.com/beego/beego-ex:$1"
ssh root@10.221.130.221 "docker push registry.travelsky.com/beego/beego-ex:$1"
oc import-image import-image beego:$1 --confirm=true --from registry.travelsky.com/beego/beego-ex:$1
