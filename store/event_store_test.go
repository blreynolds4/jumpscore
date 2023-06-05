package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jumpscore/jumpscore"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	e := jumpscore.NewEvent("Test Event!", time.Now().UTC())

	err := CreateEvent(e)
	assert.NoError(t, err)

	slug := getNameSlug(e.GetName())
	fullPath := fmt.Sprintf("%s%c%s", slug, os.PathSeparator, slug)
	actualData, err := ioutil.ReadFile(fullPath)
	assert.NoError(t, err)

	// verify the output
	expectedData, err := json.MarshalIndent(e, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, expectedData, actualData)

	// then delete it
	assert.NoError(t, os.RemoveAll(slug))
}

func TestCreateDuplicateEvent(t *testing.T) {
	e := jumpscore.NewEvent("Test Event!", time.Now().UTC())
	err := CreateEvent(e)
	assert.NoError(t, err)

	// create again and expect failure
	err = CreateEvent(e)
	assert.Error(t, err)

	// then delete it
	assert.NoError(t, os.RemoveAll(getNameSlug(e.GetName())))
}

func TestDeleteNothing(t *testing.T) {
	e := jumpscore.NewEvent("Test Event!", time.Now().UTC())
	assert.NoError(t, DeleteEvent(e))
}

func TestDeleteEvent(t *testing.T) {
	e := jumpscore.NewEvent("Test Event!", time.Now().UTC())

	err := CreateEvent(e)
	assert.NoError(t, err)

	assert.NoError(t, DeleteEvent(e))
	eventDir := getNameSlug(e.GetName())
	_, err = os.Stat(eventDir)
	assert.Error(t, err)
	assert.True(t, os.IsNotExist(err))
}

func TestListNoEvents(t *testing.T) {
	events, err := GetEvents()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(events))
}

func TestListEvents(t *testing.T) {
	event1 := jumpscore.NewEvent("Test Event1", time.Now().UTC())
	err := CreateEvent(event1)
	assert.NoError(t, err)
	event2 := jumpscore.NewEvent("Test Event2", time.Now().UTC())
	err = CreateEvent(event2)
	assert.NoError(t, err)

	// cleanup
	defer os.RemoveAll(getNameSlug(event1.GetName()))
	defer os.RemoveAll(getNameSlug(event2.GetName()))

	events, err := GetEvents()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(events))

	assert.Equal(t, event1, events[0])
	assert.Equal(t, event2, events[1])
}
