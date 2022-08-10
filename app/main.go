package app

import (
	"bytes"
	"fmt"
	"net/http"
)

func Debug(data string) {
	dataBytes := []byte(fmt.Sprintf(`{
	"debug": "%s"
	}`, data))

	http.Post("http://localhost:8080/debug", "application/json", bytes.NewBuffer(dataBytes))
}

func main() {
	Debug("Hello World")
}
