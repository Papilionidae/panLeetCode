public class Solution {
  public String convert(String s, int numRows) {
    int n = s.length();
    if (n == 1) {
        return s;
    }
    StringBuilder res = new StringBuilder();
    int cycle=2*numRows -2;
    for (int i=0;i<numRows;i++) {
        for (int j=0;i+j<n;j+=cycle) {
            res.append(s.charAt(j+i));
            if (i !=0&&i != numRows-1 && j+cycle-i<n) {
                res.append(s.charAt(j+cycle-i));
            }
        }
    }
    return res.toString();
}
}
