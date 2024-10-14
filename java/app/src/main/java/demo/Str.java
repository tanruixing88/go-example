package demo;

public class Str {

    public void Equal() {
        String str1 = "hello";
        String str2 = new String("hello");
        String str3 = "hello";
        // 使用 == 比较字符串的引用相等
        System.out.printf("str1:%s str2:%s str1 == str2:%b \r\n", str1, str2, str1 == str2);
        System.out.println(str1 == str3);
        // 使用 equals 方法比较字符串的相等
        System.out.println(str1.equals(str2));
        System.out.println(str1.equals(str3));

        // java 字面量数字跟浮点数标记是一样的
        System.out.printf("42 == 42.0 :%b\r\n", 42 == 42.0);
        System.out.printf("1000 == 1000.0 :%b\r\n", 1000 == 1000.0);

        //常量折叠,jvm会把确定的常量存放到常量池里
        String str4 = "str" + "ing";
        String str5 = "string";
        //此时是true
        System.out.printf("str4:%s str5:%s ==:%b \r\n", str4, str5, str4 == str5);
        final String str6 = "str";
        final String str7 = "ing";
        final String str8 = str6+str7;
        final String str9 = "string";
        //经过final修饰的也是true，不修饰就是false
        System.out.printf("str6:%s str7:%s str8==str9:%b \r\n", str6, str7, str8 == str9);
    }


    public static void main(String[] args) {
        Str str = new Str();
        str.Equal();
    }
}
