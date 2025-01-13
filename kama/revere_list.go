package kama

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	val  int
	next *ListNode
}

func kamaReverse() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		n, _ := strconv.Atoi(nums[0])
		if n == 0 {
			fmt.Println("list is empty")
		} else {
			var arr []int
			for i := 0; i < n; i++ {
				m, _ := strconv.Atoi(nums[i+1])
				arr = append(arr, m)

			}
			head := buildList(arr)
			print(head)
			head = reverse(head)
			print(head)
		}
	}
}

func buildList(num []int) *ListNode {
	if len(num) == 0 {
		return nil
	}
	dummy := &ListNode{
		val: -1,
	}
	p := dummy
	for _, n := range num {
		node := &ListNode{
			val: n,
		}
		p.next = node
		p = p.next

	}

	return dummy.next
}

func print(head *ListNode) {
	p := head

	for p != nil {
		fmt.Print(p.val)
		fmt.Print(" ")
		p = p.next
	}

}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		post := cur.next
		cur.next = pre
		pre = cur
		cur = post
	}
	return pre.next
}
