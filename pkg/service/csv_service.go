package service

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
	pb "github.com/Namdar1Ibrakhim/student-track-system/proto"
	"io"
	"strconv"
	"strings"
	"time"
)

type CSVService struct {
	repos    *Repositories
	mlClient pb.PredictionServiceClient
}

type Repositories struct {
	predictions   repository.Predictions
	course        repository.Course
	studentCourse repository.StudentCourse
	direction     repository.Direction
}

func NewCSVService(
	predictions repository.Predictions,
	course repository.Course,
	studentCourse repository.StudentCourse,
	direction repository.Direction,
	mlClient pb.PredictionServiceClient,
) *CSVService {
	return &CSVService{
		repos: &Repositories{
			predictions:   predictions,
			course:        course,
			studentCourse: studentCourse,
			direction:     direction,
		},
		mlClient: mlClient,
	}
}

type csvConfig struct {
	headers      []string
	hasStudentID bool
}

func getCSVConfig(isInstructor bool) csvConfig {
	headers := []string{
		"Operating System", "Analysis of Algorithm", "Programming Concept",
		"Software Engineering", "Computer Network", "Applied Mathematics",
		"Computer Security", "Hackathons attended", "Topmost Certification",
		"Personality", "Management or technical", "Leadership", "Team", "Self Ability",
	}

	if isInstructor {
		headers = append(headers, "Student_id")
	}

	return csvConfig{
		headers:      headers,
		hasStudentID: isInstructor,
	}
}

func (s *CSVService) ValidateCSVForStudent(file io.Reader) error {
	return s.validateCSV(file, getCSVConfig(false))
}

func (s *CSVService) ValidateCSVForInstructor(file io.Reader) error {
	return s.validateCSV(file, getCSVConfig(true))
}

func (s *CSVService) validateCSV(file io.Reader, config csvConfig) error {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil || len(records) == 0 {
		return fmt.Errorf("%w: failed to read CSV file", constants.ErrInvalidCSVStructure)
	}

	if !s.equalHeaders(records[0], config.headers) {
		return fmt.Errorf("%w: invalid headers", constants.ErrInvalidCSVStructure)
	}

	return s.validateRows(records[1:], config)
}

func (s *CSVService) validateRows(records [][]string, config csvConfig) error {
	for i, row := range records[1:] {
		if err := s.validateRow(row, i+2, config); err != nil {
			return err
		}
	}
	return nil
}

func (s *CSVService) validateRow(row []string, rowNum int, config csvConfig) error {
	if len(row) != len(config.headers) {
		return fmt.Errorf("invalid number of columns at row %d", rowNum)
	}

	for j := 0; j <= 6; j++ {
		gradeStr := row[j]
		if gradeStr == "" {
			return fmt.Errorf("missing grade for subject at row %d, column %d", rowNum, j+1)
		}

		grade, err := strconv.Atoi(gradeStr)
		if err != nil || grade < 0 || grade > 100 {
			return fmt.Errorf("invalid grade value at row %d, column %d, must be between 0 and 100", rowNum, j+1)
		}
	}

	hackathonsStr := row[7]
	if hackathonsStr == "" {
		return fmt.Errorf("missing value for 'hackathons' at row %d", rowNum)
	}
	hackathons, err := strconv.Atoi(hackathonsStr)
	if err != nil || hackathons < 0 {
		return fmt.Errorf("invalid 'Hackathons attended' value at row %d, must be a non-negative integer", rowNum)
	}

	for j := 8; j < 14; j++ {
		if row[j] == "" {
			return fmt.Errorf("missing value for '%s' at row %d", config.headers[j], rowNum)
		}
	}

	if config.hasStudentID {
		if len(row) < len(config.headers) {
			return fmt.Errorf("missing student_id at row %d", rowNum)
		}
		if _, err := strconv.Atoi(row[len(config.headers)-1]); err != nil {
			return fmt.Errorf("invalid student_id at row %d", rowNum)
		}
	}
	return nil
}

func (s *CSVService) PredictCSV(studentId int, file io.Reader, isInstructor bool) (map[int]*dto.PredictionResponseDto, error) {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New(constants.ErrInvalidCSVStructure.Error())
	}

	if isInstructor {
		return s.predictForMultipleStudents(records)
	} else {
		return s.predictForSingleStudent(studentId, records[1])
	}
}

func (s *CSVService) predictForSingleStudent(studentID int, row []string) (map[int]*dto.PredictionResponseDto, error) {
	if err := s.saveCourseData(studentID, row); err != nil {
		return nil, err
	}

	prediction, err := s.makePrediction(row)
	if err != nil {
		return nil, err
	}

	prediction.StudentId = studentID

	directionID, err := s.repos.direction.FindDirectionIDByName(prediction.Direction_name)
	if err != nil {
		return nil, fmt.Errorf("failed to find direction ID for predicted track: %v", err)
	}

	if err := s.repos.predictions.SavePrediction(studentID, directionID); err != nil {
		return nil, err
	}

	return map[int]*dto.PredictionResponseDto{
		studentID: prediction,
	}, nil
}

func (s *CSVService) saveCourseData(studentID int, row []string) error {
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
		courseID, err := s.repos.course.FindCourseIDByName(courseName)
		if err != nil {
			return fmt.Errorf("failed to find course ID for %s: %w", courseName, err)
		}

		if err := s.repos.studentCourse.AddStudentCourse(studentID, courseID, grade); err != nil {
			return fmt.Errorf("failed to save course data: %w", err)
		}
	}

	return nil
}

func (s *CSVService) makePrediction(row []string) (*dto.PredictionResponseDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := createPredictionRequest(row)
	resp, err := s.mlClient.Predict(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", constants.ErrMLServiceFailure, err)
	}

	return &dto.PredictionResponseDto{
		Direction_name: resp.PredictedTrack,
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

func createPredictionRequest(row []string) *pb.PredictionRequest {
	return &pb.PredictionRequest{
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
}

func parseInt(num string) int {
	result, _ := strconv.Atoi(num)
	return result
}

func parseInt32(num string) int32 {
	result, _ := (strconv.Atoi(num))
	return int32(result)
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
