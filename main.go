package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("############################################\n" +
		"#  Калькулятор расчета индекса массы тела  #\n" +
		"#                                          #\n" +
		"#         Created by Terekhov A.S.         #\n" +
		"############################################\n")

	var userHeight, userKg float64 = getInput()

	IMT := calculateIMT(userHeight, userKg)
	outputResult(IMT)
	fmt.Println("Нажмите Enter, чтобы закрыть...")
	fmt.Scanln()

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.1f", imt)
	fmt.Println(result)
	switch {
	case imt < 16.0:
		fmt.Println("У вас выраженный дефицит массы тела")
	case imt >= 16.0 && imt <= 16.9:
		fmt.Println("У вас недостаточная масса тела (умеренный дефицит)")
	case imt >= 17.0 && imt <= 18.4:
		fmt.Println("У вас недостаточная масса тела (лёгкий дефицит)")
	case imt >= 18.5 && imt <= 24.9:
		fmt.Println("Ваша масса тела в норме")
	case imt >= 25.0 && imt <= 29.9:
		fmt.Println("У вас избыточная масса тела (предожирение)")
	case imt >= 30.0 && imt <= 34.9:
		fmt.Println("У вас ожирение I степени")
	case imt >= 35.0 && imt <= 39.9:
		fmt.Println("У вас ожирение II степени")
	case imt >= 40:
		fmt.Println("У вас ожирение III степени (морбидное)")
	}
}

func calculateIMT(userHeight, userKg float64) float64 {
	const IMTPOWER = 2
	IMT := userKg / math.Pow(userHeight/100, IMTPOWER)
	return IMT
}

func getInput() (float64, float64) {
	var userHeight, userKg float64
	fmt.Print("Введите ваш рост (в сантиметрах): ")
	fmt.Scanln(&userHeight)
	fmt.Print("Введите ваш вес (в килограммах (через точку \".\"): ")
	fmt.Scanln(&userKg)
	return userHeight, userKg
}
