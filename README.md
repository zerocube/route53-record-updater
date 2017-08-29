# gorecord
[![pipeline status](https://gitlab.com/jordankueh/GoRecord/badges/master/pipeline.svg)](https://gitlab.com/jordankueh/GoRecord/commits/master)

This kind of script is almost definitely not something you want to run in a
production environment, as:

* In a production environment, you have a static IP address; or
* In a production environment, no sane person runs their sole production
infrastructure on a network that has an address lease,

Unless:

* Under some really strange circumstances, this kind of stunt might just be what
you need.

## Downloading

### Downloading a precompiled binary

[The GitLab pipelines page](https://gitlab.com/jordankueh/GoRecord/pipelines)
will allow you to download a pre-built binary for your platform.

### Downloading via the source
Cloning from the source is also possible, from:

* https://gitlab.com/jordankueh/GoRecord; or
* https://github.com/jordankueh/GoRecord.

The 'master' branch is the most stable version available.

## Usage

Best practice:

* Get a static IP address from your provider.

Next-to-best practice:

* Compile the binary with `go build` (setting GOOS and GOOARCH as necessary),
and place it in `/usr/local/sbin` (ensuring that it's executable);
* Create a wrapper script (see gorecord.sh) that exports the following
environment variables:
  * `HOSTED_ZONE_ID`
  * `RECORD_SET`
  * `RECORD_VALUE`
  * `AWS_ACCESS_KEY_ID`
  * `AWS_SECRET_ACCESS_KEY`
  * `AWS_SESSION_TOKEN` (optional)
* Run the wrapper script (ensuring that it too is executable); and
* Call the wrapper script via cron at an appropriate time interval (e.g. every
12 hours).

Worst practice (Don't do this):

* Manually compile an outdated version of golang on your piece of embedded
equipment that serves as your Internet gateway;
* Run the script via `go run`, and make sure you provide the access keys as
parameters that will be visible in your shell's history file; and
* Run this script every 0.5 seconds, resulting in API rate limits.
