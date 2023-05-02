package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	pattern := flag.String("pattern", "", "the pattern to search for")
	A := flag.Int("A", 0, "print N lines of trailing context after a match")
	B := flag.Int("B", 0, "print N lines of leading context before a match")
	C := flag.Int("C", 0, "print N lines of output context")
	c := flag.Bool("c", false, "print only a count of matching lines")
	F := flag.Bool("F", false, "interpret the pattern as a literal string")
	i := flag.Bool("i", false, "perform case-insensitive matching")
	n := flag.Bool("n", false, "prefix each line of output with its line number")
	v := flag.Bool("v", false, "select non-matching lines")
	flag.Parse()

	if *pattern == "" {
		fmt.Println("You must provide a pattern to search for using the '-pattern' flag")
		os.Exit(1)
	}

	input := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	matches := []string{}

	// сначала просто находим совпадения для F и для i
	for _, line := range input {
		var re *regexp.Regexp
		if *F {
			re = regexp.MustCompile(regexp.QuoteMeta(*pattern))
		} else {
			if *i {
				re = regexp.MustCompile("(?i)" + *pattern)
			} else {
				re = regexp.MustCompile(*pattern)
			}
		}
		if *v != re.MatchString(line) {
			matches = append(matches, line)
		}
	}

	// Кол-во совпадений из прошлого подсчета
	if *c {
		fmt.Println(len(matches))
	} else {
		// счтаем остальные флаги
		for i, line := range matches {

			if *n {
				fmt.Printf("%d:", i+1)
			}
			fmt.Println(line)

			// совпадения под флаг А
			if *A > 0 {
				for j := i + 1; j <= i+*A && j < len(matches); j++ {
					fmt.Println(matches[j])
				}
			}
		}

		for idx, line := range input {
			if *F {
				if strings.Contains(line, *pattern) {
					continue
				}
			} else {
				var re *regexp.Regexp
				if *i {
					re = regexp.MustCompile("(?i)" + *pattern)
				} else {
					re = regexp.MustCompile(*pattern)
				}
				if *v == re.MatchString(line) {
					continue
				}
			}

			if *C > 0 {
				if idx > *C {
					break
				}
			} else if *B > 0 {
				if idx <= len(input)-*B {
					continue
				}
			}

			if *n {
				fmt.Printf("%d:", idx+1)
			}
			fmt.Println(line)

			if *B > 0 {
				for j := 1; j <= *B && idx-j >= 0; j++ {
					fmt.Println(input[idx-j])
				}
			}
		}
	}
}
