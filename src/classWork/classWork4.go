package classWork


import (
"fmt"
)

type Vagon struct {
	number        int
	svCount       int
	kupeCount     int
	plazCardCount int
}

type Train struct {
	number int
	from   string
	to     string
	vagons []Vagon
	vagonCount int
}

type TrainHelper interface {
	printVagonFreePlaces()
	buySV(vagon int)
	buyKUPE(vagon int)
	buyPLAZ(vagon int)
}

func (t* Train) printVagonFreePlaces() {
	fmt.Printf("\nThe train number is %d\n",t.number)
	for _,el := range t.vagons {
		fmt.Printf("The vagon number is %d\n", el.number)
		fmt.Printf("sv %d  ", el.svCount)
		fmt.Printf("kupeCount %d ", el.kupeCount)
		fmt.Printf("plazCardCount %d\n", el.plazCardCount)
	}
}

func (t * Train)buySV(vagon int)  {
	t.vagons[vagon].svCount--
	fmt.Printf("Bought sv ticket at vagon number %d\n",vagon)
}
func (t * Train)buyKUPE(vagon int)  {
	t.vagons[vagon].kupeCount--
	fmt.Printf("Bought kupe ticket at vagon number %d\n",vagon)

}
func (t * Train)buyPLAZ(vagon int)  {
	t.vagons[vagon].plazCardCount--
	fmt.Printf("Bought plazcard ticket at vagon number %d\n",vagon)

}

func trains() {

	var train Train = Train{1,"Dnepr", "Kiev",make([]Vagon,0), 5}



	for i:= 0;i<train.vagonCount;i++{
		var vag Vagon
		vag.number = i
		vag.svCount = 10
		vag.kupeCount = 10
		vag.plazCardCount = 5
		train.vagons  = append(train.vagons, vag)
	}

	train.printVagonFreePlaces()



	train.buySV(2)
	train.buyKUPE(3)
	train.buyKUPE(3)
	train.buyKUPE(3)
	train.buyKUPE(3)
	train.buySV(1)

	train.buyPLAZ(1)
	train.buyPLAZ(1)
	train.buyPLAZ(1)
	train.buyPLAZ(1)

	train.printVagonFreePlaces()


}
