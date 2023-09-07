#!/bin/bash

input=$1
output=$2

# LOOP
for filename in $(ls $input | grep .fbx); do
    shortName=$(basename $filename .fbx)
    echo $input$filename " -> " $output$shortName.glb
    ./FBX2glTF-linux-x86_64 -i $input$filename -o $output$shortName -b -v
done
