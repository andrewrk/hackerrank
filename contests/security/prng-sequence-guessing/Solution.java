import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {
  public static void main(String[] args) {
    Scanner stdin = new Scanner(System.in);
    int testCaseCount = stdin.nextInt();
    for (int testCaseIndex = 0; testCaseIndex < testCaseCount; testCaseIndex += 1) {
      int startTime = stdin.nextInt();
      int endTime = stdin.nextInt();
      int[] values = new int[10];

      for (int j = 0; j < 10; j += 1) {
        values[j] = stdin.nextInt();
      }

      for (int seed = startTime; seed <= endTime; seed += 1) {
        Random rand = new Random(seed);
        boolean bad = false;
        for (int valueIndex = 0; valueIndex < values.length; valueIndex += 1) {
          if (rand.nextInt(1000) != values[valueIndex]) {
            bad = true;
            break;
          }
        }
        if (!bad) {
          System.out.print(seed);
          System.out.print(" ");
          for (int i = 0; i < 10; i += 1) {
            System.out.print(rand.nextInt(1000));
            System.out.print(" ");
          }
          System.out.print("\n");
        }
      }

    }
  }
}
