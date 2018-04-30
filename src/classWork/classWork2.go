package classWork

import (
	"fmt"
)

//types для своих переменных, как джеенерики

// можно привязать функции к структуре

type APP struct {
	name string
	cpu       CPU
	volumeHDD int
	ram       RAM
}

type Computer struct {
	CPU
	HDD
	RAM
	applications map[string]int // name and size
}

type CPU struct {
	name string
	MHz  int
}

type HDD struct {
	volume []int
}

type RAM struct {
	Gb int
}

var computer Computer


func classWork2() {

	computer.applications=make(map[string]int)

	fmt.Println("Enter CPU name: ")
	fmt.Scanf("%s\n",&computer.name)
	//	fmt.Println()

	fmt.Println("Enter CPU Mhz: ")
	fmt.Scanf("%d\n",&computer.MHz)
	//fmt.Println()

	var size int
	fmt.Println("Enter HDD volume: ")
	fmt.Scanf("%d\n",&size)


	fmt.Println("Enter RAM volume: ")
	fmt.Scanf("%d\n", &computer.RAM)

	computer.HDD.volume = make([]int, 0, size)

	var choise string
	for true {

		fmt.Println()
		_,err :=fmt.Scanf( "%s\n",&choise)
		if err!=nil{}

		if choise == "break" {
			break
		} else if choise == "add" {
			var newApp APP

			newApp.addApplication()
			if computer.checkNewApp(newApp) {
				fmt.Println("added")
				continue
			} else {
				fmt.Println("Cant add")
				continue
			}
		} else if choise == "show"{
			for _,el :=range computer.applications{
				fmt.Println(el)
			}

		} else if choise == "delete"{
			fmt.Println("what do u want to delete?")
			for _,el :=range computer.applications{
				fmt.Println(el)
			}
			fmt.Println()
			fmt.Scanf("%s",&choise)

			sizeToDel:=computer.applications[choise]
			computer.volume = computer.volume[0:len(computer.volume)-sizeToDel]
			delete(computer.applications, choise)
			fmt.Println("deleted")
			continue
		}else {
			fmt.Println("Wrong input")
			continue
		}
	}

}

func (a *APP) addApplication() {
	fmt.Print("Enter new App CPU mhz : ")
	fmt.Scanf("%d\n",a.cpu.MHz)
	fmt.Println()

	fmt.Print("Enter new App HDD volume: ")
	fmt.Scanf("%d\n",a.volumeHDD)
	fmt.Println()

	fmt.Print("Enter new App RAM volume: ")
	fmt.Scanf("%d\n",a.ram)
	fmt.Println()
}

func (computer *Computer) checkNewApp(newApp APP) (flag bool) {

	if newApp.ram.Gb > computer.RAM.Gb || newApp.cpu.MHz > computer.MHz {
		fmt.Println("Cant add this app because of RAM or CPU")
		return false
	}

	if cap(computer.HDD.volume)- len(computer.HDD.volume) > newApp.volumeHDD {
		for i := len(computer.HDD.volume); i < newApp.volumeHDD; i++ {
			computer.HDD.volume = append(computer.HDD.volume, 0)
		}
		computer.applications[newApp.name]= newApp.volumeHDD
		fmt.Println("added")
		return true
	} else {
		fmt.Println("cant add, size is too big")
		return false
	}

}
