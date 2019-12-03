public class HelloA {
    //构造函数
    public HelloA(){
        System.out.println("I'm A construct method code block");
    }
    //构造代码块——必须在class{}域下的
    {
        System.out.println("I'm A construct code block");
    }

    static {
        System.out.println("I'm A static code block");
    }

    public static void main(String[] args) {
        new HelloA();
        new HelloA();
    }
}

/*
I'm A static code block
I'm A construct code block
I'm A construct method code block
I'm A construct code block
I'm A construct method code block
*/