package slimuuid

import (
    "net"
)

/*
  this function generates a mac address of the machine in string format
  always call the function like => 
  macId, err := MacID()
  if err != nil {
    return "" , err
  }
  fmt.Println(macId)
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