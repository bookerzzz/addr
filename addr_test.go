package addr

import (
	"testing"
)

func TestMake(t *testing.T) {
	tests := map[string]string{
		"test.domain.com:1234": "test.domain.com:1234",
		"test.domain.com:0":    "test.domain.com",
		":1234":                ":1234",
		"":                     "",
		"iama.domain.com":      "iama.domain.com",
	}

	for input, expected := range tests {
		a, err := Make(input)
		if err != nil {
			t.Error(err.Error())
		}
		r := a.String()
		if r != expected {
			t.Errorf("Expected '%s' to match '%s'", r, expected)
		}
	}

	defaultAddr := "default:8080"
	testsDefault := map[string]string{
		"test.domain.com:1234": "test.domain.com:1234",
		"test.domain.com:0":    "test.domain.com",
		":1234":                ":1234",
		"":                     ":8080",
		"iama.domain.com":      "iama.domain.com:8080",
	}

	for input, expected := range testsDefault {
		a, err := Make(defaultAddr)
		if err != nil {
			t.Error(err.Error())
		}
		err = a.Parse(input)
		if err != nil {
			t.Error(err.Error())
		}
		r := a.String()
		if r != expected {
			t.Errorf("Expected '%s' to match '%s'", r, expected)
		}
	}

	_, err := Make("domain:port")
	if err == nil {
		t.Error("Expected error for invalid port 'port'")
	}
}
