package models

/*
	index    0  1  2   3  4   5
	nums = [-2, 0, 3, -5, 2, -1]

	preSums[0] = 0
	preSums[1] = preSums[0] + nums[0] = 0 + (-2) = -2 👉 nums[0] = preSum[1]  <=> -2 = -2
	preSums[2] = preSums[1] + nums[1] = -2 + 0 = -2
	preSums[3] = preSums[2] + nums[2] = -2 + 3 = -1
	=> preSums[n] = nums[n-1] + preSums[n-1]

	preSums[4] = nums[3] + preSum[3]
	<=> preSums[4] = nums[0] + nums[1] + nums[2] + nums[3]
	<=>	nums[1] + nums[2] + nums[3] = preSums[4] - nums[0] = preSums[4] - preSums[1] // 👉 nums[0] = preSum[1]
	<=> SumRange(1,3) = preSums[4] - preSums[1]

	   index      0   1   2  3   4   5   6
	=> preSums = [0, -2, -2, 1, -4, -2, -3]

	<=> preSums[3] = nums[0] + nums[1] + nums[2]
	<=> preSums[3] = preSums[0] + nums[0] + nums[1] + nums[2]
	<=> nums[0] + nums[1] + nums[2] = preSums[3] - preSums[0]
	<=> SumRange(0,2) = preSums[3] - preSums[0]



	numArray.sumRange(0, 5);
	preSums[6] = nums[0] + nums[1] + nums[2] + nums[3] + nums[4] + nums[5]
<=> preSums[6] = preSum[0] + nums[0] + nums[1] + nums[2] + nums[3] + nums[4] + nums[5]
<=> nums[0] + nums[1] + nums[2] + nums[3] + nums[4] + nums[5] = preSums[6] - preSum[0]
<=> SumRange(0, 5) = preSums[6] - preSum[0] = -3 - 0 = -3


	numArray.sumRange(4, 5) = nums[4] + nums[5] = 2 - 1 = 1
	preSums[6] = nums[0] + nums[1] + nums[2] + nums[3] + nums[4] + nums[5]
	nums[4] + nums[5] = preSums[6] - (nums[0] + nums[1] + nums[2] + nums[3])
	SumRange(4,5) = preSums[6] - preSum[4]

	=> SumRange(Left,Right) = preSums[Right+1] - preSums[Left]
*/

type NumArray struct {
	PreSums []int
}

func Constructor(nums []int) NumArray {
	preSums := make([]int, len(nums)+1)
	preSums[0] = 0

	for index, num := range nums {
		preSumIndex := index + 1
		preSums[preSumIndex] = num + preSums[index]
	}
	return NumArray{
		PreSums: preSums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	sumRange := this.PreSums[right+1] - this.PreSums[left]
	return sumRange
}
