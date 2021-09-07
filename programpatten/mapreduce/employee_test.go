package mapreduce

import "testing"

var testList = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}

func TestEmployeeCountIf(t *testing.T) {
	old := EmployeeCountIf(testList, func(e *Employee) bool {
		return e.Age > 40
	})
	t.Logf("old employee: %v", old)

	highPay := EmployeeCountIf(testList, func(e *Employee) bool {
		return e.Salary >= 6000
	})
	t.Logf("high pay: %v", highPay)

	totalPay := EmployeeSumIf(testList, func(e *Employee) int {
		return e.Salary
	})
	t.Logf("total pay: %v", totalPay)

	youngerPay := EmployeeSumIf(testList, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	t.Logf("younger pay: %v", youngerPay)
}
