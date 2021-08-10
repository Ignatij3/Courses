package main

import "fmt"

type funcDef func(string) string

func fun(s string) string {
	return fmt.Sprintf("from fun: %s", s) // Sprintf() - Возвращает
}

func caller(someFunc funcDef, s string) string {
	return someFunc(s)
}

func main() {
	output := caller(fun, "some string")
	fmt.Println(output)		//	from fun: some string
}
