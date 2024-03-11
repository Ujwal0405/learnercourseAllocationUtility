package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type LearningBlock struct {
	HasBlocked  bool   `json:"hasBlocked"`
	BlockedDate string `json:"blockedDate"`
	BlockDates  int64  `json:"blockDates"`
}

type LearningObject struct {
	LearnerId                     string        `json:"learnerId"`
	LearnerCode                   string        `json:"learnerCode"`
	LearnerCourseId               string        `json:"learnerCourseId"`
	LearningCenterId              string        `json:"learningCenterId"`
	LearningCenterCode            string        `json:"learningCenterCode"`
	LearningCenterName            string        `json:"learningCenterName"`
	ConfirmationMonth             int           `json:"confirmationMonth"`
	ConfirmationYear              int           `json:"confirmationYear"`
	BatchName                     string        `json:"batchName"`
	ProgramId                     string        `json:"programId"`
	ProgramName                   string        `json:"programName"`
	ECourseId                     string        `json:"eCourseId"`
	ECourseName                   string        `json:"eCourseName"`
	ECourseVersion                string        `json:"eCourseVersion"`
	ECoursePatternId              string        `json:"eCoursePatternId"`
	DateOfTransfer                int64         `json:"dateOfTransfer"`
	ECourseImage                  string        `json:"eCourseImage"`
	LearningEndDate               int64         `json:"learningEndDate"`
	LearningStartDate             int64         `json:"learningStartDate"`
	LearningLimit                 int           `json:"learningLimit"`
	PerDaySessionLimit            int           `json:"perDaySessionLimit"`
	LearnerLifeCycleEndDate       int64         `json:"learnerLifeCycleEndDate"`
	LearnerCoursePrefLanguageCode string        `json:"learnerCoursePrefLanguageCode"`
	LearningBlock                 LearningBlock `json:"learningBlock"`
}

func main() {
	var learningObjects []LearningObject

	// Create 50 objects with different learnerId and learnerCode
	for i := 1; i <= 50; i++ {
		learnerId := fmt.Sprintf("2023200%d", i)
		learnerCode := fmt.Sprintf("2023200%d", i)

		learningObject := LearningObject{
			LearnerId:                     learnerId,
			LearnerCode:                   learnerCode,
			LearnerCourseId:               "1",
			LearningCenterId:              "54547771",
			LearningCenterCode:            "54547771",
			LearningCenterName:            "HARSH COMPUTERS",
			ConfirmationMonth:             10,
			ConfirmationYear:              2023,
			BatchName:                     "October 2023",
			ProgramId:                     "cd2d2b5fcdbdd81f7737f79bc9e7b16e",
			ProgramName:                   "KLiC Courses 2023",
			ECourseId:                     "e1e47b7ec5e48481d9ea6292e72ec84d",
			ECourseName:                   "KLiC Advanced Excel (ALC Mode)",
			ECourseVersion:                "e7171985fc4510424bb863b52ba7d407",
			ECoursePatternId:              "1",
			DateOfTransfer:                time.Now().Unix(),
			ECourseImage:                  "/o/programs/d42b0856acf99bef7803b504c1efea94/6ef0f33f08779e5adaeb6887bfb708c6_1/6ef0f33f08779e5adaeb6887bfb708c6.png",
			LearningEndDate:               time.Now().AddDate(0, 0, 30).Unix(),
			LearningStartDate:             time.Now().Unix(),
			LearningLimit:                 0,
			PerDaySessionLimit:            0,
			LearnerLifeCycleEndDate:       time.Now().AddDate(0, 0, 60).Unix(),
			LearnerCoursePrefLanguageCode: "1",
			LearningBlock: LearningBlock{
				HasBlocked:  true,
				BlockedDate: time.Now().Format("2006-01-02T15:04:05-07:00"),
				BlockDates:  0,
			},
		}

		learningObjects = append(learningObjects, learningObject)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(learningObjects, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Write to a JSON file
	file, err := os.Create("learnerAllocationObjs.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data written to learnerAllocationObjs.json")
}
