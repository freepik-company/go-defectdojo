package DefectDojoApi

import "time"

type Note struct {
	Id       int       `json:"id"`
	Author   Author    `json:"author"`
	Editor   Editor    `json:"editor"`
	History  []History `json:"history"`
	Entry    string    `json:"entry"`
	Date     time.Time `json:"date"`
	Private  bool      `json:"private"`
	Edited   bool      `json:"edited"`
	EditTime time.Time `json:"edit_time"`
	NoteType int       `json:"note_type"`
}
