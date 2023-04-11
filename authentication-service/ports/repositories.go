package ports

type Logger interface {
	Log(string, string) error
}
