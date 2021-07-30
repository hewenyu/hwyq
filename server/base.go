package server

const ( 
	LimitFork = 100  // 默认的最大链接树
	DefultPort = 8950 // 端口号
)
/**
 * MakeChanInt
 */ 
func MakeChanInt(limit int) *chan int {
	var contralChan = make(chan int, limit)

	return &contralChan

}
