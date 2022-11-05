package port

import "go.bug.st/serial"

func List() ([]string, error) {
	return serial.GetPortsList()
}

func Open(portName string) (Port, error) {
	mode := &serial.Mode{
		BaudRate: 115200,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}
	return serial.Open(portName, mode)
}
