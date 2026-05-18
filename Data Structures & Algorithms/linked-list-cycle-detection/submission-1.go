/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    m := map[*ListNode]bool{}
    current := head
    for current != nil {
        if isPresent := m[current]; isPresent {
            return true
        }
        m[current] = true
        current = current.Next
    }

    return false
}
