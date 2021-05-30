package main

import (
	"fmt"

	"github.com/Wifx/gonetworkmanager"
	"github.com/gotk3/gotk3/gtk"
)

func InitNetwork() (gtk.IWidget, error) {
	/* Create new instance of gonetworkmanager */
	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		return nil, err
	}

	/* Get devices */
	devices, err := nm.GetPropertyAllDevices()
	if err != nil {
		return nil, err
	}

	/* Show each device path and interface name */
	for _, device := range devices {

		byt, err := device.MarshalJSON()
		if err != nil {
			return nil, err
		}
		fmt.Printf("%s\n", byt)
		// deviceInterface, err := device.GetPropertyInterface()
		// if err != nil {
		// 	return nil, err
		// }

		// fmt.Println(deviceInterface + " - " + string(device.GetPath()))
	}
	return nil, nil
}
