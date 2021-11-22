package terminal

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
)

type Terminal struct {
	Config *Config
	Client *ssh.Client
}

type Config struct {
	Ip string
	Port string
	User string
	Password string
	Key string
	CipherList []string
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%s", c.Ip, c.Port)
}

func (c *Config) GetKey() string {
	return c.Key
}

// 获取命令行实例
func GetInstance(c *Config) (*ssh.Client, error) {

	var auth []ssh.AuthMethod
	if c.Key == "" {
		auth = append(auth, ssh.Password(c.Password))
	} else {
		pemBytes, err := ioutil.ReadFile(c.Key)

		if err != nil {
			return nil, err
		}

		var signer ssh.Signer

		if c.Password == "" {
			signer, err = ssh.ParsePrivateKey(pemBytes)
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(c.Password))
		}

		if err != nil {
			return nil, err
		}

		auth = append(auth, ssh.PublicKeys(signer))
	}

	var sshConfig ssh.Config

	if len(c.CipherList) == 0 {
		sshConfig = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		sshConfig = ssh.Config{
			Ciphers: c.CipherList,
		}
	}

	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", c.GetAddr(), &ssh.ClientConfig{
		User:            c.User,
		Auth:            auth,
		Config: sshConfig,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	return client, nil
}

func (t *Terminal) NewSession() (*ssh.Session, error) {
	// create session
	var (
		session *ssh.Session
		err error
	)
	if session, err = t.Client.NewSession(); err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:     0,   // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err = session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}

	return session, nil
}

func (t *Terminal) Shell(session *ssh.Session, c []string) ([]string, error) {

	stdinBuf, err := session.StdinPipe()
	if err != nil {
		return []string{""}, err
	}

	var outbt, errbt bytes.Buffer
	session.Stdout = &outbt

	session.Stderr = &errbt
	err = session.Shell()
	if err != nil {
		return []string{""}, err
	}

	for _, v := range c {
		stdinBuf.Write([]byte(v+ "\n"))
	}

	session.Wait()

	return []string{outbt.String(), errbt.String()}, nil
}
