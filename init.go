package main

import (
	"flag"
)

/* The parameters we pull in via flags are:
- hostedZoneID (the Hosted Zone Domain Name, e.g. 'zerocube.com.au')
- recordSet (the FQDN, e.g. 'host01.zerocube.com.au')
- recordValue (the IP to resolve, e.g. '127.0.0.1')
- recordTTL (if provided: The TTL of the record)
- changeComment (if provided: The comment for the change)
*/
var hostedZoneID, recordSet, recordValue, changeComment *string
var recordTTL *int64

var verbose bool

func init() {
	/*  Load variables from environment if they weren't assigned a value at build
	    time.
	*/

	hostedZoneID = flag.String("zone-id", "", "The hosted zone ID")
	flag.StringVar(hostedZoneID, "zone", *hostedZoneID, "alias for -zone-id")

	recordSet = flag.String("record", "", "The record to update, e.g. host01.zerocube.com.au")
	recordValue = flag.String("value", "", "The record value, e.g. 127.0.0.1")
	changeComment = flag.String("comment", "", "Change comment (optional)")
	recordTTL = flag.Int64("ttl", int64(600), "TTL of the record")

	flag.BoolVar(&verbose, "verbose", false, "Enables verbose output")

	flag.Parse()
}
