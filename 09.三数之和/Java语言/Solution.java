public class Solution {
  public List<List<Integer>> threeSum(int[] nums) {
    int n = nums.length;
    List<List<Integer>> list = new ArrayList<List<Integer>>();
    if (n < 3) {
      return list;
    }
    Arrays.sort(nums);
    for (int left = 0; left < n - 2; left++) {
      if (nums[left] > 0) {
        return list;
      }
      if (left > 0 && nums[left] == nums[left - 1]) {
        continue;
      }
      int mid = left + 1;
      int right = n - 1;
      while (mid < right) {
        if (nums[left] + nums[mid] + nums[right] > 0) {
          right--;
        } else if (nums[left] + nums[mid] + nums[right] < 0) {
          mid++;
        } else {
          list.add(Arrays.asList(nums[left], nums[mid], nums[right]));
          mid++;
          right--;
          while (mid < right && nums[mid] == nums[mid - 1]) {
            mid++;
          }
          while (mid < right && nums[right] == nums[right + 1]) {
            right--;
          }
        }
      }
    }
    return list;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    int[] nums = new int[] { -1, 0, 1, 2, 4, -1, 3 };
    List<List<Integer>> list = solution.threeSum(nums);
    System.out.println("list:", list);
  }
}
