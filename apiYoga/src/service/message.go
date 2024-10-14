package service

type Message struct {
	IsSuccess bool
	HaveError bool
	Info      string
	Result    interface{}
}
