package tasks;


import java.util.Arrays;
import java.util.List;

public class Task13 {
    public static
    class Solution {
        public int romanToInt(String s) {
            s = s.toUpperCase();
            int i = 0;
            int result = 0;
            List<RomanNumeral> romans = RomanNumeral.asList();
            while (s.length() > 0 && i < romans.size()) {
                RomanNumeral r = romans.get(i);
                if (s.startsWith(r.name())) {
                    result += r.value;
                    s = s.substring(r.name().length());
                } else {
                    i++;
                }
            }

            if (s.length() == 0) {
                return result;
            }

            return 0;
        }

        enum RomanNumeral {

            I(1),
            V(5),
            IV(4),
            IX(9),
            X(10),
            XL(40),
            L(50),
            XC(90),
            C(100),
            CD(400),
            D(500),
            CM(900),
            M(1000);

            private final int value;

            private RomanNumeral(int value) {
                this.value = value;
            }

            public static List<RomanNumeral> asList() {
                return Arrays.asList(M, CM, D, CD, C, XC, L, XL, X, IX, V, IV, I);

            }
        }
    }
}
