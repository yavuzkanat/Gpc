package appfuncs

import (
	"fmt"
	"gpcpacks/userfuncs"
	
)

//Two Word Combanation
func TwoWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	//Sector - 1
	for _, value1 := range data.Bdate {
		Sp1 := fmt.Sprintf("%s%s",data.Name,value1)
		Sp2 := fmt.Sprintf("%s%s",data.Surname,value1)
		Sp5 := fmt.Sprintf("%s%s",value1,data.Name)
		Sp6 := fmt.Sprintf("%s%s", value1,data.Surname)
		ListMixer = append(ListMixer, Sp1)
		ListMixer = append(ListMixer, Sp2)
		ListMixer = append(ListMixer, Sp5)
		ListMixer = append(ListMixer, Sp6 )
	} 
	//Sector - 2
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
	//Sector - 3 
	for _ , value2 := range data.Bdate {
		Sp9 := fmt.Sprintf("%s%s%s",data.Name,data.Surname,value2)
		Sp10 := fmt.Sprintf("%s%s%s",value2,data.Name,data.Surname)
		Sp11 := fmt.Sprintf("%s%s%s",data.Name,value2,data.Surname)
		ListMixer = append(ListMixer, Sp9)
		ListMixer = append(ListMixer, Sp10)
		ListMixer = append(ListMixer, Sp11)

		
		
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
// It's add to speacil chars 
func SpecialTwoWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	//Sector - 6
	for _, value1 := range data.Bdate {
		Sp1 := fmt.Sprintf("%s%s",data.Name,value1)
		Sp2 := fmt.Sprintf("%s%s",data.Surname,value1)
		Sp5 := fmt.Sprintf("%s%s",value1,data.Name)
		Sp6 := fmt.Sprintf("%s%s", value1,data.Surname)
				
		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp1)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp2)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp5)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp6)...)
		
	} 
	//Sector - 7
	for _, value2 := range data.Others {
		Sp3 := fmt.Sprintf("%s%s",data.Name,value2)
		Sp4 := fmt.Sprintf("%s%s",data.Surname,value2)
		Sp7 := fmt.Sprintf("%s%s",value2,data.Name)
		Sp8 := fmt.Sprintf("%s%s",value2,data.Surname)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp3)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp4)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp7)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp8)...)
	

	}
	return ListMixer
}

func SpecialThreeWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	//Sector - 3 
	for _ , value2 := range data.Bdate {
		Sp9 := fmt.Sprintf("%s%s%s",data.Name,data.Surname,value2)
		Sp10 := fmt.Sprintf("%s%s%s",value2,data.Name,data.Surname)
		Sp11 := fmt.Sprintf("%s%s%s",data.Name,value2,data.Surname)
		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp9)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp10)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp11)...)

		

		
		
	}
	//Sector - 4
	for _ , value2 := range data.Bdate {
		Sp12 := fmt.Sprintf("%s%s%s",data.Surname,data.Name,value2)
		Sp13 := fmt.Sprintf("%s%s%s",data.Surname,value2,data.Name)
		Sp14 := fmt.Sprintf("%s%s%s",value2 ,data.Surname,data.Name)
		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp12)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp13)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp14)...)

	

	
	}
	//Sector - 5
	for _ , value2 := range data.Others {
		Sp14 := fmt.Sprintf("%s%s%s",data.Name,data.Surname,value2)
		Sp15 := fmt.Sprintf("%s%s%s",data.Name,value2,data.Surname)
		Sp16 := fmt.Sprintf("%s%s%s",value2,data.Name,data.Surname)
		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp14)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp15)...)

		ListMixer = append(ListMixer, userfuncs.AddtoSpeacilChrs(Sp16)...)

	}
	return ListMixer
}

func RemovedVowelsTwoWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	//Sector - 1
	for _, value1 := range data.Bdate {
		Sp1 := fmt.Sprintf("%s%s",userfuncs.RemoveTheVowels(data.Name),value1)
		Sp2 := fmt.Sprintf("%s%s",userfuncs.RemoveTheVowels(data.Surname),value1)
		Sp5 := fmt.Sprintf("%s%s",value1,userfuncs.RemoveTheVowels(data.Name))
		Sp6 := fmt.Sprintf("%s%s", value1,userfuncs.RemoveTheVowels(data.Surname))
		ListMixer = append(ListMixer, Sp1)
		ListMixer = append(ListMixer, Sp2)
		ListMixer = append(ListMixer, Sp5)
		ListMixer = append(ListMixer, Sp6 )
	} 
	//Sector - 2
	for _, value2 := range data.Others {
		Sp3 := fmt.Sprintf("%s%s",userfuncs.RemoveTheVowels(data.Name),value2)
		Sp4 := fmt.Sprintf("%s%s",userfuncs.RemoveTheVowels(data.Surname),value2)
		Sp7 := fmt.Sprintf("%s%s",value2,userfuncs.RemoveTheVowels(data.Name))
		Sp8 := fmt.Sprintf("%s%s",value2,userfuncs.RemoveTheVowels(data.Surname))
		ListMixer = append(ListMixer, Sp3)
		ListMixer = append(ListMixer, Sp4)
		ListMixer = append(ListMixer, Sp7)
		ListMixer = append(ListMixer, Sp8)

	}

	return ListMixer
}
func RemovedVowelsThreeWordComb(data userfuncs.PersonalInfos) []string {
	var ListMixer []string
	//Sector - 3 
	for _ , value2 := range data.Bdate {
		Sp9 := fmt.Sprintf("%s%s%s" , userfuncs.RemoveTheVowels(data.Name),userfuncs.RemoveTheVowels(data.Surname),value2)
		Sp92 := fmt.Sprintf("%s%s%s",data.Name,userfuncs.RemoveTheVowels(data.Surname),value2)
		Sp10 := fmt.Sprintf("%s%s%s",value2,userfuncs.RemoveTheVowels(data.Name),userfuncs.RemoveTheVowels(data.Surname))
		Sp102 := fmt.Sprintf("%s%s%s",value2,data.Name,userfuncs.RemoveTheVowels(data.Surname))
		Sp11 := fmt.Sprintf("%s%s%s",data.Name,value2,userfuncs.RemoveTheVowels(data.Surname))
		Sp111 := fmt.Sprintf("%s%s%s",data.Surname,userfuncs.RemoveTheVowels(data.Name),value2)
		Sp112 := fmt.Sprintf("%s%s%s",data.Surname,value2,userfuncs.RemoveTheVowels(data.Name))
		Sp113 := fmt.Sprintf("%s%s%s",value2,data.Surname,userfuncs.RemoveTheVowels(data.Name))
		
		//append the list 
		ListMixer = append(ListMixer, Sp9)
		ListMixer = append(ListMixer, Sp92)
		ListMixer = append(ListMixer, Sp10)
		ListMixer = append(ListMixer, Sp102)
		ListMixer = append(ListMixer, Sp11)
		ListMixer = append(ListMixer, Sp111)
		ListMixer = append(ListMixer, Sp112)
		ListMixer = append(ListMixer, Sp113)

		
		
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