package data

import (
	"time"

	"github.com/beevik/guid"
)

// GetSampleNotes will return sample notes
func GetSampleNotes() []Note {
	_now := time.Now().UTC()
	notes := []Note{
		{Title: "firstTitle", Contents: "# First Note!\nTest\n- Foo\n- Bar", Author: "Seb",
			LastUpdated: _now, Version: 1, Active: true, ID: guid.NewString()},
		{Title: "secondTitle", Contents: "# Second Note!\nTest\n- one\n- two", Author: "Seb",
			LastUpdated: _now, Version: 1, Active: true, ID: guid.NewString()},
		{Title: "thirdTitle", Contents: "# Third Note!\nTest\n- Nest\n\t- it", Author: "Seb",
			LastUpdated: _now, Version: 1, Active: true, ID: guid.NewString()},
	}

	return notes
}
