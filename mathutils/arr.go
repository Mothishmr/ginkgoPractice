package mathutils

import "errors"

func Arr(arr []int , val int) (bool,error){

	if len(arr) == 0 {
	   return false , errors.New("arr is empty")
	}

	for _, i := range arr{
		if val == i{
			return true , nil
		}
	}
	return false, nil

}