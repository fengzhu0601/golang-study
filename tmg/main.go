package tmg

const (
	BestLuckyContribution = 3348	// 贡献值
	BestLuckyCostPoints = 0.3		// 积分
	Regday	= 75	// 注册天数
)

func GetEveryDayContribution(curContr int32) (dayContr int32){

	return
}

func GetEveryDayReleasePer(regday int32) (per float32){
	if regday <= 30 {
		per = 0.0030
	}else if regday >= 31 && regday <= 60 {
		per = 0.0025
	}else if regday >= 61 && regday <= 90 {
		per = 0.0020
	}else if regday >= 91 && regday <= 120 {
		per = 0.0015
	}else if regday >= 121 && regday <= 150 {
		per = 0.0010
	}else if regday >= 151 && regday <= 180 {
		per = 0.0008
	}else if regday >= 181 && regday <= 210 {
		per = 0.0006
	}else if regday >= 211 && regday <= 240 {
		per = 0.0004
	}else if regday >= 241 && regday <= 270 {
		per = 0.0003
	}else if regday >= 271 && regday <= 300 {
		per = 0.0002
	}else if regday >= 301 && regday <= 400 {
		per = 0.00015
	}else if regday >= 401 && regday <= 500 {
		per = 0.00012
	}else if regday >= 501 && regday <= 600 {
		per = 0.0001
	}else if regday >= 601 && regday <= 700 {
		per = 0.00009
	}else if regday >= 701 && regday <= 800 {
		per = 0.00008
	}else if regday >= 801 && regday <= 900 {
		per = 0.00007
	}else if regday >= 901 && regday <= 1000 {
		per = 0.00006
	}else if regday >= 1001 && regday <= 1100 {
		per = 0.00005
	}else if regday >= 1101 && regday <= 1200 {
		per = 0.00004
	}else if regday >= 1201 && regday <= 1300 {
		per = 0.00003
	}else if regday >= 1301 && regday <= 1400 {
		per = 0.00002
	}else{
		per = 0.00001
	}
	return
}

func main()  {

}
