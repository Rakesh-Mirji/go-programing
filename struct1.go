// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import ("fmt"
"time")

func EmployeeById(id int) *Employee{
    return &anil
}

    type Employee struct{
        ID  int
        Name    string
        Address    string
        DOB    time.Time
        Position    string
        Salary  int
        ManagerID   int
    }                                                                 0

    var anil Employee
    
 
func main() {
    anil.Name = "Anil"
    anil.Salary += 50000
    
    //assign using pointer's
    position := &anil.Position
    *position = "Senior"+*position
    
    
    var EmplyeeOfTheMonth *Employee = &anil 
    EmplyeeOfTheMonth.Salary += 5000
    
    (*EmplyeeOfTheMonth).Position += " (Proactive Team Player)"
    fmt.Println(EmplyeeOfTheMonth)
    
    fmt.Println(EmployeeById(anil.ManagerID).Position)
}