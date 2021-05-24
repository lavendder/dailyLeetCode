package stack

import "fmt"

/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )

示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
1,2,3,4,5
5,4,3,2,1

1,2,3,4,5,6
6,5,4,3,2,1



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
**/

type CQueue struct {
	Stack        []int
	ReverseStack []int
}

func Constructor() CQueue {
	return CQueue{
		Stack:        []int{},
		ReverseStack: []int{},
	}

}

func (this *CQueue) AppendTail(value int) {
	if value < 1 || value > 1000 {
		return
	}
	this.Stack = append(this.Stack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.ReverseStack) == 0 {
		if len(this.Stack) == 0 {
			return -1
		}
		for i := len(this.Stack) - 1; i >= 0; i-- {
			this.ReverseStack = append(this.ReverseStack, this.Stack[i])
		}
	}
	if len(this.ReverseStack) == 0 {
		return -1
	}
	res := this.ReverseStack[len(this.ReverseStack)-1]
	this.ReverseStack = this.ReverseStack[:len(this.ReverseStack)-1]
	fmt.Println(this.ReverseStack)
	return res
}
