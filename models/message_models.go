package models

import()

type Message struct {
	From	  string  `json:"from, omitempty"`
	Message   string  `json:"Message, omitempty"`
}
