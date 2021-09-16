// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-08-31 14:17
// version: 1.0.0
// desc   :

package engine

const (
	baseURL = "https://api.e-learn.io" // 接口地址
)

// 等级
type level int

const (
	primary level = 1 + iota // 1 初中
	high                     // 2 高中
)

// 科目
type subject int

const (
	chinese   subject = 101 + iota // 101 语文
	math                           // 102 数学
	english                        // 103 英语
	physics                        // 104 物理学
	chemistry                      // 105 化学
	biology                        // 106 生物学
	history                        // 107 历史学
	geography                      // 108 地理学
	political                      // 109 政治学
)

const (
	apiUnitList       = "/index/unit/queryUnitList"         // 单元 /index/unit/queryUnitList?subjectId=101&levelId=2
	apiCoursePageList = "/index/course/queryCoursePageList" // 课程 /index/course/queryCoursePageList?unitId=576&current=1
	apiLectureList    = "/index/lecture/queryLectureList"   // 章节 /index/lecture/queryLectureList?courseId=4659
	apiWatch          = "/watch"                            // 视频 /watch?courseId=4659&lectureId=34908&coursewareId=59353
)

const (
	headerUserAgent = "User-Agent"
	headerForwarded = "X-Forwarded-For"
)

func api(api string) string {
	return baseURL + api
}
