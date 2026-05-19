// Package sonyflake implements Sonyflake, a distributed unique ID generator inspired by Twitter's Snowflake.
//
// By default, a Sonyflake ID is composed of
//
//	39 bits for time in units of 10 msec
//	 8 bits for a sequence number
//	16 bits for a machine id
package sonyflake

import (
	"errors"
	"net"
	"sync"
	"time"

	"github.com/sony/sonyflake/v2/types"
)

// Settings configures Sonyflake:
//
// BitsSequence is the bit length of a sequence number.
// If BitsSequence is 0, the default bit length is used, which is 8.
// If BitsSequence is 31 or more, an error is returned.
//
// BitsMachineID is the bit length of a machine ID.
// If BitsMachineID is 0, the default bit length is used, which is 16.
// If BitsMachineID is 31 or more, an error is returned.
//
// TimeUnit is the time unit of Sonyflake.
// If TimeUnit is 0, the default time unit is used, which is 10 msec.
// TimeUnit must be 1 msec or longer.
//
// StartTime is the time since which the Sonyflake time is defined as the elapsed time.
// If StartTime is 0, the start time of the Sonyflake instance is set to "2025-01-01 00:00:00 +0000 UTC".
// StartTime must be before the current time.
//
// MachineID returns the unique ID of a Sonyflake instance.
// If MachineID returns an error, the instance will not be created.
// If MachineID is nil, the default MachineID is used, which returns the lower 16 bits of the private IP address.
//
// CheckMachineID validates the uniqueness of a machine ID.
// If CheckMachineID returns false, the instance will not be created.
// If CheckMachineID is nil, no validation is done.
//
// The bit length of time is calculated by 63 - BitsSequence - BitsMachineID.
// If it is less than 32, an error is returned.
type Settings struct {
	BitsSequence   int
	BitsMachineID  int
	TimeUnit       time.Duration
	StartTime      time.Time
	MachineID      func() (int, error)
	CheckMachineID func(int) bool
}

// Sonyflake is a distributed unique ID generator.
type Sonyflake struct {
	mutex *sync.Mutex

	bitsTime     int
	bitsSequence int
	bitsMachine  int

	timeUnit    int64
	startTime   int64
	elapsedTime int64

	sequence int
	machine  int

	now func() time.Time
}

var (
	ErrInvalidBitsTime      = errors.New("bit length for time must be 32 or more")
	ErrInvalidBitsSequence  = errors.New("invalid bit length for sequence number")
	ErrInvalidBitsMachineID = errors.New("invalid bit length for machine id")
	ErrInvalidTimeUnit      = errors.New("invalid time unit")
	ErrInvalidSequence      = errors.New("invalid sequence number")
	ErrInvalidMachineID     = errors.New("invalid machine id")
	ErrStartTimeAhead       = errors.New("start time is ahead")
	ErrOverTimeLimit        = errors.New("over the time limit")
	ErrNoPrivateAddress     = errors.New("no private ip address")
)

const (
	defaultTimeUnit = 1e7 // nsec, i.e. 10 msec

	defaultBitsTime     = 39
	defaultBitsSequence = 8
	defaultBitsMachine  = 16
)

var defaultInterfaceAddrs = net.InterfaceAddrs

// New returns a new Sonyflake configured with the given Settings.
// New returns an error in the following cases:
// - Settings.BitsSequence is less than 0 or greater than 30.
// - Settings.BitsMachineID is less than 0 or greater than 30.
// - Settings.BitsSequence + Settings.BitsMachineID is 32 or more.
// - Settings.TimeUnit is less than 1 msec.
// - Settings.StartTime is ahead of the current time.
// - Settings.MachineID returns an error.
// - Settings.CheckMachineID returns false.
func New(st Settings) (*Sonyflake, error) { _ = "STUB: not implemented"; return nil, nil }

// NextID generates a next unique ID as int64.
// After the Sonyflake time overflows, NextID returns an error.
func (sf *Sonyflake) NextID() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (sf *Sonyflake) toInternalTime(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func (sf *Sonyflake) currentElapsedTime() int64 { _ = "STUB: not implemented"; return 0 }

func (sf *Sonyflake) sleep(overtime int64) { _ = "STUB: not implemented"; return }

func (sf *Sonyflake) toID() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func privateIPv4(interfaceAddrs types.InterfaceAddrs) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func isPrivateIPv4(ip net.IP) bool {
	_ = "STUB: not implemented"
	// Allow private IP addresses (RFC1918) and link-local addresses (RFC3927)
	return false
}

func lower16BitPrivateIP(interfaceAddrs types.InterfaceAddrs) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ToTime returns the time when the given ID was generated.
func (sf *Sonyflake) ToTime(id int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Compose creates a Sonyflake ID from its components.
// The time parameter should be the time when the ID was generated.
// The sequence parameter should be between 0 and 2^BitsSequence-1 (inclusive).
// The machineID parameter should be between 0 and 2^BitsMachineID-1 (inclusive).
func (sf *Sonyflake) Compose(t time.Time, sequence, machineID int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Decompose returns a set of Sonyflake ID parts.
func (sf *Sonyflake) Decompose(id int64) map[string]int64 { _ = "STUB: not implemented"; return nil }

func (sf *Sonyflake) timePart(id int64) int64 { _ = "STUB: not implemented"; return 0 }

func (sf *Sonyflake) sequencePart(id int64) int64 { _ = "STUB: not implemented"; return 0 }

func (sf *Sonyflake) machinePart(id int64) int64 { _ = "STUB: not implemented"; return 0 }
