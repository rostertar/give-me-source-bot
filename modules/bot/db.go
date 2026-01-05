package bot

type DB interface {
	GetUser(Id int64) *User
	AlterUser(*User) error
	GetChat(Id int64) *Chat
	AlterChat(*Chat) error
	GetRequest(Id int64) *Request
	AlterRequest(*Request) error
	GetMedia(Id int64) *Media
	AlterMedia(*Media) error
}
