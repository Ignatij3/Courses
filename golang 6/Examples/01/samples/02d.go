package main

import (
	"fmt"
)

type Worker interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int // provident fund
}

type Contract struct {
	empId    int
	basicpay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

type Pieceworker struct {
	empId    int
	payments []int
}

//salary of permanent employee is the sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

//salary of freelancer
func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

//salary of piecework
func (p Pieceworker) CalculateSalary() int {
	var sum int
	for _, v := range p.payments {
		sum += v
	}
	return sum
}

func main() {
	perm1 := Permanent{
		empId:    1001,
		basicpay: 2500,
		pf:       20,
	}
	perm2 := Permanent{
		empId:    1002,
		basicpay: 3000,
		pf:       30,
	}
	contr1 := Contract{
		empId:    2002,
		basicpay: 2400,
	}
	freelanc1 := Freelancer{
		empId:       4001,
		ratePerHour: 30,
		totalHours:  120,
	}
	freelanc2 := Freelancer{
		empId:       4003,
		ratePerHour: 45,
		totalHours:  80,
	}
	piece1 := Pieceworker{
		empId:    5002,
		payments: []int{450, 250, 430, 700, 315},
	}
	employees := []Worker{perm1, perm2, contr1, freelanc1, freelanc2, piece1}
	for _, emp := range employees {
		switch emp.(type) {
		case Permanent:
			fmt.Printf("Постоянный работник. ID %d\n", emp.(Permanent).empId)
		case Contract:
			fmt.Printf("Контрактник. ID %d\n", emp.(Contract).empId)
		case Freelancer:
			fmt.Printf("Фрилансер. ID %d\n", emp.(Freelancer).empId)
		case Pieceworker:
			fmt.Printf("Работник со сдельной оплатой. ID %d\n", emp.(Pieceworker).empId)
		default:
			fmt.Println("Неопределённый тип")
		}	
	}
}
/*
Постоянный работник. ID 1001
Постоянный работник. ID 1002
Контрактник. ID 2002
Фрилансер. ID 4001
Фрилансер. ID 4003
Работник со сдельной оплатой. ID 5002 
*/ 
