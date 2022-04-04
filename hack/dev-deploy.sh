#!/bin/bash

set -ex

CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-X 'main.Version=develop' -buildid=" -trimpath -o controller-linux-amd64 ./cmd/main.go

echo "# -- this file is autogenerated by dev-deploy.sh from main Dockerfile and is not version controlled" > Dockerfile.dev

run_image_start=`cat Dockerfile | grep -n "\- run image \-" | cut -d':' -f1`
kc_latest_image=`docker image ls | grep kapp-controller | sed 's/ \{2,\}/ /g' | cut -d' ' -f3 | head -n 1`
if [ -z "$kc_latest_image" ] ;
then
  echo "Error: unable to find tag for previous image of kapp-controller"
  echo "For your first deploy please use hack/deploy.sh and then try re-running this script for subsequent deploys."
  exit 1
fi
tail -n +$run_image_start Dockerfile | \
  sed 's/COPY.*kapp-controller/COPY controller-linux-amd64 kapp-controller/' | \
   sed "s/from=0/from=$kc_latest_image/g" | \
   sed 's/helm-v2-unpacked\/linux-amd64\/helm/helmv2/' | \
   sed 's/helm-unpacked\/linux-amd64\/helm/helm/' | \
   sed 's/usr\/local\/bin\///' \
   >> Dockerfile.dev

ytt -f config/ -f config-dev-deploy/ | kbld -f- | kapp deploy -a kc -f- -c -y

