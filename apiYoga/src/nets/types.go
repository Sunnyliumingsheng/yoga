package nets

// 这是service返回的信息格式,Info是一些文本信息,用来描述发生了什么,如果成功result放数据,如果失败result放失败原因string
type Message struct {
	IsSuccess bool
	Info      string
	Result    interface{}
}
