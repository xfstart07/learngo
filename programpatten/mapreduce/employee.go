package mapreduce

type Employee struct {
	Name string
	Age int
	Vacation int
	Salary int
}

func EmployeeCountIf(list []Employee, fn func(e *Employee)bool )int {
	count := 0
	for i := range list {
		if fn(&list[i]) {
			count++
		}
	}
	return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee)bool) []Employee {
	var newList []Employee
	for i := range list {
		if fn(&list[i]) {
			newList = append(newList, list[i])
		}
	}
	return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee)int)int {
	var sum = 0
	for i := range list {
		sum += fn(&list[i])
	}
	return sum
}