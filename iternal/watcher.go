package configure

import (
	"fmt"
	"log"
	"smsparser/iternal/parse"
	"smsparser/iternal/utils"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type filteredChannels struct {
	filtered_channels map[time.Time]string
	lock              sync.Mutex
}

type fileEvent struct {
	f         fsnotify.Event
	timestamp time.Time
}
type fileEvents struct {
	fileEvents []fileEvent
	lock       sync.Mutex
}

func (fe *fileEvents) Push(x fileEvent) bool {
	fe.fileEvents = append(fe.fileEvents, x)
	return true
}

func (fe *fileEvents) Pop() fileEvent {
	h := fe.fileEvents
	var el fileEvent
	l := len(h)
	el, fe.fileEvents = h[0], h[1:l]
	return el
}

func NewFileEvents() *fileEvents {
	return &fileEvents{}
}

type iWatch interface {
	configure(env *env) iStart
}

type iStart interface {
	Start() iStop
}
type iStop interface {
	Stop() error
}
type watcher struct {
	watcher *fsnotify.Watcher
	env     *env
	err     error
}

func newWatcher() iWatch {
	f_watcher, err := fsnotify.NewWatcher()
	return &watcher{
		watcher: f_watcher,
		err:     err,
	}
}

func (w *watcher) configure(env *env) iStart {
	w.env = env
	w.watcher.Add(w.env.InputFolder)
	return w
}
func (w *watcher) Start() iStop {
	var fc = filteredChannels{
		filtered_channels: make(map[time.Time]string),
	}

	go func() {
		for {
			select {
			case event := <-w.watcher.Events:

				if fc.addToFilteredChannels(event) && utils.IsStringContainsInSlice(event.Op.String(), w.env.EventsToListen) &&
					utils.IsStringContainsInSlice(utils.GetFileExtensionFromPath(event.Name), w.env.FileExtensions) {

					log.Println("-addToFilteredChannels", fc.filtered_channels)
					log.Println("-------------------", event.Op.String())

					_, err := parse.CopyFileToOutputDirectory(event.Name, w.env.OutputFolder)
					if err != nil {
						log.Println("event err:", err)
					}
					log.Println("event:", event.Op.String())
					log.Println("on file:", event.Name)
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

func (fc *filteredChannels) addToFilteredChannels(event fsnotify.Event) bool {

	fc.lock.Lock()
	defer fc.lock.Unlock()
	timeKey := time.Now().Round(time.Second)
	if _, ok := fc.filtered_channels[timeKey]; !ok {
		fc.filtered_channels[timeKey] = event.Name

		fmt.Println("--->", fc.filtered_channels[timeKey])
		return true
	}
	return false
}
