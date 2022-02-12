/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package solutions
type ListNode struct {
	Val int
	Next *ListNode
}

 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // known: 
    // 1. lengh of list <= 30
    // 2. n >= 1
    
    // transfer to an array
    arr := [30]*ListNode{}
    index := 0
    for head!=nil {
        arr[index] = head
        head = head.Next
        index++
    }
    
    // do operation on the nth node
    length := index
    deletePosition := length - n
    switch {
    case (length==1):       // only possible to remove the only one
        return nil;
    case (length==n):       // remove the head
        return arr[1];
    case (n==1):            // remove the tail
        arr[deletePosition-1].Next = nil;
        return arr[0];
    default:
        arr[deletePosition-1].Next = arr[deletePosition+1]
        return arr[0];
    }

}