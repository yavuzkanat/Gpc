package main

import (
	"flag"
	"fmt"
	"gpcpacks/appfuncs"
	"gpcpacks/userfuncs"
	"os"

)


func main()  {
	runner := flag.Bool("r",false,"it's run the app.")
	print := flag.Bool("p",false,"it's print the output.")
	basics := flag.Bool("b",false,"it's create a normal list.")
	output := flag.Bool("o",false,"it's create output as txt file.")
	info := flag.Bool("i",false,"it's print infos about the output.")
	special := flag.Bool("s",false,"it's add the special chars.")
	vowels := flag.Bool("v",false,"it's remove the vowels chars.")
	flag.Parse()


	if *runner {
		if !(*output || *print){
			fmt.Println("Please check your flags (go run main.go -r -p or go run main.go -r -o)")
			os.Exit(1)
		}
		data := userfuncs.GetUserDatas()
		basic := appfuncs.TwoWordComb(data)
		basic2  := appfuncs.ThreeWordComb(data)
		specials := appfuncs.SpecialTwoWordComb(data)
		specials2 := appfuncs.SpecialThreeWordComb(data)
		Unvowels := appfuncs.RemovedVowelsTwoWordComb(data)
		Unvowles2 := appfuncs.RemovedVowelsThreeWordComb(data)

		
		if *print && *basics {
			for _, v := range basic {
				fmt.Println(v)
			}
			for _, i := range basic2 {
				fmt.Println(i)
			}
			
			if *info {
				fmt.Println("=============Total===============")
				fmt.Println(len(basic)+len(basic2),"words")
			}
		}
		if *print && *special {
			for _, v := range basic {
				fmt.Println(v)
			}
			for _, i := range basic2 {
				fmt.Println(i)
			}
			for _, t := range specials {
				fmt.Println(t)
			}
			for _, m := range specials2 {
				fmt.Println(m)
			}
			if *info {
				fmt.Println("=============Total===============")
				fmt.Println(len(basic)+len(basic2)+len(specials2)+len(specials),"words")
			}
		
		
		}

		if *print && *vowels {
			for _, v := range basic {
				fmt.Println(v)
			}
			for _, i := range basic2 {
				fmt.Println(i)
			}
	
			for _, z := range Unvowels {
				fmt.Println(z)
			}
			for _, q := range Unvowles2 {
				fmt.Println(q)
			}
			if *info {
				fmt.Println("=============Total===============")
				fmt.Println(len(basic)+len(basic2)+len(Unvowels)+len(Unvowels),"words")
			}
		
			
		}
		if *print && *vowels && *special {
			for _, v := range basic {
				fmt.Println(v)
			}
			for _, i := range basic2 {
				fmt.Println(i)
			}
			for _, t := range specials {
				fmt.Println(t)
			}
			for _, m := range specials2 {
				fmt.Println(m)
			}	
			for _, z := range Unvowels {
				fmt.Println(z)
			}
			for _, q := range Unvowles2 {
				fmt.Println(q)
			}
			if *info {
				fmt.Println("=============Total===============")
				fmt.Println(len(basic)+len(basic2)+len(specials2)+len(specials)+len(Unvowels)+len(Unvowles2),"words")
			}
		
		}
		if *output && *basics {
			var alldata []string
			alldata = append(alldata,basic...)
			alldata = append(alldata,basic2...)

			appfuncs.OutputAsTXT(alldata)
			if *info {
				fmt.Println("=============Total===============")
				fmt.Println(len(basic)+len(basic2),"words")
			}
		}
		if *output && *special {
			var alldata2 []string
			alldata2 = append(alldata2,basic...)
			alldata2 = append(alldata2,basic2...)
			alldata2 = append(alldata2, specials...)
			alldata2 = append(alldata2, specials2...)

			appfuncs.OutputAsTXT(alldata2)
			if *info {
				fmt.Println("=============Total===============")
				fmt.Println(len(basic)+len(basic2)+len(specials2)+len(specials),"words")
			}

			
		}
		if *output && *vowels {
			var alldata3 []string
			alldata3 = append(alldata3, basic... )
			alldata3 = append(alldata3, basic2...)
			alldata3 = append(alldata3, Unvowels...)
			alldata3 = append(alldata3, Unvowles2...)

			appfuncs.OutputAsTXT(alldata3)
		}
		if *output && *vowels && *special {
			var alldata4 []string
			alldata4 = append(alldata4, basic...)
			alldata4 = append(alldata4, basic2...)
			alldata4 = append(alldata4, specials...)
			alldata4 = append(alldata4, specials2...)
			alldata4 = append(alldata4, Unvowels...)
			alldata4 = append(alldata4, Unvowles2...)

			appfuncs.OutputAsTXT(alldata4)
		}
	}else if !*runner{
		flag.Usage()

	}

}
