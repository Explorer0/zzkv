package storager

// 存储接口
type Storager interface {
	// 存储值，返回存储位置
	SetValue(Field []string, value interface{}) string

	// 根据位置获取值
	GetValue(postition string) interface{}
}
