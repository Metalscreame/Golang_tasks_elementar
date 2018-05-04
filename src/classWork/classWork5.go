package classWork

import "fmt"

type Car struct {
	name string
	maxSpeed int
	minTurn  int
}

type Trailer struct {
	name string
	maxSpeed int
	minTurn  int
}

type CarTrain struct {
	maxSpeed int
	minTurn int
	theCar Car
	vihecles []Vihecles
}

type Vihecles interface {
	moveVehicleToTrain()
	getMaxSpeed()(int)
	getMinTurn()(int)
	getName()(string)
}

func (c * Car) moveVehicleToTrain()  {
	carTrain.vihecles=append(carTrain.vihecles, c)
}

func (t * Trailer) moveVehicleToTrain()  {
	carTrain.vihecles=append(carTrain.vihecles,t)
}

func (c * Car) getMaxSpeed()(int)  {
	return c.maxSpeed
}

func (t * Trailer) getMaxSpeed()(int)  {
	return t.maxSpeed
}

func (c * Car) getMinTurn()(int)  {
	return c.minTurn
}

func (t * Trailer) getMinTurn()(int)  {
	return t.minTurn
}

func (c * Car ) getName()(string)  {
	return c.name
}

func (t * Trailer) getName()(string)  {
	return t.name
}

func (c * CarTrain) setMaxSpeed(){
	for _,el :=range carTrain.vihecles{
		if carTrain.maxSpeed > el.getMaxSpeed(){
			carTrain.maxSpeed = el.getMaxSpeed()
		}
	}
}

func (c * CarTrain) setMinTurn(){
	for _,el :=range carTrain.vihecles{
		if carTrain.minTurn < el.getMinTurn(){
			carTrain.minTurn = el.getMinTurn()
		}
	}
}


func (c * CarTrain) getAllCarsInTheCarTrain(){
	fmt.Println("The CarTrain max Speed is: " , c.maxSpeed)
	fmt.Println("The CarTrain min turn is: " , c.minTurn)
	fmt.Printf("The number of vehicles in the CarTrain are: %d\n",len(c.vihecles)+1)
	fmt.Println("The vehicels in the carTrain are: ")
	fmt.Printf("Vehicle: %s with max speed %d with min turn %d\n",c.theCar.name,c.theCar.maxSpeed,c.theCar.minTurn)

	for _,el := range c.vihecles{
		fmt.Printf("Vehicle: %s with max speed %d with min turn %d\n",el.getName(),el.getMaxSpeed(),el.getMinTurn())
	}
}


var carTrain CarTrain

func mainCar(){
	carTrain.theCar = Car{"Tavria",120,30}
	carTrain.maxSpeed=carTrain.theCar.maxSpeed
	carTrain.minTurn = carTrain.theCar.minTurn

	var carOne Car = Car{"Ferrari",300,50}
	var carTwo Car = Car{"Mazda",250,10}
	var trailerOne Trailer = Trailer{"Huston",50,90}
	var trailerTwo Trailer = Trailer{"Peterbil",90,60}
	var trailerThree Trailer = Trailer{"Lobster",75,75}

	carOne.moveVehicleToTrain()
	carTwo.moveVehicleToTrain()
	trailerOne.moveVehicleToTrain()
	trailerTwo.moveVehicleToTrain()
	trailerThree.moveVehicleToTrain()


	carTrain.setMinTurn()
	carTrain.setMaxSpeed()

	carTrain.getAllCarsInTheCarTrain()

}