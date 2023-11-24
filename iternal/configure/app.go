package configure

import (
	"os"
)

type application struct {
	env         *env
	fileWatcher iStart
	logger      iLogLevel
}

type iApplication interface {
	Start()
}

func App() iApplication {
	app := &application{}
	app.env = newEnv().configure()
	app.fileWatcher = newWatcher().configure(app.env)
	app.logger = newLogger().configure(app.env).Start()
	return app
}

func (app *application) Start() {
	fileWatcher := app.fileWatcher.Start()
	app.logger.Information("application started")
	defer func() {
		fileWatcher.Stop()
	}()
	done := make(chan bool)

	<-done

}

func (app *application) Stop() {
	app.logger.Information("application stoped")
	os.Exit(0)
}
