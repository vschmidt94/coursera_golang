package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func parseField(f string) string {
	if len(f) > 20 {
		f = f[:20]
	}

	return f
}

func (n Name) printName() {
	fmt.Println(n.fname, n.lname)
}

func printNames(na []Name) {
	fmt.Println("All the names in the file")
	for _, n := range na {
		n.printName()
	}
}

func getFile() *os.File {
	stdin := bufio.NewReader(os.Stdin)
	cwd, _ := os.Getwd()

	fmt.Println("The current working directory is: ", cwd)
	fmt.Print("Enter filename (relative to CWD): ")
	fn, err := stdin.ReadString('\n')
	if err != nil {
		fmt.Println(" > ERROR: Could not parse file name.")
		return nil
	}
	fn = strings.TrimSuffix(fn, "\n")

	var f *os.File
	f, err = os.Open(fn)
	if err != nil {
		fmt.Printf(" > ERROR opening file %s: %s", fn, err)
	}

	return f
}

func main() {
	names := make([]Name, 0)

	f := getFile()
	if f == nil {
		os.Exit(1)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			if err == io.EOF {
				fmt.Println("End of file - stopping")
				break
			}

			// fmt.Println("ERROR reading line in file, skip this line and try next one: ", err)
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("Error reading file - line does not contain expected 2 fields, skipping this line")
			break
		}

		n := Name{fname: parseField(fields[0]), lname: parseField(fields[1])}
		names = append(names, n)
	}

	printNames(names)
}