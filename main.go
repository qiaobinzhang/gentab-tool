package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// O use -o save result to file
var O = flag.String("o", "", "use -o save result to file")

// F use -f to specify a file
var F = flag.String("f", "", "use -f to specify a file")

// D use -d to specify a separator
var D = flag.String("d", "", "use -d to specify a separator")

// T use "\t" as separator
var T = flag.Bool("t", false, "use '\\t' as separator")

// C centering control
var C = flag.Bool("c", false, "centering control")

func main() {
	flag.Parse()
	var s []string
	var outstr string
	buf, err := ioutil.ReadFile(*F)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		os.Exit(-1)
	}
	str := string(buf)
	s = strings.Split(str, "\n")
	for i := 0; i < len(s); i++ {
		isEmpty := strings.TrimSpace(s[i])
		if isEmpty == "" {
			continue
		}
		var ss []string
		if *T {
			ss = strings.Split(s[i], "\t")
		} else if *D != "" {
			ss = strings.Split(s[i], *D)
		} else {
			s[i] = strings.Replace(s[i], "\t", " ", -1)
			ss = strings.Split(s[i], " ")
		}
		s[i] = ""
		for j := 0; j < len(ss); j++ {
			if ss[j] == "" {
				continue
			}
			if *C {
				ss[j] = "<td align=center>" + ss[j] + "</td>"
			} else {
				ss[j] = "<td>" + ss[j] + "</td>"
			}
			s[i] = s[i] + ss[j]
		}
		s[i] = "<tr>" + s[i] + "</tr>"
	}
	outstr = "<style type='text/css'>table{border-collapse: collapse;}  </style><table border=1>"
	for i := 0; i < len(s); i++ {
		outstr = outstr + s[i]
	}
	outstr = outstr + "</table>"
	if *O != "" {
		ioutil.WriteFile(*O, []byte(outstr), 0644)
	} else {
		fmt.Println(outstr)
	}
}
