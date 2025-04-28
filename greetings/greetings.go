package greetings

import "errors"

func Greet(name string) (string, error) {
	if(name == "") {
		return "", errors.New("name is empty")
	}
	message := "Hello "+name+"! "
	return message, nil
}