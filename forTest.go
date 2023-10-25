package main ;import ("fmt")
func main() {var num int;fmt.Print("Enter number: ");fmt.Scanln(&num);for i := 1; i < num; i++ {for j := 1; j < i; j ++ {;fmt.Print("*")};fmt.Println()}};