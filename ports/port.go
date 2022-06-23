package ports

type Coordinates struct {
	Altitude  float64
	Longitude float64
}

type Port struct {
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string // this field seems empty for test dataset, consider removing
	Coordinates Coordinates
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}
