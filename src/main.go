package main

import (
	"errors"
	"fmt"
	"net/http"
)

// A simple utility function "MyLogger"
func MyLogger(msg string) {
	fmt.Println(msg)
}

// A simple data store
type MyDataStore struct {
	someData map[string]string
}

// A method for it
func (mds MyDataStore) GetUserByID(id string) (string, bool) {
	name, ok := mds.someData[id]
	return name, ok
}

// Factory function creates an instande of MyDataStore

func NewMyDataStore() MyDataStore {
	return MyDataStore{
		someData: map[string]string{
			"1": "Michael Jackson",
			"2": "George Michael",
		},
	}
}

// Interfaces for DIP

type SomeDataStore interface {
	GetUserByID(id string) (string, bool)
}

type SomeLogger interface {
	LogMessage(message string)
}

// Adapter between Interface and Logger

type LoggerAdapter func(message string)

func (lg LoggerAdapter) LogMessage(message string) {
	lg(message)
}

// Dependencies defined, now Business Logic:

type SimpleLogic struct {
	l  SomeLogger
	ds SomeDataStore
}

func (ml SimpleLogic) GreetUser(id string) (string, error) {
	ml.l.LogMessage("UserID: " + id)
	name, ok := ml.ds.GetUserByID(id)
	if ok {
		return "Welcome " + name, nil
	} else {
		return "", errors.New("No User")
	}
}

func NewSimpleLogic(l SomeLogger, ds SomeDataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type SomeLogic interface {
	GreetUser(id string) (string, error)
}

type Controller struct {
	logger SomeLogger
	logic  SomeLogic
}

func (c Controller) HandleGreeting(w http.ResponseWriter, r *http.Request) {
	c.logger.LogMessage("Hi")
	userID := r.URL.Query().Get("id")
	message, err := c.logic.GreetUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.Write([]byte(message))
	}
}

func NewController(l SomeLogger, logic SomeLogic) Controller {
	return Controller{
		logger: l,
		logic:  logic,
	}
}

func main() {

	loggerAdapter := LoggerAdapter(MyLogger)
	dataStore := NewMyDataStore()
	logic := NewSimpleLogic(loggerAdapter, dataStore)
	controller := NewController(loggerAdapter, logic)
	http.HandleFunc("/hi", controller.HandleGreeting)
	http.ListenAndServe(":8080", nil)

	// http://localhost:8080/hi?id=1
	// Welcome Michael Jackson

	fmt.Println("Before Program End")

}
