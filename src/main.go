package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/micborens/es_cleaner/configuration"
	"github.com/micborens/es_cleaner/tools"
)

var cfg *configuration.Config

func main() {
	//	var index_pattern = "(.*)-(([0-9]{4}).([0-9]{2}).([0-9]{2}))"

	cfg = new(configuration.Config)
	if err := cfg.Init(); err != nil {
		log.Fatalf("Impossible to load configuration. err = %s", err)
	}

}
