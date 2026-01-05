package dbbbolt

import (
	"bytes"
	"encoding/binary"

	"github.com/fxamacker/cbor/v2"
	"github.com/rostertar/give-me-source-bot/modules/bot"
)

// UsersChat struct for link user ant group. Used as key serialized little endian userid:groupid pair
type UsersChat struct {
	// Временноее значение, чтобы не перегенерировать
	id []byte

	UserId      int64  `cbor:"1,keyasint" json:"user_id"`
	GroupId     int64  `cbor:"2,keyasint" json:"group_id"`
	Requests    uint64 `cbor:"3,keyasint" json:"requests"`
	Populatrity uint64 `cbor:"4,keyasint" json:"populatrity"`
	Forbiden    bool   `cbor:"10,keyasint" json:"forbiden"`
}

func MakeUsersChat(uid, cid int64, users bot.UsersChat) (id []byte, uc *UsersChat) {
	uc = &UsersChat{
		UserId:      uid,
		GroupId:     cid,
		Requests:    users.Requests,
		Populatrity: users.Populatrity,
		Forbiden:    users.Forbiden,
	}
	return uc.getKey(), uc
}

func (uc *UsersChat) getKey() []byte {
	if len(uc.id) > 0 {
		return uc.id
	}
	data := &bytes.Buffer{}
	data.Grow(16)
	binary.Write(data, binary.LittleEndian, uc.UserId)
	binary.Write(data, binary.LittleEndian, uc.GroupId)
	uc.id = data.Bytes()
	return uc.id
}

func (uc *UsersChat) asBotsUsersChar() *bot.UsersChat {
	if uc == nil {
		return nil
	}
	return &bot.UsersChat{
		Requests:    uc.Requests,
		Populatrity: uc.Populatrity,
		Forbiden:    uc.Forbiden,
	}
}

func (uc *UsersChat) Serialize() (key, value []byte, err error) {
	key = uc.getKey()
	value, err = cbor.Marshal(uc)
	return
}
