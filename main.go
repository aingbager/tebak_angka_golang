package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var lanjut, stop string
	fmt.Print("tekan y untuk menjalankan program: ")
	fmt.Scanln(&lanjut)

	if lanjut == "y" || lanjut == "Y" {
		for {
			var angka int = rand.Intn(9) + 1
			var pilih int

			fmt.Println("======= tebak angka======")
			fmt.Print("pilih angka 1-9:")
			fmt.Scanln(&pilih)

			if pilih == angka {
				fmt.Println("anda benar")
				fmt.Println("angka acak:", angka)
			} else {
				fmt.Println("anda salah")
				fmt.Println("angka acak:", angka)
			}
			//tekan y untuk break;
			fmt.Print("tekan y untuk berhenti: ")
			fmt.Scanln(&stop)

			if stop == "y" || stop == "Y" {
				fmt.Println("program selesai")
				break
			}
		}
	} else {
		fmt.Println("program selesai")
	}

}
