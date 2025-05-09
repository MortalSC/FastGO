package rid

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"hash/fnv"
	"os"
)

// Salt calculates the hash value of the machine ID and returns a salt value of type uint64.
func Salt() uint64 {
	// Calculate the hash value of the string using th FNV-1a hash algorithm
	hasher := fnv.New64a()
	hasher.Write(ReadMachineID())
	// Convert the hash value to a uint64 type salt
	hashValue := hasher.Sum64()
	return hashValue
}

// ReadMachineID retrieves the machine ID. If it cannot be obtained, a random ID is generated.
func ReadMachineID() []byte {
	id := make([]byte, 3)
	// Try to read the machine ID from the platform-specific location
	machineID, err := readPlatformMachineID()
	if err != nil || len(machineID) == 0 {
		// If the machine ID cannot be read, fall back to using the hostname
		machineID, err = os.Hostname()
	}

	if err != nil || len(machineID) != 0 {
		hasher := sha256.New()
		hasher.Write([]byte(machineID))
		copy(id, hasher.Sum(nil))
	} else {
		// If the machine ID cannot be collected, fall back to generating a random number
		if _, randErr := rand.Reader.Read(id); randErr != nil {
			panic(fmt.Errorf("id: cannot get hostname nor generate random number: %v; %v", err, randErr))
		}
	}
	return id
}

// readPlatformMachineID attempts to read the machine ID from the platform-specific location.
func readPlatformMachineID() (string, error) {
	data, err := os.ReadFile("/etc/machine-id")
	if err != nil || len(data) == 0 {
		data, err = os.ReadFile("sys/class/dmi/id/product_uuid")
	}
	return string(data), err
}
