// 多线程
import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicInteger;
import java.io.IOException;

class ThreadTest {

    public static void main(String[] args) throws InterruptedException, IOException {
        int corePoolSize = 2;
        int maximumPoolSize = 4;
        long keepAliveTime = 10; //如果线程池当前拥有超过corePoolSize的线程，那么多余的线程在空闲时间超过keepAliveTime时会被终止
        TimeUnit unit = TimeUnit.SECONDS;
        BlockingQueue<Runnable> workQueue = new ArrayBlockingQueue<>(2);
        ThreadFactory threadFactory = new NameTreadFactory(); //新线程使用ThreadFactory创建
        RejectedExecutionHandler handler = new MyIgnorePolicy();//拒绝策略
        ThreadPoolExecutor executor = new ThreadPoolExecutor(corePoolSize, maximumPoolSize, keepAliveTime, unit, workQueue, threadFactory, handler);
        executor.prestartAllCoreThreads(); // 预启动所有核心线程

        for (int i = 1; i <= 10; i++) {
            MyTask task = new MyTask(String.valueOf(i));
            executor.execute(task);//execute(Runnable)方法中提交新任务
        }

        System.in.read(); //阻塞主线程
    }

    static class NameTreadFactory implements ThreadFactory {

        private final AtomicInteger mThreadNum = new AtomicInteger(1);

        @Override
        public Thread newThread(Runnable r) {
            Thread t = new Thread(r, "my-thread-" + mThreadNum.getAndIncrement());
            System.out.println(t.getName() + " has been created");
            return t;
        }
    }

    public static class MyIgnorePolicy implements RejectedExecutionHandler {

        public void rejectedExecution(Runnable r, ThreadPoolExecutor e) {
            doLog(r, e);
        }

        private void doLog(Runnable r, ThreadPoolExecutor e) {
            // 可做日志记录等
            System.err.println( r.toString() + " rejected");
//          System.out.println("completedTaskCount: " + e.getCompletedTaskCount());
        }
    }

    static class MyTask implements Runnable {
        private String name;

        public MyTask(String name) {
            this.name = name;
        }

        @Override
        public void run() {
            try {
                System.out.println(this.toString() + " is running……");
                Thread.sleep(3000); //让任务执行慢点
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            System.out.println(this.toString() + " running end!");
        }

        public String getName() {
            return name;
        }

        @Override
        public String toString() {
            return "MyTask [name=" + name + "]";
        }
    }
}

/*
my-thread-1 has been created
my-thread-2 has been created
my-thread-3 has been created
MyTask [name=2] is running……
MyTask [name=1] is running……
MyTask [name=3] is running……
my-thread-4 has been created
MyTask [name=7] rejected
MyTask [name=6] is running……MyTask [name=8] rejected

MyTask [name=9] rejected
MyTask [name=10] rejected
MyTask [name=1] running end!
MyTask [name=2] running end!
MyTask [name=3] running end!
MyTask [name=5] is running……
MyTask [name=6] running end!
MyTask [name=4] is running……
MyTask [name=5] running end!
MyTask [name=4] running end!
……阻塞
*/