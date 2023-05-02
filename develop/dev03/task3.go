package dev03

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Sort() {
	k := flag.Int("k", -1, "Sort based on column K")
	n := flag.Bool("n", false, "Sort numerically")
	r := flag.Bool("r", false, "Reverse sort")
	u := flag.Bool("u", false, "Remove duplicates")
	flag.Parse()

	var in []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}
	sort.Slice(in, func(i, j int) bool {
		// разбирваем на колонки
		if *k != -1 {
			left := strings.Fields(in[i])
			right := strings.Fields(in[j])
			if *k <= len(left) && *k <= len(right) {
				in[i] = left[*k-1]
				in[j] = right[*k-1]
			}
		}
		// конвертим в инты
		if *n {
			n1, err1 := strconv.Atoi(in[i])
			n2, err2 := strconv.Atoi(in[j])
			if err1 == nil && err2 == nil {
				in[i] = strconv.Itoa(n1)
				in[j] = strconv.Itoa(n2)
			}
		}
		// меняем местами
		if *r {
			i, j = j, i
		}
		return in[i] < in[j]
	})

	// при помощи мапы оставляем уникальные строки
	if *u {
		var out []string
		checked := make(map[string]bool)
		for _, line := range in {
			if !checked[line] {
				checked[line] = true
				out = append(out, line)
			}
		}
		in = out
	}
	// вывод
	for _, line := range in {
		fmt.Println(line)
	}
}
