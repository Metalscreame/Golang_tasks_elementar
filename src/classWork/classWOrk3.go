package classWork

import "fmt"


//types для своих переменных, как джеенерики

// можно привязать функции к структуре
//интерфейсы для описывания поведения. для заданвания поведения для одинаковых обьектов.
// можно написать функцию для элемента интерфейса. принимать экземпрял

//интерфейс животны, понятия движегния , понятия питание, звуки, реализовать для млекопитающих, птиц, рыбы,




type Animals interface {
	move()
	eat()
	voice()
}


type Fish struct {
	name string
}

type Mammal struct {
	name string
}

type BirdIsFlying struct {
	name string
}

type BirdCantFly struct {
	name string
}



func eater(animals Animals) {
	animals.eat()
}

func speaker(animals Animals)  {
	animals.voice()
}

func movements(animals Animals)  {
	animals.move()
}


func (f * Fish) eat() {
	fmt.Printf("%s is eating",f.name)
	fmt.Println()
}

func (m * Mammal) eat() {
	fmt.Printf("%s is eating %s",m.name)
	fmt.Println()
}
func (b * BirdIsFlying)eat()  {
	fmt.Printf("%s that fly is eating",b.name)
	fmt.Println()
}
func (b* BirdCantFly) eat(){
	fmt.Printf("%s that cant fly is eating",b.name)
	fmt.Println()
}


func (b * BirdIsFlying)move()  {
	fmt.Printf("Im a %s that flying",b.name)
	fmt.Println()
}

func (b* BirdCantFly) move(){
	fmt.Printf("Im a %s that walks",b.name)
	fmt.Println()
}

func (m* Mammal) move(){
	fmt.Printf("%s walks",m.name)
	fmt.Println()
}
func (f* Fish) move(){
	fmt.Printf("%s moves",f.name)
	fmt.Println()
}

func (b * BirdIsFlying)voice()  {
	fmt.Printf("%s Flying and screams",b.name)
	fmt.Println()

}

func (b* BirdCantFly) voice(){
	fmt.Printf("%s Running and singing",b.name)
	fmt.Println()

}

func (m* Mammal) voice(){
	fmt.Printf("%s screams",m.name)
	fmt.Println()

}
func (f* Fish) voice(){
	fmt.Printf("%s doesnt do sound",f.name)
	fmt.Println()
}



func classWork3() {

	akula := Fish{"akula"}

	karas:=Fish{"karas"}

	penuin:=BirdCantFly{"penguin"}

	lastochka:=BirdIsFlying{"lastochka"}


	straus:=BirdCantFly{"straus"}

	human:=Mammal{"SuperMan"}


	slice:=[...]Animals{&akula,& karas,& penuin, &lastochka, &straus,&human}


	slice1:=make([]Animals,0)
	slice1 = append(slice1,&human)

	for _,el :=range slice{
		movements(el)
		speaker(el)
		eater(el)
	}
}

