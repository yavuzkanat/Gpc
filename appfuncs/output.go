package appfuncs

import (
	
	"fmt"
	"os"
)

func OutputAsTXT(data []string)  {
	f,err:= os.Create("wordlist.txt")
	if err != nil {
		fmt.Println("Cannot create file !")
	}
	
	for _, v := range data {
		data2 := fmt.Sprintf("%s%s",v,"\n")
		_,err := f.WriteString(data2)

		if err != nil {

			fmt.Println(err)
			f.Close()
		}
	}
}