package ssh

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"testing"
)

func Test1(t *testing.T) {
	session, err := connect("root", "xxxxx", "127.0.0.1", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Run("ls /; ls /abc")
}

//输出到命令行
func TestStdout(t *testing.T) {
	session, err := connect("root", "xxxxx", "127.0.0.1", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Run("ls /; ls /abc")
}

// 交互式命令
func TestInteractive(t *testing.T) {
	session, err := connect("root", "olordjesus", "dockers.iotalabs.io", 2210)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	// excute command
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		log.Fatal(err)
	}

	session.Run("top")
}
