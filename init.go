package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func init() {
	/*  Load variables from environment if they weren't assigned a value at build
	    time.
	*/
	if len(os.Getenv("HOSTED_ZONE")) != 0 && len(hosted_zone) == 0 {
		hosted_zone = os.Getenv("HOSTED_ZONE")
	}
	if len(os.Getenv("RECORD_SET")) != 0 && len(record_set) == 0 {
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
