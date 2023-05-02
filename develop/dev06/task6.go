package dev06

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func Cut() {
	d := flag.String("d", "\t", "Sets field delimiter")
	f := flag.String("f", "", "select only some fields")
	s := flag.Bool("s", false, "only if delimiter is present")
	cutted := make(map[int]bool)
	if *f != "" {
		for _, field := range strings.Split(*f, ",") {
			var left, right int
			if _, err := fmt.Scanf(field, "%d-%d", &left, &right); err == nil {
				for i := left; i <= right; i++ {
					cutted[i] = true
				}
			} else if _, err := fmt.Sscanf(field, "%d", &left); err == nil {
				cutted[left] = true
			} else {
				log.Println("error")
				os.Exit(1)
			}
		}
	}
	var in *bufio.Scanner
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Println("error")
			os.Exit(1)
		}
		defer file.Close()
		in = bufio.NewScanner(file)

	} else {
		in = bufio.NewScanner(os.Stdin)
	}

	for in.Scan() {
		line := in.Text()
		if *s && strings.Contains(line, *d) {
			fmt.Println(line)
			continue
		}
		fields := strings.Split(line, *d)
		var cuttedFields []string
		for i, field := range fields {
			if cutted[i+1] {
				cuttedFields = append(cuttedFields, field)
			}
		}
		fmt.Println(strings.Join(cuttedFields, *d))
	}
	if err := in.Err(); err != nil {
		log.Println("error")
		os.Exit(1)
	}
}
