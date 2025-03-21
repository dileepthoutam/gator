package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dileepthoutam/gator/internal/config"
)

type state struct {
    config *config.Config
}
type command struct {
    name string
    args []string
}

type commands struct {
    Command map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
    c.Command[name] = f
}

func (c *commands) run(s *state, cmd command) error {
    if val, ok := c.Command[cmd.name]; ok {
	return val(s, cmd)
    }
    return fmt.Errorf("Command doesn't exist")
}

func handlerLogin(s *state, cmd command) error {
    if len(cmd.args) != 1 {
	fmt.Println(fmt.Errorf("Username is required or you passed too many args."))
	os.Exit(1)
    }
    username := cmd.args[0]
    s.config.SetUser(username)
    log.Printf("The user: %s has been set.", username)
    return nil
}

func main() {

    st := &state{
	config: config.Read(),
    }

    cmds := &commands{
	Command: map[string]func(*state, command) error {
	    "login": handlerLogin,
	},
    }

    if len(os.Args) < 2 {
	log.Println("Atleast pass one argument")
	os.Exit(1)
    }


    cmd := command {
	name: os.Args[1],
	args: os.Args[2:],
    }

    cmds.run(st, cmd)

}
