// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-08-31 14:51
// version: 1.0.0
// desc   :

package res

import (
	"encoding/json"
	"errors"
)

// R ...
type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// OK ...
func (r *R) OK() bool {
	return r.Code == 0
}

// Decode ...
func Decode(bs []byte, value interface{}) error {
	r := &R{}
	if err := json.Unmarshal(bs, r); err != nil {
		return err
	}

	if !r.OK() {
		return errors.New(r.Msg)
	}

	if bs, err := json.Marshal(r.Data); err != nil {
		return err
	} else {
		return json.Unmarshal(bs, value)
	}
}

// Unit 单元
type Unit struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Counts int    `json:"counts"`
}

// Course 课程
type Course struct {
	GradeName   string `json:"gradeName"`
	CourseName  string `json:"courseName"`
	GradeId     string `json:"gradeId"`
	TeacherId   string `json:"teacherId"`
	TeacherName string `json:"teacherName"`
	CourseId    int    `json:"courseId"`
	SubjectId   int    `json:"subjectId"`
	SubjectName string `json:"subjectName"`
}

// CoursePage 课程分页
type CoursePage struct {
	Total   int      `json:"total"`
	Size    int      `json:"size"`
	Current int      `json:"current"`
	Pages   int      `json:"pages"`
	Records []Course `json:"records"`
}

// Courseware 章节
type Courseware struct {
	PlayStatus   string `json:"playStatus"`
	CoursewareId string `json:"coursewareId"`
	Name         string `json:"name"`
	CourseId     int    `json:"courseId"`
	LectureId    int    `json:"lectureId"`
	Video        string `json:"video"`
}

// Lecture 章
type Lecture struct {
	LectureId   int           `json:"lectureId"`
	LectureName string        `json:"lectureName"`
	Like        string        `json:"like"`
	CourseId    int           `json:"courseId"`
	Courseware  []*Courseware `json:"courseware"`
}
