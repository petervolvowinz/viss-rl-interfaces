#!/bin/bash

export GIT_REPO_URL=$VSS_REPO_URL
export GIT_REP_LOCAL=$LOCAL_VSS_PATH
export VSS_RUST_HOME=$RUST_HOME_VSS

# switch between these two lines depending on using the vss extensionn mechanism
# note the extension.sh does not verify and download lates version of vss
./extension.sh $GIT_REP_LOCAL
#verify, download and flatten the latest published vss versopn.
#./vss_release.sh $GIT_REPO_URL $GIT_REP_LOCAL $VSS_RUST_HOME