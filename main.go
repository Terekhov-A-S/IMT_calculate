package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("############################################\n" +
		"#  Калькулятор расчета индекса массы тела  #\n" +
		"############################################")

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
		fmt.Println("У вас сильный дефицит массы тела")
	case imt >= 16 && imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt >= 18.5 && imt < 25:
		fmt.Println("Ваша масса тела в норме")
	case imt >= 25 && imt < 30:
		fmt.Println("У вас избыточная масса тела")
	case imt >= 30 && imt < 35:
		fmt.Println("У вас первая степень ожирения")
	case imt >= 35 && imt < 40:
		fmt.Println("У вас вторая степень ожирения")
	case imt >= 40:
		fmt.Println("У вас третья степень ожирения")
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
