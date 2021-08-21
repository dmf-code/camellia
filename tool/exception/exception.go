package exception

import (
	"fmt"
	"os"
	"sync"
)

type Exception struct {
	debug bool
}

var exception *Exception
var once sync.Once

func GetIns() *Exception {
	once.Do(func() {
		exception = New()
	})

	return exception
}

func New() *Exception {

	debug := false
	if os.Getenv("debug") == "true" {
		debug = true
	}

	return &Exception{
		debug: debug,
	}
}

// 异常抛出
func (e *Exception) Throw(err error) {

	if err != nil {

		if e.debug == true {
			fmt.Println(err.Error())
		}

		panic(err)
	}
}

