package ports

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestLoadPorts(t *testing.T) {
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name      string
		args      args
		wantPorts map[string]Port
	}{
		{"one port", args{data: strings.NewReader(`{  "AEQIW": {
			"name": "Umm al Qaiwain",
			"coordinates": [
			  55.55,
			  25.57
			],
			"city": "Umm al Qaiwain",
			"country": "United Arab Emirates",
			"alias": [],
			"regions": [],
			"province": "Umm Al Quwain",
			"timezone": "Asia/Dubai",
			"unlocs": [
			  "AEQIW"
			],
			"code": "13400"
		  }}`)}, map[string]Port{"AEQIW": {
			Name:        "Umm al Qaiwain",
			Coordinates: []float64{55.55, 25.57},
			City:        "Umm al Qaiwain",
			Province:    "Umm Al Quwain",
			Country:     "United Arab Emirates",
			Timezone:    "Asia/Dubai",
			Unlocs:      []string{"AEQIW"},
			Code:        "13400",
		}}},
		{"two ports", args{data: strings.NewReader(`{"INBLR": {
			"name": "Bangalore",
			"coordinates": [
			  77.58,
			  12.98
			],
			"city": "Bangalore",
			"province": "Karnataka",
			"country": "India",
			"alias": [],
			"regions": [],
			"timezone": "Asia/Calcutta",
			"unlocs": [
			  "INBLR"
			],
			"code": "53398"
		  },  "AUADL": {
			"name": "Adelaide",
			"coordinates": [
			  138.58,
			  -34.92
			],
			"city": "Adelaide",
			"province": "South Australia",
			"country": "Australia",
			"alias": [],
			"regions": [],
			"timezone": "Australia/Adelaide",
			"unlocs": [
			  "AUADL"
			],
			"code": "60201"
		  }}`)}, map[string]Port{"INBLR": {
			Name:        "Bangalore",
			Coordinates: []float64{77.58, 12.98},
			City:        "Bangalore",
			Province:    "Karnataka",
			Country:     "India",
			Timezone:    "Asia/Calcutta",
			Unlocs:      []string{"INBLR"},
			Code:        "53398",
		}, "AUADL": {
			Name:        "Adelaide",
			Coordinates: []float64{138.58, -34.92},
			City:        "Adelaide",
			Province:    "South Australia",
			Country:     "Australia",
			Timezone:    "Australia/Adelaide",
			Unlocs:      []string{"AUADL"},
			Code:        "60201",
		},
		}},
		{"no ports", args{data: strings.NewReader("{}")}, map[string]Port{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPorts := LoadPorts(tt.args.data); !reflect.DeepEqual(gotPorts, tt.wantPorts) {
				t.Errorf("LoadPorts() = %v, want %v", gotPorts, tt.wantPorts)
			}
		})
	}
}
