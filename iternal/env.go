package configure

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type env struct {
	AppEnv         string   `mapstructure:"APP_ENV"`
	InputFolder    string   `mapstructure:"Input_Folder"`
	OutputFolder   string   `mapstructure:"Output_Folder"`
	FileExtensions []string `mapstructure:"File_Extensions"`
	EventsToListen []string `mapstructure:"Events_To_Listen"`
}

func newEnv() *env {
	env := env{}
	env.loadConfiguration()
	return &env
}

func (env *env) loadConfiguration() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		viper.Unmarshal(&env)
		fmt.Println("env.FileExtensions:", env.EventsToListen)
	})

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}
}
