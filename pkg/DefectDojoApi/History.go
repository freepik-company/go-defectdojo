package DefectDojoApi

import "time"

type History struct {
	Id            int       `json:"id"`
	CurrentEditor Editor    `json:"current_editor"`
	Data          string    `json:"data"`
	Time          time.Time `json:"time"`
	NoteType      int       `json:"note_type"`
}
