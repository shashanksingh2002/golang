package main; //Every Go program is made up of packages. Programs start running in package main.
import "fmt"; //Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

//func main is the function that gets called when the program starts. Every Go program must have a main function in the main package.
func main() { 
	fmt.Println("hello world");
}