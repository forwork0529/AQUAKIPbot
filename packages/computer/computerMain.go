package computer

import (
	"AquaBot/packages/structs"
	"strconv"
	"strings"
)

func New(input <- chan string, vars *structs.TypeVars){

/*
	type TypeVars struct{
		SystemState string      SS      состояние системы
		NumOfAlarms int         NoA		общие аварии системы
		Boiler1State string		B1S		состояние котёл 1
		Boiler1Alarm int		B1A		котёл 1 аварии
		Boiler2State string		B2S		состояние котёл 2
		Boiler2Alarm int		B2A		котёл 2 аварии
		TempHeaterNow  			THN		температура отопления сейчас
		TempGVSNow     			TGN		температура ГВС сейчас
		TempOutdoorNow 			TON		температура снаружи сейчас
		HArray
		HWSArray
		SystemPress    			SP		давление в системе сейчас
		RechargeCount  			RC		счётчик пусков подпитки
	}
	*/
	go func(){
		var inputString string
		for inputString = range input{
			res := strings.Split(inputString, " ")
			if len(res) < 2{
				continue
			}
			switch res[0]{
			case "SS" : remVarToLocalVar(res[1], nil, nil, &vars.SystemState)
			case "NoA" : remVarToLocalVar(res[1], &vars.NumOfAlarms, nil, nil)
			case "B1S" : remVarToLocalVar(res[1], nil, nil, &vars.Boiler1State)
			case "B1A" : remVarToLocalVar(res[1], &vars.Boiler1Alarm, nil, nil)
			case "B2S" : remVarToLocalVar(res[1], nil, nil, &vars.Boiler2State)
			case "B2A" : remVarToLocalVar(res[1], &vars.Boiler2Alarm, nil, nil)
			case "THN" : remVarToLocalVar(res[1], nil, &vars.TempHeaterNow, nil)
			case "TGN" : remVarToLocalVar(res[1], nil, &vars.TempGVSNow, nil)
			case "TON" : remVarToLocalVar(res[1], nil, &vars.TempOutdoorNow, nil)
			case "SP" : remVarToLocalVar(res[1], nil, &vars.SystemPress, nil)
			case "RC" : remVarToLocalVar(res[1], &vars.RechargeCount,nil, nil)
			}
			if inputString == "alarm."{
				vars.Boiler1Alarm = 1
			}
			if inputString == "norma."{
				vars.Boiler1Alarm = 0
			}
		}
	}()
}

func remVarToLocalVar(str string, var1 *int, var2 *float64, var3 *string){
	if var1 != nil {
		res, err := strconv.Atoi(str)
		if err != nil {
			return
		}
		*var1 = res
		return
	}
	if var2 != nil{
		res, err := strconv.ParseFloat(str, 64)
		if err != nil{
			return
		}
		*var2 = res
		return
	}
	if var3 != nil{
		*var3 = str
		return
	}

}

