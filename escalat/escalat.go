// declarar meu pacote principal
package main

//importar função fmt
import "fmt"

//declaração da variavel CONST do ponto de ebulição da água em K
const ebulicaoK = 373.0

//função principal
func main() {

	tempK := ebulicaoK       //variavel do valor da temperatura em K
	tempC := (tempK - 273.0) //converção de K para C

	fmt.Printf("A temperatura de ebulição da água em K = %g , temperatura de ebulição da água em °C =%g.", tempK, tempC)
	//é esperado que apareça K = 373 e C = 100

}
