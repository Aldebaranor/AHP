package ahp

import (
	"demoProject/src/AHP_Gin/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
)

func GetPoints(context *gin.Context) {
	node := &model.TreeNode{}
	cJson := make(map[string]interface{})
	context.BindJSON(&cJson)
	scheme := cJson["schema"]
	tree := cJson["tree"]
	log.Printf("%+v", scheme)
	resByre, resErr := json.Marshal(tree)
	if resErr != nil {
		log.Printf("%+v", resErr)
	}
	jsonRes := json.Unmarshal(resByre, &node)
	if jsonRes != nil {
		log.Printf("%+v", jsonRes)
	}
	leavesMap := make(map[string]float64)
	leavesMap = PreOrder(node, leavesMap)
	pointMap := make(map[string]float64)
	for key, val := range scheme.(map[string]interface{}) {
		for s, f := range val.(map[string]interface{}) {
			pointMap[key] += leavesMap[s] * f.(float64)
		}
	}
	context.JSON(http.StatusOK, pointMap)

}

/*
先序遍历计算所有节点的权重
*/
func PreOrder(node *model.TreeNode, leaves map[string]float64) map[string]float64 {

	if node.Children == nil {
		leaves[node.Name] = node.Weight
		return leaves
	}
	for _, kid := range node.Children {
		kid.Weight = kid.Weight * node.Weight
		PreOrder(kid, leaves)
	}

	return leaves
}

/*
接口1
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
	var f64Mtx [][]float64
	for _, val := range mtx.([]interface{}) {
		var col []float64
		for _, val1 := range val.([]interface{}) {
			var f64 = float64(val1.(float64))
			col = append(col, f64)
		}
		f64Mtx = append(f64Mtx, col)
	}
	mtxLen := len(f64Mtx[0])
	RI := []float64{
		0, 0, 0.58, 0.90, 1.12, 1.21, 1.32, 1.41, 1.45, 1.49,
		1.52, 1.54, 1.56, 1.58, 1.59, 1.5943, 1.6064, 1.6133, 1.6207, 1.6292,
		1.6385, 1.6403, 1.6462, 1.6497, 1.6556, 1.6587, 1.6631, 1.667, 1.6693, 1.6724,
	}
	if mtxLen > len(RI) {
		log.Printf(model.ApiCode.Message[22])
		context.JSON(500, gin.H{
			"code": model.ApiCode.ConditionExceed,
			"msg":  model.ApiCode.Message[model.ApiCode.ConditionExceed],
		})
		return
	}

	w1 := make([]float64, mtxLen)
	w2 := make([]float64, mtxLen)
	w3 := make([]float64, mtxLen)
	lam := make([]float64, mtxLen)
	lamMax := float64(0)

	sum := float64(0)
	for i := 0; i < mtxLen; i++ {
		product := float64(1)
		for j := 0; j < mtxLen; j++ {
			product *= f64Mtx[i][j]
		}
		w1[i] = math.Pow(product, 1/float64(mtxLen))
		sum += w1[i]
	}

	for i := 0; i < mtxLen; i++ {
		w2[i] = w1[i] / sum
	}

	for i := 0; i < mtxLen; i++ {
		for j := 0; j < mtxLen; j++ {
			w3[i] += f64Mtx[i][j] * w2[j]
		}
		lam[i] = w3[i] / w2[i]
		lamMax += lam[i]
	}
	lamMax /= float64(mtxLen)

	CI := (lamMax - float64(mtxLen)) / float64(mtxLen-1)

	CR := CI / RI[mtxLen-1]

	if CR > 0.1 {
		log.Printf(model.ApiCode.Message[21])
		context.JSON(500, gin.H{
			"code": model.ApiCode.CheckFAILED,
			"msg":  model.ApiCode.Message[model.ApiCode.CheckFAILED],
		})
		return
	}
	ans := w2
	result := make(map[string]float64, len(ans))
	for i, val := range attribute.([]interface{}) {
		result[val.(string)] = ans[i]
	}
	log.Printf("%+v", result)
	context.JSON(http.StatusOK, result)
}
