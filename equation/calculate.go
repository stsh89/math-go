package equation

import (
	"fmt"
	"strings"

	"github.com/Pramod-Devireddy/go-exprtk"
)

type Point struct {
	X float64
	Y float64
}

func CalculateRange(term string, args map[string]float64, logger Logger) []Point {
	var points []Point

	for x := -5.0; x <= 5.0; x++ {
		args["x"] = x
		y := Calculate(term, args, logger)
		points = append(points, Point{X: x, Y: y})
	}

	return points
}

func Calculate(term string, args map[string]float64, logger Logger) float64 {
	logger.Info("Calculating equation...")
	parts := strings.Split(term, "y=")

	if parts[1] == "" {
		logger.Error("Equation should start with 'y='")
		return 0.0
	}

	exprtkObj := exprtk.NewExprtk()

	exprtkObj.SetExpression(parts[1])

	for k := range args {
		exprtkObj.AddDoubleVariable(k)
	}

	err := exprtkObj.CompileExpression()
	if err != nil {
		fmt.Println(err.Error())
		logger.Error(err.Error())
		return 0.2
	}

	for k, v := range args {
		exprtkObj.SetDoubleVariableValue(k, v)
	}

	return exprtkObj.GetEvaluatedValue()
}
