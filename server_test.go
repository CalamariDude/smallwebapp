package main

import (
	"testing"
	"regexp"
)

// TestGetLocalAddressesValid ensures values returned by
// GetLocalAddresses are in the correct format for an ipv4 address with mask
// ie 192.168.0.2/24
func TestGetLocalAddressesValid(t *testing.T){
	// Regex doesn't catch non-valid ip addresses, just that the ip address
	// is formatted correctly ie 256.255.255.255/5000 would be valid
	// even though it is not a valid ipv4 addresse
	ipaddressmatch := regexp.MustCompile(`\d*\.\d*\.\d*\.\d*\/\d*`)
	addrs, err := GetLocalAddresses()
	for _, addr:= range addrs {
		if addr == "" || !ipaddressmatch.MatchString(addr) || err != nil {
			t.Fatalf(`GetLocalAddress() = %q, %v, want match for %#q, nil`, addr, err, ipaddressmatch)
		}
	}
}

