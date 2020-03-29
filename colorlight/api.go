package colorlight

// Response is the esphome api response for lights
type Response struct {
	ID         string `json:"id"`
	State      string `json:"state"`
	Brightness int    `json:"brightness"`
	Color      struct {
		R int `json:"r"`
		G int `json:"g"`
		B int `json:"b"`
	} `json:"color"`
}
