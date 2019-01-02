package goecharts

// 标题组件配置项
type TitleOpts struct {
	// 主标题
	Title string `json:"text,omitempty"`
	// 主标题样式配置项
	TitleStyle *TextStyle `json:"textStyle,omitempty"`
	// 副标题
	Subtitle string `json:"subtext,omitempty"`
	// 副标题样式配置项
	SubtitleStyle *TextStyle `json:"subtextStyle,omitempty"`
	// 主标题文本超链接
	Link string `json:"link,omitempty"`
	// 指定窗口打开主标题超链接
	// 'self' 当前窗口打开
	// 'blank' 新窗口打开
	Target string `json:"target,omitempty"`
	// grid 组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比，也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐
	Top string `json:"top,omitempty"`
	// 图例组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Bottom string `json:"bottom,omitempty"`
	// 图例组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// 图例组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
}

// 图例组件配置项
type LegendOpts struct {
	// 是否显示图例
	Show bool `json:"show,omitempty"`
	// 图例组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// 图例组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比，也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐。
	Top string `json:"top,omitempty"`
	// 图例组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
	// 图例组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Bottom string `json:"bottom,omitempty"`
}

// 提示框组件配置项
type TooltipOpts struct {
	// 是否显示提示框
	Show bool `json:"show,omitempty"`
	// 触发类型。
	// 'item': 数据项图形触发，主要在散点图，饼图等无类目轴的图表中使用。
	// 'axis': 坐标轴触发，主要在柱状图，折线图等会使用类目轴的图表中使用。
	// 'none': 什么都不触发。
	Trigger string `json:"trigger,omitempty"`
	// 1, 字符串模板
	// 模板变量有 {a}, {b}，{c}，{d}，{e}，分别表示系列名，数据名，数据值等。
	// 在 trigger 为 'axis' 的时候，会有多个系列的数据，此时可以通过 {a0}, {a1}, {a2}
	// 这种后面加索引的方式表示系列的索引。 不同图表类型下的 {a}，{b}，{c}，{d} 含义不一样。
	// 其中变量{a}, {b}, {c}, {d} 在不同图表类型下代表数据含义为：
	// 折线（区域）图、柱状（条形）图、K 线图 : {a}（系列名称），{b}（类目值），{c}（数值）, {d}（无）
	// 散点图（气泡）图 : {a}（系列名称），{b}（数据名称），{c}（数值数组）, {d}（无）
	// 地图 : {a}（系列名称），{b}（区域名称），{c}（合并数值）, {d}（无）
	// 饼图、仪表盘、漏斗图: {a}（系列名称），{b}（数据项名称），{c}（数值）, {d}（百分比）
	//
	// 2, 回调函数
	// 回调函数格式：
	// (params: Object|Array, ticket: string, callback: (ticket: string, html: string)) => string
	// 第一个参数 params 是 formatter 需要的数据集。格式如下：
	// {
	//    componentType: 'series',
	//    seriesType: string,	// 系列类型
	//    seriesIndex: number,	// 系列在传入的 option.series 中的 index
	//    seriesName: string,	// 系列名称
	//    name: string,			// 数据名，类目名
	//    dataIndex: number,	// 数据在传入的 data 数组中的 index
	//    data: Object,			// 传入的原始数据项
	//    value: number|Array,	// 传入的数据值
	//    color: string,		// 数据图形的颜色
	//    percent: number,		// 饼图的百分比
	// }
	Formatter string `json:"formatter,omitempty"`
}

// 工具箱组件配置项
type ToolboxOpts struct {
	// 是否显示工具栏组件
	Show      bool `json:"show"`
	TBFeature `json:"feature"`
}

type TBFeature struct {
	// 保存为图片
	SaveAsImage struct{} `json:"saveAsImage"`
	// 数据区域缩放。目前只支持直角坐标系的缩放
	DataZoom struct{} `json:"dataZoom"`
	// 数据视图工具，可以展现当前图表所用的数据，编辑后可以动态更新
	DataView struct{} `json:"dataView"`
	// 配置项还原
	Restore struct{} `json:"restore"`
}

// 字体样式配置项
type TextStyle struct {
	// 文字字体颜色
	Color string `json:"color,omitempty"`
	// 文字字体的风格
	// 可选  'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`
	// 字体大小
	FontSize int `json:"fontSize,omitempty"`
}

