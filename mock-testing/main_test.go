package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := [] struct {
		id int
		dni string
		mockFunc func()
		expectedEmployee FullTimeEmployee
	} {
		{
			id: 1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						ID: 1,
						role: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Dni: "1",
						Name: "Jesus Rivero",
						Age: 26,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person{
					Dni: "1",
					Name: "Jesus Rivero",
					Age: 26,
				},
				Employee{
					ID: 1,
					role: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()

		fte, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("Error obteniendo el empleado de tiempo completo")
		}

		if fte.Age != test.expectedEmployee.Age {
			t.Errorf("Error en el empleado tiempo comleto, se obtuvo %d, se esperaba %d", fte.Age, test.expectedEmployee.Age)
		}

	}

	GetPersonByDNI = originalGetPersonByDni
	GetEmployeeById = originalGetEmployeeById
}
