package _55__Min_Stack

type MinStack struct {
	items []int
	count int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
	this.count++
}

func (this *MinStack) Pop() {
	this.items = this.items[:this.count-1]
	this.count--
}

func (this *MinStack) Top() int {
	return this.items[this.count-1]
}

func (this *MinStack) GetMin() int {
	min := this.items[0]
	for _, v := range this.items {
		if v < min {
			min = v
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
