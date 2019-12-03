public class HelloB extends HelloA {
    {
        System.out.println("I'm B construct  code block");
    }

    public HelloB() {
        System.out.println("I'm B construct method code block");
    }

    static {
        System.out.println("I'm B static code block");
    }

    public static void main(String[] args) {
        new HelloB();
    }
}

/*

I'm A static code block
I'm B static code block
I'm A construct code block
I'm A construct method code block
I'm B construct  code block
I'm B construct method code block

*/