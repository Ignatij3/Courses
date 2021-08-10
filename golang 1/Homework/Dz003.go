package main
import ("fmt"
			  )
func main() {
	var n1, n2 float64
	fmt.Println("Какой месяц сопряжён с этой цифрой?")
	fmt.Scan(&n1)
		if n1==1 {
		fmt.Println("Это Январь")
			} 
			if n1==2 {
			fmt.Println("Это Февраль")
				} 
				if n1==3 {
				fmt.Println("Это Март")
					} 
					if n1==4 {
					fmt.Println("Это Апрель")
						} 
						if n1==5 {
						fmt.Println("Это Май")
							} 
							if n1==6 {
							fmt.Println("Это Июнь")
								} 
								if n1==7 {
								fmt.Println("Это Июль")
									} 
									if n1==8 {
									fmt.Println("Это Август")
										} 
										if n1==9 {
										fmt.Println("Это Сентябрь")
											} 
											if n1==10 {
											fmt.Println("Это Октябрь")
												} 
												if n1==11 {
												fmt.Println("Это Ноябрь")
													} 
													if n1==12 {
													fmt.Println("Это Декабрь")
													}
														if n1>12 {
														fmt.Println("Такого месяца не существует")
														}
															if n1<=0 {
															fmt.Println("Такого месяца не существует")
															} 
	fmt.Println("Какое время года сопряжено с этим числом?")
	fmt.Scan(&n2)
	if n2==12 || n2==1 || n2==2 {
		fmt.Println("Это зима")
	}
		if n2==3 || n2==4 || n2==5 {
			fmt.Println("Это весна")
		}
			if n2==6 || n2==7 || n2==8 {
				fmt.Println("Это лето")
			}
				if n2==9 || n2==10 || n2==11 {
					fmt.Println("Это осень")
				}
					if n2>12{
						fmt.Println("Такого времени года не существует")
					}
						if n2<=0 {
							fmt.Println("Такого времени года не существует")
						}
	
}
