package structs

type TypeVars struct{
	SystemState string
	NumOfAlarms int
	Boiler1State string
	Boiler1Alarm int
	Boiler2State string
	Boiler2Alarm int
	HArray []int
	HWSArray []int
	RechPress float64
	RechCount int
}

var (
	Vars = TypeVars{
		SystemState : "РАБОТАЕТ",
		NumOfAlarms : 0,
		Boiler1State :"РАБОТАЕТ",
		Boiler1Alarm : 0,
		Boiler2State : "СТОП",
		Boiler2Alarm : 0,
		HArray : []int{50, 51, 52, 51, 52},
		HWSArray : []int{60, 61, 62, 61, 62},
		RechPress : 2.0,
		RechCount : 3,
	}
)
