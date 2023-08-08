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

/*
total expense is calculated by iterating through the Worker slice and summing
the salaries of the individual employees
*/
func totalExpense(s []Worker) int {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	return expense
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
	employees := []Worker{perm1, perm2, contr1, freelanc1, freelanc2}
	fmt.Printf("Total Expense Per Month €%d\n", totalExpense(employees))
}
