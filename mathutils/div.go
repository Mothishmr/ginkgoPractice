package mathutils

import "errors"

func Div(a int , b int ) (int , error){

	if b == 0 {return -1 , errors.New("can not divide num with 0")}
	return a / b , nil
}