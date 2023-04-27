package graph

import (
	"fmt"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

func Draw(yAxisName string, values []float64 ) {
	t := time.Now()


	graph := chart.Chart{
		YAxis: chart.YAxis{
			Name: "Temperature",
			Range: &chart.ContinuousRange{
				Min: 0.0,
				Max: 4.0,
			},
			Ticks: []chart.Tick{
				{Value: 0.0, Label: "0 °C"},
				{Value: 1.0, Label: "10 °C"},
				{Value: 2.0, Label: "20 °C"},
				{Value: 3.0, Label: "30 °C"},
				{Value: 4.0, Label: "40 °C"},
				{Value: 5.0, Label: "50 °C"},
				{Value: 6.0, Label: "60 °C"},
				{Value: 7.0, Label: "70 °C"},
				{Value: 8.0, Label: "80 °C"},
				{Value: 9.0, Label: "90 °C"},
				{Value: 10.0, Label: "100 °C"},
			},
		},
		XAxis: chart.XAxis{
			Name : "Time",
			Ticks: []chart.Tick{
				{Value: 0.0, Label: timeBack(t, 60)},
				{Value: 1.0, Label: timeBack(t, 55)},
				{Value: 2.0, Label: timeBack(t, 50)},
				{Value: 3.0, Label: timeBack(t, 45)},
				{Value: 4.0, Label: timeBack(t, 40)},
				{Value: 5.0, Label: timeBack(t, 35)},
				{Value: 6.0, Label: timeBack(t, 30)},
				{Value: 7.0, Label: timeBack(t, 25)},
				{Value: 8.0, Label: timeBack(t, 20)},
				{Value: 9.0, Label: timeBack(t, 15)},
				{Value: 10.0, Label: timeBack(t, 10)},
				{Value: 11.0, Label: timeBack(t, 5)},
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 5.7, 5.9, 6.1, 6,3, 7.5, 4.9, 4.9},
			},
		},
	}
	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func timeBack(t time.Time, minutes int)string{
	rt := t.Add(time.Duration(-minutes) * time.Minute)
	return fmt.Sprintf("%v:%v", rt.Hour(), rt.Minute())
}
