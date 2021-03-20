public class solution {
  public int reverse(int x) {
    int res = 0;
    while (x != 0) {
      int tmp = x % 10;
      x /= 10;
      if (res > Integer.MAX_VALUE / 10 || res == Integer.MAX_VALUE / 10 && tmp > 7) {
        return 0;
      }
      if (res < Integer.MIN_VALUE / 10 || res == Integer.MIN_VALUE / 10 && tmp < -8) {
        return 0;
      }
      res = res * 10 + tmp;
    }
    return res;
  }

  public static void main(String[] args) {

  }
}
