package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/snowflake"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/aws_s3"
	_ "github.com/golang-migrate/migrate/v4/source/bitbucket"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/golang-migrate/migrate/v4/source/gitlab"
	_ "github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/golang-migrate/migrate/v4/source/google_cloud_storage"
	_ "github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const relativePath = "./dbmigrate"

func GetDbMigratePath() string {
	var configPath string
	_, err := os.Stat(relativePath)
	if !os.IsNotExist(err) {
		configPath, _ = filepath.Abs(relativePath)
		log.Info().Str("path", configPath).Msg("Configuration Root Folder")
	}
	return configPath
}

var Source string
var StepN uint64
var Verbose bool

func removeDuplicateValues(slice []string) []string {
	keys := make(map[string]bool)
	var list []string

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func fixupMongoConnectionString(connectionString, databaseName string) string {
	/*
				The regex `\w+\:\d+\/\?` will find a "word" followed by a ":" followed by a port #
				so if your string contains something like "mapped-dev-shard-00-02-pri.awzwl.mongodb.net:27017/?ssl=true&authSource=admin&replicaSet=atlas-uqxcpd-shard-0", it will find "net:27017/?"

				net:27017 is then replaced with net:27017/{{DATABASE_NAME}}
		                   "mapped-dev-shard-00-02-pri.awzwl.mongodb.net:27017/mydatabase/?ssl=true&authSource=admin&replicaSet=atlas-uqxcpd-shard-0"
				very specificy to MongoDB

	*/
	re := regexp.MustCompile(`\w+\:\d+\/\?`)
	foundStrings := re.FindAllString(connectionString, -1)
	foundStrings = removeDuplicateValues(foundStrings)
	if len(foundStrings) > 0 {
		log.Info().Interface("foundStrings", foundStrings).Send()

		for _, v := range foundStrings {
			vv := strings.Split(v, "/")
			newS := fmt.Sprintf("%s/%s?", vv[0], databaseName)
			connectionString = strings.ReplaceAll(connectionString, v, newS)
		}
	} else {
		re := regexp.MustCompile(`\w+\:\d+\?`)
		foundStrings := re.FindAllString(connectionString, -1)
		foundStrings = removeDuplicateValues(foundStrings)
		if len(foundStrings) > 0 {
			log.Info().Interface("foundStrings", foundStrings).Send()

			for _, v := range foundStrings {
				vv := strings.Split(v, "?")
				newS := fmt.Sprintf("%s/%s?", vv[0], databaseName)
				connectionString = strings.ReplaceAll(connectionString, v, newS)
			}
		} else {
			re := regexp.MustCompile(`\w+\:\d+`)
			foundStrings := re.FindAllString(connectionString, -1)
			foundStrings = removeDuplicateValues(foundStrings)
			if len(foundStrings) > 0 {
				log.Info().Interface("foundStrings", foundStrings).Send()

				for _, v := range foundStrings {
					newS := fmt.Sprintf("%s/%s", v, databaseName)
					connectionString = strings.ReplaceAll(connectionString, v, newS)
				}
			}
		}
	}
	return connectionString
}

// upPersistentPreRunE validateds required flags
func UpPersistentPreRunE(cmd *cobra.Command, args []string) error {
	log.Info().Msg("Validating input")
	var err error
	connectionString := viper.GetString("database")
	if len(connectionString) == 0 {
		err = fmt.Errorf("--database missing and env:DATABASE not present")
		log.Error().Err(err).Send()
		return err
	}
	log.Info().Msg("Found --database")
	databaseName := viper.GetString("database-name")
	if len(databaseName) == 0 {
		log.Info().Msg("no --database-name present")
	} else {
		if strings.Contains(connectionString, "mongodb://") {
			log.Info().Msg("Found --database == mongodb")
			connectionString = fixupMongoConnectionString(connectionString, databaseName)
			viper.Set("database", connectionString)
		}
	}
	verbose := viper.GetBool("verbose")
	if verbose {
		log.Info().Msg("Found --verbose")
		Verbose = verbose
	}
	connectionString = viper.GetString("database")

	Source = viper.GetString("source")
	if len(Source) == 0 {
		// check to see if we have an embedded dbmigrate folder
		dbMigratePath := GetDbMigratePath()
		_, err := ioutil.ReadDir(dbMigratePath)
		if err != nil {
			log.Error().Err(err).Msg("no source and no embedded dbmigrate path")
			return err
		}
		log.Info().Str("dbMigratePath", dbMigratePath).Msg("Found local dbmigration scripts")

	}
	step := viper.GetString("step")

	if len(step) > 0 {
		log.Info().Str("step", step).Msg("Found --step")

		StepN, err = strconv.ParseUint(step, 10, 64)
		if err != nil {
			log.Error().Err(err).Str("step", step).
				Msg("Invalid step argument")
			return err
		}
	}

	return nil

}

type MigrateDirection int

const (
	MigrateUp MigrateDirection = iota
	MigrateDown
	MigrateForce
)

func Migrate(md MigrateDirection) error {
	log.Info().Interface("MigrationDirection", md).Msg("Migrating.....")
	var err error
	//	verbose := viper.GetBool("verbose")
	connectionString := viper.GetString("database")
	//	fmt.Println("inside up, verbose value: ", connectionString)

	var m *migrate.Migrate
	if len(Source) > 0 {
		log.Info().Msg("Migrating from remote source")
		m, err = migrate.New(Source, connectionString)
	} else {
		log.Info().Str("Source", "file://dbmigrate").Msg("Migrating from local source")
		m, err = migrate.New(
			"file://dbmigrate",
			connectionString)
		if err == nil && Verbose {
			// read all the files in the dbmigrate folder and print them out
			dbMigratePath := GetDbMigratePath()
			files, err := ioutil.ReadDir(dbMigratePath)
			if err != nil {
				log.Error().Err(err).Msg("Failed to read dbmigrate folder")
				return err
			}
			fmt.Println("===============================================")
			fmt.Println("Contents of dbmigrate folder:")
			fmt.Println(dbMigratePath)
			fmt.Println("===============================================")

			for _, f := range files {

				// these are all text files, read them and print them out pretty like
				filePath := fmt.Sprintf("%s/%s", dbMigratePath, f.Name())
				// read the entire file into memory
				data, err := ioutil.ReadFile(filePath)
				if err != nil {
					fmt.Println("File reading error", err)
					return err
				}
				fmt.Println(">----------------------------------------------")
				fmt.Println(f.Name())
				fmt.Println(">----------------------------------------------")

				fmt.Println(string(data))
			}
			fmt.Println("===============================================")
		}
	}
	if err != nil {
		log.Error().Err(err).Msg("Failed to migrate.New")
		return err
	}

	if StepN == 0 {
		switch md {
		case MigrateDown:
			err = m.Down()
		case MigrateUp:
			err = m.Up()
		case MigrateForce:
			version := viper.GetInt("version")
			err = m.Force(version)
		default:
			err = fmt.Errorf("no valid migrate direction")
		}

	} else {
		var ss int = int(StepN)
		if md == MigrateDown {
			ss = 0 - ss
		}
		err = m.Steps(ss)
	}
	if err != nil && err != migrate.ErrNoChange {
		log.Error().Err(err).Msg("Failed to migrate")
	}

	if err == nil || err == migrate.ErrNoChange {
		err = nil
		sleepMinutes := viper.GetInt("sleep")
		log.Info().Str("sleep", fmt.Sprintf("%v", sleepMinutes)).Msg("Success")
		var sleepDuration time.Duration
		if sleepMinutes != 0 {
			if sleepMinutes < 0 {
				sleepMinutes = 60 * 24 * 365
			}
			sleepDuration = time.Minute * time.Duration(sleepMinutes)
			time.Sleep(sleepDuration)
		}
	}
	return err
}
