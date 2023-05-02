package ports

type LogEntry struct {
	Name string
	Data string
}

type LogRepository interface {
	Insert(LogEntry) error
}
