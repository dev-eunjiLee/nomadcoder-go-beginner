package something

import "fmt"

// * 동일 파일 내부에서는 사용 가능하지만, 대문자로 시작하지 않기 때문에 외부에서는 사용 불가능
func sayBye() {
	fmt.Println("Bye")
}

// SayHello * 대문자 시작 외부 사용 가능
func SayHello() {
	fmt.Println("Hello")
}
