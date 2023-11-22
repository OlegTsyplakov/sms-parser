package configure

import (
	"os"
)

type application struct {
	env         *env
	fileWatcher iStart
}

type iApplication interface {
	Start()
}

func App() iApplication {
	app := &application{}
	app.env = newEnv().configure()
	app.fileWatcher = newWatcher().configure(app.env)
	return app
}

func (a *application) Start() {
	fileWatcher := a.fileWatcher.Start()
	defer func() {
		fileWatcher.Stop()
	}()
	done := make(chan bool)

	<-done

}

func (a *application) Stop() {
	os.Exit(0)
}
