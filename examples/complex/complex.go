package complex

import (
	"fmt"
)

// Client some client interface
type Client interface {
	Do() (*Response, error)
}

// Response the clients response message
type Response struct {
	Message string
}

// Worker a worker struct containing the client
type Worker struct {
	Client Client
}

// SomeFunc some function on the client
func (w *Worker) SomeFunc(input []int) (string, error) {
	if len(input) < 2 {
		return "", fmt.Errorf("expected input length to be atleast 2")
	}
	resp, err := w.Client.Do()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s, %v", resp.Message, input), nil
}
