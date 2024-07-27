package badexample

import "fmt"
type Once struct{

}
var(
	config *Once
)
func CheckStatus(){
	if config==nil{
        config=&Once{}
        fmt.Println("New configuration has been set")
	}else{
		fmt.Println("Old configuration ")
	}
}