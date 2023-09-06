#!/bin/bash
#Running an overlay vendor extension before flattening the tree.
#path to filter vspec
currentpath=$(pwd)
extensionpath="$(pwd)"/extensions/extension.vspec
echo $extensionpath
#cd $1
tools_path="$1"vss-tools
cd $tools_path
python vspec2json.py -I ../spec -u ../spec/units.yaml ../spec/VehicleSignalSpecification.vspec -o $extensionpath out_t.json
mv out_t.json "$currentpath"/vss.json
cd $currentpath
./vssjsontoflattendjson






