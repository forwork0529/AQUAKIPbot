package structs

import "sync"

const(
	OK = "OK"
	NO = "НЕТ"
	ON = "В РАБОТЕ"
	OFF = "ВЫКЛ"
)

type TypeVars struct{
	ConnectionState	string
	AlarmsMutex		sync.Mutex
	Alarms    		map[string]int
	Boiler1State   	string
	Boiler2State  	string
	TempOutdoorNow  float64
	GVSTMutex		sync.Mutex
	GVSTArray      	[]float64
	HTMutex			sync.Mutex
	HTArray       	[]float64
	SystemPress   	float64
	ReCoMutex     	sync.Mutex
	RechargeCount 	int
}

var (

	Vars = TypeVars{
		ConnectionState :   OK,
		AlarmsMutex		:   sync.Mutex{},
		Alarms   		:   map[string]int{"Котёл1":0, "Котёл2":0, "Насос ГВС1":0, "Насос ГВС2":0, "Насос ОТ1":0, "Насос ОТ2":0  },
		Boiler1State   	:   OFF,
		Boiler2State  	:   OFF,
		TempOutdoorNow 	:   25.0,
		GVSTMutex       :	sync.Mutex{},
		GVSTArray      	:   []float64{50,50,50,50,50,50,50,50,50,50,50,50},
		HTMutex			: 	sync.Mutex{},
		HTArray			:   []float64{50,50,50,50,50,50,50,50,50,50,50,50},
		SystemPress    	: 	3.0,
		ReCoMutex		:   sync.Mutex{},
		RechargeCount  	:   0,
	}
)
