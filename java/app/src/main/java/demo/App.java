/*
 * This Java source file was generated by the Gradle 'init' task.
 */
package demo;

public class App {
    public void IncrDecr() {
        int a = 9;
        int b = a++;
        int c = ++a;
        int d = c--;
        int e = --d;
        System.out.printf("a:%d b:%d c:%d d:%d e:%d\r\n", a, b, c, d, e);

    }



    public static void main(String[] args) {
        App app = new App();
        app.IncrDecr();


    }
}
