package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dave/jennifer/jen"
)

func generateSlice(g *jen.Group, values []float64) {
	g.Index().Float64().ValuesFunc(func(g *jen.Group) {
		for _, v := range values {
			g.Lit(v)
		}
	})
}

func main() {
	f := jen.NewFile(filepath.Base(os.Args[0]))

	testCases := []struct {
		name string
		x    []float64
		y    []float64
	}{
		{"Test1", []float64{1, 2, 3}, []float64{4, 5, 6}},
		{"Test2", []float64{7, 8, 9}, []float64{10, 11, 12}},
	}

	for _, testCase := range testCases {
		f.Func().Id("TestCalculateDiffs" + testCase.name).Params(jen.Id("t").Op("*").Qual("testing", "T")).BlockFunc(func(g *jen.Group) {
			g.Id("x").Op(":=")
			generateSlice(g, testCase.x)
			g.Id("y").Op(":=")
			generateSlice(g, testCase.y)
			g.Id("meanX").Op(",").Id("_").Op(":=").Qual("github.com/montanaflynn/stats", "Mean").Call(jen.Id("x"))
			g.Id("meanY").Op(",").Id("_").Op(":=").Qual("github.com/montanaflynn/stats", "Mean").Call(jen.Id("y"))
			g.Id("diffXY").Op(",").Id("diffXX").Op(",").Id("diffYY").Op(":=").Id("CalculateDiffs").Call(jen.Id("x"), jen.Id("y"), jen.Id("meanX"), jen.Id("meanY"))
		})
	}

	err := f.Save("generated_tests.go")
	if err != nil {
		log.Fatal(err)
	}
}
