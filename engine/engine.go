// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-08-20 11:42
// version: 1.0.0
// desc   :

package engine

import (
	"e-learn/engine/res"
	"e-learn/logger"
	"e-learn/orm"
	"e-learn/orm/entity"
	"fmt"
	"math/rand"
	"time"

	"github.com/yhyzgn/goat/client"
)

// Engine ...
type Engine struct {
	retries map[string]int
}

// New ...
func New() *Engine {
	return &Engine{
		retries: make(map[string]int),
	}
}

// Start ...
func (e *Engine) Start() {
	e.loadingPlugins()

	e.fetch()
}

func (e *Engine) loadingPlugins() {
	// MySQL
	logger.Info("正在连接数据库...")
	if err := orm.Connect(); nil != err {
		panic("数据库连接失败：" + err.Error())
	}
	logger.Info("数据库连接成功！")

	logger.Info("初始化数据库...")
	if err := orm.AutoMigrate(new(entity.Subject), new(entity.Unit), new(entity.Course), new(entity.Courseware)); err != nil {
		panic(err)
	}
	logger.Info("数据库初始化完成！")
}

func (e *Engine) fetch() {
	//e.fetchAndLog(chinese, "语文")
	e.fetchAndLog(math, "数学")
}

func (e *Engine) fetchAndLog(sb subject, name string) {
	e.fetchSubject(sb, name)
	logger.InfoF("==================================== {}课程信息获取完成 ====================================", name)
}

func (e *Engine) fetchSubject(sbj subject, name string) {
	if !orm.Exists("subject", "id = ?", int(sbj)) {
		// 科目还未保存
		logger.InfoF("科目【{}】不存在，正在创建科目...", name)
		if err := orm.DB.Save(&entity.Subject{
			Id:   int(sbj),
			Name: name,
		}).Error; err != nil {
			panic(err)
		}
	} else {
		logger.InfoF("科目【{}】已存在", name)
	}

	e.fetchUnitList(sbj, name)
}

func (e *Engine) fetchUnitList(sbj subject, name string) {
	logger.InfoF("正在获取科目【{}】的单元信息...", name)
	bs, err := client.GetWithHeader(api(apiUnitList), map[string]interface{}{
		headerUserAgent: NextUserAgent(),
		headerForwarded: NextIP(),
	}, map[string]interface{}{
		"subjectId": int(sbj),
		"levelId":   2,
	})
	if err != nil {
		// 重试
		e.retries[retryKey("fetchUnitList", int(sbj))] = e.retries[retryKey("fetchUnitList", int(sbj))] + 1
		logger.ErrorF("科目【{}】的单元信息获取失败，正在第【{}】次重试...", name, e.retries[retryKey("fetchUnitList", int(sbj))])
		// 模仿延时
		delay()
		e.fetchUnitList(sbj, name)
		return
	}

	var unitList []res.Unit
	if err = res.Decode(bs, &unitList); err != nil {
		// 重试
		e.retries[retryKey("fetchUnitList", int(sbj))] = e.retries[retryKey("fetchUnitList", int(sbj))] + 1
		logger.ErrorF("科目【{}】的单元信息获取失败，正在第【{}】次重试...", name, e.retries[retryKey("fetchUnitList", int(sbj))])
		// 模仿延时
		delay()
		e.fetchUnitList(sbj, name)
		return
	}

	logger.InfoF("科目【{}】的单元信息获取成功", name)

	// 保存单元信息
	for _, unit := range unitList {
		if !orm.Exists("unit", "subject = ? and id = ?", sbj, unit.ID) {
			e.saveUnit(int(sbj), unit.ID, unit.Counts, unit.Name)
		} else {
			logger.InfoF("单元【{}】已存在", unit.Name)
		}

		// 获取每个单元的课程信息
		courses := e.fetchCourse(unit.ID, 1)
		logger.Info("courses: ", courses)
		if nil != courses {
			// 保存课程信息
			for _, course := range courses {
				if !orm.Exists("course", "subject_id = ? and course_id = ?", course.SubjectId, course.CourseId) {
					e.saveCourse(unit.ID, course)
				} else {
					logger.InfoF("课程【{}】已存在", course.CourseName)
				}

				// 获取章节列表
				e.fetchLectureList(course.CourseId, course.CourseName)
			}
		}
	}
}

