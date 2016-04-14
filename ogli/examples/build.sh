#!/bin/bash -e

cd `dirname $0`

for dir in `ls | egrep 'jason-|ex*-*'`; do
	echo $dir
	cd $dir
	go build
	cd ..
done
