package system

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

type DBConfig struct {
	Drivername string
	DataSource string
}

type SessionCfg struct {
	Salt   string
	Name   string
	MaxAge int
	Path   string
}

type AdminCfg struct {
	Email string
}

type Config struct {
	Downnet DownnetCfg
	Metric  MetricConfig
	Pprof   string
	DB      DBConfig
	Session SessionCfg
	Admin   AdminCfg
}

var QasConfig *Config = new(Config)

func LoadConfig(appcfgfile *string) error {

	if len(*appcfgfile) > 0 {
		data, err := ioutil.ReadFile(*appcfgfile)
		if err != nil {
			l4g.Error(err)
			data = []byte(*appcfgfile)
		}

		err = json.Unmarshal(data, QasConfig)
		if err != nil {
			return l4g.Error(err)
		}

	} else {
		err := l4g.Error("No CommandLine Param")
		return err
	}

	return nil
}
