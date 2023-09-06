#!/bin/bash

function return_latest_release_tag(){
  curl "https://api.github.com/repos/$1/releases/latest" | # Get latest release from GitHub api
  grep '"tag_name":' |                                            # Get tag line
  sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value

}


echo getting last release tag for $1 ...
tag_name=$(return_latest_release_tag $1)
echo current release version is : $tag_name
cd $2
git checkout $tag_name -b latest
#run vss-tools generating json
cd $2
tools_path="$2"vss-tools
cd $tools_path
python vspec2json.py  -I ../spec -u ../spec/units.yaml ../spec/VehicleSignalSpecification.vspec vss.json
mv vss.json "$3"vss.json
cd $3
./vssjsontoflattendjson # go script that flattens the vss tree, result in vss-flat-json
