package serve

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	grpcdotnetgocore "github.com/fluffy-bunny/grpcdotnetgo/core"

	"github.com/spf13/cobra"
)

type EnvType int

const (
	String EnvType = iota // String = 0
	Int                   // Int = 1
)

type EnvDefinition struct {
	Type  EnvType
	Name  string
	Regex string
}
type env struct {
	Name  string
	Value string
}

var wellKnownEnvDefinitions = []EnvDefinition{
	{
		Type: String,
		Name: "APPLICATION_ENVIRONMENT",
	},
	{
		Type: Int,
		Name: "GRPC_PORT",
	},
}

func setAdditionalEnvDefs(additional ...EnvDefinition) {
	wellKnownEnvDefinitions = append(wellKnownEnvDefinitions, additional...)
}

func validateOsEnvs() error {
	for _, e := range wellKnownEnvDefinitions {
		env, ok := os.LookupEnv(e.Name)
		if !ok {
			return fmt.Errorf("ENV:%v is not present", e.Name)
		}
		switch e.Type {
		case Int:
			_, err := strconv.ParseInt(env, 0, 64)
			if err != nil {
				return fmt.Errorf("ENV:%v=%v must be an int", e, env)
			}
		case String:
			if len(env) == 0 {
				return fmt.Errorf("ENV:%v is nil or empty", e)
			}
		}
	}
	return nil
}
func validateEnvString(envS string) (*env, error) {
	idx := strings.Index(envS, "=")
	if idx <= 0 {
		return nil, fmt.Errorf("invalid env:[%v]", envS)
	}
	return &env{
		Name:  envS[:idx],
		Value: envS[idx+1:],
	}, nil
}
func validateEnvs() ([]*env, error) {
	var result []*env
	for _, env := range envsets {
		env, err := validateEnvString(env)
		if err != nil {
			return nil, err
		}
		result = append(result, env)
	}
	return result, nil
}

var command = &cobra.Command{
	Use:   "serve",
	Short: "run the grpc server",
	Long:  `run the grpc server`,

	PreRunE: func(cmd *cobra.Command, args []string) error {

		var err error
		envs, err = validateEnvs()
		if err != nil {
			return err
		}
		printEnvs()
		for _, env := range envs {
			os.Setenv(env.Name, env.Value)
		}
		err = validateOsEnvs()
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		//		env := os.Getenv("GRPC_PORT")
		//		port, _ := strconv.ParseInt(env, 0, 64)

		//	appStartup.SetPort(int(port))
		grpcdotnetgocore.Start()

	},
}

var envsets []string
var envs []*env

func printEnvs() {
	for _, env := range envs {
		fmt.Printf("ENV: %v='%v'\n", env.Name, env.Value)
	}
	/*
		for _, envSet := range os.Environ() {
			fmt.Println(envSet)
		}
	*/
}

func Init(rootCmd *cobra.Command, additional ...EnvDefinition) {
	setAdditionalEnvDefs(additional...)
	command.Flags().StringArrayVarP(&envsets, "env", "e", []string{}, "override any env variable")
	rootCmd.AddCommand(command)
}
