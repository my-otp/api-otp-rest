package main

import (
	"github.com/my-otp/otp-rest/internal/config"
	"github.com/restartfu/gophig"
)

var yaml = gophig.YAMLMarshaler{}

func main() {
	conf, err := gophig.LoadConf[config.Config]("", yaml)
	if err != nil {
		panic(err)
	}

	_ = conf
}
