package medianoftwosortedarrays

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX, minRightX := math.MinInt32, math.MaxInt32
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}
		if partitionX < x {
			minRightX = nums1[partitionX]
		}

		maxLeftY, minRightY := math.MinInt32, math.MaxInt32
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}
		if partitionY < x {
			minRightY = nums2[partitionY]
		}

		if maxLeftY <= minRightY && maxLeftY <= minRightX {
			max := math.Max(float64(maxLeftX), float64(maxLeftY))
			min := math.Min(float64(minRightX), float64(minRightY))

			if (x+y)&1 != 0 {
				return max
			}

			return (max + min) / 2
		}

		if maxLeftX > minRightY {
			high = partitionX - 1
			continue
		}

		low = partitionX + 1
	}
	return -1
}
