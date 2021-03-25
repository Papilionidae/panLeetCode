import java.util.Arrays;
import java.util.List;

public class Solution {
  public List<List<Integer>> fourSum(int[] nums, int target) {
    List<List<Integer>> list = new ArrayList<List<Integer>>();
    int n = nums.length;
    if (n < 4) {
      return list;
    }
    Arrays.sort(nums);
    for (int i = 0; i < n - 3; i++) {
      if (i > 0 && nums[i] == nums[i - 1]) {
        continue;
      }
      if (nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target) {
        break;
      }
      if (nums[i] + nums[n - 3] + nums[n - 2] + nums[n - 1] < target) {
        continue;
      }
      for (int j = i + 1; j < n - 2; j++) {
        if (j > i + 1 && nums[j] == nums[j - 1]) {
          continue;
        }
        if (nums[i] + nums[j] + nums[j + 1] + nums[j + 2] > target) {
          break;
        }
        if (nums[i] + nums[j] + nums[n - 2] + nums[n - 1] < target) {
          continue;
        }
        int left = 0, right = 0;
        while (left < right) {
          if (nums[i] + nums[j] + nums[left] + nums[right] < target) {
            left++;
          } else if (nums[i] + nums[j] + nums[left] + nums[right] > target) {
            right--;
          } else {
            list.add(Arrays.asList(nums[i], nums[j], nums[left], nums[right]));
            left++;
            right--;
            while (left < right && nums[left] == nums[left - 1]) {
              left++;
            }
            while (left < right && nums[right] == nums[right + 1]) {
              right--;
            }
          }
        }
      }
    }
    return list;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    int[] nums = new int[] { -1, 0, 1, 2, 4, -1, 3 };
    List<List<Integer>> list = solution.fourSum(nums, 1);
    System.out.println("list:", list);
  }
}
