package demo;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Arr {
    public void ListDeclared() {
        //最简单标准的方式来声明ArrayList
        List<String> myList = new ArrayList<>(Arrays.asList("a", "b", "c"));
        System.out.printf("list declared myList:%s\r\n", myList);
        //stream方式
        String[] strList = {"a", "b", "c"};
        List<String> streamList = Arrays.stream(strList).toList();
        System.out.printf("list declared streamList:%s\r\n", streamList);

        //错误的声明方式, 此时不能执行add,remove,clear等操作，会抛出异常，因为这种声明返回的类是
        //java.util.Arrays的一个内部类，而非java.util.ArrayList
        List<String> errorList = Arrays.asList("a", "b", "c");
        System.out.printf("list declared errorList:%s errorList:%s\r\n", errorList, errorList.getClass());

    }
    public static void main(String[] args) {
        Arr arr = new Arr();
        arr.ListDeclared();
    }
}
