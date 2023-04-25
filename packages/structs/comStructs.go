package structs

type TypeVars struct{
	SystemState    string
	NumOfAlarms    int
	Boiler1State   string
	Boiler1Alarm   int
	Boiler2State   string
	Boiler2Alarm   int
	TempHeaterNow  float64
	TempGVSNow     float64
	TempOutdoorNow float64
	HArray         []int
	HWSArray       []int
	SystemPress    float64
	RechargeCount  int
}

var (
	Vars = TypeVars{
		SystemState :  "РАБОТАЕТ",
		NumOfAlarms :  0,
		Boiler1State : "РАБОТАЕТ",
		Boiler1Alarm : 0,
		Boiler2State : "СТОП",
		Boiler2Alarm : 0,
		HArray :       []int{50, 51, 52, 51, 52},
		HWSArray :     []int{60, 61, 62, 61, 62},
		SystemPress:   2.0,
		RechargeCount: 3,
	}
)
