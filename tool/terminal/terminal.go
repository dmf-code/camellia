package terminal

import (
	"golang.org/x/crypto/ssh"
	"log"
)

type Terminal struct {
	Ip string
	Port string
	User string
	Password string
}

func (t *Terminal) GetAddr() string {
	return t.Ip + ":" + t.Port;
}

// 获取命令行实例
func GetInstance(t *Terminal) (*ssh.Client, error) {
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", t.GetAddr(), &ssh.ClientConfig{
		User:            t.User,
		Auth:            []ssh.AuthMethod{ssh.Password(t.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	return client, nil
}

