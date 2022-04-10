package gox

import (
	"fmt"
	"strconv"
	"strings"
)

type Pairs [2]string

var Pairs_Brackets = Pairs([2]string{"{", "}"})
var Pairs_Parenthesis = Pairs([2]string{"(", ")"})

func Tab(code ...string) []string {
	arr := make([]string, len(code))
	for i, s := range code {
		arr[i] = "\t" + s
	}
	return arr
}

func Block(pars Pairs, title string, code ...string) string {
	return title + " " + pars[0] + "\n" + strings.Join(Tab(code...), "\n") + "\n" + pars[1]
}

func Func(name string, recv string, args string, retTypes string, code ...string) string {
	if recv == "" {
		return Block(
			Pairs_Brackets,
			fmt.Sprintf("func %s(%s) (%s)", name, args, retTypes),
			code...,
		)
	}
	return Block(
		Pairs_Brackets,
		fmt.Sprintf("func (%s) %s(%s) (%s)", recv, name, args, retTypes),
		code...,
	)
}

func Struct(name string, fields ...string) string {
	return Block(
		Pairs_Brackets,
		"type "+name+" struct",
		fields...,
	)
}

func Imports(code ...string) string {
	return Block(
		Pairs_Parenthesis,
		"import",
		code...,
	)
}

func Interface(name string, code ...string) string {
	return Block(
		Pairs_Brackets,
		"type "+name+" interface",
		code...,
	)
}

func Enum(name string, enumVars ...string) string {
	ens := make([]string, len(enumVars))
	ecount := 1
	for i, en := range enumVars {
		ens[i] = "var " + en + " = " + name + "(" + strconv.Itoa(ecount) + ")"
		ecount += 1
	}
	return strings.Join([]string{
		"type " + name + " int",
		"",
		strings.Join(ens, "\n"),
	}, "\n")
}

func CodeBlock(code ...string) string {
	return Block(Pairs_Brackets, "", code...)
}
