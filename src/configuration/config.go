package configuration

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Elasticsearch struct {
		Url  string `yaml:"url"`
		Port string `yaml:"port"`
	} `yaml:"elasticsearch"`
	BaseRetentionDays    string   `yaml:"base_retention_days"`
	LongRetentionDays    string   `yaml:"long_retention_days"`
	LongRetentionIndices []string `yaml:"long_retention_indices"`
	Verbose              bool     `yaml:"verbose"`
	ConfigFile           string
}

func (c *Config) setDefaults() {
	if c.Elasticsearch.Url == ""{
		panic("elasticsearch url should be specified and is missing from configuration file")
	}

	if c.Elasticsearch.Port == ""{
		panic("elasticsearch port should be specified and is missing from configuration file")
	}

	if c.BaseRetentionDays == "" {
		log.Warnf("Loading default base retention days: 30 days")
	}

	if c.LongRetentionDays == "" {
		log.Warnf("Loading default long retention days: 90 days")
	}

	if len(c.LongRetentionIndices) == 0 {
		log.Warnf("Empty long retention indices list")
	}
}

func (c *Config) Init() (err error) {
	flag.StringVar(&c.ConfigFile, "config", "", "Configuration File")
	flag.BoolVar(&c.Verbose, "verbose", false, "Enable debug output.")
	flag.Parse()

	if c.ConfigFile != "" {
		_, err = os.Stat(c.ConfigFile)
		if os.IsNotExist(err) {
			return fmt.Errorf("non existent configuration file %v", c.ConfigFile)
		}
		source, _ := ioutil.ReadFile(c.ConfigFile)
		err = yaml.Unmarshal(source, &c)
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("configuration filename not specifed")
	}
	c.setDefaults()

	return err
}
