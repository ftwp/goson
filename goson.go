package goson

import (
	"encoding/json"
	"fmt"
)

type goson struct {
	v map[string]json.RawMessage
}

func emptyGOSON() *goson {
	return &goson{
		v: make(map[string]json.RawMessage),
	}
}
func unmarshalGOSON(data []byte) (*goson, error) {
	g := &goson{}
	err := g.UnmarshalJSON(data)
	return g, err
}

func (g *goson) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.v)
}

func (g *goson) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &g.v)
}

func (g *goson) Value() map[string]json.RawMessage {
	for s, message := range g.v {
		fmt.Println("key", s, "value", string(message))
	}
	return g.v
}
