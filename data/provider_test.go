// provider_test.go - test sample_provider.go
package data

import (
	"reflect"
	"testing"
)

func TestCRUDProvider(t *testing.T) {
	_title := "Test Title"
	_author := "Test Author"
	_contents := "Test Contents"

	note := NewNote(_title, _author, _contents)

	// Initialize the SampleProvider for testing
	sampleProvider := NewSampleProvider()

	// Test SaveNote
	err := sampleProvider.SaveNote(note)
	if err != nil {
		t.Errorf("SaveNote failed: %v", err)
	}

	// Test ListNotes
	notes, err := sampleProvider.ListNotes("")
	if err != nil {
		t.Errorf("ListNotes failed: %v", err)
	}
	if len(notes) != 1 {
		t.Errorf("Expected 1 note, got %d", len(notes))
	}

	// Test LoadNote
	loadedNote, err := sampleProvider.LoadNote(*note.ID)
	if err != nil {
		t.Errorf("LoadNote failed: %v", err)
	}
	if !reflect.DeepEqual(note, loadedNote) {
		t.Errorf("Loaded note does not match the original note")
	}

	// TODO: additional tests
}
