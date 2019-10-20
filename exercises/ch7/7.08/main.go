package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Employee struct {
	FirstName  string
	LastName   string
	EmployeeID int
	StartDate  time.Time
}

var (
	employees = []*Employee{
		{"Denis", "Angulo", 1, time.Now().Add(144 * time.Hour)},
		{"Example", "Exampleson", 3, time.Now().Add(96 * time.Hour)},
		{"John", "Smith", 10, time.Now().Add(2 * time.Hour)},
		{"Alice", "Of Wonderland", 99, time.Now().Add(24 * time.Hour)},
		{"Dreymond", "Green", 101, time.Now().Add(1231 * time.Hour)},
		{"Geralt", "Of Rivia", 323, time.Now().Add(1212 * time.Hour)},
		{"Adam", "Jensen", 451, time.Now().Add(365 * 24 * 8 * time.Hour)},
	}
)

func printEmployees(employees []*Employee) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "First name", "Last name", "Employee ID", "Start Date")
	fmt.Fprintf(tw, format, "----------", "---------", "-----------", "----------")
	for _, e := range employees {
		fmt.Fprintf(tw, format, e.FirstName, e.LastName, e.EmployeeID, e.StartDate.Format(time.RFC3339))
	}
	tw.Flush()
}

func main() {
	firstName := func(e1, e2 *Employee) bool {
		return e1.FirstName < e2.FirstName
	}
	lastName := func(e1, e2 *Employee) bool {
		return e1.LastName < e2.LastName
	}
	employeeID := func(e1, e2 *Employee) bool {
		return e1.EmployeeID < e2.EmployeeID
	}
	startDate := func(e1, e2 *Employee) bool {
		return e1.StartDate.Before(e2.StartDate)
	}
	fmt.Println("original order")
	printEmployees(employees)
	fmt.Println("")

	fmt.Println("Sorted by: Last name")
	OrderedBy(lastName).Sort(employees)
	printEmployees(employees)
	fmt.Println("")

	fmt.Println("Sorted by: Last name, first name")
	OrderedBy(lastName, firstName).Sort(employees)
	printEmployees(employees)
	fmt.Println("")

	fmt.Println("Sorted by: start date, last name")
	OrderedBy(startDate, lastName).Sort(employees)
	printEmployees(employees)
	fmt.Println("")

	fmt.Println("Sorted by: start date, employee ID")
	OrderedBy(startDate, employeeID).Sort(employees)
	printEmployees(employees)
	fmt.Println("")

	fmt.Println("Sorted by: start date, last name, first name")
	OrderedBy(startDate, lastName, firstName).Sort(employees)
	printEmployees(employees)
	fmt.Println("")

}

type lessFunc func(p1, p2 *Employee) bool

type multiSorter struct {
	employees []*Employee
	less      []lessFunc
}

func (ms *multiSorter) Sort(employees []*Employee) {
	ms.employees = employees
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.employees)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.employees[i], ms.employees[j] = ms.employees[j], ms.employees[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := ms.employees[i], ms.employees[j]

	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}

/*
```sh
~$ go run exercises/ch7/7.08/main.go
original order
First name  Last name      Employee ID  Start Date
----------  ---------      -----------  ----------
Denis       Angulo         1            2019-10-26T08:30:27-04:00
Example     Exampleson     3            2019-10-24T08:30:27-04:00
John        Smith          10           2019-10-20T10:30:27-04:00
Alice       Of Wonderland  99           2019-10-21T08:30:27-04:00
Dreymond    Green          101          2019-12-10T15:30:27-04:00
Geralt      Of Rivia       323          2019-12-09T20:30:27-04:00
Adam        Jensen         451          2027-10-18T08:30:27-04:00

Sorted by: Last name
First name  Last name      Employee ID  Start Date
----------  ---------      -----------  ----------
Denis       Angulo         1            2019-10-26T08:30:27-04:00
Example     Exampleson     3            2019-10-24T08:30:27-04:00
Dreymond    Green          101          2019-12-10T15:30:27-04:00
Adam        Jensen         451          2027-10-18T08:30:27-04:00
Geralt      Of Rivia       323          2019-12-09T20:30:27-04:00
Alice       Of Wonderland  99           2019-10-21T08:30:27-04:00
John        Smith          10           2019-10-20T10:30:27-04:00

Sorted by: Last name, first name
First name  Last name      Employee ID  Start Date
----------  ---------      -----------  ----------
Denis       Angulo         1            2019-10-26T08:30:27-04:00
Example     Exampleson     3            2019-10-24T08:30:27-04:00
Dreymond    Green          101          2019-12-10T15:30:27-04:00
Adam        Jensen         451          2027-10-18T08:30:27-04:00
Geralt      Of Rivia       323          2019-12-09T20:30:27-04:00
Alice       Of Wonderland  99           2019-10-21T08:30:27-04:00
John        Smith          10           2019-10-20T10:30:27-04:00

Sorted by: start date, last name
First name  Last name      Employee ID  Start Date
----------  ---------      -----------  ----------
John        Smith          10           2019-10-20T10:30:27-04:00
Alice       Of Wonderland  99           2019-10-21T08:30:27-04:00
Example     Exampleson     3            2019-10-24T08:30:27-04:00
Denis       Angulo         1            2019-10-26T08:30:27-04:00
Geralt      Of Rivia       323          2019-12-09T20:30:27-04:00
Dreymond    Green          101          2019-12-10T15:30:27-04:00
Adam        Jensen         451          2027-10-18T08:30:27-04:00

Sorted by: start date, employee ID
First name  Last name      Employee ID  Start Date
----------  ---------      -----------  ----------
John        Smith          10           2019-10-20T10:30:27-04:00
Alice       Of Wonderland  99           2019-10-21T08:30:27-04:00
Example     Exampleson     3            2019-10-24T08:30:27-04:00
Denis       Angulo         1            2019-10-26T08:30:27-04:00
Geralt      Of Rivia       323          2019-12-09T20:30:27-04:00
Dreymond    Green          101          2019-12-10T15:30:27-04:00
Adam        Jensen         451          2027-10-18T08:30:27-04:00

Sorted by: start date, last name, first name
First name  Last name      Employee ID  Start Date
----------  ---------      -----------  ----------
John        Smith          10           2019-10-20T10:30:27-04:00
Alice       Of Wonderland  99           2019-10-21T08:30:27-04:00
Example     Exampleson     3            2019-10-24T08:30:27-04:00
Denis       Angulo         1            2019-10-26T08:30:27-04:00
Geralt      Of Rivia       323          2019-12-09T20:30:27-04:00
Dreymond    Green          101          2019-12-10T15:30:27-04:00
Adam        Jensen         451          2027-10-18T08:30:27-04:00
```
*/
