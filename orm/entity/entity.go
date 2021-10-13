// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-08-31 15:50
// version: 1.0.0
// desc   :

package entity

// Subject 学科
type Subject struct {
	Id   int    `gorm:"column:id"`               //  ID
	Name string `gorm:"varchar(32);column:name"` // 名称
}

// TableName ...
func (s *Subject) TableName() string {
	return "subject"
}

// Unit 单元
type Unit struct {
	Id      int    `gorm:"column:id"`                // ID
	Subject int    `gorm:"int;column:subject"`       // 科目
	Name    string `gorm:"varchar(255);column:name"` // 名称
	Counts  int    `gorm:"int;column:counts"`        // 课程数量
}

// TableName ...
func (s *Unit) TableName() string {
	return "unit"
}

// Course 课程
type Course struct {
	GradeName   string `gorm:"varchar(255);column:grade_name"`   // 年级
	CourseName  string `gorm:"varchar(255);column:course_name"`  // 课程名称
	GradeId     string `gorm:"varchar(255);column:grade_id"`     // 年级ID
	TeacherId   string `gorm:"varchar(255);column:teacher_id"`   // 教师ID
	TeacherName string `gorm:"varchar(255);column:teacher-Name"` // 教师名称
	CourseId    int    `gorm:"column:course_id"`                 // 课程ID
	SubjectId   int    `gorm:"int;column:subject_id"`            // 学科ID
	SubjectName string `gorm:"varchar(32);column:subject_name"`  // 学科名称
	Unit        int    `gorm:"int;column:unit"`                  // 单元
}

// TableName ...
func (s *Course) TableName() string {
	return "course"
}

// Courseware 章节
type Courseware struct {
	CourseId     int    `gorm:"column:course_id"`                  // 课程ID
	LectureId    int    `gorm:"column:lecture_id"`                 // 章ID
	LectureName  string `gorm:"varchar(255);column:lecture_name"`  // 章名称
	CoursewareId string `gorm:"varchar(255);column:courseware_id"` // 章节ID
	Name         string `gorm:"varchar(255);column:name"`          // 章节名称
	Video        string `gorm:"varchar(1024);column:video"`        // 视频地址
}

// TableName ...
func (s *Courseware) TableName() string {
	return "courseware"
}
