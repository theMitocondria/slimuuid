package slimuuid

import (
    "net"
)

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