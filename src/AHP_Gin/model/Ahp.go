package model

var ApiCode = &Error{
	SUCCESS:         10,
	CheckFAILED:     21,
	ConditionExceed: 22,
}

type TreeNode struct {
	Name     string      `json:"name"`
	Weight   float64     `json:"weight"`
	Children []*TreeNode `json:"children"`
}

type Error struct {
	SUCCESS         uint            `json:"success"`
	CheckFAILED     uint            `json:"check_failed"`
	ConditionExceed uint            `json:"condition_exceed"`
	Message         map[uint]string `json:"message"`
}
