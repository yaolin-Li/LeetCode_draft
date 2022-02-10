package draft

func SubarraySum(nums []int, k int) (ans int) {
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				ans++
			}
		}
	}
	return
}

//  通过大小判断是不对的
// 	for idx < n{
// 		windows = append(windows, nums[idx])
// 		sum += nums[idx]
// 		if sum > k {
// 			sum -= windows[0]
// 			windows = windows[1:]
// 		}
// 		if sum == k{
// 			ans++
// 			sum -= windows[0]
// 			windows = windows[1:]
// 		}
// 		idx++
// 	}
// 	return
// }
