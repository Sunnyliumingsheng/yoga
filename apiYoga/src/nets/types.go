package nets

type Message struct {
	IsSuccess bool
	Info      string
	Result    interface{}
}
