package computer

import (
	"AquaBot/packages/structs"
	"strconv"
	"strings"
	"sync"
	"time"
)

func New(input <- chan string, vars *structs.TypeVars){

/*
	type TypeVars struct{

	ConnectionState	string				  	соединение с ПЛК
	Alarms    		map[string]int			словарь аварий
									B1A		котёл 1 авария
									B2A		котёл 2 авария
	Boiler1State   	string			B1S		состояние котёл 1
	Boiler2State  	string			B2S		состояние котёл 2
									THN		температура отопления сейчас
									TGN		температура ГВС сейчас
	TempOutdoorNow 	float64			TON		температура снаружи сейчас
	GVSTArray      	[]float64
	HTArray       	[]float64
	SystemPress    	float64			SP		давление в системе
	RechargeCount  	int				RC		счётчик запусков подпитки

	}
	*/
	go func(){
		for {
			select{
			case inputString := <- input:
				res := strings.Split(inputString, " ")
				if len(res) < 2{
					continue
				}
				if vars.ConnectionState == structs.NO{
					vars.ConnectionState = structs.OK
				}
				switch res[0]{

				case "B1S" : StrToStrWork(res[1],&vars.Boiler1State)
				case "B1A" : StrToIntAlarm(res[0], res[1], vars)
				case "B2S" : StrToStrWork(res[1],&vars.Boiler2State)
				case "B2A" : StrToIntAlarm(res[0], res[1], vars)
				case "THN" : StrToFloat64Save(res[1], &vars.HTMutex, &vars.HTArray[11])
				case "TGN" : StrToFloat64Save(res[1],&vars.GVSTMutex, &vars.GVSTArray[11])
				case "TON" : StrToFloat64(res[1],&vars.TempOutdoorNow)
				case "SP" : StrToFloat64(res[1], &vars.SystemPress)
				case "RC" : RechargeCount(vars)
				}

			case <- time.After(time.Second * 15) :
				if	vars.ConnectionState == structs.OK{
					vars.ConnectionState = structs.NO
					vars.TempOutdoorNow = 0
					vars.GVSTMutex.Lock()
					vars.GVSTArray[11] = 0
					vars.GVSTMutex.Unlock()
					vars.HTMutex.Lock()
					vars.HTArray[11] = 0
					vars.HTMutex.Unlock()
					vars.SystemPress = 0
				}
			}

		}
	}()
	go RotationArraysSave(vars)
}

func StrToIntAlarm(command, str string, vars *structs.TypeVars){
	res, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	// "Котёл1":0, "Котёл2":0, "Насос ГВС1":0, "Насос ГВС2":0, "Насос ОТ1":0, "Насос ОТ2":0
	switch command {
	case "B1A" :
		vars.AlarmsMutex.Lock()
		if res > 0 {vars.Alarms["Котёл1"] = 1}else{vars.Alarms["Котёл1"] = 0}
		vars.AlarmsMutex.Unlock()
	case "B2A" :
		vars.AlarmsMutex.Lock()
		if res > 0 {vars.Alarms["Котёл2"] = 1}else{vars.Alarms["Котёл2"] = 0}
		vars.AlarmsMutex.Unlock()
	}
	return
}

func StrToStrWork(str string, var1 *string){
	res, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	if res > 0{
		if *var1 != structs.ON{
			*var1 = structs.ON
		}
	}else{
		if *var1 != structs.OFF{
			*var1 = structs.OFF
		}
	}
}


func StrToFloat64(str string, var1 *float64){
	res, err := strconv.ParseFloat(str, 64)
	if err != nil{
		return
	}
	if res < 0 || res > 100{
		return
	}
	*var1 = res
	return
}

func StrToFloat64Save(str string, mu *sync.Mutex, var1 *float64){
	res, err := strconv.ParseFloat(str, 64)
	if err != nil{
		return
	}
	if res < 0 || res > 100{
		return
	}
	mu.Lock()
	*var1 = res
	mu.Unlock()
	return

}

func RechargeCount(vars *structs.TypeVars){
	vars.ReCoMutex.Lock()
	vars.RechargeCount += 1
	vars.ReCoMutex.Unlock()
	go func(){
		time.Sleep(time.Minute)
		vars.ReCoMutex.Lock()
		vars.RechargeCount -= 1
		vars.ReCoMutex.Unlock()
	}()
}

func RotationArraysSave(vars *structs.TypeVars){
	for{
		<- time.After(time.Second * 10)
		vars.GVSTMutex.Lock()
		RotationArray(vars.GVSTArray)
		vars.GVSTMutex.Unlock()
		vars.HTMutex.Lock()
		RotationArray(vars.HTArray)
		vars.HTMutex.Unlock()
	}
}

func RotationArray(array []float64){
	if len(array) < 2{
		return
	}
	if len(array) == 2{
		array[0] = 1
		array[0],array[1] = array[0], array[1]
	}
	for i := 0; i < len(array)-1; i++ {
		array[i] = array[i+1]
	}
}