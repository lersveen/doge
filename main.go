package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	flag.Bool("h", false, "Show help")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: ", os.Args[0], "filename")
		os.Exit(1)
	}

	filename, _ := filepath.Abs(os.Args[1])
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("very nope")
		os.Exit(1)
	}

	switch {
	case info.Mode().IsRegular():
		processRegularFile(info, filename)
	case info.Mode()&os.ModeSymlink != 0:
		fmt.Println("very symlink, much wow")
	case info.Mode().IsDir():
		fmt.Println("such directory")
	case info.Mode()&os.ModeSocket != 0:
		fmt.Println("Such socket, many wow")
	case info.Mode()&os.ModeType != 0:
		fmt.Println("Such file, very mystery")
	default:
		fmt.Println("very nope")
		os.Exit(1)
	}
}

func processRegularFile(info os.FileInfo, filename string) {
	switch {
	case info.Mode().Perm()&0400 == 0:
		fmt.Println("such file, no read")
	case info.Size() > 0:
		if info.Mode().Perm()&0100 != 0 {
			absPathFilename, _ := filepath.Abs(os.Args[1])
			absPathProgram, _ := filepath.Abs(os.Args[0])
			if absPathProgram == absPathFilename {
				self()
			} else {
				fmt.Println("much file, very execute")
			}
		} else if info.Mode().Perm()&0200 != 0 {
			fmt.Println("much file, many writes")
		} else {
			fmt.Println("much file")
		}
	default:
		fmt.Println("such file, very empty")
	}
}

func clearScreen() {
	cmd := exec.Command("clear") // Use "cls" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func walkDog() {
	frames := []string{`


                       


           ▄              ▄
          ▌▒█           ▄▀▒▌
          ▌▒▒█        ▄▀▒▒▒▐
         ▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐
       ▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐
     ▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌
    ▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒▌
    ▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐
   ▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌
   ▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌
  ▌▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒▐
  ▐▒▒▐▀▐▀▒░▄▄▒▄▒▒▒▒▒▒░▒░▒░▒▒▒▒▌
  ▐▒▒▒▀▀▄▄▒▒▒▄▒▒▒▒▒▒▒▒░▒░▒░▒▒▐
   ▌▒▒▒▒▒▒▀▀▀▒▒▒▒▒▒░▒░▒░▒░▒▒▒▌
   ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▄▒▒▐
    ▀▄▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▄▒▒▒▒▌
      ▀▄▒▒▒▒▒▒▒▒▒▒▄▄▄▀▒▒▒▒▄▀
        ▀▄▄▄▄▄▄▀▀▀▒▒▒▒▒▄▄▀
           ▒▒▒▒▒▒▒▒▒▒▀▀`, `
  Wow

                       


           ▄              ▄
          ▌▒█           ▄▀▒▌
          ▌▒▒█        ▄▀▒▒▒▐
         ▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐
       ▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐
     ▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌
    ▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒▌
    ▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐
   ▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌
   ▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌
  ▌▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒▐
  ▐▒▒▐▀▐▀▒░▄▄▒▄▒▒▒▒▒▒░▒░▒░▒▒▒▒▌
  ▐▒▒▒▀▀▄▄▒▒▒▄▒▒▒▒▒▒▒▒░▒░▒░▒▒▐
   ▌▒▒▒▒▒▒▀▀▀▒▒▒▒▒▒░▒░▒░▒░▒▒▒▌
   ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▄▒▒▐
    ▀▄▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▄▒▒▒▒▌
      ▀▄▒▒▒▒▒▒▒▒▒▒▄▄▄▀▒▒▒▒▄▀
        ▀▄▄▄▄▄▄▀▀▀▒▒▒▒▒▄▄▀
           ▒▒▒▒▒▒▒▒▒▒▀▀`, `


                       


           ▄              ▄
          ▌▒█           ▄▀▒▌
          ▌▒▒█        ▄▀▒▒▒▐
         ▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐
       ▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐
     ▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌
    ▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒▌
    ▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐    Very doge
   ▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌
   ▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌
  ▌▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒▐
  ▐▒▒▐▀▐▀▒░▄▄▒▄▒▒▒▒▒▒░▒░▒░▒▒▒▒▌
  ▐▒▒▒▀▀▄▄▒▒▒▄▒▒▒▒▒▒▒▒░▒░▒░▒▒▐
   ▌▒▒▒▒▒▒▀▀▀▒▒▒▒▒▒░▒░▒░▒░▒▒▒▌
   ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▄▒▒▐
    ▀▄▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▄▒▒▒▒▌
      ▀▄▒▒▒▒▒▒▒▒▒▒▄▄▄▀▒▒▒▒▄▀
        ▀▄▄▄▄▄▄▀▀▀▒▒▒▒▒▄▄▀
           ▒▒▒▒▒▒▒▒▒▒▀▀`, `

                            Such self
                       


           ▄              ▄
          ▌▒█           ▄▀▒▌
          ▌▒▒█        ▄▀▒▒▒▐
         ▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐
       ▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐
     ▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌
    ▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒▌
    ▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐
   ▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌
   ▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌
  ▌▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒▐
  ▐▒▒▐▀▐▀▒░▄▄▒▄▒▒▒▒▒▒░▒░▒░▒▒▒▒▌
  ▐▒▒▒▀▀▄▄▒▒▒▄▒▒▒▒▒▒▒▒░▒░▒░▒▒▐
   ▌▒▒▒▒▒▒▀▀▀▒▒▒▒▒▒░▒░▒░▒░▒▒▒▌
   ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▄▒▒▐
    ▀▄▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▄▒▒▒▒▌
      ▀▄▒▒▒▒▒▒▒▒▒▒▄▄▄▀▒▒▒▒▄▀
        ▀▄▄▄▄▄▄▀▀▀▒▒▒▒▒▄▄▀
           ▒▒▒▒▒▒▒▒▒▒▀▀`, `


                       
            Much recursion

           ▄              ▄
          ▌▒█           ▄▀▒▌
          ▌▒▒█        ▄▀▒▒▒▐
         ▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐
       ▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐
     ▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌
    ▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒▌
    ▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐
   ▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌
   ▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌
  ▌▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒▐
  ▐▒▒▐▀▐▀▒░▄▄▒▄▒▒▒▒▒▒░▒░▒░▒▒▒▒▌
  ▐▒▒▒▀▀▄▄▒▒▒▄▒▒▒▒▒▒▒▒░▒░▒░▒▒▐
   ▌▒▒▒▒▒▒▀▀▀▒▒▒▒▒▒░▒░▒░▒░▒▒▒▌
   ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▄▒▒▐
    ▀▄▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▄▒▒▒▒▌
      ▀▄▒▒▒▒▒▒▒▒▒▒▄▄▄▀▒▒▒▒▄▀
        ▀▄▄▄▄▄▄▀▀▀▒▒▒▒▒▄▄▀
           ▒▒▒▒▒▒▒▒▒▒▀▀`,
	}

	for _, frame := range frames {
		clearScreen()
		fmt.Println(frame)
		time.Sleep(500 * time.Millisecond)
	}
	clearScreen()
}

func self() {
	walkDog()
	fmt.Println(`Such self`)
}
