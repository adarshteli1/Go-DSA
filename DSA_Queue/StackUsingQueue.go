package dsaqueue

type Mystack struct {
	queue []int
}

func Constructer() Mystack {
	return Mystack{}
}
func (this *Mystack) Push(x int) {
	this.queue = append(this.queue, x)
	for i := 0; i < len(this.queue)-1; i++ {
		front := this.queue[len(this.queue)-1]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, front)
	}
}
func (this *Mystack) Pop() int {
	top := this.queue[len(this.queue)-1]
	this.queue = this.queue[1:]
	return top
}
func (this *Mystack) Top() int {
	return this.queue[len(this.queue)-1]
}
func (this *Mystack) IsEmpty() bool {
	return len(this.queue) == 0
}
