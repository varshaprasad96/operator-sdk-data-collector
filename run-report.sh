#!/bin/bash

echo The following data will be collected and report can be found at reports/ folder:
echo operator,created,company,sdkversion,operatortype, csvName

go run indexdump.go \
"db/index.db.4.7.prod:prod:4.7" \
"db/index.db.4.7.redhat-operators:redhat:4.7" \
"db/index.db.4.7.community-operators:community:4.7" \
"db/index.db.4.7.redhat-marketplace-operators:marketplace:4.7" \
"db/index.db.4.7.certified-operators:certified:4.7" \
"db/index.db.operatorhub.io:operatorhub:4.7"

exit