func (e *Engine) saveUnit(subject, id, counts int, name string) {
	logger.InfoF("单元【{}】不存在，正在创建单元...", name)
	if err := orm.DB.Save(&entity.Unit{
		Id:      id,
		Subject: subject,
		Name:    name,
		Counts:  counts,
	}).Error; err != nil {
		panic(err)
	}
	logger.InfoF("单元【{}】保存成功", name)
}

func (e *Engine) fetchCourse(unitId, current int) []res.Course {
	result := make([]res.Course, 0)

	logger.InfoF("正在获取单元【{}】的课程信息...", unitId)
	bs, err := client.GetWithHeader(api(apiCoursePageList), map[string]interface{}{
		headerUserAgent: NextUserAgent(),
		headerForwarded: NextIP(),
	}, map[string]interface{}{
		"unitId":  unitId,
		"current": current,
	})
	if err != nil {
		// 重试
		e.retries[retryKey("fetchCourse", unitId)] = e.retries[retryKey("fetchCourse", unitId)] + 1
		logger.ErrorF("单元【{}】的课程信息获取失败，正在第【{}】次重试...", unitId, e.retries[retryKey("fetchCourse", unitId)])
		// 模仿延时
		delay()
		temp := e.fetchCourse(unitId, current)
		if nil != temp {
			result = append(result, temp...)
		}
		return result
	}

	var coursePage res.CoursePage
	if err = res.Decode(bs, &coursePage); err != nil {
		// 重试
		e.retries[retryKey("fetchCourse", unitId)] = e.retries[retryKey("fetchCourse", unitId)] + 1
		logger.ErrorF("单元【{}】的课程信息获取失败，正在第【{}】次重试...", unitId, e.retries[retryKey("fetchCourse", unitId)])
		// 模仿延时
		delay()
		temp := e.fetchCourse(unitId, current)
		if nil != temp {
			result = append(result, temp...)
		}
		return result
	}

	if nil != coursePage.Records {
		result = append(result, coursePage.Records...)
	}

	if current < coursePage.Pages {
		// 获取下一页
		temp := e.fetchCourse(unitId, current+1)
		if nil != temp {
			result = append(result, temp...)
		}
	}
	return result
}

func (e *Engine) saveCourse(unitId int, course res.Course) {
	logger.InfoF("正在保存科目【{}】的课程【{}】...", course.SubjectName, course.CourseName)
	if err := orm.DB.Create(&entity.Course{
		GradeName:   course.GradeName,
		CourseName:  course.CourseName,
		GradeId:     course.GradeId,
		TeacherId:   course.TeacherId,
		TeacherName: course.TeacherName,
		CourseId:    course.CourseId,
		SubjectId:   course.SubjectId,
		SubjectName: course.SubjectName,
		Unit:        unitId,
	}).Error; err != nil {
		panic(err)
	}
	logger.InfoF("保存科目【{}】的课程【{}】保存成功", course.SubjectName, course.CourseName)
}

