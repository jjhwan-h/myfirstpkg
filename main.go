/*
기본 패키지들은 go를 설치할때 자동으로 같이 설치된다.
go env | grep GOROOT //go root 경로 확인
go mod init github.com/{깃허브아이디}/myfirstpkg //패키지모듈생성
*/

package myfirstpkg

import "fmt"

func myfirstFunc() { //첫글자가 소문자는 외부에서 사용할 수 없다.
	fmt.Println("firt pkg")
}

func MyfirstFunc() { //함수명의 첫글자가 대문자여야 외부에서 가져와서 사용할 수 있다.
	fmt.Println("firt pkg")
}
