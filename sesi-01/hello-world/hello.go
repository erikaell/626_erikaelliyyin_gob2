package main

import "fmt"

func main() {
	// Learn hello
	fmt.Println("Erika", "Elliyyin", "Belajar Ngoding")
	fmt.Println("=====================")
	/*
		Learn to print
		part 2
	*/
	fmt.Println("Hello")
	fmt.Println("Hello", "Erika")
	fmt.Println("=====================")
	fmt.Print("Hello Erika\n")
	fmt.Print("Hello", " Erika\n")
	fmt.Print("Hello", " ", "Erika\n")
	fmt.Println("=====================")

	// Variable with data type
	var namaLengkap string
	var umur int = 99

	namaLengkap = "Erika"
	umur = 23

	fmt.Println("Ini adalah namanya ==> ", namaLengkap)
	fmt.Println("Ini adalah umurnya ==> ", umur)
	fmt.Println("=====================")

	// Variable without data type
	var namaAnda = "Erika"
	var umurAnda = true

	fmt.Printf("%T, %T\n", namaAnda, umurAnda)
	fmt.Println("=====================")

	// Variable without data type (short declaration)
	namaSaya := "Erika"
	umurSaya := 23

	fmt.Printf("%T, %T\n", namaSaya, umurSaya)
	fmt.Println("=====================")

	// Multiple variable declaration
	var one, two, three string = "satu", "dua", "tiga"
	var angkaOne, angkaTwo, angkaThree int = 1, 2, 3
	fmt.Println("test data string multiple > ", one, two, three)
	fmt.Println("test data int multiple > ", angkaOne, angkaTwo, angkaThree)
	fmt.Println("=====================")

	// Multiple variable (short declaration)
	var oneData, twoData, threeData = 1, true, "Erika"
	dataOne, dataTwo, dataThree := true, 55.5, "Elliyyin"
	fmt.Println("test data short multiple > ", oneData, twoData, threeData)
	fmt.Println("test data short multiple > ", dataOne, dataTwo, dataThree)
	fmt.Println("=====================")

	// Underscore variable
	var testVariabel string
	var oneName, twoName, alamatSaya, dataUmur = "Erika", "Elliyyin", "Jakarta", 23
	_, _, _, _, _ = testVariabel, oneName, twoName, alamatSaya, dataUmur

	// fmt.Printf Usage
	fmt.Printf("test underscore variabel > %T\n", testVariabel)
	fmt.Printf("halo nama saya %s, umur saya adalah %d, saya beralamat di %s\n", oneName, dataUmur, alamatSaya)
}
