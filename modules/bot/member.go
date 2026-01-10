package bot

type Member struct {
	// Chat Указатель на чат
	Chat *Chat
	// User пользователь для данного чата
	User *User
	// CAP права пользователя на взаимодействие с данным чатом
	CAP ChatCAP
	// Requests количиество запросов от пользователя в этом чате
	Requests uint64
	// Populatity сколбкл раз запросили контент размещённый пользователем в этом чате
	Popularity uint64
}
