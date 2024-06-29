package main

import (
	"fmt"
	"os"
)

type Employee struct {
	name string
	grade int16
	pref_work_times []string
	pref_working_hours float32
}

func (*Employee) new_employee (
	name string,
	grade int16,
	work_times []string,
	work_hours float32,
) Employee {
	if grade < 0 && 4 < grade {
		fmt.Printf("grade number input error")
		os.Exit(1)
	}
	return Employee{
		name: name,
		grade: grade,
		pref_work_times: work_times,
		pref_working_hours: work_hours,
	}	
}
