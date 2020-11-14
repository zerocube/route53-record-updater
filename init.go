package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	hostedZoneID = flag.String("zone-id", "", "The hosted zone ID")
	flag.StringVar(hostedZoneID, "zone", *hostedZoneID, "alias for -zone-id")

	recordSet = flag.String("record", "", "The record to update, e.g. host01.zerocube.com.au")
	recordValue = flag.String("value", "", "The record value, e.g. 127.0.0.1")
	changeComment = flag.String("comment", "", "Change comment (optional)")
	recordTTL = flag.Int64("ttl", int64(600), "TTL of the record (Default: 600 seconds)")

	flag.BoolVar(&verbose, "verbose", false, "Enables verbose output")

	outputVersion := flag.Bool("version", false, "Outputs version information and exits.")

	flag.Parse()

	if *outputVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	// Some sanity checks
	var missingVars bool

	if *hostedZoneID == "" {
		log.Println("Missing: Hosted Zone ID (--zone-id)")
		missingVars = true
	}
	if *recordSet == "" {
		log.Println("Missing: Record Set (--record)")
		missingVars = true
	}
	if *recordValue == "" {
		log.Println("Missing: Record Value (--value)")
		missingVars = true
	}

	if missingVars {
		os.Exit(1)
	}
}
