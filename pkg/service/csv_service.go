package service

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
	pb "github.com/Namdar1Ibrakhim/student-track-system/proto"
	"io"
	"strconv"
	"strings"
	"time"
)

type CSVService struct {
	repo     repository.Predictions
	repo2    repository.Course
	repo3    repository.StudentCourse
	repo4    repository.Direction
	mlClient pb.PredictionServiceClient
}

func NewCSVService(repo repository.Predictions, repo2 repository.Course, repo3 repository.StudentCourse, repo4 repository.Direction, mlClient pb.PredictionServiceClient) *CSVService {
	return &CSVService{
		repo:     repo,
		repo2:    repo2,
		repo3:    repo3,
		repo4:    repo4,
		mlClient: mlClient,
	}
}

func parseInt(num string) int {
	result, _ := strconv.Atoi(num)
	return result
}

func parseInt32(num string) int32 {
	result, _ := (strconv.Atoi(num))
	return int32(result)
}

func (s *CSVService) ValidateCSVForStudent(file io.Reader) error {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return errors.New("invalid CSV structure")
	}

	expectedHeaders := []string{
		"Operating System", "Analysis of Algorithm", "Programming Concept", "Software Engineering",
		"Computer Network", "Applied Mathematics", "Computer Security", "Hackathons attended",
		"Topmost Certification", "Personality", "Management or technical", "Leadership", "Team", "Self Ability"}

	if len(records) == 0 || !s.equalHeaders(records[0], expectedHeaders) {
		return errors.New("invalid CSV structure, expected columns: check the required format")
	}

	return s.validateRows(records, expectedHeaders, false)
}

func (s *CSVService) ValidateCSVForInstructor(file io.Reader) error {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return errors.New("invalid CSV structure")
	}

	expectedHeaders := []string{
		"Operating System", "Analysis of Algorithm", "Programming Concept", "Software Engineering",
		"Computer Network", "Applied Mathematics", "Computer Security", "Hackathons attended",
		"Topmost Certification", "Personality", "Management or technical", "Leadership", "Team", "Self Ability", "Student_id"}

	if len(records) == 0 || !s.equalHeaders(records[0], expectedHeaders) {
		return errors.New("invalid CSV structure, expected columns: check the required format(Instructor)")
	}

	return s.validateRows(records, expectedHeaders, true)
}

func (s *CSVService) validateRows(records [][]string, expectedHeaders []string, hasStudentID bool) error {
	for i, row := range records[1:] {
		if len(row) != len(expectedHeaders) {
			return fmt.Errorf("invalid number of columns at row %d", i+2)
		}

		for j := 0; j <= 6; j++ {
			gradeStr := row[j]
			if gradeStr == "" {
				return fmt.Errorf("missing grade for subject at row %d, column %d", i+2, j+1)
			}

			grade, err := strconv.Atoi(gradeStr)
			if err != nil || grade < 0 || grade > 100 {
				return fmt.Errorf("invalid grade value at row %d, column %d, must be between 0 and 100", i+2, j+1)
			}
		}

		hackathonsStr := row[7]
		if hackathonsStr == "" {
			return fmt.Errorf("missing value for 'hackathons' at row %d", i+2)
		}
		hackathons, err := strconv.Atoi(hackathonsStr)
		if err != nil || hackathons < 0 {
			return fmt.Errorf("invalid 'Hackathons attended' value at row %d, must be a non-negative integer", i+2)
		}

		for j := 8; j < 14; j++ {
			if row[j] == "" {
				return fmt.Errorf("missing value for '%s' at row %d", expectedHeaders[j], i+2)
			}
		}

		if hasStudentID {
			if len(row) < len(expectedHeaders) {
				return fmt.Errorf("missing student_id at row %d", i+2)
			}
			if _, err := strconv.Atoi(row[len(expectedHeaders)-1]); err != nil {
				return fmt.Errorf("invalid student_id at row %d", i+2)
			}
		}
	}

	return nil
}

