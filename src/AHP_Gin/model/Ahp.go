package model

type TreeNode struct {
	Name     string      `json:"name" bson:"bson_name"`
	Weight   float64     `json:"weight" bson:"bson_weight"`
	Children []*TreeNode `json:"children" bson:"bson_children"`
}

type Error struct {
	code    int    `json:"json_code" bson:"bson_code" custom:"my_code"`
	message string `json:"json_message" bson:"bson_message" custom:"my_message"`
}
