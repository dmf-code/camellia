package main

import (
	"camellia/tool/exception"
	"camellia/tool/terminal"
	"fmt"
)

func ssh() {
	conf := terminal.Config{
		Ip:       "192.168.3.9",
		Port:     "22",
		User: "root",
		Password: "root",
	}

	client, err := terminal.GetInstance(&conf)

	exception.GetIns().Throw(err)

	t := terminal.Terminal{
		Config: &conf,
		Client: client,
	}

	session, err := t.NewSession()

	str, err := t.Shell(session, []string{
		"cd /work/dockerfile",
		"mkdir domain",
		"exit",
	})
	exception.GetIns().Throw(err)
	for _, msg := range str {
		fmt.Println(msg)
	}

	defer session.Close()
}

func main()  {
	ssh()
}
