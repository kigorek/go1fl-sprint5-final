package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info принимает слайс строк с данными о тренировках или прогулках и экземпляр одной из ваших структур Training или DaySteps.
func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Println(err)
		}
		outStr, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(outStr)
	}

}
