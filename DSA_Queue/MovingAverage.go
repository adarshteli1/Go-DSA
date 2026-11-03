package dsaqueue

type MovingAverage struct {
	Data  []int
	Front int
	Rear  int
	Size  int
	Cap   int
	Sum   int
}

func Constructor3(size int) MovingAverage {
	return MovingAverage{
		Data:  make([]int, size),
		Front: 0,
		Rear:  -1,
		Size:  0,
		Cap:   size,
		Sum:   0,
	}
}

func (q *MovingAverage) Next(value int) float64 {
	if q.Size == q.Cap {
		q.Sum -= q.Data[q.Front]
		q.Front = (q.Front + 1) % q.Cap
		q.Size--
	}
	q.Rear = (q.Rear + 1) % q.Cap
	q.Data[q.Rear] = value
	q.Sum += q.Data[q.Rear]
	q.Size++
	return float64(q.Sum) / float64(q.Size)
}

/*func main(){
	obj:=Constructor3(3)
	fmt.Println(obj.Next(1))
    fmt.Println(obj.Next(10))
    fmt.Println(obj.Next(3))
    fmt.Println(obj.Next(5))
}

*/
