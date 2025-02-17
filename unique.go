package slimuuid

import (
    "net"
)

/*
  this function generates a mac address of the machine in string format
  always call the function like this so you can get it in string format  => 
  ```
	package main
	import (
		"encoding/hex"
		"fmt"
		"strings"
		"github.com/theMitocondria/slimuuid"
	)
	
	func formatMAC(macRaw string) string {
		// Convert the string to bytes, then get a hex string.
		hexStr := hex.EncodeToString([]byte(macRaw))
		// Insert ":" every 2 characters.
		var parts []string
		for i := 0; i < len(hexStr); i += 2 {
			parts = append(parts, hexStr[i:i+2])
		}
		return strings.Join(parts, ":")
	}
	
	func main() {
		macRaw, err := slimuuid.MacID()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println( formatMAC(macRaw))
	}
  ```
  This function is not supported on windows machines 
*/
func MacID()  ( string , error ) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return  "" , err
    }
    for _, iface := range interfaces {
        if iface.HardwareAddr != nil {
            return  string(iface.HardwareAddr) , nil 
        }
    }
	return "" , nil
}


// MacID returns the MAC address of a machine on Windows.
func MacIDForWindows() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	
	for _, iface := range interfaces {
		// Skip loopback and interfaces with no valid hardware address.
		if iface.Flags&net.FlagLoopback != 0 || len(iface.HardwareAddr) != 6 {
			continue
		}
		return iface.HardwareAddr.String(), nil
	}
	
	return "", nil
}

// MacID returns the MACOS address of a machine on macOS.
func MacIDForDarwin() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		// Skip loopback and interfaces with no valid hardware address.
		if iface.Flags&net.FlagLoopback != 0 || len(iface.HardwareAddr) != 6 {
			continue
		}
		return iface.HardwareAddr.String(), nil
	}

	return "", nil
}