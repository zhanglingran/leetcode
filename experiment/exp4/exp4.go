package exp4

import (
	"math/rand"
	"time"
)

func InertSort(A []int, p, r int) {
	for i := p + 1; i <= r; i++ {
		for j := i; j > p && A[j] < A[j-1]; j-- {
			A[j], A[j-1] = A[j-1], A[j]
		}
	}
}

func QuickSortAfter(A []int, p, r int) {

	//判断是不是有序，若有序直接返回？
	left, right := p, r
	for left < right {
		if A[left] > A[left+1] {
			break
		}
		left++
	}

	if left >= right {
		return
	}

	//如果 r-p < 10 进行插入排序
	if r-p < 10 {
		InertSort(A, p, r)
	} else {
		if p < r {
			q := RandPartitionAfter(A, p, r)
			QuickSortAfter(A, p, q-1)
			QuickSortAfter(A, q+1, r)
		}
	}
}

func QuickSortBefore(A []int, p, r int) {

	if p < r {
		q := RandPartitionBefore(A, p, r)
		QuickSortBefore(A, p, q-1)
		QuickSortBefore(A, q+1, r)
	}
}

func RandPartitionBefore(A []int, p int, r int) int {

	// 1、随机数计算pivot的方法
	pivot := rand.Intn(r-p+1) + p
	A[pivot], A[r] = A[r], A[pivot]

	x := A[r]
	i := p - 1

	for j := p; j <= r-1; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func RandPartitionAfter(A []int, p int, r int) int {

	// 1、随机数计算pivot的方法
	pivot := rand.Intn(r-p+1) + p
	A[pivot], A[r] = A[r], A[pivot]

	x := A[r]
	i := p - 1

	for j := p; j <= r-1; j++ {
		if A[j] < x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

// GenerateData for QuickSort 生成大小为n=10^6, 其中包含十分之rate的重复元素的整数数据集
func GenerateData(n, rate int) []int {

	rand.Seed(time.Now().Unix())
	length := n
	n = length - (length / 10 * rate)

	// 生成 1 到 n 总共 n 个数
	data := make([]int, length)
	for i := 0; i < n; i++ {
		data[i] = i + 1
	}

	for i := n; i < length; i++ {
		data[i] = n + 1
	}

	// 随机生成一个随机数，并将其作为下标，使该下标元素与第i个元素交换
	end := length - 1
	for i := 0; i < length; i++ {
		pivot := rand.Intn(end + 1)
		data[end], data[pivot] = data[pivot], data[end]
		end--
	}

	return data
}
