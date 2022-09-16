package ahp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
)

/*
全局变量
*/
var KVmap = make(map[string]float32)
var LeavesMap = make(map[string]float32)

/*
输入：
方案中各特征的分数
输出：
方案得分
*/
func GetPoints(context *gin.Context) {
	cJson := make(map[string]interface{})
	context.BindJSON(&cJson)
	scheme := cJson["schema"]
	pointMap := make(map[string]float32)
	for key, val := range scheme.(map[string]interface{}) {
		for s, f := range val.(map[string]interface{}) {
			pointMap[key] += LeavesMap[s] * float32(f.(float64))
		}
	}
	context.JSON(http.StatusOK, pointMap)
}

/*
输入：
特征名称：长度为N的一维数组，
判断矩阵：N*N的二维数组
输出：
特征名称和特征权重的键值对
*/
func GetWeight(context *gin.Context) {
	cJson := make(map[string]interface{})
	context.BindJSON(&cJson)
	mtx := cJson["matrix"]
	attribute := cJson["attribute"]
	var f32Mtx [][]float32
	for _, val := range mtx.([]interface{}) {
		var col []float32
		for _, val1 := range val.([]interface{}) {
			var f32 = float32(val1.(float64))
			col = append(col, f32)
		}
		f32Mtx = append(f32Mtx, col)
	}
	ans := Weight(f32Mtx)
	result := make(map[string]float32, len(ans))
	for i, val := range attribute.([]interface{}) {
		result[val.(string)] = ans[i]
		KVmap[val.(string)] = ans[i]
	}
	log.Printf("%+v", KVmap)
	context.JSON(http.StatusOK, result)
}

/*
输入：N*N的判断矩阵
输出：对应权重
*/
func Weight(judgeMtx [][]float32) []float32 {

	mtxLen := len(judgeMtx[0])
	RI := []float32{
		0, 0, 0.58, 0.90, 1.12, 1.21, 1.32, 1.41, 1.45, 1.49, 1.52, 1.54, 1.56, 1.58,
	}
	if mtxLen > len(RI) {
		fmt.Println(error(20))
		return nil
	}

	w1 := make([]float32, mtxLen)
	w2 := make([]float32, mtxLen)
	w3 := make([]float32, mtxLen)
	lam := make([]float32, mtxLen)
	lamMax := float32(0)

	sum := float32(0)
	for i := 0; i < mtxLen; i++ {
		product := float32(1)
		for j := 0; j < mtxLen; j++ {
			product *= judgeMtx[i][j]
		}
		w1[i] = float32(math.Pow(float64(product), float64(1/float32(mtxLen))))
		sum += w1[i]
	}

	for i := 0; i < mtxLen; i++ {
		w2[i] = w1[i] / sum
	}

	for i := 0; i < mtxLen; i++ {
		for j := 0; j < mtxLen; j++ {
			w3[i] += judgeMtx[i][j] * w2[j]
		}
		lam[i] = w3[i] / w2[i]
		lamMax += lam[i]
	}
	lamMax /= float32(mtxLen)

	CI := (lamMax - float32(mtxLen)) / float32(mtxLen-1)

	CR := CI / RI[mtxLen-1]

	if CR > 0.1 {
		fmt.Println(error(10))
		return nil
	}

	return w2
}

func error(errnum int) string {
	var errorCode = map[int]string{
		10: "一致性校验未通过，请重新输入！",
		20: "评估指标数目超过上限，请减少指标数量再重新输入！",
	}
	return errorCode[errnum]
}
