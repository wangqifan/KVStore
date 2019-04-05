package util

import "KVStore/errors"

func StingTounin64(s string) (uint64 ,error) {

    chars := []byte(s)
    if len(chars) ==0 || len(chars) >8 {
        return 0, errors.NewKeyInvaildError()
    } 

    var num uint64 = 0 
    for _,val := range chars {
        num = num*128 + uint64(val)
	}
	return num, nil
}

func StringToArray(s string) ([256]byte ,error){

	var result [256]byte

    var value []byte = []byte(s)
    if len(value) > 256 || len(value) ==0 {
        return result, errors.NewValueInvaildError()
    }
    for i, item := range value {
        result[i] = item
	}
	return result, nil
}