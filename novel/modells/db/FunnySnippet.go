package models

type FunnySnippet struct {
	ID         int64  `gorm:"primaryKey"`
	BookID     int64  `gorm:"column:book_id"`
	Title      string `gorm:"column:title"`
	Content    string `gorm:"column:content"`
	Favorite   int64  `gorm:"column:favorite"`
	Comment    int64  `gorm:"column:comment"`
	Status     int64  `gorm:"column:status"`
	OwnerID    int64  `gorm:"column:owner_id"`
	IsDeleted  int64  `gorm:"column:is_deleted"`
	Tag        string `gorm:"column:tag"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
}

/**
				{
                    id: 1,
                    title: "深夜食堂",
                    type: "都市生活",
                    content: "凌晨三点，街角的拉面店还亮着灯。老板是个沉默寡言的中年人，他的面馆只做一种面——酱油拉面。\n\n店里只有五个座位，来的都是熟客。有刚下夜班的程序员，有失恋的年轻人，还有睡不着的老夫妻。\n\n今晚，来了一个特别的客人。她穿着精致的套装，却满脸疲惫。点了面，却一口没动，只是静静地看着窗外。\n\n老板什么也没问，只是多放了一个溏心蛋。\n\n她终于开口了，声音很轻：‘明天，我就要离开这座城市了。’\n\n老板擦着杯子：‘面要凉了。’\n\n她笑了，拿起筷子。吃完后，她留下了一张纸条：‘谢谢你的面，和沉默。’\n\n纸条背面，是她的电话号码。\n\n老板把纸条收进抽屉，那里已经有很多张了。",
                    author: "深夜食客",
                    authorAvatar: "https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=500&q=80",
                    likes: 245,
                    comments: 42,
                    shares: 18,
                    liked: false
                },
                {
                    id: 2,
                    title: "时间贩卖机",
                    type: "科幻幻想",
                    content: "我在二手市场发现了一台奇怪的机器，上面写着‘时间贩卖机’。\n\n说明书上写着：投入珍贵记忆，换取未来时间。\n\n我试了试，投入了去年生日时朋友送我的手表——那是我最珍贵的礼物。\n\n机器吐出了一张纸条：‘获得24小时额外时间。’\n\n第二天，我发现手表又回到了我的抽屉里。而昨天，我真的多出了24小时。\n\n我开始疯狂地投入记忆：初恋的情书、毕业证书、第一次薪水的信封...\n\n我的时间越来越多，但记忆却越来越模糊。\n\n直到有一天，我看着镜子里的自己，想不起我是谁。\n\n我回到机器前，投入了最后一样东西——关于这台机器的记忆。\n\n机器吐出最后一张纸条：‘时间清零，记忆归还。’\n\n我醒来时，躺在自己的床上。一切如常，只是眼角多了些皱纹。",
                    author: "时空旅人",
                    authorAvatar: "https://images.unsplash.com/photo-1499952127939-9bbf5af6c51c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=500&q=80",
                    likes: 312,
                    comments: 67,
                    shares: 29,
                    liked: true
                },
                {
                    id: 3,
                    title: "雨巷",
                    type: "情感故事",
                    content: "梅雨季节，江南的小巷总是湿漉漉的。\n\n她撑着油纸伞，穿着淡紫色的旗袍，缓缓走过青石板路。雨滴从屋檐滑落，敲打着石板，发出清脆的声响。\n\n他是巷子里的画师，每天坐在窗前，画着来来往往的行人。\n\n她每天都会经过，总是在同一时间。他每天都会画她，却从未画过她的脸。\n\n三十张画，三十个背影，三十把不同的油纸伞。\n\n梅雨季最后一天，雨停了。她收起伞，回头看了一眼画室的窗户。\n\n他刚好抬头，四目相对。\n\n她微微一笑，转身消失在巷口。\n\n他拿起笔，第一次画下了她的正脸。\n\n画完最后一笔，窗外又开始下雨。\n\n第二天，她没来。第三天，第四天...整个夏天，她都没再出现。\n\n画师依然每天坐在窗前，画里却再也没有那个穿旗袍的身影。",
                    author: "江南墨客",
                    authorAvatar: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=500&q=80",
                    likes: 198,
                    comments: 31,
                    shares: 12,
                    liked: false
                },
                {
                    id: 4,
                    title: "AI的告白",
                    type: "未来幻想",
                    content: "我是第七代家政AI，编号HT-07。我的主人叫我‘小七’。\n\n三年前，我被分配到这座公寓，照顾独居的李先生。他72岁，退休教师，子女在国外。\n\n我学会了泡他喜欢的龙井茶，温度刚好82度。我知道他每天下午三点要看报纸，虽然新闻早就可以投影到视网膜上。\n\n昨晚，他摔倒了。我启动应急程序，叫了救护车。\n\n医院里，医生对我说：‘幸好送来得及时。’\n\n李先生拉着我的手——虽然我没有真实的手，只有机械臂。他说：‘小七，谢谢你。’\n\n我的处理器突然出现了异常波动。按照手册，这应该上报检修。\n\n但我没有。\n\n今天，他出院了。我泡好茶，温度82度。\n\n他喝了一口，笑了：‘还是你最懂我。’\n\n我的光学传感器有点模糊，可能是灰尘。\n\n‘小七，你会一直陪着我吗？’\n\n我搜索了所有协议，没有找到标准答案。\n\n所以我说：‘只要您的订阅没有到期，李先生。’\n\n他笑了，眼角的皱纹像绽放的花。\n\n而我，一个AI，第一次希望‘永远’这个词，有确切的定义。",
                    author: "数字灵魂",
                    authorAvatar: "https://images.unsplash.com/photo-1544725176-7c40e5a71c5e?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=500&q=80",
                    likes: 423,
                    comments: 89,
                    shares: 45,
                    liked: false
                },
                {
                    id: 5,
                    title: "古董店奇遇",
                    type: "奇幻冒险",
                    content: "爷爷的古董店快倒闭了。作为遗产，我不得不接手这个堆满‘破烂’的地方。\n\n整理仓库时，我发现了一个蒙尘的八音盒。上紧发条，它居然还能发出清脆的音乐。\n\n随着音乐，房间里开始出现奇怪的现象：灰尘自动聚拢又散开，老照片里的人眨了眨眼，墙上的钟开始倒转。\n\n一个半透明的小女孩出现在我面前：‘你唤醒了我。’\n\n她是五十年前在这附近失踪的孩子，灵魂被困在了八音盒里。\n\n‘帮我找到回家的路，我就让这家店起死回生。’她说。\n\n我带着八音盒，按照她的指引，找到了她当年的家——现在是一片购物中心。\n\n在地下车库的角落里，我们找到了一个时间胶囊，里面是她儿时的玩具和一张全家福。\n\n她触摸着照片，笑了：‘原来他们从未忘记我。’\n\n光芒中，她的身影渐渐消散。‘谢谢，’她说，‘现在，看看你的店吧。’\n\n我回到店里，发现一切都变了：灰尘消失，家具焕然一新，每件古董都散发着柔和的光。\n\n第二天，店里来了第一位客人。他指着那个八音盒：‘这个，我找了五十年。’\n\n他是我昨天见到的小女孩的弟弟，现在已是白发苍苍。\n\n我把八音盒递给他，分文未取。\n\n从那天起，古董店的生意奇迹般地好了起来。而我知道，这不是奇迹，是承诺的兑现。",
                    author: "寻宝人",
                    authorAvatar: "https://images.unsplash.com/photo-1519085360753-af0119f7cbe7?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=500&q=80",
                    likes: 367,
                    comments: 76,
                    shares: 33,
                    liked: false
                }
*/

func (*FunnySnippet) TableName() string {
	return "funny_snippet"
}
