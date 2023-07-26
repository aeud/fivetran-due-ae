package fivetranio

import (
	"encoding/json"
	"log"
)

type Secret struct {
	APIKey   string   `json:"api_key"`
	Entities []string `json:"entities"`
}

func (s *Secret) MarshalForce() []byte {
	bs, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return bs
}
func (s *Secret) LogContent() {
	log.Printf("%s\n", s.MarshalForce())
}
