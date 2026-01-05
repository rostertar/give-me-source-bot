package dbbbolt

import (
	"bytes"
	"encoding/binary"

	"github.com/fxamacker/cbor/v2"
	"github.com/rostertar/give-me-source-bot/modules/bot"
)

type Chat struct {
	id []byte

	Id   int64  `cbor:"1,keyasint" json:"id"`
	Name string `cbor:"2,keyasint,omitempty" json:"name,omitempty"`
	Anon int64  `cbor:"3,keyasint,omitempty" json:"anon,omitempty"`
}

func MakeChat(c *bot.Chat) *Chat {
	id := &bytes.Buffer{}
	binary.Write(id, binary.LittleEndian, c.Id)
	return &Chat{
		id:   id.Bytes(),
		Id:   c.Id,
		Name: c.Name,
		Anon: c.Anon,
	}
}

func (c *Chat) getKey() []byte {
	if len(c.id) > 0 {
		return c.id
	}
	id := &bytes.Buffer{}
	binary.Write(id, binary.LittleEndian, c.Id)
	c.id = id.Bytes()
	return c.id
}

func (c *Chat) Serialize() (key, value []byte, err error) {
	key = c.getKey()
	value, err = cbor.Marshal(c)
	return
}
