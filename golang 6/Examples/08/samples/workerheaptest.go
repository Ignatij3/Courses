package main

import (
	"container/heap"
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

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func (p Pieceworker) CalculateSalary() int {
	var sum int
	for _, v := range p.payments {
		sum += v
	}
	return sum
}

type WorkerHeap []Worker

func (w WorkerHeap) Len() int           { return len(w) }
func (w WorkerHeap) Less(i, j int) bool { return w[i].CalculateSalary() < w[j].CalculateSalary() }
func (w WorkerHeap) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }

func (w *WorkerHeap) Push(x interface{}) {
	*w = append(*w, x.(Worker))
}

func (w *WorkerHeap) Pop() interface{} {
	n := len(*w)
	x := (*w)[n-1]
	*w = (*w)[:n-1]
	return x
}

func main() {
	perm1 := Permanent{empId: 1001, basicpay: 2500, pf: 20}
	perm2 := Permanent{empId: 1002, basicpay: 3000, pf: 30}
	contr1 := Contract{empId: 2002, basicpay: 2400}
	freelanc1 := Freelancer{empId: 4001, ratePerHour: 30, totalHours: 120}
	freelanc2 := Freelancer{empId: 4003, ratePerHour: 45, totalHours: 80}
	piece1 := Pieceworker{empId: 5002, payments: []int{450, 250, 430, 700, 315}}

	empHeap := &WorkerHeap{perm1, perm2, contr1, freelanc1, freelanc2, piece1}
	heap.Init(empHeap)

	var employees []Worker
	for empHeap.Len() > 0 {
		employees = append(employees, heap.Pop(empHeap).(Worker))
	}
	for _, emp := range employees {
		fmt.Println(emp, emp.CalculateSalary())
	}
}

// {5002 [450 250 430 700 315]} 2145
// {2002 2400} 2400
// {1001 2500 20} 2520
// {1002 3000 30} 3030
// {4003 45 80} 3600
// {4001 30 120} 3600
