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

func (a *application) Start() {
	fileWatcher := a.fileWatcher.Start()
	a.logger.Information("application started")
	defer func() {
		fileWatcher.Stop()
	}()
	done := make(chan bool)

	<-done

}

func (a *application) Stop() {
	a.logger.Information("application stoped")
	os.Exit(0)
}