func (e *Engine) fetchLectureList(courseId int, courseName string) {
	logger.InfoF("正在获取课程【{}】的章节信息...", courseName)
	bs, err := client.GetWithHeader(api(apiLectureList), map[string]interface{}{
		headerUserAgent: NextUserAgent(),
		headerForwarded: NextIP(),
	}, map[string]interface{}{
		"courseId": courseId,
	})
	if err != nil {
		// 重试
		e.retries[retryKey("fetchLectureList", courseId)] = e.retries[retryKey("fetchLectureList", courseId)] + 1
		logger.ErrorF("课程【{}】的章节信息获取失败，正在第【{}】次重试...", courseName, e.retries[retryKey("fetchLectureList", courseId)])
		// 模仿延时
		delay()
		e.fetchLectureList(courseId, courseName)
		return
	}

	var lectureList []*res.Lecture
	if err = res.Decode(bs, &lectureList); err != nil {
		// 重试
		e.retries[retryKey("fetchLectureList", courseId)] = e.retries[retryKey("fetchLectureList", courseId)] + 1
		logger.ErrorF("课程【{}】的章节信息获取失败，正在第【{}】次重试...", courseName, e.retries[retryKey("fetchLectureList", courseId)])
		// 模仿延时
		delay()
		e.fetchLectureList(courseId, courseName)
		return
	}

	for _, lecture := range lectureList {
		// 获取每个章节每个课时的视频链接
		for _, ware := range lecture.Courseware {
			if !orm.Exists("courseware", "course_id = ? and lecture_id = ? and courseware_id = ?", lecture.CourseId, lecture.LectureId, ware.CoursewareId) {
				// 模仿延时
				delay()
				e.fetchVideoURL(lecture, ware)
				e.saveCourseware(lecture, ware)
			} else {
				logger.InfoF("章节【{}】已存在", lecture.LectureName)
			}
		}
	}
}

func (e *Engine) fetchVideoURL(lecture *res.Lecture, ware *res.Courseware) {
	logger.InfoF("正在获取章节【{}】课时【{}】的视频信息...", lecture.LectureName, ware.Name)
	bs, err := client.GetWithHeader(api(apiWatch), map[string]interface{}{
		headerUserAgent: NextUserAgent(),
		headerForwarded: NextIP(),
	}, map[string]interface{}{
		"courseId":     lecture.CourseId,
		"lectureId":    lecture.LectureId,
		"coursewareId": ware.CoursewareId,
	})
	if err != nil {
		// 重试
		e.retries[retryKey("fetchVideoURL", lecture.LectureId)] = e.retries[retryKey("fetchVideoURL", lecture.LectureId)] + 1
		logger.ErrorF("章节【{}】课时【{}】的视频信息获取失败，正在第【{}】次重试...", lecture.LectureName, ware.Name, e.retries[retryKey("fetchVideoURL", lecture.LectureId)])
		// 模仿延时
		delay()
		e.fetchVideoURL(lecture, ware)
		return
	}

	var videoURL string
	if err = res.Decode(bs, &videoURL); err != nil {
		// 重试
		e.retries[retryKey("fetchVideoURL", lecture.LectureId)] = e.retries[retryKey("fetchVideoURL", lecture.LectureId)] + 1
		logger.ErrorF("章节【{}】课时【{}】的视频信息获取失败，正在第【{}】次重试...", lecture.LectureName, ware.Name, e.retries[retryKey("fetchVideoURL", lecture.LectureId)])
		// 模仿延时
		delay()
		e.fetchVideoURL(lecture, ware)
		return
	}
	ware.Video = videoURL
	logger.InfoF("章节【{}】课时【{}】的视频信息获取成功", lecture.LectureName, ware.Name)
}

func (e *Engine) saveCourseware(lecture *res.Lecture, ware *res.Courseware) {
	logger.InfoF("正在保存章节【{}】的课时【{}】信息...", lecture.LectureName, ware.Name)
	if err := orm.DB.Create(&entity.Courseware{
		CourseId:     ware.CourseId,
		LectureId:    ware.LectureId,
		LectureName:  lecture.LectureName,
		PlayStatus:   ware.PlayStatus,
		CoursewareId: ware.CoursewareId,
		Name:         ware.Name,
		Video:        ware.Video,
	}).Error; err != nil {
		panic(err)
	}
	logger.InfoF("保存章节【{}】的课时【{}】信息成功", lecture.LectureName, ware.Name)
}

func retryKey(fn string, recordId int) string {
	return fmt.Sprintf("%s-%d", fn, recordId)
}

func delay() {
	time.Sleep(time.Duration(rand.Intn(3)+2) * time.Second)
	//time.Sleep(time.Second)
}
