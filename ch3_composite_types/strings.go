package main
// strings in go are imutable
// rune type is an alais for int32 type and byte is an alias for int8
// if you are referring to a character, it is recommended to use rune type and not
// int32, even though this is the same for compiler, it show the intent of your code better for you.
import "fmt"
func main() {
    fmt.Println(len("Hello World!"));
    fmt.Println("Hellow World!"[1]);
    // you can use + for string concatentaton
    fmt.Println("Hello" + "World!");
    var x string;
    x = "first";
    x = x + " second";
    fmt.Println(x)
    // strings are compared for equality using == and difference with !=
    var m string = "Hello"
    var n string = "Hello"
    fmt.Println(n == m)
    n = "hello"
    fmt.Println(n == m)
}
