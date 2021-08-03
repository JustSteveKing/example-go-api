package photos

import (
	"net/http"
	"sync"
	"time"
)

type Service struct {
	Client *http.Client
}

func NewService() *Service {
	return &Service{
		Client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (service *Service) Get(url string, waitGroup *sync.WaitGroup) (*http.Response, error) {
	response, err := service.Client.Get(url)

	waitGroup.Done()

	return response, err
}
