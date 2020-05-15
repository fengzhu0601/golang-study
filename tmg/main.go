package main

import (
	"fmt"
	"os"
)

const (
	BestLuckyContribution = 3348	// 贡献值
	BestLuckyCostPoints = 0.3		// 积分

	perContribution int32 = 5
	Contribution2Points = 2500     // 贡献值转换成积分的比例 2500：1

	OverDay = 77			// 计算结束天数
	EndLuckDay = 90		// 结束抽奖天数
)

// 当前贡献值
const CurContribution int32 = 437905
// 登陆第几天
const LoginId int32 = 7
// 注册天数
const RegisterDay = 76
// 锁仓绩效积分
const LockedPoint float32 = 187.85
// 可用绩效积分
const UsablePoint float32 = 0.52
// 养老绩效积分
const OldPoint float32 = 2.91

// 每天能正常获取的贡献值
func GetEveryDayContribution(curContr, loginId int32) (dayContr int32){
	// 签到贡献
	loginContr := GetLoginContribution(loginId)
	fmt.Println("签到贡献 = ", loginContr)
	// 阅读贡献
	readContrPer := GetReadContributionPer(curContr)
	readContr := readContrPer * 12 * perContribution + perContribution*5*12*2
	fmt.Println("阅读贡献 = ", readContr)
	// 分享贡献
	shareContr := perContribution*10+ perContribution*20
	fmt.Println("分享贡献 = ", shareContr)

	dayContr = loginContr + readContr + shareContr
	return
}

// 每天的释放比例
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

func GetReadContributionPer( curContr int32)(per int32){
	if curContr < 60000{
		per = 48
	}else if curContr >= 60000 && curContr < 120000{
		per = 44
	}else if curContr >= 120000 && curContr < 180000{
		per = 40
	}else if curContr >= 180000 && curContr < 240000{
		per = 36
	}else if curContr >= 240000 && curContr < 300000{
		per = 32
	}else if curContr >= 300000 && curContr < 360000{
		per = 28
	}else if curContr >= 360000 && curContr < 420000 {
		per = 24
	}else if curContr >= 420000 && curContr < 480000{
		per = 20
	}else if curContr >= 480000 && curContr < 540000{
		per = 16
	}else if curContr >= 540000 && curContr < 600000{
		per = 12
	}else if curContr >= 600000 && curContr < 660000{
		per = 8
	}else if curContr >= 660000 && curContr < 720000{
		per = 4
	}else {
		per = 1
	}
	return
}

// 签到获取的贡献值
func GetLoginContribution(loginId int32) (result int32){
	switch loginId {
	case 1:
		result = 6
	case 2:
		result = 12
	case 3:
		result = 18
	case 4:
		result = 24
	case 5:
		result = 30
	case 6:
		result = 36
	case 7:
		result = 688
	}
	// 双倍
	result *= 2
	return
}
// 获取下一个签到序号
func GetNextLoginId(loginId int32) int32{
	if loginId >= 7{
		return 1
	}
	return loginId+1
}

func ContributionToPoints(contr int32) float32{
	return  float32(contr/ Contribution2Points)
}

func main()  {
	fileNameClient := "./tmg.csv"
	os.Remove(fileNameClient)
	fileClient, err := os.OpenFile(fileNameClient, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic("file not exsit! %!s(MISSING)")
	}
	defer fileClient.Close()

	for _, areaId := range areaList {
		birthGroupList := GenBirthPoint(areaId)
		//l := RandBirthPointList(core, areaId, birthPointMap)
		fmt.Fprintf(fileServer, "%d", areaId)
		for _, birthGroup := range birthGroupList {
			fmt.Fprintf(fileServer, "|%d,%d", birthGroup.GroupId, birthGroup.Circle)
			for _, i := range birthGroup.Indexes {
				fmt.Fprintf(fileClient, "%d,", i)
				fmt.Fprintf(fileServer, ",%d", i)
			}
		}
		fmt.Fprintf(fileServer, "\n")
	}

	fmt.Println("不使用转盘抽奖 =====")
	noUseBestLuck(CurContribution, LoginId, RegisterDay, LockedPoint, UsablePoint, OldPoint)
	//fmt.Println("使用转盘抽奖 =====")
	//useBestLuck(curContribution, loginId, regDay, lockedPoint, usedPoint, oldPoint)
}

func noUseBestLuck(curContribution, loginId, regDay int32, lockedPoint, usedPoint, oldPoint float32){
	var releaseContribution float32
	for regDay < OverDay{
		// 今天获取的贡献值
		todayContribution := GetEveryDayContribution(curContribution, loginId)
		lockedPoint = lockedPoint + ContributionToPoints(todayContribution)
		fmt.Println(todayContribution)
		// 释放
		releasePer := GetEveryDayReleasePer(regDay)
		releaseContribution = lockedPoint * releasePer

		lockedPoint = lockedPoint - releaseContribution
		usedPoint = usedPoint + releaseContribution * 0.8
		oldPoint = oldPoint + releaseContribution * 0.2

		curContribution += todayContribution
		regDay += 1
		loginId = GetNextLoginId(loginId)
	}
	fmt.Printf("锁仓绩效积分 = %f\n" +
		"可用绩效积分 = %f \n" +
		"养老绩效积分 = %f \n" +
		"注册天数 = %d \n" +
		"当前贡献值 = %d " +
		"每天释放值 = %f\n", lockedPoint, usedPoint, oldPoint, regDay, curContribution, releaseContribution)
}

func GetBestLuckContribution(regDay int32, usePoint float32) (float32, int32){
	var contri int32
	if regDay <= EndLuckDay {
		for usePoint >= BestLuckyCostPoints {
			usePoint -= BestLuckyCostPoints
			contri += BestLuckyContribution
		}
	}
	return usePoint, contri
}

func useBestLuck(curContribution, loginId, regDay int32, lockedPoint, usedPoint, oldPoint float32){
	var releaseContribution float32
	for regDay <= OverDay{
		// 先抽奖
		var bestLock int32
		usedPoint, bestLock = GetBestLuckContribution(regDay, usedPoint)
		// 今天获取的贡献值
		todayContribution := GetEveryDayContribution(curContribution, loginId) + bestLock


		lockedPoint = lockedPoint + ContributionToPoints(todayContribution)
		// 释放
		releasePer := GetEveryDayReleasePer(regDay)
		releaseContribution = lockedPoint * releasePer

		lockedPoint = lockedPoint - releaseContribution
		usedPoint = usedPoint + releaseContribution * 0.8
		oldPoint = oldPoint + releaseContribution * 0.2

		curContribution += todayContribution
		regDay += 1
		loginId = GetNextLoginId(loginId)
	}
	fmt.Printf("锁仓绩效积分 = %f\n" +
		"可用绩效积分 = %f \n" +
		"养老绩效积分 = %f \n" +
		"注册天数 = %d \n" +
		"当前贡献值 = %d " +
		"每天释放值 = %f\n", lockedPoint, usedPoint, oldPoint, regDay, curContribution, releaseContribution)
}