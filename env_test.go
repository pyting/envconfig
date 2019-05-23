package envconfig

import (
	"fmt"
	"os"
	"testing"
)

type Config struct {
	String string `env:"STRING" default:"abc"`
	Int    int    `env:"INT" default:"90"`
	Int8   int8   `env:"INT8" default:"91"`
	Bool   bool   `env:"BOOL" default:"true"`
}

func TestParse(t *testing.T) {
	v := Config{}
	err := os.Setenv("EX_STRING", "")
	if err != nil {
		t.Error(err)
	}
	err = os.Setenv("EX_INT", "12")
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
	fmt.Println(v)
}
