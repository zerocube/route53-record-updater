package main

import (
  "fmt"
  "os"
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

var hosted_zone string;
var record_set string;
var record_value string;

func init() {
  hosted_zone = os.Getenv("HOSTED_ZONE")
  record_set = os.Getenv("RECORD_SET")
  record_value = os.Getenv("RECORD_VALUE")
  if len(hosted_zone) == 0 {
    fmt.Fprintf(os.Stderr, "[CRITICAL] HOSTED_ZONE not defined.\n")
    os.Exit(1)
  }
  if len(record_set) == 0 {
    fmt.Fprintf(os.Stderr, "[CRITICAL] RECORD_SET not defined.\n")
    os.Exit(2)
  }
  if len(record_value) == 0 {
    fmt.Fprintf(os.Stderr, "[CRITICAL] RECORD_VALUE not defined.\n")
    os.Exit(3)
  }
}

func main() {
  // Build the session object
  sess, err := session.NewSession()
  if err != nil {
    fmt.Fprintf(
      os.Stderr,
      "An error occurred while trying to create a new AWS session:\n%s\n",
      err.Error(),
    )
    os.Exit(4)
  }
  // Get a Route53 API wrapper from the session
  svc := route53.New(sess)

  // Create the HostedZoneInput struct
  hosted_zone_input := route53.GetHostedZoneInput{Id: &hosted_zone}

  // Ensure that the hosted zone exists first
  hosted_zone_output, err := svc.GetHostedZone(&hosted_zone_input)
  if err != nil {
    fmt.Fprintf(
      os.Stderr,
      "[CRITICAL] Unable to find zone with ID '%s'.\n%v",
      hosted_zone,
      err.Error(),
    )
    os.Exit(5)
  }

  // Define the record that we're going to be UPSERTing
  record_type := "A"
  record_ttl := int64(600)
  resource_record_set := route53.ResourceRecordSet{
    Name: &record_set,
    Type: &record_type,
    TTL: &record_ttl,
  }

  // Define the change struct
  change_action := "UPSERT"
  change := route53.Change{
    Action: &change_action,
    ResourceRecordSet: &resource_record_set,
  }
  var change_array []*route53.Change
  change_array = append(change_array, &change)
  change_batch := route53.ChangeBatch{Changes: change_array }
  change_set := route53.ChangeResourceRecordSetsInput{
    ChangeBatch: &change_batch,
    HostedZoneId: hosted_zone_output.HostedZone.Id,
  }
  // Make the change
  change_output, err := svc.ChangeResourceRecordSets(&change_set)
  if err != nil {
    fmt.Fprintf(
      os.Stderr,
      "[ERROR] An error occurred while trying to update the record set.\n%v\n",
      err,
    )
    os.Exit(6)
  } else {
    fmt.Printf(
      "Change submitted. Current status: %s",
      change_output.ChangeInfo.Status,
    )
  }
}
