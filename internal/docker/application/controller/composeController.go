package controller

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type ComposeController struct {
}

func (ctr *ComposeController) Start() {

	err := os.Chdir("")
	cmd := exec.Command("docker-compose", "up", "-d")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
