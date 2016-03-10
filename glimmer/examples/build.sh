#!/bin/bash -e

cd `dirname $0`

for dir in `ls | grep jason-`; do
	echo $dir
	cd $dir
	go build
	cd ..
done
