package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
)

func main() {
	File, err := os.Create("Simple.go")
	if err != nil {fmt.Println(err)}
	var temp []byte
	
	/*for k := 1; k <= 19; k++ {																									Более старая версия того, что я делал
		fmt.Println(k, "               ", temp)
		if k == 1 {temp = []byte("package main\n\n")}
		if k == 2 {temp = []byte("import (\n")}
		if k == 3 {temp = []byte("	\"fmt\"\n")}
		if k == 4 {temp = []byte(")\n\n")}
		if k == 5 {temp = []byte("func main() {\n")}
		if k == 6 {temp = []byte("	var (\n")}
		if k == 7 {temp = []byte("		num int\n")}
		if k == 8 {temp = []byte("		n int\n")}
		if k == 9 {temp = []byte("		sum int\n")}
		if k == 10 {temp = []byte("	)\n")}
		if k == 11 {temp = []byte("	fmt.Print(\"Введите кол-во чисел:\")\n")}
		if k == 12 {temp = []byte("	fmt.Scan(&num)\n")}
		if k == 13 {temp = []byte("	for i := 0; i == num; i++ {\n")}
		if k == 14 {temp = []byte("		fmt.Print(\"Введите число:\")\n")}
		if k == 15 {temp = []byte("		fmt.Scan(&n)\\n\n")}
		if k == 16 {temp = []byte("		sum += n\n")}
		if k == 17 {temp = []byte("	}\n")}
		if k == 18 {temp = []byte("	fmt.Println(\"Вот их сумма:\", sum)\n")}
		if k == 19 {temp = []byte("}"); break}
		err := ioutil.WriteFile("Result\\Simple.go", temp, 7777)
		if err != nil {fmt.Println(err)}
	}*/
	
	temp = []byte("package main\n\nimport (\n	\"fmt\"\n)\n\nfunc main() {\n	var (\n		num int\n		n int\n		sum int\n	)\n	fmt.Print(\"Введите кол-во чисел:\")\n	fmt.Scan(&num)\n	for i := 0; i < num; i++ {\n		fmt.Print(\"Введите число:\")\n		fmt.Scan(&n)\n		sum += n\n	}\n	fmt.Println(\"Вот их сумма:\", sum)\n}")
	err1 := ioutil.WriteFile("Simple.go", temp, 7777)
	if err1 != nil {fmt.Println(err1)}
	File.Close()
	
	exec.Command("go", "build", "Simple.go").Run()
	exec.Command("Simple.exe").Run()
}