type LineStyleOpts struct {
	// 线的颜色
	Color string `json:"color,omitempty"`
	// 线的宽度
	// 默认 1
	Width float32 `json:"width,omitempty"`
	// 线的类型，可选 "solid", "dashed", "dotted"
	// 默认 "solid"
	Type string `json:"type,omitempty"`
	// 线的透明度。支持从 0 到 1 的数字，为 0 时不绘制线
	Opacity float32 `json:"opacity,omitempty"`
	// 线的曲度，支持从 0 到 1 的值，值越大曲度越大
	// 默认 0
	Curveness float32 `json:"curveness,omitempty"`
}

type AreaStyleOpts struct {
	// 填充区域的颜色
	Color string `json:"color,omitempty"`
	// 填充区域的透明度。支持从 0 到 1 的数字，为 0 时不填充区域
	Opacity float32 `json:"opacity,omitempty"`
}

// 区域缩放组件配置项
type DataZoomOpts struct {
	// 缩放类型，可选 "inside", "slider"
	Type string `json:"type" default:"inside"`
	// 数据窗口范围的起始百分比。范围是：0 ~ 100。表示 0% ~ 100%。
	// 默认: 0
	Start float32 `json:"start,omitempty"`
	// 数据窗口范围的结束百分比。范围是：0 ~ 100。
	// 默认: 100
	End float32 `json:"end,omitempty"`
	// 触发视图刷新的频率。单位为毫秒（ms）。
	// 默认: 100
	Throttle float32 `json:"throttle,omitempty"`
	// 设置 dataZoom 组件控制的 X 轴
	// 不指定时，当 dataZoom-inside.orient 为 'horizontal'时，默认控制和 dataZoom 平行的第一个 xAxis。
	// 但是不建议使用默认值，建议显式指定。
	// 如果是 number 表示控制一个轴，如果是 Array 表示控制多个轴。
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`
	// 设置 dataZoom 组件控制的 Y 轴
	// 不指定时，当 dataZoom-slider.orient 为 'vertical'时，默认控制和 dataZoom 平行的第一个 yAxis。
	// 但是不建议使用默认值，建议显式指定。
	// 如果是 number 表示控制一个轴，如果是 Array 表示控制多个轴。
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`
}

type DataZoomOptsList []DataZoomOpts

func (dz DataZoomOptsList) Len() int {
	return len(dz)
}

// 视觉映射组件配置项
// 用于进行『视觉编码』，也就是将数据映射到视觉元素（视觉通道）
type VisualMapOpts struct {
	// 映射类型，可选 "continuous", "piecewise"
	Type string `json:"type,omitempty" default:"continuous"`
	// 是否显示拖拽用的手柄（手柄能拖拽调整选中范围）
	Calculable bool `json:"calculable"`
	// VisualMap 组件的允许的最小值
	Mix float32 `json:"mix,omitempty"`
	// VisualMap 组件的允许的最大值
	Max float32 `json:"max,omitempty"`
	// 指定手柄对应数值的位置。range 应在 min max 范围内
	Range []float32 `json:"range,omitempty"`
	// 两端的文本，如 ['High', 'Low']
	Text []string `json:"text,omitempty"`
	// 定义在选中范围中的视觉元素
	VMInRange `json:"inRange,omitempty"`
}

type VisualMapOptsList []VisualMapOpts

func (vm VisualMapOptsList) Len() int {
	return len(vm)
}

type VMInRange struct {
	// 图元的颜色
	Color []string `json:"color,omitempty"`
	// 图元的图形类别
	// 可选 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	Symbol string `json:"symbol,omitempty"`
	// 图元的大小
	SymbolSize float32 `json:"symbolSize,omitempty"`
}

// X 轴配置项组件
type XAxisOpts struct {
	// X 轴名称
	Name string `json:"name,omitempty"`
	// 是否显示 X 轴
	Show bool `json:"show,omitempty"`
	// X 轴数据项
	Data interface{} `json:"data,omitempty"`
}

// Y 轴配置项组件
type YAxisOpts struct {
	// Y 轴名称
	Name string `json:"name,omitempty"`
	// 是否显示 Y 轴
	Show bool `json:"show,omitempty"`
	// Y 轴数据项
	Data interface{} `json:"data,omitempty"`
}

// 地理坐标系组件配置项
type GeoOpts struct {
	Map  string `json:"map,omitempty"`
	Roam bool   `json:"roam,omitempty"`
}

// 处理 function 类型配置项
// TODO: 处理 \n
func FuncOpts(fn string) string {
	return "__x__" + fn + "__x__"
}