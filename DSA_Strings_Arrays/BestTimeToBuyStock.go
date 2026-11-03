package dsastringsarrays

func BestTimeToBuySellStock(nums []int) int {
	minprices := nums[0]
	maxprofit := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < minprices {
			minprices = nums[i]
		}
		profit := nums[i] - minprices
		if maxprofit < profit {
			maxprofit = profit
		}
	}
	return maxprofit
}
