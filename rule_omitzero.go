package sample

import "time"

// ExampleOmitZero demonstrates the "omitzero" rule.
// go fix replaces `json:",omitempty"` with `json:",omitzero"` for struct/time fields.
type ExampleOmitZero struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Nested    struct {
		Value int `json:"value"`
	} `json:"nested,omitempty"`
}
