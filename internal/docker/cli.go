package main

import (
	"camellia/tool/exception"
	"camellia/tool/terminal"
	"fmt"
)

func ssh() {
	ter := terminal.Terminal{
		Ip:       "192.168.3.9",
		Port:     "22",
		User: "root",
		Password: "root",
	}

	client, err := terminal.GetInstance(&ter)

	exception.GetIns().Throw(err)

	session, err := client.NewSession()

	exception.GetIns().Throw(err)

	result, err := session.Output("ls -al")

	exception.GetIns().Throw(err)

	fmt.Println(string(result))

}

func main()  {
	ssh()
}
