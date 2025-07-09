package models

type StudentByIDPathParams struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type StudentResponse struct {
	Data Student `json:"data"`
}

type Student struct {
	ID                 int     `json:"id" label:"ID"`
	Name               string  `json:"name" label:"Name"`
	Email              string  `json:"email" label:"Email"`
	SystemAccess       bool    `json:"systemAccess" label:"System Access"`
	Phone              string  `json:"phone" label:"Phone"`
	Gender             string  `json:"gender" label:"Gender"`
	DOB                string  `json:"dob" label:"Date Of Birth"`
	Class              *string `json:"class" label:"Class"`
	Section            *string `json:"section" label:"Section"`
	Roll               *int    `json:"roll" label:"Roll"`
	FatherName         string  `json:"fatherName" label:"Father Name"`
	FatherPhone        *string `json:"fatherPhone" label:"Father Phone"`
	MotherName         string  `json:"motherName" label:"Mother Name"`
	MotherPhone        *string `json:"motherPhone" label:"Mother Phone"`
	GuardianName       *string `json:"guardianName" label:"Guardian Name"`
	GuardianPhone      *string `json:"guardianPhone" label:"Guardian Phone"`
	RelationOfGuardian *string `json:"relationOfGuardian" label:"Relation Of Guardian"`
	CurrentAddress     *string `json:"currentAddress" label:"Current Address"`
	PermanentAddress   *string `json:"permanentAddress" label:"Permanent Address"`
	AdmissionDate      *string `json:"admissionDate" label:"Admission Date"`
	ReporterName       *string `json:"reporterName" label:"Reporter Name"`
}
