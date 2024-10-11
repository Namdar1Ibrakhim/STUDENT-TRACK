package repository

import (
	"fmt"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) GetAllCourse() ([]dto.CourseResponse, error) {
	query := "SELECT id, course_name, description FROM course"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []dto.CourseResponse
	for rows.Next() {
		var course dto.CourseResponse
		err := rows.Scan(&course.Id, &course.Course_name, &course.Description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *CourseRepository) GetCourseById(courseId int) (dto.CourseResponse, error) {
	var course dto.CourseResponse
	query := "SELECT id, course_name, description FROM course WHERE id = $1"

	err := r.db.QueryRow(query, courseId).Scan(&course.Id, &course.Course_name, &course.Description)
	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *CourseRepository) GetCourseByName(courseName string) (dto.CourseResponse, error) {
	var course dto.CourseResponse
	query := "SELECT id, course_name, description FROM course WHERE course_name = $1"

	err := r.db.QueryRow(query, courseName).Scan(&course.Id, &course.Course_name, &course.Description)
	if err != nil {
		return course, err
	}

	return course, nil
}

func (r *CourseRepository) FindCourseIDByName(courseName string) (int, error) {
	var courseID int
	query := fmt.Sprintf("SELECT id FROM %s WHERE course_name = $1", courseTable)
	err := r.db.Get(&courseID, query, courseName)
	if err != nil {
		return 0, fmt.Errorf("course not found -> %v", courseName)
	}
	return courseID, nil
}
