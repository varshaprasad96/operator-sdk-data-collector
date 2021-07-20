#!/bin/bash
set -x

echo "Creating folder to save DBs if it does not exist"
mkdir -p db/

oc image extract registry.redhat.io/redhat/redhat-operator-index:v4.8 --file=/database/index.db --filter-by-os='linux/amd64'
mv index.db db/index.db.4.8.redhat-operators
oc image extract registry.redhat.io/redhat/certified-operator-index:v4.8 --file=/database/index.db --filter-by-os='linux/amd64'
mv index.db db/index.db.4.8.certified-operators
oc image extract registry.redhat.io/redhat/community-operator-index:v4.8 --file=/database/index.db --filter-by-os='linux/amd64'
mv index.db db/index.db.4.8.community-operators
oc image extract registry.redhat.io/redhat/redhat-marketplace-index:v4.8 --file=/database/index.db --filter-by-os='linux/amd64'
mv index.db db/index.db.4.8.redhat-marketplace-operators
oc image extract quay.io/operatorhubio/catalog:latest --file=/database/index.db
mv index.db db/index.db.operatorhub.io
oc image extract registry-proxy.engineering.redhat.com/rh-osbs/iib-pub:v4.8 --file=/database/index.db --filter-by-os='linux/amd64'
mv index.db db/index.db.4.8.prod