package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/pteich/go-timeout-httpclient"
	"github.com/pteich/logger"

	"github.com/pteich/esphomekit/colorlight"
	"github.com/pteich/esphomekit/config"
	"github.com/pteich/esphomekit/devices"
	"github.com/pteich/esphomekit/sensor"
)

func main() {
	// read command line flags and init logger
	cfg := config.New()
	log := logger.New(logger.WithLogLevel(cfg.LogLevel), logger.WithLogConsole(cfg.LogConsole))
	httpClient := timeouthttp.New(timeouthttp.WithTimeout(15))

	// read existing esphome devices config from file
	accConfigs, err := config.NewEsphomeDevicesFromFile(cfg.ConfigFile)
	if err != nil {
		log.Fatal().Err(err).Str("file", cfg.ConfigFile).Msg("could not load accessories config")
	}

	log.Info().Interface("config", cfg).Interface("accessories", accConfigs).Msg("starting esphomekit")

	// add bridge
	var accID uint64 = 1
	bridge := accessory.NewBridge(accessory.Info{Name: "EspHomekit Bridge", ID: accID})

	// every esphome device need to have an accessory in HomeKit and and internal service that takes care of updates
	deviceList := make(devices.List, 0)
	for _, accConfig := range accConfigs {
		accID++
		switch accConfig.Type {
		case config.TypeColorLight:
			acc := accessory.NewColoredLightbulb(accessory.Info{
				ID:               accID,
				Name:             accConfig.Name,
				Manufacturer:     accConfig.Manufacturer,
				SerialNumber:     accConfig.SerialNumber,
				FirmwareRevision: accConfig.FirmwareRevision,
				Model:            accConfig.Model,
			})
			light := colorlight.New(accConfig.ID, accConfig.Addr, acc, httpClient, log)
			light.Init()
			deviceList = append(deviceList, light)

		case config.TypeTemperature:
			acc := accessory.NewTemperatureSensor(accessory.Info{
				ID:               accID,
				Name:             accConfig.Name,
				Manufacturer:     accConfig.Manufacturer,
				SerialNumber:     accConfig.SerialNumber,
				FirmwareRevision: accConfig.FirmwareRevision,
				Model:            accConfig.Model,
			}, 25, -15, 85, 0.1)
			tempsensor := sensor.NewTemperature(accConfig.ID, accConfig.Addr, acc, httpClient, log)
			tempsensor.Init()
			deviceList = append(deviceList, tempsensor)

		}
	}

	log.Info().Int("count", len(deviceList)).Msg("add accessories")
	// init HomeKit ip connection with pin
	hcConfig := hc.Config{
		Pin: cfg.Pin,
	}

	hcTransport, err := hc.NewIPTransport(hcConfig, bridge.Accessory, deviceList.GetAccessories()...)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating transport")
	}

	hc.OnTermination(func() {
		<-hcTransport.Stop()
	})

	// start HomeKit connection
	hcTransport.Start()
}
