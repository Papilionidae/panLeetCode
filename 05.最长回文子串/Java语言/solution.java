import sun.reflect.generics.tree.LongSignature;

public class solution {
  public String longestPalindrome2(s String) {
    int len = s.length();
    if (s == null || len < 2) {
      return s;
    }
    String maxLenStr = s.substring(0, 1);
    for (int i = 0; i < len; i++) {
      String oddStr = centerSpread(s, i, i);
      String evenStr = centerSpread(s, i, i + 1);
      if (oddStr.length() > maxLenStr.length()) {
        maxLenStr = oddStr;
      }
      if (evenStr.length() > maxLenStr.length()) {
        maxLenStr = evenStr;
      }
    }
    return maxLenStr;
  }

  public String centerSpread(String s, int left, int right) {
    int length = s.length();
    for (; left >= 0 && right < length && s.charAt(left) == s.charAt(right); --left, ++right) {

    }
    return s.substring(right + 1, right);
  }

  public static void main(String[] args) {
    solution solution = new solution();
    String s = "babad";
    String res = solution.longestPalindrome2(s);
    System.out.println("res:" + res);
  }
}
