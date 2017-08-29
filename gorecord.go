package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
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

var hosted_zone string;
var record_set string;
var record_value string;
var record_type string;
var record_ttl int64;
var change_comment string;
var verbose bool;

func init() {
  /*  Load variables from environment if they weren't assigned a value at build
      time.
  */
  if len(os.Getenv("HOSTED_ZONE")) != 0 && len(hosted_zone) == 0 {
    hosted_zone = os.Getenv("HOSTED_ZONE")
  }
  if len(os.Getenv("RECORD_SET")) != 0 &&len(record_set) == 0 {
    record_set = os.Getenv("RECORD_SET")
  }
  if len(os.Getenv("RECORD_VALUE")) != 0 && len(record_value) == 0 {
    record_value = os.Getenv("RECORD_VALUE")
  }
  if len(os.Getenv("RECORD_TYPE")) != 0 && len(record_type) == 0 {
    record_type = os.Getenv("RECORD_TYPE")
  }
  if len(os.Getenv("RECORD_TTL")) != 0 && record_ttl == 0 {
    env_ttl, err := strconv.ParseInt(os.Getenv("RECORD_TTL"), 10, 64)
    if err != nil {
      fmt.Fprintf(
        os.Stderr,
        "[CRITICAL] Unsupported value for RECORD_TTL: %v\n",
        err,
      )
      os.Exit(1)
    }
    record_ttl = env_ttl
  }
  if len(os.Getenv("CHANGE_COMMENT")) != 0 && len(change_comment) == 0 {
    change_comment = os.Getenv("CHANGE_COMMENT")
  } else {
    change_comment = fmt.Sprintf(
      "GoRecord change at %s",
      time.Now().UTC(),
    )
  }

  /*  Complain if we still don't have the values we need, or set a sensible
      default.
  */
  if len(hosted_zone) == 0 {
    fmt.Fprintf(os.Stderr, "[CRITICAL] HOSTED_ZONE not defined.\n")
    os.Exit(2)
  }
  if len(record_set) == 0 {
    fmt.Fprintf(os.Stderr, "[CRITICAL] RECORD_SET not defined.\n")
    os.Exit(3)
  }
  if len(record_value) == 0 {
    fmt.Fprintf(os.Stderr, "[CRITICAL] RECORD_VALUE not defined.\n")
    os.Exit(4)
  }
  if len(record_type) == 0 {
    record_type = "A"
  }
  if record_ttl == 0 {
    record_ttl = int64(600)
  }
  if len(os.Getenv("VERBOSE")) == 0 {
    verbose = false
  } else {
    verbose = true
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
    Name: &record_set,
    Type: &record_type,
    TTL: &record_ttl,
    ResourceRecords: []*route53.ResourceRecord{ &resource_record },
  }

  // Define the Change struct and ChangeBatch
  change := route53.Change{
    Action: aws.String("UPSERT"),
    ResourceRecordSet: &resource_record_set,
  }
  var change_array []*route53.Change
  change_array = append(change_array, &change)
  change_batch := route53.ChangeBatch{
    Changes: change_array,
    Comment: &change_comment,
  }
  change_input := route53.ChangeResourceRecordSetsInput{
    ChangeBatch: &change_batch,
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
