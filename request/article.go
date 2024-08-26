package request

type ArticleCreateReq struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type ArticleUpdateReq struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
}

type ArticleDeleteReq struct {
	IDs []uint `json:"ids"`
}
