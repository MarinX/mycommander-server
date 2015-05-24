// main.go
package main

import (
	"config"
	"flag"
	"log"
	"service"
)

var path *string
var voice *string

func init() {

	path = flag.String("config", "", "Path for command configuration file")
	voice = flag.String("voice", "", "Path for voice configuration file, default is disabled")
	flag.Parse()
}

func main() {

	if path == nil || len(*path) == 0 {
		log.Println("Invalid configuration for path. Cmd configuration is mandatory")
		flag.Usage()
		return
	}

	cmdConf, err := getConfiguration(*path)
	if err != nil {
		log.Println("Error reading command config", err)
		return
	}

	ctrl := service.New(cmdConf)

	if voice != nil && len(*voice) > 0 {
		voiceConf, err := getConfiguration(*voice)
		if err != nil {
			log.Println("Error reading voice config", err)
			return
		}

		ctrl.UseVoice(voiceConf)
	}

	if err := ctrl.Run(); err != nil {
		log.Println(err)
	}
}

func getConfiguration(cfgPath string) (map[string]string, error) {
	cfg, err := config.New(cfgPath)
	if err != nil {
		return nil, err
	}
	return cfg.ReadConfig()
}
