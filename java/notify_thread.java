import java.util.concurrent.TimeUnit;
class Communication {

    public static void main(String[] args) {
        Student s=new Student();

        //设置和获取的类
        setThread st=new setThread(s);
        getThread gt=new getThread(s);

        Thread t1=new Thread(st,"t1");
        Thread t2=new Thread(gt,"t2");
        Thread t3=new Thread(st,"t3");
        // Thread t4=new Thread(gt);
        System.out.println("t1："+t1.toString());
        System.out.println("t2："+t2.toString());
        System.out.println("t3："+t3.toString());

        t1.start();
        t2.start();
        t3.start();
    }

}

class Student{
    private String name;
    private int age;
    private int x;
    private boolean flag;//默认为没有数据

    public synchronized void set(String name,int age,int x){
        //如果有数据就等着
        if(this.flag){
            try{
                System.out.println("set wait for:"+this.flag);
                this.wait();//wait是指在一个已经进入了同步锁的线程内，让自己暂时让出同步锁，
                //以便其他正在等待此锁的线程可以得到同步锁并运行
            }catch(InterruptedException e){
                e.printStackTrace();
            }
        }

        this.name=name;
        this.age=age;
        this.x=x;

        this.flag=true;
        this.notify();//用于通知一个处在wait状态的线程，并使他获得该对象上的锁，
    }

    public synchronized void get(){
        //如果没有数据就等着
        if(!this.flag){
            try{
                System.out.println("get wait for:"+this.flag);
                this.wait();
            }catch(InterruptedException e){
                e.printStackTrace();
            }
        }

        System.out.println("获取名字是"+this.name+"---"+this.age+"-----x:"+this.x);

        this.flag=false;
        this.notify();
    }

}

class setThread implements Runnable{
    private Student s;
    private int x=0;
    private final String name="setThread";

    public setThread(Student s){
        this.s=s;
    }
    public void run(){
        while(true){
            try{
                TimeUnit.SECONDS.sleep(1);
            }catch(InterruptedException e){
                System.out.println(e);
            }

            if(x%2==1){
                x++;
                System.out.println("正在运行的线程名称："+this.name+":"+Thread.currentThread().getName()+" 开始林青霞设置线程x:"+x);
                s.set("林青霞", 23,x);
            }else{
                x++;
                System.out.println("正在运行的线程名称："+this.name+":"+Thread.currentThread().getName()+" 开始东方不败设置线程x:"+x);
                s.set("东方不败",44,x);
            }
        }
    }
}

class getThread implements Runnable{
    private Student s;
    private final String name="getThread";
    public getThread(Student s){
        this.s=s;
    }
    public void run(){
        while(true){
            // Thread.currentThread().getName();
            try{
                TimeUnit.SECONDS.sleep(1);
            }catch(InterruptedException e){
                System.out.println(e);
            }

            System.out.println("正在运行的线程名称："+this.name+":"+Thread.currentThread().getName()+" 开始");
            s.get();
            System.out.println("\n\n");
        }
    }
}

/*
t1：Thread[Thread-0,5,main]
t2：Thread[Thread-1,5,main]
t3：Thread[Thread-2,5,main]
正在运行的线程名称：getThread:Thread-1 开始
get wait for:false
正在运行的线程名称：setThread:Thread-0 开始林青霞设置线程x:2
正在运行的线程名称：setThread:Thread-2 开始东方不败设置线程x:2
获取名字是林青霞---23-----x:2



正在运行的线程名称：setThread:Thread-0 开始东方不败设置线程x:3
set wait for:true
正在运行的线程名称：getThread:Thread-1 开始
获取名字是东方不败---44-----x:2
正在运行的线程名称：setThread:Thread-2 开始林青霞设置线程x:4
set wait for:true



正在运行的线程名称：setThread:Thread-0 开始东方不败设置线程x:5
set wait for:true
正在运行的线程名称：getThread:Thread-1 开始
获取名字是东方不败---44-----x:3



正在运行的线程名称：setThread:Thread-0 开始林青霞设置线程x:6
set wait for:true
正在运行的线程名称：setThread:Thread-2 开始东方不败设置线程x:7
set wait for:true
正在运行的线程名称：getThread:Thread-1 开始
获取名字是东方不败---44-----x:5



正在运行的线程名称：getThread:Thread-1 开始
获取名字是东方不败---44-----x:7



正在运行的线程名称：setThread:Thread-2 开始林青霞设置线程x:8
正在运行的线程名称：setThread:Thread-0 开始林青霞设置线程x:8
set wait for:true
正在运行的线程名称：getThread:Thread-1 开始
获取名字是林青霞---23-----x:8



正在运行的线程名称：setThread:Thread-2 开始东方不败设置线程x:9
set wait for:true
正在运行的线程名称：setThread:Thread-0 开始林青霞设置线程x:10
set wait for:true
正在运行的线程名称：getThread:Thread-1 开始
获取名字是林青霞---23-----x:8




*/