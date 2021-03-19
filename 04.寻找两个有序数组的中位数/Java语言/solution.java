import java.lang.Math;
import java.lang.System;

public class solution {
  public double findMedianSortedArrays(int[] nums1, int[] nums2) {
    int m = nums1.length;
    int n = nums2.length;
    int total = m + n;
    if (total % 2 == 1) {
      int k = total / 2;
      double medain = getKthElement(k, nums1, nums2);
      return medain;
    } else {
      int k1 = total / 2 - 1, k2 = total / 2;
      double medain = (getKthElement(k1, nums1, nums2) + getKthElement(k2, nums1, nums2)) / 2.0;
      return medain;
    }
  }

  public int getKthElement(int k, int[] nums1, int[] nums2) {
    int length1 = nums1.length;
    int length2 = nums2.length;
    int index1 = 0, index2 = 1;
    while (true) {
      // 特殊情况
      if (index1 >= length1) {
        return nums2[index2 + k - 1];
      }
      if (index2 >= length2) {
        return nums1[index1 + k - 1];
      }
      if (k == 1) {
        return Math.min(nums1[index1], nums2[index2]);
      }
      // 正常情况
      int mid = k / 2;
      int newIndex1 = Math.min(index1 + mid, length1) - 1;
      int newIndex2 = Math.min(index2 + mid, length2) - 1;
      if (nums1[newIndex1] <= nums2[newIndex2]) {
        k -= newIndex1 - index1 + 1;
        index1 = newIndex1 + 1;
      } else {
        k -= newIndex2 - index2 + 1;
        index2 = newIndex2 + 1;
      }
    }
  }

  public static void main(String[] args) {
    int[] input1 = new int[] { 1, 2 };
    int[] input2 = new int[] { 3, 4 };

    double res = new solution().findMedianSortedArrays(input1, input2);
    System.out.println("result:" + res);
  }
}
