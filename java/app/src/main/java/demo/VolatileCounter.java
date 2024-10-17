package demo;

import java.util.concurrent.CountDownLatch;

public class VolatileCounter {
    public volatile static int num = 0;
    static CountDownLatch countDownLatch = new CountDownLatch(50);
    public static void main(String[] args) throws InterruptedException {
        for (int i = 0; i < 50; i++) {
            new Thread(() -> {
                for (int j = 0; j < 10000; j++) {
                    num++; //volatile 无法保证原子操作
                }
                countDownLatch.countDown();
            }).start();
        }

        countDownLatch.await();
        System.out.println(num);
    }
}
