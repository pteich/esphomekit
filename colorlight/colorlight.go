package colorlight

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/brutella/hc/accessory"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/pteich/logger"
)

type Colorlight struct {
	brightness int // brightness in percent 0..100
	color      colorful.Color
	on         bool
	addr       string
	id         string
	accessory  *accessory.ColoredLightbulb
	httpClient *http.Client
	log        *logger.Logger
	scheduler  *time.Timer
}

func New(id string, addr string, acc *accessory.ColoredLightbulb, httpClient *http.Client, log *logger.Logger) *Colorlight {
	return &Colorlight{
		accessory:  acc,
		httpClient: httpClient,
		log:        log,
		addr:       addr,
		id:         id,
	}
}

func (cl *Colorlight) Init() {

	req, _ := http.NewRequest(http.MethodGet, cl.GetDeviceUrl(), nil)

	resp, err := cl.httpClient.Do(req)
	if err != nil {
		cl.log.Error().Err(err).Msg("esphome api")
		return
	}

	apiResponse := Response{}
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		cl.log.Error().Err(err).Msg("esphome api")
		return
	}

	if apiResponse.State == "ON" {
		cl.on = true
	}

	cl.brightness = apiResponse.Brightness / 255 * 100
	cl.color = colorful.Color{
		R: float64(apiResponse.Color.R) / 255,
		G: float64(apiResponse.Color.G) / 255,
		B: float64(apiResponse.Color.B) / 255,
	}

	cl.log.Debug().Interface("response", apiResponse).Msg("api status response")

	cl.accessory.Lightbulb.On.OnValueRemoteGet(func() bool {
		cl.log.Info().Bool("value", cl.GetState()).Msg("switch get")
		return cl.GetState()
	})

	cl.accessory.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		cl.log.Info().Bool("on", on).Msg("switch update")
		cl.State(on)
		cl.ScheduleLightOn()
	})

	cl.accessory.Lightbulb.Brightness.OnValueRemoteGet(func() int {
		cl.log.Info().Int("brightness", cl.GetBrightness()).Msg("brightness get")
		return cl.GetBrightness()
	})

	cl.accessory.Lightbulb.Brightness.OnValueRemoteUpdate(func(value int) {
		cl.log.Info().Int("value", value).Msg("brightness update")
		cl.Brightness(value)
		cl.ScheduleLightOn()
	})

	cl.accessory.Lightbulb.Hue.OnValueRemoteUpdate(func(hue float64) {
		cl.log.Info().Float64("hue", hue).Msg("hue update")
		_, s, l := cl.GetColor().Hsl()
		cl.SetColor(colorful.Hsl(hue, s, l))
		cl.ScheduleLightOn()
	})

	cl.accessory.Lightbulb.Hue.OnValueRemoteGet(func() float64 {
		h, _, _ := cl.GetColor().Hsl()
		cl.log.Info().Float64("hue", h).Msg("hue get")
		return h
	})

	cl.accessory.Lightbulb.Saturation.OnValueRemoteUpdate(func(sat float64) {
		cl.log.Info().Float64("saturation", sat).Msg("saturation update")
		h, _, l := cl.GetColor().Hsl()
		cl.SetColor(colorful.Hsl(h, sat/100, l))
		cl.ScheduleLightOn()
	})

	cl.accessory.Lightbulb.Saturation.OnValueRemoteGet(func() float64 {
		cl.log.Info().Msg("saturation get")
		_, s, _ := cl.GetColor().Hsl()
		return s * 100
	})
}

func (cl *Colorlight) Brightness(brightness int) {
	cl.brightness = brightness
}

func (cl *Colorlight) ScheduleLightOn() {
	if cl.scheduler != nil {
		return
	}

	cl.scheduler = time.AfterFunc(20*time.Millisecond, func() {
		cl.LightOn()
		cl.scheduler = nil
	})
}

func (cl *Colorlight) LightOn() {
	if !cl.on {
		cl.LightOff()
		return
	}

	req, _ := http.NewRequest(http.MethodPost, cl.GetDeviceUrl()+"/turn_on", nil)
	params := req.URL.Query()
	params.Add("transition", "1")
	r, g, b := cl.color.RGB255()
	params.Add("r", strconv.Itoa(int(r)))
	params.Add("g", strconv.Itoa(int(g)))
	params.Add("b", strconv.Itoa(int(b)))
	params.Add("brightness", strconv.Itoa(int(float64(cl.brightness)/100*255)))
	req.URL.RawQuery = params.Encode()

	cl.log.Info().Str("url", req.URL.String()).Msg("esphome request")
	start := time.Now()
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		cl.log.Error().Err(err).Msg("esphome api")
		return
	}

	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	cl.log.Info().Dur("duration", time.Now().Sub(start)).Msg("esphome request finished")
}

func (cl *Colorlight) LightOff() {
	req, _ := http.NewRequest(http.MethodPost, cl.GetDeviceUrl()+"/turn_off?transition=1", nil)

	resp, err := cl.httpClient.Do(req)
	if err != nil {
		cl.log.Error().Err(err).Msg("esphome api")
		return
	}

	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	cl.log.Info().Str("url", req.URL.String()).Msg("esphome response")
}

func (cl *Colorlight) State(state bool) {
	cl.on = state
}

func (cl *Colorlight) GetState() bool {
	return cl.on
}

func (cl *Colorlight) GetBrightness() int {
	return cl.brightness
}

func (cl *Colorlight) GetColor() colorful.Color {
	return cl.color
}

func (cl *Colorlight) SetColor(color colorful.Color) {
	cl.color = color
}

func (cl *Colorlight) GetDeviceUrl() string {
	return fmt.Sprintf("http://%s/light/%s", cl.addr, cl.id)
}

func (cl *Colorlight) GetAccessory() *accessory.Accessory {
	return cl.accessory.Accessory
}
