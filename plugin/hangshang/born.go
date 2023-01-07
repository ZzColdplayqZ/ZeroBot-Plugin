package hangshang

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
		wr.Choice{Item: "圣遗物里面放着音乐，虽然这盒子里并没有任何显而易见的类似播放器的东西。\n圣遗物马上被关上，并且被遗忘在某个角落里面", Weight: 10000},
		wr.Choice{Item: "圣遗物装满了美丽的蓝色大理石。\n获得消费品+2000", Weight: 5000},
		wr.Choice{Item: "圣遗物里存放着一整套的制服。可惜的是，这些衣服的设计对象是某种比人类更多肢体的生物。\n不过，这些衣服的质料还是可以循环利用的。\n获得消费品+867", Weight: 5000},
		wr.Choice{Item: "圣遗物装着满满的军粮。虽\n然上面的所有标记都已经磨损殆尽了，但是里面的东西看起来尚且安全,而且适合人类的生理结构。\n食物+26663", Weight: 5000},
		wr.Choice{Item: "圣遗物里面装着数不清的财富,如果你对不在场的人说起的话，没有人会相信这么巨量的财富会真实存在。\n不过财务报表上的赤字正宣告着这个事实。\n能量币+4000，食物+4000", Weight: 5000},
		wr.Choice{Item: "圣遗物很难打开，感觉里面似乎是有一个小型的重力并阻止它打开似的。\n物理学研究+114514		", Weight: 10000},
		wr.Choice{Item: "圣遗物里面装着几个外星设备，看上去没什么反应，应该是坏了的。\n获得工程学研究+114514		", Weight: 10000},
		wr.Choice{Item: "圣遗物里存放着一把未知来历的枪械， 其发射扳机明显已被破坏掉。", Weight: 5000},
		wr.Choice{Item: "圣遗物里面什么都没有。\n原本用于贮藏的空间被一块平滑的合金占据了 ,与圣遗物的内壁无缝地衔接在一起。\n获得合金+7260", Weight: 5000},
		wr.Choice{Item: "圣遗物里面装着几个外星设备，看上去没什么反应，应该是坏了的。", Weight: 5000},
		wr.Choice{Item: "虽然是空的，但是这个圣遗物的内壁上到底都是各种印记。\n在更深处的地方追踪到了-些生物物质的踪迹。\n获得社会学研究+114514", Weight: 10000},
		wr.Choice{Item: "圣遗物装着一篇关于 爱与万物自然循环的论文。\n它由海星语写成，而且文风十分现代。\n获得社会学研究+114514", Weight: 10000},
		wr.Choice{Item: "圣遗物里有炸弹!\n别担心，已经拆掉了。", Weight: 6000},
		wr.Choice{Item: "圣遗物装满了美丽的蓝色大理石。\n消费品+2000", Weight: 5000},
		wr.Choice{Item: "圣遗物里存放着一个只能以波频的方式存在于空灵界面的生灵。\n它在被发现后瞬间逃离到更高阶的位置，只在你心里留下深刻的启示与不安。\n你拥有了灵媒特性", Weight: 1500},
		wr.Choice{Item: "这是银河之星，行商组织的基石。\n我们的先祖曾在此宣誓，组件出行商联盟。\n早先，行商联盟的发展全都仰赖银河之星的神秘力量，而随着时间的推移，这种力量逐渐消散了。\n获得群头衔“银河之星", Weight: 500},
		wr.Choice{Item: "获得远古贸易路线 \n贸易额+10%", Weight: 2000},
	)
)

func randcoun() string {
	return areac.Pick().(string)
}

func randgen() string {
	return gender.Pick().(string)
}
