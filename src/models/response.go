package models

type HttpResponse struct {
	Status      string
	Description string
	Code        int
	Data        interface{}
}
