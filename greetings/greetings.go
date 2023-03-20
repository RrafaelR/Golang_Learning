package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//	  |          |         |
	//	func	  parameter   Return
	//  name       type  	   type
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	//creating a message using random
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hello associate with a map of names with a greeting message
func Hellos(names []string) (map[string]string, error) {
	//A map to associate names with messages
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//In the map, associate the retrived message with the map
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying
	//random index fo rthe slice of formats.
	return formats[rand.Intn(len(formats))]
}
