# operator-sdk-data-collector
Use this script to get data on the reachability of operator-sdk.

## Pre-requisites:
- Authenticate yourself to https://catalog.redhat.com/software/containers/explore. You should be able to successfully login to `registry.redhat.io`.

## How to use this project:

Step 1: Clone this repository
- Run `git clone https://github.com/varshaprasad96/operator-sdk-data-collector.git`.

Step 2: Download the databases
- Run `oc-get-index-redhat-operators.sh` to download the databases for indices. The DBs will be downloaded under `db/`.

Step 3: Generate report
- Run `run-report.sh` to generate report from the data extracted from the DBs.

The generated report will be saved in xlsx format in the folder `report/` with the timestamp.