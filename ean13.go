package ean13

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidCode = errors.New("erro: código fornecido inválido");

func GetEan13VerficationDigit(eanCode string) (string, error) {
	if len(eanCode) < 12 {
		return "", ErrInvalidCode;
	}

	var soma uint

	arrCode := strings.Split(eanCode, "")

	for i, elem := range arrCode {
		num, error := strconv.ParseInt(elem, 0, 64)

		if error != nil {
			return "", ErrInvalidCode
		}


		if i % 2 == 0 {
			soma += uint(num) * 1
		}else{
			soma += uint(num) * 3
		}
	}


	result := (soma / 10) + 1
	result *= 10

	result -= soma
	
	if result % 10 == 0 {
		return fmt.Sprintf("%s-%d",eanCode,0), nil
	}else{
		return fmt.Sprintf("%s-%d",eanCode,result), nil
	}

}


