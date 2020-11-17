#!/usr/bin/env bash  
export AWS_ACCESS_KEY_ID
export AWS_SECRET_ACCESS_KEY
export AWS_SESSION_TOKEN

AWS_ACCESS_KEY_ID="AWS_ACCESS_KEY_ID"
AWS_SECRET_ACCESS_KEY="AWS_SECRET_ACCESS_KEY"
AWS_SESSION_TOKEN="ThisIsOptional"

CURL_INTERFACE="eth0"
HOSTED_ZONE_ID="HOSTED_ZONE_ID"
RECORD_SET="your.fancy.tld"
RECORD_VALUE="$(curl --interface ${CURL_INTERFACE} https://api.ipify.org)"

/path/to/gorecord \
  --zone-id "${HOSTED_ZONE_ID}" \
  --record "${RECORD_SET}" \
  --value "${RECORD_VALE}"
