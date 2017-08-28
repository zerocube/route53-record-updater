package main

import (
  "flag"
  "fmt"
  "os"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/route53"
)

/* Environment variables that the SDK will look for:
(https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)
  - AWS_ACCESS_KEY_ID
  - AWS_SECRET_ACCESS_KEY
  - AWS_SESSION_TOKEN (optional)
*/

var hosted_zone string    // The hosted zone (e.g. jordankueh.com)
var record_set string     // The record set (e.g. remoteserver.jordankueh.com)
var record_value string   // The record value (e.g. 127.0.1.1)
/*
  While the record_value variable is specified here, if using build arguments
  to set these variables, record_value should be left untouched as this data
  is prone to change, especially on residential connections where the DHCP lease
  renewal may occur frequently.
*/

func init() {
  hosted_zone := flag.String("z", "", "The hosted zone / domain name.")
  record_set := flag.String("r", "", "The record set / record to update.")
  record_value := flag.String("a", "", "The new A record value.")
  flag.Parse()
}

func main() {
  /*  Temporary - Because I haven't actually used these variables anywhere yet,
      the golang compiler is (rightfully) complaining that I haven't used them
      yet.
  */

  hosted_zone_ptr = &hosted_zone
  record_set_ptr = &record_set
  record_value_ptr = &record_value
  // Build the session object
  sess, err := session.NewSession()
  if err != nil {
    fmt.Fprintf(
      os.Stderr,
      "An error occurred while trying to create a new AWS session:\n%s\n",
      err.Error(),
    )
    os.Exit(1)
  }
  // Get a Route53 API wrapper from the session
  svc := route53.New(sess)

}