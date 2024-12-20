package appfuncs

import (
	"fmt"

)




//https://www.reddit.com/r/dataisbeautiful/comments/2vfgvh/most_frequentlyused_special_characters_in_10/ depens on this information
var SpecialChrs = []int{35,36,37,38,43,45,46,63,64,95}


var Vowels = []int{65,69,73,79,85,97,101,105,111,117}




//Add to special charachters end of the words and return a array.
func AddtoSpecialChrs(word string) []string {
	var array_words []string

	for _, c := range SpecialChrs {
		
		data := fmt.Sprintf("%s%s",word,string(rune(c)))
		data2 := fmt.Sprintf("%s%s",string(rune(c)),word)
		array_words = append(array_words, data)
		array_words = append(array_words, data2)
	}


	return array_words
}



//Remove the vowels from the word  and return a array 
func RemoveTheVowels(word string) string {
		
		sesliHarfler := map[byte]bool{
			'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
			'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
			'ö': true, 'Ö': true, 'ü': true, 'Ü': true, 
		}
	
		
		var sonuc []byte
	

		for i := 0; i < len(word); i++ {
			if !sesliHarfler[word[i]] { 
				sonuc = append(sonuc, word[i]) 
			}
		}
	

		return string(sonuc)
	}
	
