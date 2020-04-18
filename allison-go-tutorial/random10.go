package main
import (
    "fmt" //libreria para accesar Print()
    "math/rand" //libreria para accesar rand()
    "time" //libreria para accesar unix time()
)
func main() {
 for bucle := 1; bucle <= 8; bucle++ { //bucle se repite 8 veses
  for bucle2 := 1; bucle2 <= 8; bucle2++ { //bucle se repite otras 8 veses
	  for numero := 1; numero <= 8; numero++ { //bucle genera 8 numeros aleatorios 8 veses
    semilla := rand.NewSource(time.Now().UnixNano()) //agarra nueva semilla para generar aleatorios basado en hora actual del cpu
    nueva_semilla := rand.New(semilla) //asigna la semilla semilla al semilla generador de numeros
    fmt.Print(nueva_semilla.Intn(2)) //genera un numero aleatorio entre 0 y 1, imprimelo
   } //fin bucle numero
   fmt.Print(" ")
	 } //fin bucle2
		fmt.Println() //imprime linea en blanco
	} //fin bucle
} //fin func main()

