package config

import (
	"encoding/json"
	"os"
)

const (
	TypeAccessoryColorLight  = "colorlight"
	TypeAccessoryTemperature = "temperature"
)

// EsphomeDevices is the list of esphome devices in config file
type EsphomeDevices []EsphomeDevice

// EsphomeDevice holds the configuration for one device
type EsphomeDevice struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Addr             string `json:"addr"`
	Type             string `json:"type"`
	SerialNumber     string `json:"serial_number"`
	Manufacturer     string `json:"manufacturer"`
	Model            string `json:"model"`
	FirmwareRevision string `json:"firmware_revision"`
}

// NewEsphomeDevicesFromFile reads a JSON config file and returns all devices
func NewEsphomeDevicesFromFile(filename string) (EsphomeDevices, error) {
	devices := make(EsphomeDevices, 0)

	f, err := os.Open(filename)
	if err != nil {
		return devices, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&devices)
	if err != nil {
		return devices, err
	}

	return devices, nil
}
