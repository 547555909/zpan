package bind

type QueryPage struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit,default=20"`
}

type QueryFiles struct {
	QueryPage
	Dir     string `form:"dir"`
	Type    string `form:"type"`
	Search  bool   `form:"search"`
	Keyword string `form:"keyword"`
}

type QueryMatter struct {
	Name string `form:"name" binding:"required"`
	Type string `form:"type"`
	Size int64  `form:"size"`
	Dir  string `form:"dir"`
}

type QueryFolder struct {
	QueryPage
	Parent string `form:"parent"`
}

type BodyFolder struct {
	Name string `json:"name" binding:"required"`
	Dir  string `json:"dir"`
}

type BodyFile struct {
	Uid    int64  `json:"uid"`
	Name   string `json:"name" binding:"required"`
	Type   string `json:"type" binding:"required"`
	Size   int64  `json:"size" binding:"required"`
	Dir    string `json:"dir"`
	Object string `json:"object" binding:"required"`
}

type BodyFileOperation struct {
	Id     int64  `json:"id" binding:"required"`
	Dest   string `json:"dest"`
	Action int64  `json:"action" binding:"required"`
}

type BodyShare struct {
	Id        int64 `json:"id"`
	MId       int64 `json:"mid"`
	Private   bool  `json:"private"`
	ExpireSec int64 `json:"expire_sec"`
}

// 	bucket=callback-test&object=test.txt&etag=D8E8FCA2DC0F896FD7CB4CB0031BA249&size=5&mimeType=text%2Fplain&imageInfo.height=&imageInfo.width=&imageInfo.format=&x:var1=for-callback-test
