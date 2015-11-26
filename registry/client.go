package registry

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/viniciuschiele/decker/config"
)

var (
	currentToken  string
	currentServer string
)

// Login ...
func Login(username, password string) (string, error) {
	var jsonStr = []byte(`""`)

	request, err := http.NewRequest("POST", config.GetServer(), bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return "", nil
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(body)

	return "token 1", nil
}

// Init ...
func Init(token, server string) {
	currentToken = token
	currentServer = server
}
