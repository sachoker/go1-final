package actioninfo

import "fmt"

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, rawInfo := range dataset {
		if err := dp.Parse(rawInfo); err != nil {
			fmt.Println(err)
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(info)
	}
}
