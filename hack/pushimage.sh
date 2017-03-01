#!/bin/bash
#sudo docker tag beego-ex:latest registry.travelsky.com/beego/beego-ex:$1
#sudo docker push registry.travelsky.com/beego/beego-ex:$1 
version=v1-`git rev-parse HEAD | cut -b 1-6`

echo "Save And Push Images"
sudo docker save beego-ex:$version > .output/beego-ex.tar
scp .output/beego-ex.tar root@10.221.130.221:/root/
rm -rf .output/beego-ex.tar
ssh root@10.221.130.221 "docker load < beego-ex.tar"
ssh root@10.221.130.221 "docker tag beego-ex:$version registry.travelsky.com/beego/beego-ex:$version"
ssh root@10.221.130.221 "docker push registry.travelsky.com/beego/beego-ex:$version"


echo "Import Images"
oc import-image beego:$version --confirm=true --from registry.travelsky.com/beego/beego-ex:$version

echo "Begin Clean......."
for i in `sudo docker images | grep beego | awk '{print $3}'`;do sudo docker rmi -f $i;done | grep -v "IMAGE"
