package bootstrap

type application struct {
	env         *env
	fileWatcher iStart
}

type iApplication interface {
	Start()
}

func App() iApplication {
	app := &application{}
	app.env = newEnv()
	app.fileWatcher = newWatcher().configurePath(app.env.InputFolder).configureEvents(app.env.EventsToListen)
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
