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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    if list1==nil && list2==nil {
        return nil
    }
    if list1==nil {
        return list2
    }
    if list2==nil{
        return list1
    }
    
    var root = &ListNode{}
    var result = root;
    var head1 = list1;
    var head2 = list2

    for head1 != nil && head2 != nil {
        if head1.Val < head2.Val {
            root.Next = head1;
            head1 = head1.Next;
            root = root.Next;
        }else{
            root.Next = head2;
            head2 = head2.Next
            root = root.Next;
        }
    }
    
    // remaining lists are already sorted
    if head1 != nil {
        root.Next = head1;
    }else{
        root.Next = head2;
    }
    
    
    return result.Next;
}