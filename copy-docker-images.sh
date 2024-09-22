#!/bin/bash
docker pull $*
docker tag $* hub.superlee.top/$*
docker push hub.superlee.top/$*
