package bot

type Media struct {
	// Id Идентификатор медиа — уникальный и ссылающийся сразу на группу
	Id int64

	Content []*Content
}

type Content struct {
	// UnitId — внутренний идентификатор единицы контента
	UnitId int64
	// FileId идентификатор для файла в Telegram
	FileId string
	// Class класс медиа
	Class string
	// StorageId идентификатор в локальном хранилище
	StorageId string
}
