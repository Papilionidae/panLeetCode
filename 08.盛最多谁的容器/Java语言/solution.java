public class solution {
  public int maxArea(int[] height) {
    int n = height.length;
    int maxArea = 0;
    for (int left = 0, right = n - 1; left < right;) {
      int minHeight = Math.min(height[left], height[right]);
      if (minHeight * (right - left) > maxArea) {
        maxArea = minHeight * (right - left);
      }
      if (minHeight == height[left]) {
        left++;
      } else {
        right--;
      }
    }
    return maxArea;
  }

  public static void main(String[] args) {
    solution solution = new solution();
    int[] height = new int[] { 1, 8, 7, 6, 7, 4, 3, 7, 8 };
    int maxArea = solution.maxArea(height);
    System.out.println("maxArea:" + maxArea);
  }
}
