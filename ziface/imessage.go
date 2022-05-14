package ziface

/*
	将请求的一个消息封装到message中，定义抽象层接口
*/
type IMessage interface {
	//获取消息数据段长度
	GetDataLen() uint32
	// 获取消息ID
	GetMsgId() uint32
	// 获取消息内容
	GetData() []byte
	// 设计消息ID
	SetMsgId(uint32)
	// 设计消息内容
	SetData([]byte)
	// 设置消息数据段长度
	SetDataLen(uint32)
}
