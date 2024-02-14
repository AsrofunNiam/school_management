// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"
)

type Class struct {
	ID            int64  `json:"id"`
	ClassGroupsID string `json:"class_groups_id"`
	Name          string `json:"name"`
}

type ClassGroup struct {
	// kode jurusan
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Student struct {
	ID  int64  `json:"id"`
	Nis string `json:"nis"`
}

type StudentClass struct {
	ID        int64 `json:"id"`
	StudentID int64 `json:"student_id"`
	ClassID   int64 `json:"class_id"`
	// tahun awal TA
	Period string `json:"period"`
}

type Subject struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type SubjectReport struct {
	ID                int64         `json:"id"`
	CreatedAt         time.Time     `json:"created_at"`
	TeacherSubjectID  int64         `json:"teacher_subject_id"`
	StudentClassID    int64         `json:"student_class_id"`
	MidExamResult     sql.NullInt32 `json:"mid_exam_result"`
	FinalExamResult   sql.NullInt32 `json:"final_exam_result"`
	MonthlyExamResult []int32       `json:"monthly_exam_result"`
}

type Teacher struct {
	ID  int64  `json:"id"`
	Nip string `json:"nip"`
}

type TeacherSubject struct {
	ID        int64 `json:"id"`
	TeacherID int64 `json:"teacher_id"`
	SubjectID int64 `json:"subject_id"`
	// tahun awal TA
	Period string `json:"period"`
}

type User struct {
	// from nip or nis
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	// enum teacher, student
	Role string `json:"role"`
}
