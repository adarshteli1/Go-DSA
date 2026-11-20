package dsaqueue

type MovingAverage struct {
	Data      []int
	MaxWindow int
	Sum       int
}

func Constructor3(size int) MovingAverage {
	return MovingAverage{
		Data:      make([]int, 0, size),
		MaxWindow: size,
		Sum:       0,
	}
}

func (q *MovingAverage) Next(value int) float64 {
	q.Data = append(q.Data, value)
	q.Sum += value
	if len(q.Data) > q.MaxWindow {
		q.Sum -= q.Data[0]
		q.Data = q.Data[1:]
	}
	return float64(q.Sum) / float64(len(q.Data))
}

/*func main(){
	obj:=Constructor3(3)
	fmt.Println(obj.Next(1))
    fmt.Println(obj.Next(10))
    fmt.Println(obj.Next(3))
    fmt.Println(obj.Next(5))
}

*/
