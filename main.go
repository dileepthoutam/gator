package main

import (
	"fmt"

	"github.com/dileepthoutam/gator/internal/config"
)

type state struct {

}

type command struct {

}

func main() {

    cfg := config.Read()
    cfg.SetUser("dileep")
    fmt.Println(cfg)

}
