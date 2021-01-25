//package main

//import "fmt"
//import stephencrypt "github.com/bradfield-csi-2020/szerfas_CSI/advanced_programming/6_packages_and_modules"
//
//func main() {
//	input := "aaa"
//	fmt.Printf("input text: %s\n", input)
//	fmt.Printf("ciphered text: %s\n", stephencrypt.StephenCrypt(input))
//	fmt.Printf("decciphereed text: %s\n", stephencrypt.StephenDecrypt(stephencrypt.StephenCrypt(input)))
//}

package stephencrypt

//const KEY = `STEPHEN`

func StephenCrypt(clearText string) string {
	var result string
	//var keyIndex int
	//for i, v := range clearText {
	for _, v := range clearText {
		//keyIndex = i % len(KEY)
		//result = result + string(v + int32(KEY[keyIndex]))
		result = result + string(v + 1)
	}
	return result
}

func StephenDecrypt(cipherText string) string {
	var result string
	for _, v := range cipherText {
		result = result + string(v - 1)
	}
	return result
}
