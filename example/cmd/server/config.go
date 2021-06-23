package main

import (
	"bytes"

	pflag "github.com/spf13/pflag"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type flagContainer struct {
	StringFlag string
	BoolFlag   bool
	IntFlag    int
}

// ReadViperConfig initial read
func ReadViperConfig(rootYaml []byte, dst interface{}) error {
	var err error
	viper.SetConfigType("yaml")
	// Environment Variables override everything.
	viper.AutomaticEnv()
	fc := &flagContainer{}

	pflag.StringVarP(&fc.StringFlag, "Mode", "m", "1234", "help message for flagname")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	// 1. Read in as buffer to set a default baseline.
	err = viper.ReadConfig(bytes.NewBuffer(rootYaml))
	if err != nil {
		log.Err(err).Msg("ConfigDefaultYaml did not read in")
		return err
	}
	err = viper.Unmarshal(dst)
	return nil
}
