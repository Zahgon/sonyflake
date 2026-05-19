// Package awsutil provides utility functions for using Sonyflake on AWS.
package awsutil

import (
	"net"
	"time"
)

func amazonEC2PrivateIPv4() (net.IP, error) { _ = "STUB: not implemented"; return *new(net.IP), nil }

// AmazonEC2MachineID retrieves the private IP address of the Amazon EC2 instance
// and returns its lower 16 bits.
// It works correctly on Docker as well.
func AmazonEC2MachineID() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// TimeDifference returns the time difference between the localhost and the given NTP server.
func TimeDifference(server string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
