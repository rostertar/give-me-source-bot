package bot

type Invite struct {
	Chat      *Chat
	MessageId int
	Request   uint64
}

type UsersChat struct {
	Requests    uint64
	Populatrity uint64
	Forbiden    bool
}

// User Структура для представления отдельного пользователя
type User struct {
	// ID идентификатор пользователяв Telegram
	ID int64
	// Nick имя пользователя для отобрадения
	Nick string
	// Invites активные приглашения для пользоватея c его запросами на контент
	Invites []*Invite
	// Chats список чатов в которых участвует пользователь и в которых он делал запросы
	Chats map[int64]*UsersChat
	// AdminCap возможности по управлению ботом
	AdminCap CAP
}
