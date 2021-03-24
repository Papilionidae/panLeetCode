public class Solution {
  public int threeSumClosest(int[] nums, int target) {
    int n = nums.length;
    if (n == 3) {
      return nums[0] + nums[1] + nums[2];
    }
    Arrays.sort(nums);
    int sum = nums[0] + nums[1] + nums[2];
    int larger = sum, smaller = sum;
    for (int i = 0; i < n - 2; i++) {
      if (i > 1 && nums[i] == target) {
        break;
      }
      int mid = i + 1;
      int right = n - 1;
      while (mid < right) {
        sum = nums[i] + nums[mid] + nums[right];
        if (sum == target) {
          return sum;
        }
        if (sum > target) {
          if (Math.abs(sum - target) < Math.abs(larger - target)) {
            larger = sum;
          }
          right--;
        } else {
          if (Math.abs(target - sum) < Math.abs(target - smaller)) {
            smaller = sum;
          }
          mid++;
        }
      }
    }
    if (Math.abs(larger - target) > Math.abs(target - smaller)) {
      return smaller;
    }
    return larger;
  }
}
