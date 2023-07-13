func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := append(nums1, nums2...)
    result := 0.0

    sort.Ints(merged)

    if len(merged) % 2 == 1 {
        result = float64(merged[(len(merged)-1)/2])
    } else {
        result = (float64(merged[(len(merged)-1)/2]) + float64(merged[(len(merged))/2]))/2
    }
    return result
}