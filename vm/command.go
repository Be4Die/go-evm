// vm/command.go
package vm

type Command interface {
    Execute() error
}

type CommandFunc func() error

func (f CommandFunc) Execute() error {
    return f()
}