func scanString() string {
	in:=bufio.NewReader(os.Stdin)
	str,err:=in.ReadString('\n')
	if err!=nil && err!=io.EOF {
		fmt.Println("Ошибка ввода:", err)
	}
	return str
}
// Import "io" "bufio"
