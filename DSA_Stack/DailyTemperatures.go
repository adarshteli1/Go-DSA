package dsastack

func Dailytemperatures(temperatures []int)[]int{
	stack:=[]int{}
	res:=make([]int,len(temperatures))
	for i:=0;i<len(temperatures);i++{
		for len(stack)>0 && temperatures[i]>temperatures[stack[len(stack)-1]]{
			top:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			res[top]=i-top
		}
		stack=append(stack,i)
	}
	return res
}