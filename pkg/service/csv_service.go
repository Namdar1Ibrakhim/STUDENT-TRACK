package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/repository"
	"io"
	"strconv"
	"strings"
)

type CSVService struct {
	repo repository.CSV
}

func NewCSVService(repo repository.CSV) *CSVService {
	return &CSVService{repo: repo}
}

func (s *CSVService) ValidateCSV(file io.Reader) error {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return errors.New("invalid CSV structure")
	}

	expectedHeaders := []string{"subjectname", "grade"}
	if len(records) == 0 || !s.equalHeaders(records[0], expectedHeaders) {
		return errors.New("invalid CSV structure, expected columns: subjectName, Grade")
	}

	for i, row := range records[1:] {
		if len(row) != 2 {
			return fmt.Errorf("invalid number of columns at row %d", i+2)
		}

		subjectName := row[0]
		gradeStr := row[1]

		if subjectName == "" {
			return fmt.Errorf("missing subjectName at row %d", i+2)
		}
		if gradeStr == "" {
			return fmt.Errorf("missing grade at row %d", i+2)
		}

		grade, err := strconv.Atoi(gradeStr)
		if err != nil {
			return fmt.Errorf("invalid grade value at row %d, must be a number", i+2)
		}

		if grade < 0 || grade > 100 {
			return fmt.Errorf("invalid grade value at row %d, must be between 0 and 100", i+2)
		}
	}

	return nil
}

func (s *CSVService) equalHeaders(headers, expectedHeaders []string) bool {
	if len(headers) != len(expectedHeaders) {
		return false
	}

	for i, header := range headers {
		if strings.TrimSpace(strings.ToLower(header)) != expectedHeaders[i] {
			return false
		}
	}
	return true
}
