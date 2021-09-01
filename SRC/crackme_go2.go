package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"os"
)

func EncryptDecrypt(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}
	return output
}

func main() {

	fmt.Print("Enter Password: ")
	user_input, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		fmt.Println("Some wrong input error, tralala", err)
	}
	key := string([]byte{38, 13, 87, 34, 20, 12, 36, 11, 85, 36, 22, 10, 80, 80, 80, 80})
	decrypted := EncryptDecrypt(string(user_input), key)
	sEncoded := b64.StdEncoding.EncodeToString([]byte(decrypted))

	if string(sEncoded) == "YWJjZGFiY2RhYmNkYWJjZA==" {
		fmt.Println("You Win!!!")
	} else {
		fmt.Println("Wrong, Try Harder!!!")
	}
}
