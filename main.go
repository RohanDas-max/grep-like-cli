package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 2 {
		var input []string
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
			present := strings.Contains(input[i], os.Args[1])
			if present {
				fmt.Println(input[i])
			}
		}

	} else if len(os.Args) <= 3 {
		args := os.Args[1]
		filename := os.Args[2]

		searched, fn := search(filename, args)
		if searched != "" {
			fmt.Printf("%s found in %s \n ", searched, fn)
		}
	} else if len(os.Args) >= 4 {
		args := os.Args[1]
		filename := os.Args[2]
		var data string

		searched, fn := search(filename, args)
		if searched != "" {
			fmt.Printf("%s found in %s \n ", searched, fn)
		}

		OutputFilename := os.Args[4]
		if OutputFilename != "" && data != "" {
			go oflag(OutputFilename, data)
			time.Sleep(time.Millisecond)
		}
	}
}

func search(filename, args string) (string, string) {
	var searched string
	arg := strings.Split(args, " ")
	data, err := os.ReadFile(filename)

	if err != nil {
		fs, err := os.ReadDir(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			var data []byte
			var err error
			for _, fn := range fs {
				ds := fn.Name()
				fmt.Println(data)
				data, err = os.ReadFile(ds)
				if err != nil {
					log.Fatal(err)
				}
			}
			fmt.Println(string(data))
			d := strings.Split(string(data), " ")
			fmt.Println(d)
			fmt.Println(err)
			// 	for i := range d {
			// 		for j := range arg {
			// 			if arg[j] == d[i] {
			// 				// fmt.Println(arg[j], d[i])
			// 			}
			// 		}
			// 	}

		}
	}

	d := strings.Split(string(data), " ")
	for i := range d {
		for j := range arg {
			if arg[j] == d[i] {
				searched = arg[j]
			}
		}
	}

	return searched, filename
}

func oflag(args string, data string) string {
	s := flag.String("o", args, "asdadas")
	flag.Parse()

	if err := os.WriteFile(*s, []byte(data), 0400); err != nil {
		fmt.Println(err)
	}
	return *s
}
