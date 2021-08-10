package main
import (
			"fmt"
			"os"
			)
func main() {
fin, err:=os.Open("numbers01.txt")
	var k int
	if err!=nil {return}
		fmt.Fscan(fin,&k)
		m:=make([]int, k)
		for n:=range m{
			if m[n]==m[n+1] {m[n]=0}/*я пытался сделать с помощью delete(m[n]), но не получилось*/
			fmt.Fprintln(fin, m[n])
			n++
		}
fin.Close()
}
