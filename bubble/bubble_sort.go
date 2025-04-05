package bubble

// Sort uses Bubble Sort algorithm to sort an array of integers in ascending order
// Time Complexity:
//   Best Case: O(n)       -> when the array is already sorted (can be optimized with a flag)
//   Average Case: O(n^2)  -> typical case with unordered elements
//   Worst Case: O(n^2)    -> when the array is in reverse order
//
// Space Complexity:
//   O(1) -> in-place sorting, uses a constant amount of extra memory
func Sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return arr
}
