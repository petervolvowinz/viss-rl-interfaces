# VSS manipulation scripts

We have scripts that lets us download the latest VSS version from github as a json file. We also have a script that
flattens and normalizes the json to be more human readable and as the same time prepare for a linear data structure 
such as a hashmap.

**paths.sh**
```typescript
...
export GIT_REPO_URL="COVESA/vehicle_signal_specification"
export GIT_REP_LOCAL=<LOCAL-PATH-TO-VSS-REPO>
export VSS_RUST_HOME=<PATH-TO-SRC-VSS-SCRIPTS>"
...

```

The above script sets the VSS github URL, the path to where the local repo is and where the vss related scripts are located in 
relation to the RUST project.

VSS has the notion of overlays which is used to add vendor/oem specific extensions or changes to the signal tree.
```typescript
./extension.sh $GIT_REP_LOCAL
```
See the following link on VSS extensions: https://covesa.github.io/vehicle_signal_specification/rule_set/overlay/

In paths.sh we can comment the above call and instead uncomment:
```typescript
#./vss_release.sh $GIT_REPO_URL $GIT_REP_LOCAL $VSS_RUST_HOME
```

This downloads the latest vss version and generates a flattened json structure. Note that vss-tools: https://github.com/COVESA/vss-tools
is a sub repo and needs to be downloaded separately.

This is an executable written in Golang that simply generates a simple json structure.
```typescript
vssjsontoflattenedjson
```

```json
[
 {
  "signaltype": 0,
  "signame": "Vehicle.CurrentLocation.Heading",
  "lextokens": [
   "Vehicle",
   "Current",
   "Location",
   "Heading"
  ],
  "description": "Current heading relative to geographic north. 0 = North, 90 = East, 180 = South, 270 = West.",
  "datatype": "double",
  "unit": "degrees"
 },
 ...
]
```

**RUN**

make sure the scripts are executable:
```commandline
chmod +x <script>.sh
source <script>.sh
./script.sh
```

run ./paths.sh and flattened json is generated to be used to verify the client filter. The client filter being defined in
config.json.





