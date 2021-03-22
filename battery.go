package main

import (
	"fmt"
	"github.com/distatus/battery"
)

func batteryDisp() {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		return
	}
	for i, battery := range batteries {
		fmt.Printf("Bat%d: ", i)
		fmt.Printf("state: %s, ", battery.State.String())
		fmt.Printf("current capacity: %f mWh, ", battery.Current)
		fmt.Printf("last full capacity: %f mWh, ", battery.Full)
		fmt.Printf("design capacity: %f mWh, ", battery.Design)
		fmt.Printf("charge rate: %f mW, ", battery.ChargeRate)
		fmt.Printf("voltage: %f V, ", battery.Voltage)
		fmt.Printf("design voltage: %f V\n", battery.DesignVoltage)
	}
}
