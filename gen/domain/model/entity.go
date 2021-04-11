package model

type ReferInfo struct {
	ReferEntityName string `json:"referEntityName"`
	IsDetail        bool   `json:"isDetail"`
}

type SourceInfo struct {
	SourceEntityName string `json:"sourceEntityName"`
	IsDetail         bool   `json:"isDetail"`
}

type Field struct {
	Name       string      `json:"name"`
	Title      string      `json:"title"`
	ReferInfo  *ReferInfo  `json:"referInfo"`
	SourceInfo *SourceInfo `json:"sourceInfo"`
	Type       string      `json:"type"`
	Valid      string      `json:"valid"`
	Value      interface{} `json:"value"`
}

type Entity struct {
	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Fields      []Field `json:"fields"`
	Type        string  `json:"type"`
	Description string  `json:"Description"` // 描述文字
}
