package main

import (
	"fmt"
	"os"
)

func main() {
	finfo, err := os.Stat("..//Files//Rainbow.bik")
	if err != nil {
		fmt.Println(err)
		return
	}
	var (
		temp int
		kiloBytes int
		path string
		dir string
	)	
	
	temp = int(finfo.Size())
	kiloBytes = temp / 1000
	path, _ = os.Getwd()
	
	fmt.Println("Название:", finfo.Name())
	fmt.Println("Размер:", finfo.Size(), "байт, примерно", kiloBytes, "килобайт")
	fmt.Println("Путь:", path)
	fmt.Println("Разрешения:", finfo.Mode())
	fmt.Println("Дата последнего изменения:", finfo.ModTime())
	if finfo.IsDir() {
		dir = "Да"
	} else {
		dir = "Нет"
	}
	fmt.Println("Является ли директорией? ", dir)
	
	
	
	
	
	
	
	
	
	
	
}
