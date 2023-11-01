package main

import (
	"testing"
)

func TestEmployeeSalaries(t *testing.T) {
	employeeSalaries := map[string]int{
		"John Doe": 50000,
		"Jane Smith": 60000,
	}

	johnSalary := employeeSalaries["John Doe"]
	janeSalary := employeeSalaries["Jane Smith"]

	if johnSalary != 50000 {
		t.Errorf("Expected John Doe's salary to be 50000, but got %d", johnSalary)
	}

	if janeSalary != 60000 {
		t.Errorf("Expected Jane Smith's salary to be 60000, but got %d", janeSalary)
	}
}