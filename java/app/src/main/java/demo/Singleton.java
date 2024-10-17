package demo;

public class Singleton {

    private volatile static Singleton instance = null;
    private Singleton() {

    }

    public static Singleton getInstance() {
        /*
         * instance类变量如果没有用volatile关键字修饰的，会导致这样一个问题：代码读取到instance不为null时，
         * instance引用的对象有可能还没有完成初始化。造成这种现象主要的原因是重排序，
         * 重排序是指编译器和处理器为了优化程序性能而对指令序列进行重新排序的一种手段。
         * 如下是赋值
         * emory = allocate();　　// 1：分配对象的内存空间
         * ctorInstance(memory);　// 2：初始化对象
         * instance = memory;　　// 3：设置instance指向刚分配的内存地址
         * 如果上面的指令指向到2和3之间被重排过了，就有可能在 if (instance == null) {这条语句
         * 在某个线程获取到的instance值是非空且未初始化。
         * */
        if (instance == null) {
            synchronized (Singleton.class) {
                if (instance == null) {
                    instance = new Singleton();
                }
            }
        }

        return instance;
    }
}
