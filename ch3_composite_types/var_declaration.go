package main
import "fmt"
import "reflect"

func main() {
    // the most verbose way to declare a variable is:
    var x int = 10
    fmt.Println("var x is", x)
    // if the type on the right hand side of the = is the expected type of the variable, you can deleter
    // the type from the left side
    var y = 2
    fmt.Println("type of y is", reflect.TypeOf(y))

    // if you drop the = on the righthand side, the value will be "zero" value
    var zero int
    fmt.Println("value of zero is", zero)
    // multiple declarations
    var a, b int = 10, 20
    fmt.Println("a, b", a, b)
    // multiple declarations with zero values
    var c, d string
    fmt.Printf("c=%v, d=%v\n", c, d)
    // mutiple declaration wrapped
    var (
        e int
        f   = 20
        g, h  = 40, "hello"
        i, j string
    )
    fmt.Printf("e=%v, f=%v, g=%v, h=%v, i=%v, j=%v\n",e,f,g,h,i,j)
    // := operator, a short declaration format
    // Can only be used when you are within a function, not for declaring a variable at the package level
    // var x = 10 and x := 10, do exactly the same
    o := 10
    fmt.Printf(":= operator does the type inference, type of o is %v", reflect.TypeOf(o))
    // you can do multiple declaration
    o, p := 20, "hello"
    fmt.Println("o, p", o, p)

}
