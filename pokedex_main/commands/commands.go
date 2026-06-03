package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
}
