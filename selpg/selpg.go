package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/pflag"
)

// CliArgs todo
type CliArgs struct {
	startPage      int
	endPage        int
	lineNumPerPage int
	isFtype        bool
	printDest      string
	inFilename     string
}

// const usage string = "Usage: selpg <-s> <-e> [-f|-l] [-d] file_name"
const specialNum = -10086
const defaultLineNum = 72

func main() {

	args := new(CliArgs)

	args.init()

	args.checkArgs()

	args.execCommand()
}

func (args *CliArgs) init() {
	pflag.IntVarP(&args.startPage, "start", "s", 0, "define the start page number")
	pflag.IntVarP(&args.endPage, "end", "e", 0, "end the start page number")
	// define the line number per page
	pflag.IntVarP(&args.lineNumPerPage, "linenum", "l", specialNum, "specify the number of line per page")
	pflag.StringVarP(&args.printDest, "printdes", "d", "", "define the printer name")
	pflag.BoolVarP(&args.isFtype, "othertype", "f", false, "define the page type, -f mean the pages are seperated  by special symbol")
	pflag.Parse()
}

// checkArgs , check the valiadtion of all arguments
func (args *CliArgs) checkArgs() {
	// print all filed of the object
	if !(args.startPage > 0 && args.endPage > 0 && args.endPage-args.startPage >= 0) {
		fmt.Fprintf(os.Stderr, "start page and end page should be positive and endpage should be bigger than startpage")
		os.Exit(1)
	}

	if args.isFtype {
		if args.lineNumPerPage != specialNum {
			fmt.Fprintln(os.Stderr, "Fatal: setting -f and -l simultaneously is not allowed")
			os.Exit(1)
		}
	} else {
		if args.lineNumPerPage == specialNum {
			args.lineNumPerPage = defaultLineNum
		} else if args.lineNumPerPage < 0 {
			fmt.Fprintln(os.Stderr, "Fatal: the linenum should be positive")
			os.Exit(1)
		}
	}

	// set default number of linenumber
	if pflag.NArg() != 0 {
		args.inFilename = pflag.Args()[0]
	}

	fmt.Printf("%+v", args)
}

func (args *CliArgs) execCommand() {
	// stderr := os.Stderr
	// open the file

	// initialize the input, output
	input := os.Stdin
	output := os.Stdout
	errorOutput := os.Stderr
	var inpipe io.WriteCloser
	var err error
	havePrinter := false
	curLineIdx := 0
	curPageIdx := 0

	if args.inFilename != "" {
		// can't not define a new variable here, must using global input and error
		// if don't have input file, then get input from keyboard
		input, err = os.Open(args.inFilename)
		if err != nil {
			// log.Fatal(err)
			fmt.Fprint(output, err)
			fmt.Fprint(output, err)
			os.Exit(1)
		}
		defer input.Close()
	}

	// use pipe write data to detinaltion
	if args.printDest != "" {
		cmd := exec.Command("echo", "output direct to destination : "+args.printDest)
		inpipe, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		havePrinter = true
		// the command will be executed when the pipe is closed
		defer inpipe.Close()
		cmd.Stdout = output
		cmd.Start()
	}

	if args.isFtype {
		// read as pages
		// using bufio.Reader wrap the original reader to support reading with buffer
		rd := bufio.NewReader(input)
		for {
			page, ferr := rd.ReadString('\f')
			if ferr != nil || ferr == io.EOF || curPageIdx > args.endPage {
				if ferr == io.EOF {
					// output the remaining content
					if curPageIdx >= args.startPage && curPageIdx <= args.endPage {
						output.Write([]byte(page))
						if havePrinter {
							fmt.Fprint(inpipe, page)
						}
					}
				}
				break
			}
			page = strings.Replace(page, "\f", "", -1)
			curPageIdx++
			if curPageIdx >= args.startPage && curPageIdx <= args.endPage {
				fmt.Fprint(output, page)
				if havePrinter {
					fmt.Fprint(inpipe, page)
				}
			}
			fmt.Printf("log info : page = %d", curPageIdx)
		}
	} else {
		// read the pages seperated by '\f'
		// how can I read from start to the end
		scanner := bufio.NewScanner(input)
		fmt.Fprintln(output, "test for output")
		for scanner.Scan() {
			// Scan function return false when either error happen or encounter a EOF
			// if the page is what your want, then oputput every line in this page
			curline := scanner.Text()
			// fmt.Fprintf(output, curline)
			if curPageIdx >= args.startPage && curPageIdx <= args.endPage {
				if curPageIdx > args.endPage {
					break
				}
				fmt.Fprintln(output, curline)
				if havePrinter {
					fmt.Fprint(inpipe, curline)
				}
			}
			curLineIdx++
			if curLineIdx%args.lineNumPerPage == 0 {
				curPageIdx++
			}
		}

		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
		}
	}

	if curPageIdx < args.endPage {
		fmt.Fprintf(errorOutput, "invaild end page, which is bigger than the total page")
	}

}
