package ports

// TODO: implement coordinates struct in the json decoder. Using a slice for time constraints
type Coordinates struct {
	Altitude  float64 `json:"altitude"`
	Longitude float64 `json:"longitude"`
}

type Port struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
	// Alias       []string `json:"alias"` // this field seems empty for test dataset, removing to save some memory
	// Regions     []string `json:"regions"` // this field seems empty for test dataset, removing to save some memory
	// Coordinates Coordinates `json:"coordinates"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}
