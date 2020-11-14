package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

/*  Environment variables that this script needs:
      - HOSTED_ZONE
      - RECORD_SET
      - RECORD_VALUE
    Environment variables that the SDK will look for:
      - AWS_ACCESS_KEY_ID
      - AWS_SECRET_ACCESS_KEY
      - AWS_SESSION_TOKEN (optional)
(https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)
*/

var hosted_zone string
var record_set string
var record_value string
var record_type string
var record_ttl int64
var change_comment string
var verbose bool

func main() {
	// Build the session object
	sess, err := session.NewSession()
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"An error occurred while trying to create a new AWS session:\n%s\n",
			err.Error(),
		)
		os.Exit(5)
	}
	// Get a Route53 API wrapper from the session
	svc := route53.New(sess)

	// Create the HostedZoneInput struct
	hosted_zone_input := route53.GetHostedZoneInput{Id: &hosted_zone}

	// Ensure that the hosted zone exists first
	_, err = svc.GetHostedZone(&hosted_zone_input)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"[CRITICAL] Unable to find zone with ID '%s'.\n%v\n",
			hosted_zone,
			err.Error(),
		)
		os.Exit(6)
	}

	// Define the record that we're going to be UPSERTing
	resource_record := route53.ResourceRecord{
		Value: &record_value,
	}
	resource_record_set := route53.ResourceRecordSet{
		Name:            &record_set,
		Type:            &record_type,
		TTL:             &record_ttl,
		ResourceRecords: []*route53.ResourceRecord{&resource_record},
	}

	// Define the Change struct and ChangeBatch
	change := route53.Change{
		Action:            aws.String("UPSERT"),
		ResourceRecordSet: &resource_record_set,
	}
	var change_array []*route53.Change
	change_array = append(change_array, &change)
	change_batch := route53.ChangeBatch{
		Changes: change_array,
		Comment: &change_comment,
	}
	change_input := route53.ChangeResourceRecordSetsInput{
		ChangeBatch:  &change_batch,
		HostedZoneId: &hosted_zone,
	}
	if verbose {
		fmt.Printf("ChangeResourceRecordSetsInput: %v\n", change_input)
	}
	// Make the change
	change_output, err := svc.ChangeResourceRecordSets(&change_input)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"[ERROR] An error occurred while trying to update the record set.\n%v\n",
			err,
		)
		os.Exit(7)
	} else {
		fmt.Printf(
			"Change submitted. Current status: %v\n",
			change_output.ChangeInfo.Status,
		)
	}
}
