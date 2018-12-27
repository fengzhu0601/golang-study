package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"encoding/json"
	"slg_game_server/server/goredis"
	"slg_game_server/server/util"
)

func main()  {
	client := createClient()
	defer client.Close()
	testStruct()
}

// 创建 redis 客户端
func createClient() *redis.Client {
	client := redis.NewClient( &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 5,
	} )

	pong, err := client.Ping().Result()
	fmt.Println( pong, err )

	return client
}

type Land struct {
	LandId		int32		// 地块ID
	OwnerId 	int32		// 归属玩家
	MarchMap	map[MarchKey]int32	// 行军列表：[玩家id+部队id]部队行军信息
}

type MarchKey struct {
	UserId		int32		// 玩家ID
	ArmyId		int32		// 部队ID
}
func testStruct() {
	landObj := &Land{LandId:100, OwnerId:200, MarchMap:make(map[MarchKey]int32)}
	marchKey := MarchKey{UserId:200, ArmyId:300}
	landObj.MarchMap[marchKey] = 300
	marchJson, _ := json.Marshal(landObj.MarchMap)

	m := map[string]interface{}{"OwnerId":landObj.OwnerId, "LandId":landObj.LandId, "MarchMap": marchJson}
	goredis.ClientRedis.HMSet("land" + util.ToStr(landObj.LandId), m)
	fmt.Println("-------------put land : ", landObj)
}
