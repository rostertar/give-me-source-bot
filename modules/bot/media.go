package bot

type MediaType int

const (
	TypePhoto = MediaType(iota)
	TypeAudio
	TypeVideo
	TypeAnimation
	TypeDocument
)

type Media struct {
	// Id Идентификатор медиа — уникальный и ссылающийся сразу на группу
	Id int64

	// CommonType тип всех записей в еонтенте
	CommonType MediaType

	// Список единиц контента
	Content []*Content
}

type Content struct {
	// UnitId — внутренний идентификатор единицы контента
	UnitId uint64
	// FileId идентификатор для файла в Telegram
	FileId string
	// Class класс медиа
	Type MediaType
	// Path идентификатор в локальном хранилище
	Path string
}
