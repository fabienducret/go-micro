package ports

type Log struct {
	Name string
	Data string
}

type Logger interface {
	Log(Log) error
}
