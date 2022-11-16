package medianoftwosortedarrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	type arg struct {
		nums1 []int
		nums2 []int
	}

	testCase := []arg{
		{nums1: []int{1, 3}, nums2: []int{2}},
		{nums1: []int{1, 2}, nums2: []int{3, 4}},
		{nums1: []int{}, nums2: []int{3, 4, 5}},
		{nums1: []int{4, 5, 6}, nums2: []int{}},
		{nums1: []int{2, 3, 4, 5}, nums2: []int{1}},
		{nums1: []int{1}, nums2: []int{2, 3, 4, 5, 6, 7}},
		{nums1: []int{1, 2, 3, 4, 5, 6, 8, 9, 23, 34}, nums2: []int{7}},
	}

	expected := []float64{2, 2.5, 4, 5, 3, 4, 6}

	for idx, data := range testCase {
		if res := findMedianSortedArrays(data.nums1, data.nums2); res != expected[idx] {
			t.Errorf("expected %f, got %f", expected[idx], res)
		}
	}
}
