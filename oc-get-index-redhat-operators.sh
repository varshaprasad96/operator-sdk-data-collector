#!/bin/bash
set -x

echo "Creating folder to save DBs if it does not exist"
mkdir -p db/

oc image extract registry-proxy.engineering.redhat.com/rh-osbs/iib-pub-pending:v4.7 --file=/database/index.db --filter-by-os='linux/amd64'
mv index.db db/index.db.4.7.prod
# oc image extract registry.redhat.io/redhat/certified-operator-index:v4.6 --file=/database/index.db --filter-by-os='linux/amd64'
# mv index.db db/index.db.4.6.certified-operators
# oc image extract registry.redhat.io/redhat/community-operator-index:v4.6 --file=/database/index.db --filter-by-os='linux/amd64'
# mv index.db db/index.db.4.6.community-operators
# oc image extract registry.redhat.io/redhat/redhat-marketplace-index:v4.6 --file=/database/index.db --filter-by-os='linux/amd64'
# mv index.db db/index.db.4.6.redhat-marketplace-operators
# oc image extract quay.io/operatorhubio/catalog:latest --file=/database/index.db
# mv index.db db/index.db.operatorhub.io
# oc image extract registry-proxy.engineering.redhat.com/rh-osbs/iib-pub:v4.6 --file=/database/index.db --filter-by-os='linux/amd64'
# mv index.db db/index.db.4.6.prod