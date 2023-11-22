package configure

import (
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
	return &env
}

func (env *env) configure() *env {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}
	viper.OnConfigChange(func(event fsnotify.Event) {
		if event.Has(fsnotify.Write) {
			viper.Unmarshal(&env)
		}
	})
	viper.WatchConfig()
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}
	return env
}
