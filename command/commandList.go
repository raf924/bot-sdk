package command

import (
	"strings"
	"sync"
)

type List interface {
	Range(f func(command Command) bool)
	Get(i int) Command
	Find(command string) Command
	Add(command Command)
}

var _ List = (*immutableCommandList)(nil)

var _ List = (*commandList)(nil)

type immutableCommandList struct {
	List
}

func (i *immutableCommandList) Add(Command) {
	panic("cannot modify list")
}

func (i *immutableCommandList) Append(List) {
	panic("cannot modify list")
}

func ImmutableCommandList(cl List) List {
	return &immutableCommandList{List: cl}
}

var _ = ImmutableCommandList

type commandList struct {
	rwm            *sync.RWMutex
	commands       []Command
	commandIndexes map[string]int
}

func (l *commandList) Range(f func(command Command) bool) {
	for _, command := range l.commands {
		if !f(command) {
			break
		}
	}
}

func (l *commandList) Get(i int) Command {
	l.rwm.RLock()
	command := l.commands[i]
	l.rwm.RUnlock()
	return command
}

func (l *commandList) Find(command string) Command {
	l.rwm.RLock()
	actualCommand := func() Command {
		i, ok := l.commandIndexes[command]
		if !ok {
			return nil
		}
		return l.commands[i]
	}()
	l.rwm.RUnlock()
	return actualCommand
}

func (l *commandList) Add(command Command) {
	aliases := append(command.Aliases(), command.Name())
	if len(strings.TrimSpace(command.Name())) == 0 {
		return
	}
	l.rwm.Lock()
	func() {
		l.commands = append(l.commands, command)
		for _, alias := range aliases {
			if len(strings.TrimSpace(alias)) == 0 {
				return
			}
			l.commandIndexes[alias] = len(l.commands) - 1
		}
	}()
	l.rwm.Unlock()
}

func NewCommandList(commands ...Command) List {
	ul := &commandList{
		commands:       []Command{},
		commandIndexes: map[string]int{},
		rwm:            &sync.RWMutex{},
	}
	for _, command := range commands {
		ul.Add(command)
	}
	return ul
}
