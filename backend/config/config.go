package config

import (
	"github.com/charmbracelet/log"
	"github.com/invopop/validation"
	"github.com/spf13/viper"
)

type App struct {
	Name        string
	Port        int
	Environment string
}

type Database struct {
	Host      string
	Port      int
	User      string
	Pass      string
	Schema    string
	Charset   string
	Collation string
}

type Config struct {
	App
	Database
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("config file not found")
		} else {
			log.Fatal(err.Error())
		}
	}

	app := App{
		Name:        viper.GetString("app.name"),
		Port:        viper.GetInt("app.port"),
		Environment: viper.GetString("app.env"),
	}

	if err := app.Validate(); err != nil {
		log.Fatal(err)
	}

	database := Database{
		Host:      viper.GetString("database.host"),
		Port:      viper.GetInt("database.port"),
		User:      viper.GetString("database.user"),
		Pass:      viper.GetString("database.pass"),
		Schema:    viper.GetString("database.schema"),
		Charset:   viper.GetString("database.charset"),
		Collation: viper.GetString("database.collation"),
	}

	if err := database.Validate(); err != nil {
		log.Fatal(err)
	}

	config := Config{
		App:      app,
		Database: database,
	}

	return &config
}

func (a App) Validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Port, validation.Required),
		validation.Field(&a.Environment, validation.Required),
	)
}

func (d Database) Validate() error {
	return validation.ValidateStruct(
		&d,
		validation.Field(&d.Host, validation.Required),
		validation.Field(&d.Port, validation.Required),
		validation.Field(&d.User, validation.Required),
		validation.Field(&d.Pass, validation.Required),
		validation.Field(&d.Schema, validation.Required),
		validation.Field(&d.Charset, validation.Required),
		validation.Field(&d.Collation, validation.Required),
	)
}
