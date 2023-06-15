package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"jumpscore/jumpscore"
	"os"
	"strings"
	"unicode"
)

type EventStore interface {
	CreateEvent(event jumpscore.Event) error
	GetEvent(name string) (jumpscore.Event, error)
	UpdateEvent(event jumpscore.Event) error
	DeleteEvent(event jumpscore.Event) error
	GetEvents() ([]jumpscore.Event, error)
}

type fileEventStore struct {
	Root string
}

func NewFileEventStore(root string) (EventStore, error) {
	rootDir := strings.TrimSpace(root)
	if len(rootDir) == 0 {
		return nil, fmt.Errorf("event store rood directory is required")
	}

	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &fileEventStore{
		Root: rootDir,
	}, nil
}

func getNameSlug(name string) string {
	// remove any chars that are not leters or numbers
	var b bytes.Buffer
	for _, c := range name {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			b.WriteRune(c)
		}
	}

	return b.String()
}

func (fec *fileEventStore) notFoundError(name string) error {
	return fmt.Errorf("event not found: %s", name)
}

func (fes *fileEventStore) getEventDir(slug string) string {
	return fmt.Sprintf("%s%c%s", fes.Root, os.PathSeparator, slug)
}

func (fes *fileEventStore) getEventPath(name string) string {
	eventDir := fes.getEventDir(getNameSlug(name))
	return fmt.Sprintf("%s%c%s", eventDir, os.PathSeparator, getNameSlug(name))
}

func (fes *fileEventStore) eventExists(name string) bool {
	eventFullPath := fes.getEventPath(name)
	if _, err := os.Stat(eventFullPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func (fes *fileEventStore) CreateEvent(event jumpscore.Event) error {
	eventDir := fes.getEventDir(getNameSlug(event.Name))
	if _, err := os.Stat(eventDir); os.IsNotExist(err) {
		// create the event directory
		if err := os.Mkdir(eventDir, os.ModePerm); err != nil {
			return err
		}

		// create the event
		data, err := json.MarshalIndent(event, "", "  ")
		if err != nil {
			return err
		}

		eventFullPath := fes.getEventPath(event.Name)
		err = ioutil.WriteFile(eventFullPath, data, fs.FileMode(0660))
		if err != nil {
			return err
		}

		return nil
	}

	// don't overwrite an event, return an error
	return fmt.Errorf("event %s already exists", event.Name)
}

func (fes *fileEventStore) GetEvent(name string) (jumpscore.Event, error) {
	if !fes.eventExists(name) {
		return jumpscore.Event{}, fes.notFoundError(name)
	}

	var result jumpscore.Event
	eventPath := fes.getEventPath(name)
	eventData, err := ioutil.ReadFile(eventPath)
	if err != nil {
		return jumpscore.Event{}, err
	}

	err = json.Unmarshal(eventData, &result)
	if err != nil {
		return jumpscore.Event{}, err
	}

	return result, nil
}

func (fes *fileEventStore) UpdateEvent(event jumpscore.Event) error {
	if !fes.eventExists(event.Name) {
		return fes.notFoundError(event.Name)
	}

	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return err
	}

	eventFullPath := fes.getEventPath(event.Name)
	return ioutil.WriteFile(eventFullPath, data, fs.FileMode(0660))
}

func (fes *fileEventStore) DeleteEvent(event jumpscore.Event) error {
	var err error
	eventDir := fes.getEventDir(getNameSlug(event.Name))
	if _, err = os.Stat(eventDir); os.IsNotExist(err) {
		// not found
		return nil
	}

	// return other errors
	if err != nil {
		return err
	}

	return os.RemoveAll(eventDir)
}

func (fes *fileEventStore) GetEvents() ([]jumpscore.Event, error) {
	// create an events directory locally if it doesn't exist
	// get an event list from dirs (event slugs)
	// get the full name from event json in the dir
	// event json immediate
	result := make([]jumpscore.Event, 0, 5)
	files, err := ioutil.ReadDir(fes.Root)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		// events are in directories, skip over files
		if file.IsDir() {
			// load the event json from the event dir
			eventData, err := ioutil.ReadFile(fmt.Sprintf("%s%s%s", fes.getEventDir(file.Name()), string(os.PathSeparator), file.Name()))
			if err != nil {
				return nil, err
			}

			var event jumpscore.Event
			err = json.Unmarshal(eventData, &event)
			if err != nil {
				return nil, err
			}

			result = append(result, event)
		}
	}

	return result, nil
}
