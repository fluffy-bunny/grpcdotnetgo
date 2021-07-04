package main

import (
	"bytes"

	"github.com/fluffy-bunny/viperEx"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// ReadViperConfig initial read
func ReadViperConfig(rootYaml []byte, dst interface{}) error {
	var err error
	viper.SetConfigType("yaml")
	// Environment Variables override everything.
	viper.AutomaticEnv()

	// 1. Read in as buffer to set a default baseline.
	err = viper.ReadConfig(bytes.NewBuffer(rootYaml))
	if err != nil {
		log.Err(err).Msg("ConfigDefaultYaml did not read in")
		return err
	}
	allSettings := viper.AllSettings() // normal viper stuff
	myViperEx, err := viperEx.New(allSettings, func(ve *viperEx.ViperEx) error {
		ve.KeyDelimiter = "__"
		return nil
	})
	if err != nil {
		return err
	}
	myViperEx.UpdateFromEnv()
	err = myViperEx.Unmarshal(dst)
	if err != nil {
		return err
	}
	//err = viper.Unmarshal(dst)
	return nil
}
