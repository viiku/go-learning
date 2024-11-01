# structs
A struct is an aggregate data type that groups together zero or more named values of arbitrary
types as a single entity. Each value is called a field.

<!-- Employee defenintion of struct type, a variable called dilbert that is an instance of an Employee: -->
```
type Employee struct {
    ID int
    Name string
    Address string
    DoB time.Time
    Position string
    Salary int
    ManagerID int
}

var dilbert Employee
```
// Accessing indivisual

The individual fields of dilbert are accessed using dot notation like dilbert.Name and dilbert.DoB. Because dilbert is a variable, its fields are variables too, so we may assign to a field:

dilbert.Salary -= 5000 // demoted, for writing too few lines of code or take its address and access it through a pointer:

position := &dilbert.Position
*position = "Senior " + *position // promoted, for outsourcing to Elbonia

The dot notation also works with a pointer to a struct:

var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"

The last statement is equivalent to
(*employeeOfTheMonth).Position += " (proactive team player)"

Given an employee’s unique ID, the function EmployeeByID returns a pointer to an Employee
struct. We can use the dot notation to access its fields:
```
func EmployeeByID(id int) *Employee { /* ... */ }
fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
id := dilbert.ID
EmployeeByID(id).Salary = 0 // fired for... no real reason
```