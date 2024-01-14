package main
import "fmt"
func main(){
    // go does not perform any implicit type conversion
    var x int = 10
    var y float64 = 30.2
    var z float64 = float64(x) + y
    var d int = x + int(y)
    fmt.Println(z)
    fmt.Println(d)
    // all type conversions in go are explicit, this means unlike some other programming languages
    // you cannot treat another Go type as booelan
    // in fact no other type can be converted to boolean in go, implicitly or explicitly
    // if you need to do some comparison you need to use the comparison operators
    // ==, !=, >, <, <= or >=
    var s string = "Not empty"
    if s == ""{
        fmt.Println("s is empty")
    } else{fmt.Println("s is not empty!")}
}
