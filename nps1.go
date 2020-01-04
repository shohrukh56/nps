package main

import "fmt"

func main()  {

	score1 :=10
	score2 :=10
	score3 :=10

	promoters := 0
	promoter := promoters
	detractors := promoters

	if score1 >=9{
		promoter=promoter+1
	}
	if score2 <=6{
	detractors=detractors+1
	}
	if score3 >=9{
	promoter=promoter+1
	}
nps := (promoter-detractors)/3*100
fmt.Println(nps)
}