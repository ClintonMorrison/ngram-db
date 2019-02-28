package commands

type Command interface {
	Apply() string
	String() string
}


type ExistsCommand struct {
	Key string
}
