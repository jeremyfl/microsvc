package lib

import "os"

var Config map[string]string

func init() {
	Config = map[string]string{
		"APP_NAME":    os.Getenv("APP_NAME"),
		"DB_HOST":     os.Getenv("DB_HOST"),
		"JAEGER_HOST": os.Getenv("JAEGER_HOST"),
	}
}

