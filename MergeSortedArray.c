void merge(int* nums1, int nums1Size, int m, int* nums2, int nums2Size, int n){
    int all = m + n - 1;
    while(n > 0) nums1[ all-- ] = (m > 0 && nums1[m-1] >= nums2[n-1] ? nums1[--m] : nums2[--n]);
}
