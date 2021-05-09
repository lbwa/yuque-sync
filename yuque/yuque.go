package yuque

// https://www.yuque.com/yuque/developer/docdetailserializer
type DocDetailSerializer struct {
	Id              string         `json:"id"`                 // 文档编号
	Slug            string         `json:"slug"`               //文档路径
	Title           string         `json:"title"`              //标题
	BookId          string         `json:"book_id"`            //仓库编号，就是 repo_id
	Book            BookSerializer `json:"book"`               //仓库信息 <BookSerializer>，就是 repo 信息
	UserId          string         `json:"user_id"`            //用户/团队编号
	User            UserSerializer `json:"user"`               //用户/团队信息 <UserSerializer>
	Format          string         `json:"format"`             //描述了正文的格式 [lake , markdown]
	Body            string         `json:"body"`               //正文 Markdown 源代码
	BodyDraft       string         `json:"body_draft"`         //草稿 Markdown 源代码
	BodyHTML        string         `json:"body_html"`          //转换过后的正文 HTML （重大变更，详情请参考：https://www.yuque.com/yuque/developer/yr938f）
	BodyLake        string         `json:"body_lake"`          //语雀 lake 格式的文档内容
	CreatorId       string         `json:"creator_id"`         //文档创建人 User Id
	Public          string         `json:"public"`             //公开级别 [0 - 私密, 1 - 公开]
	Status          string         `json:"status"`             //状态 [0 - 草稿, 1 - 发布]
	LikesCount      string         `json:"likes_count"`        //赞数量
	CommentsCount   string         `json:"comments_count"`     //评论数量
	ContentUpdateAt string         `json:"content_updated_at"` //文档内容更新时间
	DeletedAt       string         `json:"deleted_at"`         //删除时间，未删除为 null
	CreatedAt       string         `json:"created_at"`         //创建时间
	UpdatedAt       string         `json:"updated_at"`         //更新时间
}

// 一般在列表的场景返回的仓库信息。
// https://www.yuque.com/yuque/developer/bookserializer
type BookSerializer struct {
	Id           string         `json:"id"`            // 仓库编号
	Type         string         `json:"type"`          // 类型 [Book - 文档]
	Slug         string         `json:"slug"`          // 仓库路径
	Name         string         `json:"name"`          // 名称
	Namespace    string         `json:"namespace"`     // 仓库完整路径 user.login/book.slug
	UserId       string         `json:"user_id"`       // 所属的团队/用户编号
	User         UserSerializer `json:"user"`          // <UserSerializer>
	Description  string         `json:"description"`   // 介绍
	CreatorId    string         `json:"creator_id"`    // 创建人 User Id
	Public       string         `json:"public"`        // 公开状态 [1 - 公开, 0 - 私密]
	LikesCount   string         `json:"likes_count"`   // 喜欢数量
	WatchesCount string         `json:"watches_count"` // 订阅数量
	CreatedAt    string         `json:"created_at"`    // 创建时间
	UpdatedAt    string         `json:"updated_at"`    // 更新时间
}

// 一般在列表的场景返回的用户信息。
// https://www.yuque.com/yuque/developer/userserializer
type UserSerializer struct {
	Id        string `json:"id"`         // 用户编号
	Type      string `json:"type"`       // 类型 [`User`  - 用户, Group - 团队]
	Login     string `json:"login"`      // 用户个人路径
	Name      string `json:"name"`       // 昵称
	AvatarUrl string `json:"avatar_url"` // 头像 URL
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间
}
