package Solution;

public class Solution {
  public boolean isPalindrome(int x) {
    if (x < 0 || (x%10==0 && x!=0)) {
      return false;
    }
    int rev = 0;
    while( x > rev) {
      rev = rev *10 + x%10;
      x /= 10;
    }
    return x== rev || x==rev/10;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    int x = 12321;
    int res = solution.isPalindrome(x);
    System.out.println("res:"+res);
  }
}
