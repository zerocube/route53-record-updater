# route53-record-updater

This script is designed to update a single Route53 record.

Our primary use case for it is running it on an EdgeRouter or two in order to implement some kind of Dynamic DNS™.

## Usage

Best practice:

* Get a static IP address from your provider.

Next-to-best practice:

* Get the binary: `go get -v github.com/zerocube/route53-record-updater`
* Create a wrapper script (see gorecord.sh) that exports the following environment variables:
  * `HOSTED_ZONE_ID`
  * `RECORD_SET`
  * `RECORD_VALUE`
  * `AWS_ACCESS_KEY_ID`
  * `AWS_SECRET_ACCESS_KEY`
  * `AWS_SESSION_TOKEN` (optional)
* Run the wrapper script (ensuring that it is also executable)
* Call the wrapper script via cron at an appropriate time interval (e.g. every 12 hours)

Worst practice (Don't do this):

* Manually compile an outdated version on your piece of embedded equipment that serves as your Internet gateway
* Run the script via `go run`, providing the access keys as parameters that will be visible in your shell's history file
* Run this script every 0.5 seconds, resulting in API rate limits
