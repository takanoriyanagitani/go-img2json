#!/bin/sh

input=~/Downloads/PNG_transparency_demonstration_1.png

file "${input}"

cat "${input}" |
	./img2json |
	wc -l

cat "${input}" |
	./img2json |
	tail -1 |
	jq -c '.[]' |
	wc -l
