import java.util.HashMap;

class towNumsSum {
  /**
   * 方法一：暴力法
   */
  public static int[] twoSum1(int[] nums, int target) {
    for (int i = 0; i < nums.length - 1; i++) {
      for (int j = i + 1; j < nums.length; j++) {
        if (nums[i] + nums[j] == target) {
          return new int[] { i, j };
        }
      }
    }
    return new int[] {};
  }

  /**
   * 方法二： 哈希表法
   */
  public static int[] twoSum2(int[] nums, int target) {
    HashMap<Integer, Integer> map = new HashMap<>();
    for (int i = 0; i < nums.length; i++) {
      if (map.containsKey(target - nums[i])) {
        return new int[] { map.get(target - nums[i]), i };
      }
      map.put(nums[i], i);
    }
    return new int[] {};
  }

  public static void main(String[] args) {
    int[] nums = { 2, 7, 11, 15 };
    int target = 18;
    int[] result1 = twoSum1(nums, target);
    System.out.println("[" + result1[0] + "," + result1[1] + "]");
    int[] result2 = twoSum2(nums, target);
    System.out.println("[" + result2[0] + "," + result2[1] + "]");
  }
}
