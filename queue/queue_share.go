package queue

type QueueShard []*Queue

var total = 32

// 减少锁
func NewQueueShard() QueueShard {

	_shard := make(QueueShard, total)

	for i := 0; i < total; i++ {
		_shard = append(_shard, New())
	}

	return _shard

}

func (q QueueShard) GetShard(key string) *Queue {
	return q[uint(fnv32(key))%uint(total)]
}

// 输入
func (q QueueShard) Push(k string, value interface{}) {

	_shard := q.GetShard(k)

	_shard.Push(value)

}

// 输入
func (q QueueShard) Pop(k int) interface{} {

	_shard := q[uint(k)%uint(total)]

	return _shard.Pop()

}

// 筛选
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
