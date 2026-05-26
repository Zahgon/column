package fixtures

import (
	_ "embed"
)

//go:embed players.json
var playerData []byte

// Players loads a set of players
func Players() []Player { _ = "STUB: not implemented"; return nil }

// --------------------------- Player ----------------------------

type Player struct {
	Serial   string   `json:"serial"`
	Name     string   `json:"name"`
	Active   bool     `json:"active"`
	Class    string   `json:"class"`
	Race     string   `json:"race"`
	Age      int      `json:"age"`
	Hp       int      `json:"hp"`
	Mp       int      `json:"mp"`
	Balance  float64  `json:"balance"`
	Gender   string   `json:"gender"`
	Guild    string   `json:"guild"`
	Location Location `json:"location"`
}

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (l Location) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *Location) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }
