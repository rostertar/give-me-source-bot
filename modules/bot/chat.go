package bot

import "time"

type Chat struct {
	Id      int64
	Name    string
	TTL     time.Duration
	Members []*Member
}
