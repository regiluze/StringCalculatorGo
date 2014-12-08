package src

import "strconv"
import "strings"
import "errors"
import "fmt"


const (

   DEFAULT_DELIMETER = ","

)

type StringCalculator struct {

}

func (cal *StringCalculator) Add(numbers string) (int, error){
		delimiter , numbersList := cal.setupNumbers(numbers)
		total := 0
		negatives := []int{}
		for _ , value := range cal.getNumbers(numbersList,delimiter){
		  number, _ := strconv.Atoi(value)
		  total += number
		  if number < 0{
			append(negatives,number)
		  }
		}
		if len(negatives) == 0 {
			return total, nil
		}else{
			msg := fmt.Sprintf("negative numbers not allowed : %s",negatives)
			return 0, errors.New(msg)
		}
}

func (cal *StringCalculator) setupNumbers(numbers string) (string,string){
	if strings.HasPrefix(numbers,"//"){
		return numbers[3:4] , numbers[5:]
	} else {
	    return	DEFAULT_DELIMETER , numbers
	}
}

func (cal * StringCalculator) getNumbers(numbers,delimiter string) ([]string){
	pnumbers := strings.Replace(numbers,"/n",delimiter,-1)
	return strings.Split(pnumbers,delimiter)
}
