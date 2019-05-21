package main

import (
	"fmt"

	"github.com/importcjj/sensitive"
)

func main() {
	filter := sensitive.New()
	filter.LoadWordDict("../dict/dict.txt")
	filter.AddWord("长者")

	fmt.Println(filter.Filter("我为长者续一秒"))       // 我为续一秒
	fmt.Println(filter.Replace("我为长者续一秒", '*')) // 我为**续一秒
	fmt.Println(filter.FindIn("我为长者续一秒"))       // true, 长者
	fmt.Println(filter.Validate("我为长者续一秒"))     // False, 长者
	fmt.Println(filter.FindAll("我为长者续一秒"))      // [长者]

	fmt.Println(filter.FindIn("我为长x者续一秒")) // false
	filter.UpdateNoisePattern(`x`)
	fmt.Println(filter.FindIn("我为长x者续一秒"))   // true, 长者
	fmt.Println(filter.Validate("我为长x者续一秒")) // False, 长者
}
