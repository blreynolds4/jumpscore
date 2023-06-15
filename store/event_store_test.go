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

func TestCreateFileStore(t *testing.T) {
	_, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	info, err := os.Stat(t.Name())
	assert.NoError(t, err)
	assert.True(t, info.IsDir())

	// then delete the store
	assert.NoError(t, os.RemoveAll(t.Name()))
}

func TestCreateFileStoreBadRoot(t *testing.T) {
	_, err := NewFileEventStore(" ")
	assert.Error(t, err)
}

func TestCreateEvent(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	e := jumpscore.NewEvent("Test Event!", time.Now().UTC())

	err = store.CreateEvent(e)
	assert.NoError(t, err)

	slug := getNameSlug(e.Name)
	fullPath := fmt.Sprintf("%s%c%s%c%s", t.Name(), os.PathSeparator, slug, os.PathSeparator, slug)
	actualData, err := ioutil.ReadFile(fullPath)
	assert.NoError(t, err)

	// verify the output
	expectedData, err := json.MarshalIndent(e, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, expectedData, actualData)

	// then delete the store
	assert.NoError(t, os.RemoveAll(t.Name()))
}

func TestCreateDuplicateEvent(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	e := jumpscore.NewEvent("Test Event!", time.Now().UTC())
	err = store.CreateEvent(e)
	assert.NoError(t, err)

	// create again and expect failure
	err = store.CreateEvent(e)
	assert.Error(t, err)

	// then delete the store
	assert.NoError(t, os.RemoveAll(t.Name()))
}

func TestDeleteNothing(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)
	// then delete the store
	defer os.RemoveAll(t.Name())

	e := jumpscore.NewEvent("Test Event!", time.Now().UTC())

	assert.NoError(t, store.DeleteEvent(e))
}

func TestDeleteEvent(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	// cleanup
	defer os.RemoveAll(t.Name())

	e := jumpscore.NewEvent("Test Event!", time.Now().UTC())

	err = store.CreateEvent(e)
	assert.NoError(t, err)

	assert.NoError(t, store.DeleteEvent(e))

	slug := getNameSlug(e.Name)
	eventDir := fmt.Sprintf("%s%c%s", t.Name(), os.PathSeparator, slug)
	_, err = os.Stat(eventDir)
	assert.Error(t, err)
	assert.True(t, os.IsNotExist(err))

	// then delete the store
	assert.NoError(t, os.RemoveAll(t.Name()))
}

func TestListNoEvents(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	// cleanup
	defer os.RemoveAll(t.Name())

	events, err := store.GetEvents()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(events))
}

func TestListEvents(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	// cleanup
	defer os.RemoveAll(t.Name())

	event1 := jumpscore.NewEvent("Test Event1", time.Now().UTC())
	err = store.CreateEvent(event1)
	assert.NoError(t, err)
	event2 := jumpscore.NewEvent("Test Event2", time.Now().UTC())
	err = store.CreateEvent(event2)
	assert.NoError(t, err)

	events, err := store.GetEvents()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(events))

	assert.Equal(t, event1, events[0])
	assert.Equal(t, event2, events[1])
}

func TestGetMissingEvent(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	// cleanup
	defer os.RemoveAll(t.Name())

	event := jumpscore.NewEvent("not found", time.Now().UTC())
	_, err = store.GetEvent(event.Name)
	assert.Error(t, err)
	assert.Equal(t, "event not found: not found", err.Error())
}

func TestGetEvent(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	// cleanup
	defer os.RemoveAll(t.Name())

	event := jumpscore.NewEvent("get me", time.Now().UTC())
	assert.NoError(t, store.CreateEvent(event))

	gotEvent, err := store.GetEvent(event.Name)
	assert.NoError(t, err)
	assert.Equal(t, event, gotEvent)
}

func TestUpdateMissingEvent(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	// cleanup
	defer os.RemoveAll(t.Name())

	event := jumpscore.NewEvent("not found", time.Now().UTC())
	err = store.UpdateEvent(event)
	assert.Error(t, err)
	assert.Equal(t, "event not found: not found", err.Error())
}

func TestUpdateEvent(t *testing.T) {
	store, err := NewFileEventStore(t.Name())
	assert.NoError(t, err)

	// cleanup
	defer os.RemoveAll(t.Name())

	event := jumpscore.NewEvent("update me", time.Now().UTC())
	assert.NoError(t, store.CreateEvent(event))

	newDate := event.Date.Add(time.Hour)
	event.Date = newDate
	err = store.UpdateEvent(event)
	assert.NoError(t, err)
	updatedEvent, err := store.GetEvent(event.Name)
	assert.NoError(t, err)

	assert.Equal(t, newDate, updatedEvent.Date)
}
