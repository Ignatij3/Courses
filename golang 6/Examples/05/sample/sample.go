package main

import (
	"fmt"
	"time"
)

const day = 24 * time.Hour

type Payee interface {
	CalculateSalary() int
}

type Worker interface {
	Payee
	ID() int
}

type VacationBenefit interface {
	SetVacation(duration time.Duration) VacationBenefit
	GetVacation(duration time.Duration) (v VacationBenefit, possible bool) // use vacation in whole or in part;
	// if the remainder is enough if the remainder is sufficient, then returns VacationBenefit with the modified
	// used part of vacation and true, else returns returns an unchanged value of VacationBenefit and false.
	VacationLeft() (remainder time.Duration) // vacation remainder.
}

type PrivilegedWorker interface {
	Worker
	VacationBenefit
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int           // provident fund
	vacation time.Duration // vacation duration
	used     time.Duration // used part of vacation
}

func (p Permanent) ID() int {
	return p.empId
}

// salary of permanent employee is the sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// set vacation duration of permanent worker
func (p Permanent) SetVacation(duration time.Duration) VacationBenefit {
	p.vacation = duration
	var v VacationBenefit = p
	return v
}

// get vacation in whole or partially of permanent worker
func (p Permanent) GetVacation(duration time.Duration) (v VacationBenefit, possible bool) {
	possible = p.vacation-p.used >= duration
	if possible {
		p.used += duration
	}
	v = p
	return
}

// vacation remainder of permanent worker
func (p Permanent) VacationLeft() (remainder time.Duration) {
	remainder = p.vacation - p.used
	return
}

type Contract struct {
	empId    int
	basicpay int
	vacation time.Duration // vacation duration
	used     time.Duration // used part of vacation
}

func (c Contract) ID() int {
	return c.empId
}

//salary of contract worker is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// set vacation duration of contract worker
func (c Contract) SetVacation(a time.Duration) VacationBenefit {
	c.vacation = a
	var v VacationBenefit = c
	return v
}

// get vacation in whole or in part of contract worker
func (c Contract) GetVacation(a time.Duration) (v VacationBenefit, possible bool) {
	if c.vacation < a+c.used {
		v, possible = c, false
		return
	}
	c.used += a
	v, possible = c, true
	return
}

// vacation remainder of contract worker
func (c Contract) VacationLeft() time.Duration {
	return c.vacation - c.used
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (f Freelancer) ID() int {
	return f.empId
}

//salary of freelancer
func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

type Pieceworker struct {
	empId    int
	payments []int
}

func (p Pieceworker) ID() int {
	return p.empId
}

//salary of pieceworker
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
		vacation: 21 * day,
		used:     0 * day,
	}
	if v, ok := perm1.GetVacation(3 * day); ok {
		perm1 = v.(Permanent)
	}
	perm2 := Permanent{
		empId:    1002,
		basicpay: 3000,
		pf:       30,
		vacation: 15 * day,
		used:     15 * day,
	}
	contr1 := Contract{
		empId:    2002,
		basicpay: 2400,
		vacation: 10 * day,
		used:     5 * day,
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

	fmt.Printf("Total Expense Per Month €%d\n\n", totalExpense(employees))
	// Total Expense Per Month €17295

	// Interface compatibility

	fmt.Println(HasVacationLeft(employees))
	// [{1001 2500 20 1814400000000000 259200000000000} {2002 2400 864000000000000 432000000000000}]

	// var _ VacationBenefit = freelanc2
	/*	compile-time error:
		cannot use freelanc2 (variable of type Freelancer) as
		VacationBenefit value in variable declaration:
		missing method GetVacation
	*/

	// var _ VacationBenefit = Freelancer{}
	/* compile-time error:
	cannot use (Freelancer literal) (value of type Freelancer)
	as VacationBenefit value in variable declaration:
	missing method GetVacation
	*/

	// Interface value
	for _, v := range employees {
		fmt.Printf("Type = %T. Value = %v.\n", v, v)
	}
	// Type = main.Permanent. Value = {1001 2500 20 1814400000000000 259200000000000}.
	// Type = main.Permanent. Value = {1002 3000 30 1296000000000000 1296000000000000}.
	// Type = main.Contract. Value = {2002 2400 864000000000000 432000000000000}.
	// Type = main.Freelancer. Value = {4001 30 120}.
	// Type = main.Freelancer. Value = {4003 45 80}.
	// Type = main.Pieceworker. Value = {5002 [450 250 430 700 315]}.

	var w Worker
	fmt.Println(w) //	<nil>
	fmt.Printf("w is %v and has value of type %T\n", w, w)
	//	w is <nil> and has value of type <nil>

	fmt.Println(w.ID())
	/* panic: runtime error: invalid memory address or nil pointer
		dereference [signal 0xc0000005 code=0x0 addr=0x0 pc=0xb3ec79]

		goroutine 1 [running]:
		main.main()
	        ./sample.go:222 +0x639
		exit status 2
	*/
	PrintHasVacation(perm1)          //	true
	PrintHasVacation(piece1)         //	false
	PrintHasVacation("Vasja Pupkin") //	ignore
}

// total expense is calculated by iterating through the SalaryCalculator
// slice and summing the salaries of the individual employees
func totalExpense(s []Worker) int {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}
	return expense
}

// list of workers who haven't used the whole vacation yet
func HasVacationLeft(s []Worker) (list []VacationBenefit) {
	for _, v := range s {
		if pv, ok := v.(VacationBenefit); ok {
			if pv.VacationLeft() > 0 {
				list = append(list, pv)
			}
		}
	}
	return list
}

// get empty interface as a parameter
func PrintHasVacation(w interface{}) {
	if _, ok := w.(Worker); ok {
		_, ok = w.(VacationBenefit)
		fmt.Println(ok)
	}
}
