package reborn

import (
	wr "github.com/mroth/weightedrand"
)

type rate []struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

var (
	areac     *wr.Chooser
	gender, _ = wr.NewChooser(
		wr.Choice{Item: "统治者", Weight: 1000},
		wr.Choice{Item: "专家", Weight: 2000},
		wr.Choice{Item: "劳工", Weight: 10000},
		wr.Choice{Item: "复合子个体", Weight: 5000},
		wr.Choice{Item: "低级子个体", Weight: 10000},
		wr.Choice{Item: "不可接触者", Weight: 10000},
		wr.Choice{Item: "机器仆从", Weight: 10000},
		wr.Choice{Item: "奴隶", Weight: 10000},
		wr.Choice{Item: "活体陈设", Weight: 10000},
		wr.Choice{Item: "罪犯", Weight: 5000},
		wr.Choice{Item: "正在同化", Weight: 5000},
		wr.Choice{Item: "反常工蜂", Weight: 5000},
		wr.Choice{Item: "腐化子个体", Weight: 5000},
		wr.Choice{Item: "土著物种", Weight: 2000},
		wr.Choice{Item: "未开智物种", Weight: 2000},
		wr.Choice{Item: "科学家", Weight: 2000},
		wr.Choice{Item: "舰队司令", Weight: 2000},
		wr.Choice{Item: "陆军将领", Weight: 2000},
		wr.Choice{Item: "总督", Weight: 2000},
	)
)

func randcoun() string {
	return areac.Pick().(string)
}

func randgen() string {
	return gender.Pick().(string)
}
