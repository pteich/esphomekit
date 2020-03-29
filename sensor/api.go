package sensor

// Response is the esphome api response for sensors
type Response struct {
	ID    string  `json:"id"`
	State string  `json:"state"`
	Value float64 `json:"value"`
}
