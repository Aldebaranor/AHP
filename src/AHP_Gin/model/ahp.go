package model

import "demoProject/src/AHP_Gin/handlers/ahp"

type TreeNode struct {
	Name     string      `json:"name" bson:"bson_name"`
	Weight   float32     `json:"weight" bson:"bson_weight"`
	Children []*TreeNode `json:"children" bson:"bson_children"`
}

/*
先序遍历计算所有节点的权重，保存叶子节点权重到全局变量LeavesMap
*/
func (node *TreeNode) PreOrder() {
	if node == nil {
		return
	}
	if ahp.KVmap[node.Name] == 0 {
		node.Weight = 1
	} else {
		node.Weight = ahp.KVmap[node.Name]
	}
	for _, kid := range node.Children {
		kid.Weight = ahp.KVmap[kid.Name] * node.Weight
		if kid.Children == nil {
			ahp.LeavesMap[kid.Name] = kid.Weight
		}
		kid.PreOrder()
	}

	return
}

type Error struct {
	code    int    `json:"json_code" bson:"bson_code" custom:"my_code"`
	message string `json:"json_message" bson:"bson_message" custom:"my_message"`
}
