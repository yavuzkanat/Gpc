package userfuncs

import (
	"fmt"
	"strings"
)

//create a struct for infos
type PersonalInfos struct {
	Name string
	Surname string
	Bdate []string
	Others []string
}


//get informations with input
func GetUserDatas() PersonalInfos{
	var name string
	fmt.Print("#Name :")
	fmt.Scan(&name)
	var surname string
	fmt.Print("#Surname :")
	fmt.Scan(&surname)
	var bdate string
	fmt.Print("#Enter birthday (dd-mm-yyyy) :")
	fmt.Scan(&bdate)
	var others string
	fmt.Print("Enter other infos (w1,w2,w3) :")
	fmt.Scan(&others)

	data := PersonalInfos{
		Name: name,
		Surname: surname,
		Bdate: strings.Split(bdate,"-"),
		Others: strings.Split(others,","),
	}
	return data
}