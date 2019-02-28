package commands

func ParseCommand(message string) (Command, error) {
	tokens := Tokenize(message)
	// Make sure message not empty
	if len(tokens) == 0 {
		return nil, nil
	}

	// Parse into command struct
	return nil, nil
}
