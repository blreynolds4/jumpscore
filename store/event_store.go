package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"jumpscore/jumpscore"
	"os"
	"unicode"
)

func CreateEvent(event jumpscore.Event) error {
	eventFileName := getNameSlug(event.GetName())
	if _, err := os.Stat(eventFileName); os.IsNotExist(err) {
		// create the event directory
		if err := os.Mkdir(eventFileName, os.ModePerm); err != nil {
			return err
		}

		// create the event
		data, err := json.MarshalIndent(event, "", "  ")
		if err != nil {
			return err
		}

		eventFullPath := fmt.Sprintf("%s%c%s", eventFileName, os.PathSeparator, eventFileName)
		err = ioutil.WriteFile(eventFullPath, data, fs.FileMode(0660))
		if err != nil {
			return err
		}

		return nil
	}

	// don't overwrite an event, return an error
	return fmt.Errorf("event %s already exists", event.GetName())
}

func UpdateEvent(event jumpscore.Event) error {
	eventFileName := getNameSlug(event.GetName())
	if _, err := os.Stat(eventFileName); os.IsNotExist(err) {
		// not found, can't update
		return err
	}

	// get updated data
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return err
	}

	eventFullPath := fmt.Sprintf("%s%c%s", eventFileName, os.PathSeparator, eventFileName)
	return ioutil.WriteFile(eventFullPath, data, fs.FileMode(0660))
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

func DeleteEvent(event jumpscore.Event) error {
	var err error
	eventDirName := getNameSlug(event.GetName())
	if _, err = os.Stat(eventDirName); os.IsNotExist(err) {
		// not found
		return nil
	}

	// return other errors
	if err != nil {
		return err
	}

	return os.RemoveAll(eventDirName)
}

func GetEvents() ([]jumpscore.Event, error) {
	// create an events directory locally if it doesn't exist
	// get an event list from dirs (event slugs)
	// get the full name from event json in the dir
	// event json immediate
	result := make([]jumpscore.Event, 0, 5)
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fmt.Println("current file", file.Name())
		if file.IsDir() {
			fmt.Println(file.Name(), "is a dir")
			// load the event json from the event dir
			eventData, err := ioutil.ReadFile(fmt.Sprintf("%s%s%s", file.Name(), string(os.PathSeparator), file.Name()))
			if err != nil {
				return nil, err
			}
			fmt.Println("event data", string(eventData))
			var event jumpscore.Event
			err = json.Unmarshal(eventData, &event)
			fmt.Println("err", err)
			if err != nil {
				return nil, err
			}

			result = append(result, event)
		}
	}

	return result, nil
}

// func NewJsonEventStore(dir string) (EventStore, error) {
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		if err := os.MkdirAll(dir, os.ModeDir); err != nil {
// 			return nil, err
// 		}
// 	}

// 	result := &jsonEventStore{
// 		dataDir: dir,
// 		Jumps:   make(map[string]jumpscore.SkiJump),
// 	}

// 	content, err := ioutil.ReadFile(storeFileName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = json.Unmarshal(content, result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func CreateJump(sj jumpscore.SkiJump) error {
// 	jes.Jumps[sj.JumpName] = sj

// 	return jes.saveStore()
// }

// func GetJump(name string) (jumpscore.SkiJump, error) {
// 	sj, found := jes.Jumps[name]
// 	if !found {
// 		return sj, errors.New("jump not found for name " + name)
// 	}

// 	return sj, nil
// }

// func DeleteJump(sj jumpscore.SkiJump) error {
// }
