package server

const LimitFork = 100

func MakeChanInt(limit int) *chan int {
	var contralChan = make(chan int, LimitFork)

	return &contralChan

}
