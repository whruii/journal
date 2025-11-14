package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Grades []float64
}

type Students map[string]Student

func (s Student) average() float64 {
	if len(s.Grades) == 0 {
		return 0.0
	}
	var sum float64
	for _, grade := range s.Grades {
		sum += grade
	}
	return sum / float64(len(s.Grades))
}

func (ss *Students) addStudent(name string, grades []float64) {
	(*ss)[name] = Student{Name: name, Grades: grades}
}

func (ss Students) printAll() {
	if len(ss) == 0 {
		fmt.Println("Список студентов пуст.")
		return
	}
	fmt.Println("\n--- Список студентов ---")
	for _, student := range ss {
		fmt.Printf("ФИО: %s | Оценки: %v | Средний балл: %.2f\n",
			student.Name, student.Grades, student.average())
	}
}

func (ss Students) filterByAverage(threshold float64) []Student {
	var result []Student
	for _, student := range ss {
		if student.average() < threshold {
			result = append(result, student)
		}
	}
	return result
}

func readGrades(input string) ([]float64, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("оценки не введены")
	}
	var grades []float64
	for _, part := range parts {
		grade, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, fmt.Errorf("некорректная оценка: '%s'", part)
		}
		if grade < 2 || grade > 5 {
			return nil, fmt.Errorf("оценка '%.1f' должна быть от 2 до 5", grade)
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	students := make(Students)

	fmt.Println("🎓 Система учёта успеваемости студентов")
	fmt.Println("Поддерживаемые команды:")
	fmt.Println("  add      — добавить/обновить студента")
	fmt.Println("  list     — показать всех студентов")
	fmt.Println("  filter   — показать студентов со средним баллом ниже заданного")
	fmt.Println("  exit     — выйти")

	for {
		fmt.Print("\nВведите команду (add/list/filter/exit): ")
		if !scanner.Scan() {
			break
		}
		cmd := strings.TrimSpace(strings.ToLower(scanner.Text()))

		switch cmd {
		case "add":
			fmt.Print("Введите ФИО студента: ")
			if !scanner.Scan() {
				break
			}
			name := strings.TrimSpace(scanner.Text())
			if name == "" {
				fmt.Println("ФИО не может быть пустым.")
				continue
			}

			fmt.Print("Введите оценки через пробел (например: 5 4 5 3): ")
			if !scanner.Scan() {
				break
			}
			gradesInput := scanner.Text()
			grades, err := readGrades(gradesInput)
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				continue
			}

			students.addStudent(name, grades)
			avg := Student{Name: name, Grades: grades}.average()
			fmt.Printf("Студент '%s' добавлен. Средний балл: %.2f\n", name, avg)

		case "list":
			students.printAll()

		case "filter":
			fmt.Print("Введите порог среднего балла (например: 4): ")
			if !scanner.Scan() {
				break
			}
			thresholdStr := strings.TrimSpace(scanner.Text())
			threshold, err := strconv.ParseFloat(thresholdStr, 64)
			if err != nil {
				fmt.Println("Некорректное число.")
				continue
			}

			filtered := students.filterByAverage(threshold)
			if len(filtered) == 0 {
				fmt.Printf("Нет студентов со средним баллом ниже %.2f\n", threshold)
			} else {
				fmt.Printf("\n--- Студенты со средним баллом ниже %.2f ---\n", threshold)
				for _, s := range filtered {
					fmt.Printf("ФИО: %s | Оценки: %v | Средний балл: %.2f\n",
						s.Name, s.Grades, s.average())
				}
			}

		case "exit":
			fmt.Println("До свидания!")
			os.Exit(0)

		default:
			fmt.Println("Неизвестная команда. Доступные: add, list, filter, exit")
		}
	}
}
