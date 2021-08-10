package main
import "fmt"

func main() {
	var a,b,c, Min, Mid, Max int
	fmt.Print("a=")
	fmt.Scan(&a)
	fmt.Print("b=")
	fmt.Scan(&b)
	fmt.Print("c=")
	fmt.Scan(&c)
	if a<b && a<c && b<c {
		Min=a 
		Mid=b
		Max=c
	}
	if a<b && a<c && c<b {
		Min=a 
		Mid=c
		Max=b
	}
	if b<a && b<c && a<c {
		Min=b
		Mid=a
		Max=c
	}
	if b<a && b<c && c<a {
		Min=b
		Mid=c
		Max=a
	}
	if c<a && c<b && a<b {
		Min=c
		Mid=a
		Max=b
	}
	if c<a && c<b && b<a {
		Min=c
		Mid=b
		Max=a
	}
	for a==c || a==b || b==c {
		fmt.Println("Ошибка, некоторые числа равны, попробуйте заново:")
		fmt.Print("a=")
		fmt.Scan(&a)
		fmt.Print("b=")
		fmt.Scan(&b)
		fmt.Print("c=")
		fmt.Scan(&c)	
		}
	fmt.Println("Минимальное число:", Min)
	fmt.Println("Среднее число:", Mid)
	fmt.Println("Максимальное число:", Max)
}
