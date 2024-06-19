package cli

type Command interface {
	Name() string
	Description() string
	Run()
}
