package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {
	re := regexp.MustCompile(`^ex-\S*[^_][^t][^e][^s][^t].go$`)
	match := re.MatchString(path)
	if match {
		tmp := strings.Split(path, ".")
		cmd := exec.Command("go", "build", "-tags", "local", "-o", "./"+tmp[len(tmp)-2], "./"+path)
		fmt.Println(cmd)
		out, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		fmt.Print(string(out))

	}

	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
