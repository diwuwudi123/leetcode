package main

import "log"

func main() {
	LRUCache := Constructor(2)
	LRUCache.Put(1, 1)
	LRUCache.Put(2, 2)
	val := LRUCache.Get(1)
	log.Println("get 1", val)
	log.Printf("lru cache :%+v", LRUCache)
	LRUCache.Put(3, 3)
	val = LRUCache.Get(3)
	log.Println("get 3", val)

}

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type Node struct {
	Key  int
	Val  int
	pre  *Node
	next *Node
}
type DoubleList struct {
	head, tail *Node
}

// 把这个节点插入到头结点
func (d *DoubleList) InsertHead(node *Node) {
	next := d.head.next
	d.head.next = node
	node.pre = d.head
	node.next = next
	next.pre = node
}

// 把这个接点从头结点拿掉
func (d *DoubleList) RemoveHead() *Node {
	next := d.head.next
	d.head.next = next.next
	next.next.pre = d.head
	//删除这个节点
	next.next = nil
	next.pre = nil
	return next
}

// 把这个节点从末尾拿掉
func (d *DoubleList) RemoveBack() *Node {
	node := d.tail.pre
	d.RemoveNode(node)
	return node
}

// 把指定节点拿掉
func (d *DoubleList) RemoveNode(node *Node) *Node {

	node.pre.next = node.next
	node.next.pre = node.pre
	//删除这个节点
	node.pre = nil
	node.next = nil
	return node
}

type LRUCache struct {
	cap  int
	dict map[int]*Node
	list *DoubleList
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		cap:  capacity,
		dict: map[int]*Node{},
		list: &DoubleList{head: head, tail: tail},
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.dict[key]; ok {
		//把这个节点对应的node 移到最前面
		node := this.list.RemoveNode(n)
		this.list.InsertHead(node)
		return n.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.dict[key]; ok {
		this.dict[key].Val = value
		//把这个节点移到最前面
		this.Get(key)
		return
	}
	//超出容量了
	if this.cap <= len(this.dict) {
		//删除尾节点.因为尾节点代表最久没使用的
		tail := this.list.RemoveBack()
		delete(this.dict, tail.Key)
	}
	//插入节点 插入到头节点,因为头结点代表最近使用
	node := &Node{Key: key, Val: value}
	this.list.InsertHead(node)
	this.dict[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
