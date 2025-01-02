package main

import (
	"fmt"
	"strconv"
)

type Company struct {
    Name        string
    Address     string
    PhoneNumber string
    NumberOfEmployees int
}

func main() {
    fmt.Println("====Ex0====")
    ex0()
    fmt.Printf("\n")

    fmt.Println("====Ex1====")
    ex1()
    fmt.Printf("\n")

    fmt.Println("====Ex1.2====")
    ex1_2()
    fmt.Printf("\n")

    fmt.Println("====Ex2====")
    ex2()
    fmt.Printf("\n")

    fmt.Println("====Ex3====")
    ex3()
    fmt.Printf("\n")

    fmt.Println("====Ex3.1====")
    ex3_1()
    fmt.Printf("\n")

    fmt.Println("====Ex4====")
    ex4()
    fmt.Printf("\n")

    fmt.Println("====Ex4.1====")
    ex4_1()
    fmt.Printf("\n") 

    fmt.Println("====Ex5====")
    ex5()
    fmt.Printf("\n")

    fmt.Println("====Ex6====")
    ex6()
    fmt.Printf("\n")

    fmt.Println("====Ex Special====")
    special()
    fmt.Printf("\n")
}

//  <!====================== ข้อ 0 =======================>
func ex0() {
    i := 2
    if i == 0 {
            fmt.Println("Zero")
    } else if i == 1 {
            fmt.Println("One")
    } else if i == 2 {
            fmt.Println("Two")
    } else if i == 3 {
            fmt.Println("Three")
    } else {
            fmt.Println("เจ้าไม่อยู่ในเงื่ิอนไข.")
    }
}

//  <!====================== ข้อ 1 =======================>
func ex1() {
    count := 0 
    for i := 0; i < 100; i++ {
        if i%3 == 0 {
            fmt.Println(i)
            count++
        }
    }
    fmt.Printf("\n")
    fmt.Printf("จำนวนที่หาร 3 ลงตัวมีทั้งหมด %d\n", count)
}

//  <!====================== ข้อ 1.2 =======================>
func ex1_2() {
    base := 20
    exponent := 2
    result := pow(base, exponent)
    fmt.Println(result)
}

func pow(base, exponent int) int {
    result := 1
    for i := 0; i < exponent; i++ {
        result *= base
    }
    return result
}

//  <!====================== ข้อ 2 =======================>
func ex2() {
    x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
    min, max := findMinAndMax(x)
    fmt.Printf("min is %d\n", min)
    fmt.Printf("max is %d\n", max)
}

func findMinAndMax(x []int) (int, int) {
    min := x[0]
    max := x[0]
    for _, v := range x {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    return min, max
}

//  <!====================== ข้อ 3 =======================> 
func ex3() {
	count := 0 
	for i := 1; i <= 1000; i++ {
		str := strconv.Itoa(i) 
		for _, ch := range str {
			if ch == '9' { 
				count++
			}
		}
	}
	fmt.Printf("จำนวนทั้งหมดที่มีเลข 9 จาก 1 ถึง 1000: %d\n", count)
}
   

//  <!====================== ข้อ 3.1 =======================>
func ex3_1() {
    {
        num := 10000
        result := countNumberNines(num)
        fmt.Printf("จำนวนทั้งหมดที่มีเลข 9 จาก 1 ถึง %d: %d\n", num, result)
    }
}

func countNumberNines(n int) int {
    count := 0
    for i := 1; i <= n; i++ {
        str := strconv.Itoa(i) 
        for _, ch := range str {
            if ch == '9' { 
                count++
            }
        }
    }
    return count
}

//  <!====================== ข้อ 4 =======================>
func ex4() {
    var myWords = "AW SOME GO!"
    var result string
    for _, char := range myWords {
            if char != ' ' {
                    result += string(char)
            }
    }
    fmt.Println(result)
}

//  <!====================== ข้อ 4.1 =======================>
func ex4_1() {
    myWords := "ine t"
    result := cutText(myWords)
    fmt.Println(result)
}

func cutText(s string) string {
    result := ""
    for _, char := range s {
        if char != ' ' {
            result += string(char)
        }
    }
    return result
}
//  <!====================== ข้อ 5 =======================>
func ex5() {
    people := map[string]map[string]string {
        "emp_01": {"name": "Marry", "lname": "Jane"},
        "emp_02": {"name": "Gwen", "lname": "Stefrane"},
        "emp_03": {"name": "John", "lname": "Doe"},
        "emp_04": {"name": "Alice", "lname": "Wonderland"},
}
    for empID, person := range people {
            fmt.Printf("Employee ID: %s\n", empID)
            fmt.Printf("Name: %s %s\n", person["name"], person["lname"])
            fmt.Println()
    }
}

//  <!====================== ข้อ 6 =======================>
func ex6() {
    company1 := Company{
        Name:        "บริษัท ABC จำกัด",
        Address:     "123 กรุงเทพฯ",
        PhoneNumber: "02-123-4567",
        NumberOfEmployees: 100,
}

fmt.Println("ข้อมูลบริษัท")
fmt.Println("ชื่อบริษัท:", company1.Name)
fmt.Println("ที่อยู่:", company1.Address)
fmt.Println("เบอร์โทรศัพท์:", company1.PhoneNumber)
fmt.Println("จำนวนพนักงาน:", company1.NumberOfEmployees)
}


//  <!====================== ข้อพิเศษ =======================>
func special () {
    rows := 6
    for i := 1; i <= rows; i++ {
            for j := 1; j <= i; j++ {
                fmt.Print("*")
            }
        fmt.Println() 
    }
}
