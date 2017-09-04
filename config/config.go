package config

import (
	"encoding/json"
	l4g "github.com/alecthomas/log4go"
	"io/ioutil"
)

type DownnetCfg struct {
	HttpIP string
	Port   string
}

type MetricConfig struct {
	Addr string
}

type Config struct {
	Downnet DownnetCfg
	Metric  MetricConfig
	Pprof   string
}

var QasConfig *Config = new(Config)

func LoadConfig(appcfgfile *string) error {

	if len(*appcfgfile) > 0 {
		data, err := ioutil.ReadFile(*appcfgfile)
		if err != nil {
			data = []byte(*appcfgfile)
		}

		err = json.Unmarshal(data, QasConfig)
		if err != nil {
			return err
		}

	} else {
		err := l4g.Error("No CommandLine Param")
		return err
	}

	return nil
}
