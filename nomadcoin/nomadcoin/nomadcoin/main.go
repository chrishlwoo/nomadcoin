package main

import (
	"github.com/chrishlwoo/nomadcoin/cli"
	"github.com/chrishlwoo/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
	}



