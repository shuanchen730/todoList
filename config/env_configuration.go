package config

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var EnvConfig envConfiguration

type envConfiguration struct {
	Env               string            `env:"env" validate:"required"`
	Port              string            `env:"port" validate:"required"`
	ConnectionStrings ConnectionStrings `env:"connectionstrings" validate:"required"`
}

type App struct {
	Port string `env:"port"`
}

type ConnectionStrings struct {
	Mysql MysqlConfig `env:"Mysql"`
}

type MysqlConfig struct {
	Host     string `env:"host" validate:"required"`
	Database string `env:"database" validate:"required"`
	User     string `env:"user" validate:"required"`
	Password string `env:"password" validate:"required"`
	Port     string `env:"port"`
}

func InitialEnvConfiguration() (err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	bindEnvs(EnvConfig)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		switch err.(type) {
		default:
			panic(fmt.Errorf("Fatal error loading config file: %s \n", err))
		case viper.ConfigFileNotFoundError:
			log.Print("No config file found. Using defaults and environment variables")
		}
	}
	err = viper.Unmarshal(&EnvConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	validate := validator.New()
	err = validate.Struct(&EnvConfig)
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	log.Println("EnvConfig:", EnvConfig)

	return
}

func bindEnvs(iFace interface{}, parts ...string) {
	ifv := reflect.ValueOf(iFace)
	ift := reflect.TypeOf(iFace)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("env")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			_ = viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