func (s *CSVService) PredictCSV(studentId int, file io.Reader, isInstructor bool) (map[int]*dto.PredictionResponseDto, error) {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("invalid CSV structure")
	}

	if isInstructor {
		return s.predictForMultipleStudents(records)
	} else {
		return s.predictForSingleStudent(studentId, records[1])
	}
}

func (s *CSVService) predictForSingleStudent(studentId int, row []string) (map[int]*dto.PredictionResponseDto, error) {
	courses := map[string]int{
		"Operating System":      parseInt(row[0]),
		"Analysis of Algorithm": parseInt(row[1]),
		"Programming Concept":   parseInt(row[2]),
		"Software Engineering":  parseInt(row[3]),
		"Computer Network":      parseInt(row[4]),
		"Applied Mathematics":   parseInt(row[5]),
		"Computer Security":     parseInt(row[6]),
	}
	for courseName, grade := range courses {
		courseID, err := s.repo2.FindCourseIDByName(courseName)
		if err != nil {
			return nil, fmt.Errorf("failed to find course ID for %s: %v", courseName, err)
		}

		err = s.repo3.AddStudentCourse(studentId, courseID, grade)
		if err != nil {
			return nil, fmt.Errorf("failed to save course data for student: %v", err)
		}
	}

	prediction, err := s.sendPredictionRequest(row)
	if err != nil {
		return nil, err
	}

	prediction.StudentId = studentId

	directionID, err := s.repo4.FindDirectionIDByName(prediction.Direction_name)
	if err != nil {
		return nil, fmt.Errorf("failed to find direction ID for predicted track: %v", err)
	}

	err = s.repo.SavePrediction(studentId, directionID)
	if err != nil {
		return nil, errors.New("failed to save prediction")
	}

	return map[int]*dto.PredictionResponseDto{
		studentId: prediction,
	}, nil
}

func (s *CSVService) predictForMultipleStudents(records [][]string) (map[int]*dto.PredictionResponseDto, error) {
	predictions := make(map[int]*dto.PredictionResponseDto)

	for i, row := range records[1:] {
		studentId, err := strconv.Atoi(row[14])
		if err != nil {
			return nil, fmt.Errorf("invalid student_id at row %d", i+2)
		}

		prediction, err := s.predictForSingleStudent(studentId, row[:14])
		if err != nil {
			return nil, fmt.Errorf("failed to predict for student_id %d: %v", studentId, err)
		}

		predictions[studentId] = prediction[studentId]
	}

	return predictions, nil
}

func (s *CSVService) sendPredictionRequest(row []string) (*dto.PredictionResponseDto, error) {
	predictionRequest := &pb.PredictionRequest{
		OperatingSystem:      parseInt32(row[0]),
		AnalysisOfAlgorithm:  parseInt32(row[1]),
		ProgrammingConcept:   parseInt32(row[2]),
		SoftwareEngineering:  parseInt32(row[3]),
		ComputerNetwork:      parseInt32(row[4]),
		AppliedMathematics:   parseInt32(row[5]),
		ComputerSecurity:     parseInt32(row[6]),
		HackathonsAttended:   parseInt32(row[7]),
		TopmostCertification: row[8],
		Personality:          row[9],
		ManagementTechnical:  row[10],
		Leadership:           row[11],
		Team:                 row[12],
		SelfAbility:          row[13],
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	resp, err := s.mlClient.Predict(ctx, predictionRequest)
	if err != nil {
		return nil, fmt.Errorf("Error while calling ML service: %v", err)
	}

	return &dto.PredictionResponseDto{
		Direction_name: resp.PredictedTrack,
	}, nil
}

func (s *CSVService) equalHeaders(headers, expectedHeaders []string) bool {
	if len(headers) != len(expectedHeaders) {
		return false
	}

	for i, header := range headers {
		if strings.TrimSpace(strings.ToLower(header)) != strings.TrimSpace(strings.ToLower(expectedHeaders[i])) {
			return false
		}
	}
	return true
}
