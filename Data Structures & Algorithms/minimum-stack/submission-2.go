type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	value := this.stack[len(this.stack)-1]
	return value
}

func (this *MinStack) GetMin() int {
	minValue := this.stack[0]
	for _, v := range this.stack {
		if minValue > v {
			minValue = v
		}
	}

	return minValue
}
