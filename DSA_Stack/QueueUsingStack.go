package dsastack

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructer() MyQueue {
	return MyQueue{}
}
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}
func (this *MyQueue) Pop(x int) int {
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			top := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, top)
		}
	}
	top := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return top
}
func (this *MyQueue) Peek(x int) int {
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			top := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, top)
		}
	}
	return this.stack2[len(this.stack2)-1]
}
func (this *MyQueue) IsEmpty(x int) bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}
