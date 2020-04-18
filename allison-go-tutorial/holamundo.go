package main

import (
//    "bufio"
    "fmt"
 //   "os"
    //"strconv"

)

func main() {

    //nombre := bufio.NewReader(os.Stdin)
    //fmt.Print("Como te llamas: ")
    //sunombre, _ := nombre.ReadString('\n')

    //edad := bufio.NewReader(os.Stdin)
    //fmt.Print("Que edad tienes: ")
    //suedad, _ := edad.ReadString('\n')

    fmt.Print("Que tabla de dividir quieres ver: ")
    var tabla float64
    fmt.Scan(&tabla)
    fmt.Println("Generando tabla ", tabla, " de dividir")


    var float64 suma
    for suma < 10 {
        suma += 1
        fmt.Println(tabla," / ",suma," = ",suma/tabla )

   }
	
    //fmt.Print("Hola " + sunombre)
    //fmt.Println("Tienes " + suedad)

    //if tabla == 1 {
    //   fmt.Println("1 x 1 = 1")
    //   fmt.Println("1 x 2 = 2")
    //   fmt.Println("1 x 3 = 3")

    //} else if tabla == 2 {
    //   fmt.Println("2 x 1 = 2")
    //  fmt.Println("2 x 2 = 4")
    //  fmt.Println("2 x 3 = 6")

    //} else {
    //    fmt.Println("No me la tabla del ", tabla)
    //}

}

