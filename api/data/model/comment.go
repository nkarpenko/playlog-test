package model

import "time"

type Comment struct {
	ID      int64     `json:"id"`
	Comment string    `json:"comment"`
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
}
