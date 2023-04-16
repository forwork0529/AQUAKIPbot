package computer

import (
	"AquaBot/packages/structs"
	"fmt"
)

func New(input <- chan string, vars *structs.TypeVars){
	go func(){
		for str := range input{
			if str == "alarm."{
				vars.Boiler1Alarm = 1
			}
			if str == "norma."{
				vars.Boiler1Alarm = 0
			}
			fmt.Println(str)
		}
	}()
}