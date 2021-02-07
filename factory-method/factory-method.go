package main

import "fmt"

// ModeOfTransfer : required methods interface of transfer mode
type ModeOfTransfer interface {
	Diliver() bool
}

// Truck : Different modes of transfer
type Truck struct {
	Name string `default:"truck"`
}

// Diliver : must implemented by all the modes
func (t *Truck) Diliver() bool {
	fmt.Println("Dilivering via Truck")
	return true
}

// Ship : Different modes of transfer
type Ship struct {
	Name string `default:"ship"`
}

// Diliver : must implemented by all the modes
func (t *Ship) Diliver() bool {
	fmt.Println("Dilivering via Ship")
	return true
}

// ModeArray : Contains all types of modes
type ModeArray struct {
	Truck
	Ship
}

// LogisticsAppInterface : Factory method
type LogisticsAppInterface interface {
	DiliveryMode() ModeOfTransfer
}

// LogisticsApp : struct of the app
type LogisticsApp struct {
	Name          string
	TransferModes ModeArray // Just to store all types of modes not required
	DefaultMode   string
}

// DiliveryMode : factory method implementation in creator class
func (l *LogisticsApp) DiliveryMode() ModeOfTransfer {
	var obj ModeOfTransfer
	switch l.DefaultMode {
	case "truck":
		obj = ModeOfTransfer(&Truck{})
	case "ship":
		obj = ModeOfTransfer(&Ship{})
	default:
		obj = ModeOfTransfer(&Truck{})

	}
	return obj
}

func main() {
	Logistics := LogisticsApp{Name: "BigLogistics"}
	Logistics.DefaultMode = "truck" // this can be set according to the business requirement
	mode := Logistics.DiliveryMode()
	mode.Diliver()
}
