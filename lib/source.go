package slownews

type Source interface {
	GetNews() (error, []Article)
}
