package main

import "fmt"
import "io/ioutil"
import "os/exec"
import "os"
import "syscall"

func main() {

	execing()
	//return

	// 我们将从一个简单的命令开始，没有参数或者输入，仅打印一些信息到标准输出流。
	// exec.Command 函数帮助我们创建一个表示这个外部进程的对象。
	dateCmd := exec.Command("date")

	// .Output 是另一个帮助我们处理运行一个命令的常见情况的函数，它等待命令运行完成，并收集命令的输出。
	// 如果没有出错，dateOut 将获取到日期信息的字节。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 下面我们将看看一个稍复杂的例子，我们将从外部进程的stdin 输入数据并从 stdout 收集结果。
	grepCmd := exec.Command("grep", "hello")

	// 这里我们明确的获取输入/输出管道，运行这个进程，写入一些输入信息，读取输出的结果，最后等待程序运行结束。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// 上面的例子中，我们忽略了错误检测，但是你可以使用if err != nil 的方式来进行错误检查，
	// 我们也只收集StdoutPipe 的结果，但是你可以使用相同的方法收集StderrPipe 的结果。
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 注意，当我们需要提供一个明确的命令和参数数组来生成命令，和能够只需要提供一行命令行字符串相比，
	// 你想使用通过一个字符串生成一个完整的命令，那么你可以使用 bash命令的 -c 选项：
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

func execing() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。我们给 ls 一些基本的参数。注意，第一个参数需要是程序名。
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec 同样需要使用环境变量。
	// 这里我们仅提供当前的环境变量。
	env := os.Environ()

	// 这里是 os.Exec 调用。
	// 如果这个调用成功，那么我们的进程将在这里被替换成 /bin/ls -a -l -h 进程。
	// 如果存在错误，那么我们将会得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
