package console

type Command interface {
	Exec([]string)
}
