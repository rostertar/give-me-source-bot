package bot

// Request
type Request struct {
	// MediaId идентификатор медиа для пересылки
	MediaId int64

	// RequesterId идентификатор пользователя, который запаросил медиа
	RequesterId int64

	// Executed признак того что запрос выполнен
	Executed bool
}
