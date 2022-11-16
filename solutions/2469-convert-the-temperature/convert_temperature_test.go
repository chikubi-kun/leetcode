package convertthetemperature

import (
	"reflect"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	type arg struct {
		celcius  float64
		expected []float64
	}

	testCase := []arg{
		{celcius: 36.50, expected: []float64{309.65000, 97.70000}},
		{celcius: 122.11, expected: []float64{395.26000, 251.79800}},
	}

	for _, data := range testCase {
		if res := convertTemperature(data.celcius); !reflect.DeepEqual(res, data.expected) {
			t.Errorf("expected %v, got %v", data.expected, res)
		}
	}
}
