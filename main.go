package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 2 {
		var input []string
		arg := os.Args[1]
		Searchstdin(input, arg)

	} else if len(os.Args) <= 3 {
		args := os.Args[1]
		filename := os.Args[2]

		found := search(filename, args)
		if found {
			fmt.Printf("%s found in %s \n ", args, filename)
		}
	} else if len(os.Args) >= 4 {
		args := os.Args[1]
		filename := os.Args[2]
		found := search(filename, args)
		if found {
			output := fmt.Sprintf("%s found in %s \n ", args, filename)
			OutputFilename := os.Args[4]
			if OutputFilename != "" {
				go oflag(OutputFilename, output)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}

//function to search string from standard input
func Searchstdin(input []string, arg string) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			input = append(input, text)
		} else {
			break
		}
	}
	for i := range input {
		present := strings.Contains(input[i], arg)
		if present {
			fmt.Println(input[i])
		}
	}
	return nil
}

//function to search string from a file/folder
func search(filename, args string) bool {
	fs, error := ioutil.ReadDir(filename)
	if error != nil {
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		return strings.Contains(string(data), args)
	} else {
		for _, fn := range fs {
			fn := fn.Name()
			fmt.Println(fn)
			// search(fn.Name(), args)
		}
	}
	return false
}

//function to invoke -o option to write in a file specified
func oflag(filename string, data string) error {
	s := flag.String("o", filename, "asdadas")
	flag.Parse()
	if err := os.WriteFile(*s, []byte(data), 0400); err != nil {
		return err
	}
	return nil
}
