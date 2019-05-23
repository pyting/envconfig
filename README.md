# envconfig

go config with environment.

- setting defaults

- reading from environment variables

- setting environment variables prefix

## envconfig supports these struct field types:
- string

- bool

- int8, int16, int32, int64, int

- uint8, uint16, uint32, uint64, uint

- float32, float64

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/pyting/envconfig"
)

// environment prefix
const SERVICE = "app"

// k8s pod env
type Config struct{
	Port string `env:"PORT" default:"8080" json:"port"`
	MysqlUser string `env:"MYSQL_USER" default:"root" json:"mysql_user"`
	MysqlPWD string `env:"MYSQL_PWD" default:"root" json:"mysql_pwd"`
	MysqlHost string `env:"MYSQL_HOST" default:"localhost" json:"mysql_host"`
	MysqlPort string `env:"MYSQL_PORT" default:"3306" json:"mysql_port"`
}


func main() {
	
	os.Setenv("APP_MYSQL_USER","mysql_user")
	
	c := Config{}
	err := envconfig.Parse(SERVICE,&c)
	if err != nil {
		panic(err)
	}

	b,err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
```

print out

```json
{"port":"8080","mysql_user":"mysql_user","mysql_pwd":"root","mysql_host":"localhost","mysql_port":"3306"}
```

