package sensor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/brutella/hc/accessory"
	"github.com/pteich/logger"
)

type Temperature struct {
	addr       string
	id         string
	accessory  *accessory.Thermometer
	httpClient *http.Client
	log        *logger.Logger
	scheduler  *time.Ticker
}

func NewTemperature(id string, addr string, acc *accessory.Thermometer, httpClient *http.Client, log *logger.Logger) *Temperature {
	return &Temperature{
		accessory:  acc,
		httpClient: httpClient,
		log:        log,
		addr:       addr,
		id:         id,
	}
}

func (t *Temperature) GetDeviceUrl() string {
	return fmt.Sprintf("http://%s/sensor/%s", t.addr, t.id)
}

func (t *Temperature) Init() {
	t.UpdateTemp()

	t.scheduler = time.NewTicker(15 * time.Second)
	go func() {
		for range t.scheduler.C {
			t.UpdateTemp()
		}
	}()
}

func (t *Temperature) UpdateTemp() {
	req, _ := http.NewRequest(http.MethodGet, t.GetDeviceUrl(), nil)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		t.log.Error().Err(err).Msg("esphome api")
		return
	}

	apiResponse := Response{}
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		t.log.Error().Err(err).Msg("esphome api")
		return
	}

	t.log.Info().Float64("temp", apiResponse.Value).Msg("update temp")
	t.accessory.TempSensor.CurrentTemperature.SetValue(apiResponse.Value)
}
