package main

import "fmt"

func main() {
    //there are few declaration types:
    var x [3]int
    fmt.Println("x", x)
    fmt.Println("len(x):", len(x))
    // no values are specified so inititated to zero value for int
    // or with array literals:
    var y = [3]int{10, 20, 30}
    fmt.Println("y", y)
    fmt.Println("len(y):", len(y))
    // if you have sparse arrays:
    // value for index 5 is 4 and four index 10 is 100,
    // unspecified values are zero.
    var z = [12]int{1, 5:4, 6, 10:100, 15}
    fmt.Println("z:", z)
    // comparing arrays
    var a = [...]int{1, 2, 3}
    var b = [3]int{1, 2, 3}
    fmt.Println("a == b?", a == b)
    // you can use [] to index into arrays
    fmt.Println("a[2]", a[2])
    fmt.Println("----------------------")
    // fmt.Println("a[3]", a[3])
    // you should get this error:
    // ./arrays.go:27:27: invalid argument: index 3 out of bounds [0:3]

}
