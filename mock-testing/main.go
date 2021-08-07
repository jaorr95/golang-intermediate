package main

import "time"

type Person struct {
	Dni  string
	Name string
	Age  int
}

type Employee struct {
	ID int
	role string
}

type FullTimeEmployee struct {
	Person
	Employee
}

var GetPersonByDNI =  func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * person where ....
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// SELECT * employee where ....
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

   e, err := GetEmployeeById(id)
   if err != nil {
   	return ftEmployee, err
   }

   p, err := GetPersonByDNI(dni)
   if err != nil {
   	return ftEmployee, err
   }
   ftEmployee.Person = p
   ftEmployee.Employee = e

   return ftEmployee, nil
}

func main() {

}
