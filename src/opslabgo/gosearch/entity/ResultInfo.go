package entity

type ResultInfo struct {
	//来源接口
	From  string    `json:"-"`;
	//标题
	Title string;
	//资源时间
	Info  string;
	//资源链接
	Href  string;
	//资源描述
	Desc  string;
	//资源排序
	Order int;
}

type ResultSlice [] ResultInfo

// 重写 Len() 方法
func (a ResultSlice) Len() int {

	return len(a)
}
// 重写 Swap() 方法
func (a ResultSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
// 重写 Less() 方法， 从大到小排序
func (a ResultSlice) Less(i, j int) bool {
	return a[j].Order < a[i].Order
}