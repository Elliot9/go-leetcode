package main

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{
		k: k,
		nums: []int{
			0,
		},
	}

	for _, num := range nums {
		kthLargest.Add(num)
	}

	return kthLargest
}

// 解題思路 root 也就是 items[1] 會是第 k 大的數值, 而底下的數值都會是比他還大的數值
// 例如: [1,3,5,7,2] , k = 3
//  3
// 5 7
// 注意的是 heap 規則是 子節點一定比父節點大而已, 左右並沒有比較關係
//       3
//      / \
//     5   8
//    / \
//   9   6

// 此時 k == len 時, 代表數組已滿, 如果 加入的數值有大於 items[1] 則該數值該進入數組, => items[1] = val, 後向下檢查是否需要 swap
// 例如加入的數字為 6 => ˊ6
//                    5 7  <-  此時 6 需要和 5 swap

//	而如果 k = 5, [1,3,5] 代表數組還沒有滿, 此時加入數值後 要從尾巴開始檢查數組是否合法
// 假入加入數字為 2 => 1
//                 3  5
//                2         從尾開始 2 是否比 3小, 如果有就交換 1
// 													      2 5
// 														 3

func (this *KthLargest) Add(val int) int {
	if len(this.nums) <= this.k {
		this.nums = append(this.nums, val)
		this.heapifyUp()
	} else {
		// 加入的數值有大於 items[1] 則該數值該進入數組
		if this.nums[1] < val {
			this.nums[1] = val
			this.heapifyDown()
		}
	}

	return this.nums[1]
}

func (this *KthLargest) heapifyUp() {
	// [0] 是哨兵元素不需要計算
	index := len(this.nums) - 1

	for this.parentIndex(index) >= 1 && this.parent(index) > this.nums[index] {
		this.nums[this.parentIndex(index)], this.nums[index] = this.nums[index], this.nums[this.parentIndex(index)]
		index = this.parentIndex(index)
	}
}

func (this *KthLargest) parentIndex(index int) int {
	return index / 2
}

func (this *KthLargest) parent(index int) int {
	return this.nums[this.parentIndex(index)]
}

func (this *KthLargest) heapifyDown() {
	index := 1
	for this.hasLeft(index) {
		smallestIndex := this.leftChildIndex(index)

		if this.hasRight(index) && this.right(index) < this.left(index) {
			smallestIndex = this.rightChildIndex(index)
		}

		// 目前比子小 就不用動
		if this.nums[index] < this.nums[smallestIndex] {
			break
		} else {
			this.nums[index], this.nums[smallestIndex] = this.nums[smallestIndex], this.nums[index]
			index = smallestIndex
		}
	}
}

func (this *KthLargest) leftChildIndex(index int) int {
	return index * 2
}

func (this *KthLargest) rightChildIndex(index int) int {
	return index*2 + 1
}

func (this *KthLargest) hasLeft(index int) bool {
	return this.leftChildIndex(index) < len(this.nums)
}

func (this *KthLargest) hasRight(index int) bool {
	return this.rightChildIndex(index) < len(this.nums)
}

func (this *KthLargest) left(parentIndex int) int {
	return this.nums[this.leftChildIndex(parentIndex)]
}

func (this *KthLargest) right(parentIndex int) int {
	return this.nums[this.rightChildIndex(parentIndex)]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
