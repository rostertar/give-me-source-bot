package bot

// CAP матрица прав для пользователя. Без капа можно только
// запрашивать контент, если админ этого не запретит
type CAP uint64

const (
	CAP_GOD = ^CAP(0)
)

func (c CAP) String() string {
	if c != 0 {
		return "god"
	}
	return ""
}

type ChatCAP uint64

const (
	// ChatCAP_Request Права запросить медиа
	ChatCAP_Request = ChatCAP(0x1)
	// ChatCAP_Disable Запретить пересылку конкретного медиа
	ChatCAP_Disable = ChatCAP(0x2)
	// ChatCAP_ALL взвести все разрешения
	ChatCAP_ALL = ^ChatCAP(0)
)

func (c ChatCAP) String() string {
	cap := ""
	if c&ChatCAP_Request != 0 {
		cap = "request"
	}
	if c&ChatCAP_Disable != 0 {
		if len(cap) > 0 {
			cap = cap + ";"
		}
		cap = cap + "disable"
	}
	return cap
}
