package bootstrap

import (
	"log"
	"smsparser/iternal/utils"

	"github.com/fsnotify/fsnotify"
)

type iWatch interface {
	configurePath(path string) iEvents
}
type iEvents interface {
	configureEvents(events []string) iStart
}
type iStart interface {
	Start() iStop
}
type iStop interface {
	Stop() error
}
type watcher struct {
	watcher *fsnotify.Watcher
	events  []string
	err     error
}

func newWatcher() iWatch {
	f_watcher, err := fsnotify.NewWatcher()

	return &watcher{
		watcher: f_watcher,
		err:     err,
	}
}

func (w *watcher) configurePath(path string) iEvents {
	w.err = w.watcher.Add(path)
	return w
}
func (w *watcher) configureEvents(events []string) iStart {
	w.events = events
	return w
}

func (w *watcher) Start() iStop {
	go func() {
		for {
			select {
			case event := <-w.watcher.Events:
				if utils.IsStringContainsInSlice(event.Op.String(), w.events) {
					log.Println("event:", event.Op.String())
				}
			case err := <-w.watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	return w
}
func (w *watcher) Stop() error {
	return w.watcher.Close()
}
