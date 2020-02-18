package tasks;

import java.util.HashMap;
import java.util.Stack;

public class Task20 {
    public static
    class Solution {
        public boolean isValid(String s) {
            if ("".equals(s)) {
                return true;
            }

            Stack<Character> stack = new Stack<>();
            HashMap<Character, Character> brackets = new HashMap<>();

            brackets.put(')', '(');
            brackets.put('}', '{');
            brackets.put(']', '[');
            System.out.println("test");
            for (int i = 0; i < s.length(); i++) {
                Character currentChar = s.charAt(i);
                switch (currentChar) {
                    case '(':
                    case '{':
                    case '[':
                        stack.push(s.charAt(i));
                        break;
                    case ')':
                    case '}':
                    case ']':
                        if (stack.size() == 0) {
                            return false;
                        }

                        if (stack.pop() != brackets.get(currentChar)) {
                            return false;
                        }

                        break;
                }
            }

            return stack.size() == 0;
        }
    }
}
