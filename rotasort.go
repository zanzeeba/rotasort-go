package main

import (
	"errors"
	"fmt"
	"time"
)

// mine

// transfer qb from db to template

type Staff struct {
	Id          int64
	Email       string
	Username    string
	Password    string
	Phone       string
	Address     string
	Postcode    string
	Firstname   string
	Lastname    string
	CompaniesId int64
	StoresId    int64
	DeptId      int64
	Siteowner   bool
	Active      bool
	BreaksId    int64
	CreatedOn   time.Time
	UpdatedOn   time.Time
}

type Tasks struct {
	Id          int64
	TaskName    string
	TaskType    string
	Weighting   string
	TimeWhen    time.Time
	Time_Offset int64
	TimeFloat   bool
	TimeLength  int64
	NoOfJobs    int64
	TimeMin     int64
	TimeMax     int64
	DayOfWeek   string
	Active      bool
	CompaniesId int64
	StoresId    int64
	DeptId      int64
	UpdatedOn   time.Time
	CreatedOn   time.Time
}

type Holidays struct {
	Id           int64
	HolidayName  string
	HolidayStart time.Time
	HolidayEnd   time.Time
	CompaniesId  int64
	StoresId     int64
	DeptId       int64
	CreatedOn    time.Time
	UpdatedOn    time.Time
}

type Breaks struct {
	iId         int64
	BreakName   string
	BreakLength int64
	CompaniesId int64
	StoresId    int64
	DeptId      int64
	CreatedOn   time.Time
	UpdatedOn   time.Time
}

// depoels

var (
	// ErrEmptyName - name is required
	ErrEmptyName = errors.New("name not set")
)

// Shift contains job tree hierarchy.
type Task struct {
	ID       string `json:"taskId"`
	TaskName string `json:"taskName"`
}

// BusinessUnitJob contains display name attributes.
type BusinessUnitJob struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

// BusinessUnitJobs contain list of jobs available at a given Business Unit.
type BusinessUnitJobs []*BusinessUnitJob

// Job contains job tree hierarchy.
type Job struct {
	ID                     int64    `json:"id"`
	ClientID               int64    `json:"clientId"`
	Name                   string   `json:"name"`
	CategoryCode           string   `json:"categoryCode,omitempty"`
	GoldStandard           *float64 `json:"goldStandard,omitempty"`
	CostSavingsDepoelShare *float64 `json:"costSavingsDepoelShare,omitempty"`
	AgencyFee              *float64 `json:"agencyFee,omitempty"`
	PostCostSavingFee      *float64 `json:"postCostSavingFee,omitempty"`
	PotentialVATExemption  bool     `json:"potentialVATExemption"` // todo E5-731 VAT exemption is not a bool, but - 0 no, 1 yes, 2 agency decides
	FullSearchName         string   `json:"fullSearchName"`
	Children               Jobs     `json:"children,omitempty"`
}

// Shift contains job tree hierarchy.
type Shift struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Shift contains job tree hierarchy.
type JobName struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Validate validates the Group struct.
func (j *Job) Validate() error {
	if j.Name == "" {
		return ErrEmptyName
	}

	return nil
}

// GetFullSearchName builds the full search name for a job.
func (j *Job) GetFullSearchName() string {
	return fmt.Sprintf("%s (#%d)", j.Name, j.ID)
}

// Jobs contain list of jobs.
type Jobs []*Job

// Nominal code structure
type NominalStructure struct {
	JobID          int64  `json:"jobId"`
	BusinessUnitID int64  `json:"businessUnitId"`
	Format         string `json:"format"`
}

// A list of nominal code structures.
type NominalStructures []*NominalStructure
