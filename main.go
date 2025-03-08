package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

/* Environment variables that the SDK will look for:
      - AWS_ACCESS_KEY_ID
      - AWS_SECRET_ACCESS_KEY
      - AWS_SESSION_TOKEN (optional)
(https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)
*/

const version = "1.1.0"

func main() {
	// Build the session object
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln("An error occurred while trying to create an AWS session:", err)
	}
	// Get a Route53 API wrapper from the session
	svc := route53.New(sess)

	// Ensure that the hosted zone exists first
	_, err = svc.GetHostedZone(&route53.GetHostedZoneInput{Id: hostedZoneID})
	if err != nil {
		log.Fatalln("Unable to find hosted zone with ID", *hostedZoneID, "-", err)
	}

	// Define the record that we're going to be UPSERTing
	resourceRecordSet := route53.ResourceRecordSet{
		Name: recordSet,
		Type: aws.String("A"),
		TTL:  recordTTL,
		ResourceRecords: []*route53.ResourceRecord{
			{Value: recordValue},
		},
	}

	// Define the Change struct and ChangeBatch
	changeBatch := &route53.ChangeBatch{
		Changes: []*route53.Change{
			{
				Action:            aws.String("UPSERT"),
				ResourceRecordSet: &resourceRecordSet,
			},
		},
	}
	if *changeComment != "" {
		changeBatch.Comment = changeComment
	}

	changeInput := route53.ChangeResourceRecordSetsInput{
		ChangeBatch:  changeBatch,
		HostedZoneId: hostedZoneID,
	}
	if verbose {
		log.Println("ChangeResourceRecordSetsInput", changeInput)
	}
	// Make the change
	changeOutput, err := svc.ChangeResourceRecordSets(&changeInput)
	if err != nil {
		log.Fatalln("An error occurred while trying to update the record set:", err)
	} else {
		log.Println("Change submitted. Current status:", *changeOutput.ChangeInfo.Status)
	}
}
