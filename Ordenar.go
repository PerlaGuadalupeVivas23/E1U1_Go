import (
	"log"
	"os"
    "bufio"
	"fmt")
func main() {
	var palabras [5]string
	palabras[0] = "Perla"
	palabras[1] = "estudia"
	palabras[2] = "en"
    palabras[3] = "el"
    palabras[4] = "Tecnologico"
	var n [5]string

	file, err := os.Open("datos.txt")
	if err != nil {
		log.Fatal(err)
	}defer file.Close()
	scanner := bufio.NewScanner(file)
	var contador= 0
    fmt.Println("             O R D E N A R  O R A C I O N")
	fmt.Println("=============Palabras de Archivo txt=============")
	for scanner.Scan() {
fmt.Println(scanner.Text())
		n[contador] = scanner.Text()
		contador++
		//fmt.Println(scanner.Bytes())	}
	var palabraCorrecta [5]string
	for x := 0; x < 5; x++ {
		for o:= 0; o < 5; o++ {
			if a[x] == b[o] {
				palabraCorrecta[x] = a[x]}
                }	}
fmt.Println("=========== Palabras Ordenadas=============")
	for i := 0; i < 5; i++ {
		fmt.Println(palabraCorrecta[i])
	}
}
