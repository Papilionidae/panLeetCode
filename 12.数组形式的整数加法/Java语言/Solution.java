import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collector;

public class Solution {
  public List<Integer> addToArrayForm(int[] A, int K) {
    List<Integer> list = new ArrayList<>();
    int n = A.length;
    for (int i = n - 1; i >= 0 || K != 0; i--, K /= 10) {
      if (i >= 0) {
        K += A[i];
      }
      list.add(K % 10);
    }
    Collections.reverse(list);
    return list;
  }
}
