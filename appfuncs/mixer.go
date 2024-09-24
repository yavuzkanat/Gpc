package appfuncs

import (
	"fmt"
	"gpcpacks/userfuncs"
	
)

//Two Word Combanation
func TwoWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	
	//2! + 2! = 4 |  4*3 = 12 different strings
	for _, value1 := range data.Bdate {
		Fac1 := fmt.Sprintf("%s%s",data.Name,value1)   
		Fac2 := fmt.Sprintf("%s%s",value1,data.Name)  
		Fac3 := fmt.Sprintf("%s%s",data.Surname,value1) 
		Fac4 := fmt.Sprintf("%s%s", value1,data.Surname) 
		ListMixer = append(ListMixer, Fac1)

		ListMixer = append(ListMixer, Fac2)

		ListMixer = append(ListMixer, Fac3)

		ListMixer = append(ListMixer, Fac4)
	} 
	//2! + 2! = 4 4*n(others) = 4n different strings
	for _, value2 := range data.Others {
		Sp3 := fmt.Sprintf("%s%s",data.Name,value2)
		Sp4 := fmt.Sprintf("%s%s",data.Surname,value2)
		Sp7 := fmt.Sprintf("%s%s",value2,data.Name)
		Sp8 := fmt.Sprintf("%s%s",value2,data.Surname)
		ListMixer = append(ListMixer, Sp3)

		ListMixer = append(ListMixer, Sp4)

		ListMixer = append(ListMixer, Sp7)

		ListMixer = append(ListMixer, Sp8)

	}
	return ListMixer
}

//Three Word  Combanation	
func ThreeWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	// 3! x 3 = 18 different strings
	for _ , value2 := range data.Bdate {
		Fac1 := fmt.Sprintf("%s%s%s",data.Name,data.Surname,value2)
		Fac2 := fmt.Sprintf("%s%s%s",data.Name,value2,data.Surname)
		Fac3 := fmt.Sprintf("%s%s%s",value2,data.Name,data.Surname)
		Fac4 := fmt.Sprintf("%s%s%s",data.Surname,data.Name,value2)
		Fac5 := fmt.Sprintf("%s%s%s",data.Surname,value2,data.Name)
		Fac6 := fmt.Sprintf("%s%s%s",value2,data.Surname,data.Name)
		ListMixer = append(ListMixer, Fac1)

		ListMixer = append(ListMixer, Fac2)

		ListMixer = append(ListMixer, Fac3)

		ListMixer = append(ListMixer, Fac4)

		ListMixer = append(ListMixer, Fac5)

		ListMixer = append(ListMixer, Fac6)

		
		
	}
	
	// 3! x n = 6n different strings
	for _ , value2 := range data.Others {
		Fac1 := fmt.Sprintf("%s%s%s",data.Name,data.Surname,value2)
		Fac2 := fmt.Sprintf("%s%s%s",data.Name,value2,data.Surname)
		Fac3 := fmt.Sprintf("%s%s%s",value2,data.Name,data.Surname)
		Fac4 := fmt.Sprintf("%s%s%s",data.Surname,data.Name,value2)
		Fac5 := fmt.Sprintf("%s%s%s",data.Surname,value2,data.Name)
		Fac6 := fmt.Sprintf("%s%s%s",value2,data.Surname,data.Name)

		ListMixer = append(ListMixer, Fac1)

		ListMixer = append(ListMixer, Fac2)

		ListMixer = append(ListMixer, Fac3)

		ListMixer = append(ListMixer, Fac4)

		ListMixer = append(ListMixer, Fac5)

		ListMixer = append(ListMixer, Fac6)

	}
	return ListMixer
}
// It's add to speacil chars 
func SpecialTwoWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	// 2! + 2! = 4 * 3 = 12 different strings
	for _, value1 := range data.Bdate {
		Fac1 := fmt.Sprintf("%s%s",data.Name,value1)
		Fac2 := fmt.Sprintf("%s%s",data.Surname,value1)
		Fac3 := fmt.Sprintf("%s%s",value1,data.Name)
		Fac4 := fmt.Sprintf("%s%s", value1,data.Surname)
				
		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac1)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac2)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac3)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac4)...)
		
	} 
	// 2! + 2! = 4 * n = 4n different strings
	for _, value2 := range data.Others {
		Fac1 := fmt.Sprintf("%s%s",data.Name,value2)
		Fac2 := fmt.Sprintf("%s%s",data.Surname,value2)
		Fac3 := fmt.Sprintf("%s%s",value2,data.Name)
		Fac4 := fmt.Sprintf("%s%s",value2,data.Surname)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac1)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac2)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac3)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac4)...)
	

	}
	return ListMixer
}

func SpecialThreeWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	// 3! x 3 = 18 different strings 
	for _ , value2 := range data.Bdate {
		Fac1 := fmt.Sprintf("%s%s%s",data.Name,data.Surname,value2)
		Fac2 := fmt.Sprintf("%s%s%s",data.Name,value2,data.Surname)
		Fac3 := fmt.Sprintf("%s%s%s",value2,data.Name,data.Surname)
		Fac4 := fmt.Sprintf("%s%s%s",data.Surname,data.Name,value2)
		Fac5 := fmt.Sprintf("%s%s%s",data.Surname,value2,data.Name)
		Fac6 := fmt.Sprintf("%s%s%s",value2,data.Surname,data.Name)
		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac1)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac2)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac3)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac4)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac5)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac6)...)


		

		
		
	}
	//Sector - 5
	for _ , value2 := range data.Others {
		Fac1 := fmt.Sprintf("%s%s%s",data.Name,data.Surname,value2)
		Fac2 := fmt.Sprintf("%s%s%s",data.Name,value2,data.Surname)
		Fac3 := fmt.Sprintf("%s%s%s",value2,data.Name,data.Surname)
		Fac4 := fmt.Sprintf("%s%s%s",data.Surname,data.Name,value2)
		Fac5 := fmt.Sprintf("%s%s%s",data.Surname,value2,data.Name)
		Fac6 := fmt.Sprintf("%s%s%s",value2,data.Surname,data.Name)
		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac1)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac2)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac3)...)
		
		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac4)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac5)...)

		ListMixer = append(ListMixer, AddtoSpecialChrs(Fac6)...)

	}
	return ListMixer
}

func RemovedVowelsTwoWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	// 2! + 2! = 4 |  4*3 = 12 different strings
	for _, value1 := range data.Bdate {
		Fac1 := fmt.Sprintf("%s%s",RemoveTheVowels(data.Name),value1)   
		Fac2 := fmt.Sprintf("%s%s",value1,RemoveTheVowels(data.Name))  
		Fac3 := fmt.Sprintf("%s%s",RemoveTheVowels(data.Surname),value1) 
		Fac4 := fmt.Sprintf("%s%s", value1,RemoveTheVowels(data.Surname)) 
		ListMixer = append(ListMixer, Fac1)

		ListMixer = append(ListMixer, Fac2)

		ListMixer = append(ListMixer, Fac3)

		ListMixer = append(ListMixer, Fac4)
	} 
	//2! + 2! = 4 |  4*n = 4n different strings
	for _, value2 := range data.Others {
		Fac1 := fmt.Sprintf("%s%s",RemoveTheVowels(data.Name),value2)
		Fac2 := fmt.Sprintf("%s%s",value2,RemoveTheVowels(data.Name))
		Fac3 := fmt.Sprintf("%s%s",RemoveTheVowels(data.Surname),value2)
		Fac4 := fmt.Sprintf("%s%s",value2,RemoveTheVowels(data.Surname))
		ListMixer = append(ListMixer, Fac1)

		ListMixer = append(ListMixer, Fac2)

		ListMixer = append(ListMixer, Fac3)

		ListMixer = append(ListMixer, Fac4)

	}

	return ListMixer
}
func RemovedVowelsThreeWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	// General Passowrd Examples
	for _ , value2 := range data.Bdate {
		Fac1 := fmt.Sprintf("%s%s%s" , RemoveTheVowels(data.Name),RemoveTheVowels(data.Surname),value2)
		Fac2 := fmt.Sprintf("%s%s%s",data.Name,RemoveTheVowels(data.Surname),value2)
		Fac3 := fmt.Sprintf("%s%s%s",value2,RemoveTheVowels(data.Name),RemoveTheVowels(data.Surname))
		Fac4 := fmt.Sprintf("%s%s%s",value2,data.Name,RemoveTheVowels(data.Surname))
		Fac5 := fmt.Sprintf("%s%s%s",data.Name,value2,RemoveTheVowels(data.Surname))
		Fac6 := fmt.Sprintf("%s%s%s",data.Surname,RemoveTheVowels(data.Name),value2)
		
		
		//append the list 
		ListMixer = append(ListMixer, Fac1)

		ListMixer = append(ListMixer, Fac2)

		ListMixer = append(ListMixer, Fac3)
		
		ListMixer = append(ListMixer, Fac4)

		ListMixer = append(ListMixer, Fac5)

		ListMixer = append(ListMixer, Fac6)


		
		
	}
	//Sector - 4
	for _ , value2 := range data.Bdate {
		Sp12 := fmt.Sprintf("%s%s%s",data.Surname,data.Name,value2)
		Sp13 := fmt.Sprintf("%s%s%s",data.Surname,value2,data.Name)
		Sp14 := fmt.Sprintf("%s%s%s",value2 ,data.Surname,data.Name)
		ListMixer = append(ListMixer, Sp12)

		ListMixer = append(ListMixer, Sp13)

		ListMixer = append(ListMixer, Sp14)

	
	}
	//Sector - 5
	for _ , value2 := range data.Others {
		Sp14 := fmt.Sprintf("%s%s%s",data.Name,data.Surname,value2)
		Sp15 := fmt.Sprintf("%s%s%s",data.Name,value2,data.Surname)
		Sp16 := fmt.Sprintf("%s%s%s",value2,data.Name,data.Surname)
		ListMixer = append(ListMixer, Sp14)

		ListMixer = append(ListMixer, Sp15)

		ListMixer = append(ListMixer, Sp16)

	}
	return ListMixer
}