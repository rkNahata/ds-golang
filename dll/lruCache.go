package dll

type LRUCache struct {
	keyMap   map[int]*Node
	list     *LinkedList
	capacity int
}

type Pair struct {
	key, value int
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		keyMap:   make(map[int]*Node, capacity),
		list:     New(),
		capacity: capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.keyMap[key]; ok {
		cache.list.DeleteNode(node)
		newNode := cache.list.Prepend(node.value)
		cache.keyMap[key] = newNode
		return node.value.(*Pair).value
	}
	return -1
}

func (cache *LRUCache) Set(key, value int) {
	pair := &Pair{key: key, value: value,}
	if node, ok := cache.keyMap[key]; ok {
		cache.list.DeleteNode(node)
		newNode := cache.list.Prepend(pair)
		cache.keyMap[key] = newNode

	} else {
		var newNode *Node
		if cache.capacity == cache.list.Size() {
			node := cache.list.DeleteLast()
			delete(cache.keyMap, node.value.(*Pair).key)
		}
		newNode = cache.list.Prepend(pair)
		cache.keyMap[key] = newNode
	}
}
