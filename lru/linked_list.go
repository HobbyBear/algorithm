package lru

// 最近最少被使用
type CacheNode struct {
	Next *CacheNode
	Val  string
}

type Cache struct {
	Head        *CacheNode
	MaxSize     int
	CurrentSize int
}

func (c *Cache) Put(n *CacheNode) {
	n.Next = c.Head
	c.Head = n
	c.CurrentSize++
	if c.CurrentSize > c.MaxSize {
		c.delEnd()
	}
}

func (c *Cache) delEnd() {
	end := c.Head
	if end == nil {
		return
	}
	if end.Next == nil {
		c.Head = nil
		return
	}

	var cur = new(CacheNode)
	// 找到最后一个节点的前一个节点
	for end.Next != nil {
		cur = end
		end = end.Next
	}
	cur.Next = nil
}

func (c *Cache) remove(val string) *CacheNode {
	// 找到删除节点的前一个节点
	cur := c.Head
	if cur == nil {
		return nil
	}

	if cur.Val == val {
		c.Head = c.Head.Next
		return nil
	}

	next := c.Head
	pre := c.Head
	for next != nil && next.Val != val {
		pre = next
		next = next.Next
	}
	if next != nil {
		pre.Next = next.Next
	}
	return next
}
