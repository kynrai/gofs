package cli

type helpCommand struct {
}

func (h helpCommand) Name() string {
	return "help"
}

func (h helpCommand) Description() string {
	return "Display help message"
}

func (h helpCommand) Run() {

}
