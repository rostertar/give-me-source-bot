package bot

type Invite struct {
	Id        uint64
	ChatId    int64
	MessageId int
}

// User Структура для представления отдельного пользователя
type User struct {
	// ID идентификатор пользователяв Telegram
	ID int64
	// Username — имя пользователя как кнр видит telegram
	Username string
	// Nick имя пользователя для отобрадения
	Nick string
	// Invites активные приглашения для пользоватея c его запросами на контент
	Invites []*Invite
	// AdminCap возможности по управлению ботом
	AdminCap CAP
}
