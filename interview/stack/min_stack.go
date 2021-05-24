package stack

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/queue-stack/g5l7d/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

type MinStack struct {
	StackData []int
	StackMin  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		StackMin:  []int{},
		StackData: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.StackData = append(this.StackData, val)
	if len(this.StackMin) > 0 {
		if this.StackMin[len(this.StackMin)-1] >= val {
			this.StackMin = append(this.StackMin, val)
		}
	} else {
		this.StackMin = append(this.StackMin, val)
	}
	return
}

func (this *MinStack) Pop() {
	if len(this.StackMin) == 0 {
		return
	}
	value := this.StackData[len(this.StackData)-1]
	this.StackData = this.StackData[:len(this.StackData)-1]
	if value == this.StackMin[len(this.StackMin)-1] {
		this.StackMin = this.StackMin[:len(this.StackMin)-1]
	}
	return
}

func (this *MinStack) Top() int {
	if len(this.StackData) == 0 {
		return 0
	}
	return this.StackData[len(this.StackData)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.StackMin) == 0 {
		return 0
	}
	return this.StackMin[len(this.StackMin)-1]
}
