package envconfig

import (
	"os"
	"testing"
)

type Config struct {
	String  string  `env:"STRING" default:"abc"`
	Int     int     `env:"INT" default:"90"`
	Bool    bool    `env:"BOOL" default:"true"`
	Float32 float32 `env:"FLOAT32" default:"9.9"`
}

func TestParse(t *testing.T) {
	v := Config{}
	err := os.Setenv("EX_STRING", "")
	if err != nil {
		t.Error(err)
	}
	err = os.Setenv("EX_INT", "")
	if err != nil {
		t.Error(err)
	}
	err = os.Setenv("EX_BOOL", "false")
	if err != nil {
		t.Error(err)
	}
	err = Parse("EX", &v)
	if err != nil {
		t.Error(err)
	}

	if v.String != "abc" {
		t.Error("parse string error")
	}

	if v.Bool {
		t.Error("parse bool error")
	}

	if v.Int != 90 {
		t.Error("parse int error")
	}

	if v.Float32 != 9.9 {
		t.Error("parse float32 error")
	}
}
