#!/bin/bash

echo The following data will be collected and report can be found at reports/ folder:
echo operator,created,company,sdkversion,operatortype, csvName

go run indexdump.go \
"db/index.db.4.6.prod:prod:4.6" \
"db/index.db.4.6.redhat-operators:redhat:4.6" \
"db/index.db.4.6.community-operators:community:4.6" \
"db/index.db.4.6.redhat-marketplace-operators:marketplace:4.6" \
"db/index.db.4.6.certified-operators:certified:4.6" \
"db/index.db.operatorhub.io:operatorhub:4.6"

exit