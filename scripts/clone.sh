#!/bin/bash

org=nonofficial

repos=(`gh repo list ${org} --json name | jq ".[] | .name" -r`)
for repo in "${repos[@]}"
do
   echo $repo
   git clone git@github.com:${org}/${repo}
done

