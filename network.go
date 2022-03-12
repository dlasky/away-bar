package main

import (
	"fmt"

	"github.com/Wifx/gonetworkmanager"
	"github.com/gotk3/gotk3/gtk"
)

type NetworkData struct {
	Connection string
}

func InitNetwork() (gtk.IWidget, error) {

	module, err := NewModule("network", "{{.Connection}}", "", "./feather/wifi.svg")
	if err != nil {
		return nil, err
	}

	nm, err := gonetworkmanager.NewNetworkManager()
	if err != nil {
		return nil, err
	}

	devices, err := nm.GetPropertyAllDevices()
	if err != nil {
		return nil, err
	}

	for _, device := range devices {

		state, err := device.GetPropertyState()
		if err != nil {
			return nil, err
		}
		if state == gonetworkmanager.NmDeviceStateActivated {
			path := device.GetPath()
			name, err := device.GetPropertyInterface()
			if err != nil {
				return nil, err
			}
			fmt.Println(path, name)
			typ, err := device.GetPropertyDeviceType()
			if err != nil {
				return nil, err
			}
			switch typ {
			case gonetworkmanager.NmDeviceTypeEthernet:
				data := NetworkData{Connection: "wired"}
				module.Render(data)
			case gonetworkmanager.NmDeviceTypeWifi:
				data := NetworkData{Connection: "wifi"}
				module.Render(data)
			case gonetworkmanager.NmDeviceTypeTun:
				data := NetworkData{Connection: "vpn"}
				module.Render(data)
			}
		}

		// byt, err := device.MarshalJSON()
		// if err != nil {
		// 	return nil, err
		// }
		// fmt.Printf("%s\n", byt)
		// deviceInterface, err := device.GetPropertyInterface()
		// if err != nil {
		// 	return nil, err
		// }

		// fmt.Println(deviceInterface + " - " + string(device.GetPath()))
	}
	return module.GetWidget(), nil
}
