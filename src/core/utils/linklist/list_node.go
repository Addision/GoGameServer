package linklist

type ListNode struct {
	Data interface{}
	Prev *ListNode
	Next *ListNode
}

func newlinkNode(Data interface{}) *ListNode {
	return &ListNode{
		Data: Data,
		Prev: nil,
		Next: nil,
	}
}
