package types

import (
	"encoding/json"
)

// Employee contains all the fields of employees provided by the API.
// (https://www.bamboohr.com/api/documentation/employees.php)
type EmployeeUpdate struct {
	Address1                *string  `json:"address1,omitempty"`                    // The employee's first address line.
	Address2                *string  `json:"address2,omitempty"`                    // The employee's second address line.
	Age                     *int     `json:"age,string,omitempty"`                  // The employee's age. To change age, update dateOfBirth field.
	BestEmail               *string  `json:"bestEmail,omitempty"`                   // The employee's work email if set, otherwise their home email.
	Birthday                *string  `json:"birthday,omitempty"`                    // The employee's month and day of birth. To change birthday, update dateOfBirth field.
	City                    *string  `json:"city,omitempty"`                        // The employee's city.
	Country                 *string  `json:"country,omitempty"`                     // The employee's country.
	DateOfBirth             *string  `json:"dateOfBirth,omitempty"`                 // The date the employee was born.
	Department              *string  `json:"department,omitempty"`                  // The employee's CURRENT department.
	Division                *string  `json:"division,omitempty"`                    // The employee's CURRENT division.
	Eeo                     *string  `json:"eeo,omitempty"`                         // The employee's EEO job category. These are defined by the U.S. Equal Employment Opportunity Commission.
	EmployeeNumber          *string  `json:"employeeNumber,omitempty"`              // Employee number (assigned by your company).
	EmploymentHistoryStatus *string  `json:"employmentHistoryStatus,omitempty"`     // The employee's CURRENT employment status. Options are customized by account. Read-only starting with version 1.1; update using the employmentStatus table.
	Ethnicity               *string  `json:"ethnicity,omitempty"`                   // The employee's ethnicity.
	Exempt                  *string  `json:"exempt,omitempty"`                      // The FLSA Overtime status (Exempt or Non-exempt).
	FirstName               *string  `json:"firstName,omitempty"`                   // The employee's first name.
	FullName1               *string  `json:"fullName1,omitempty"`                   // The employee's first and last name. (e.g., John Doe). Read only.
	FullName2               *string  `json:"fullName2,omitempty"`                   // The employee's last and first name. (e.g., Doe, John). Read only.
	FullName3               *string  `json:"fullName3,omitempty"`                   // The employee's full name and their preferred name. (e.g., Doe, John Quentin (JDog)). Read only.
	FullName4               *string  `json:"fullName4,omitempty"`                   // The employee's full name without their preferred name, last name first. (e.g., Doe, John Quentin). Read only.
	FullName5               *string  `json:"fullName5,omitempty"`                   // The employee's full name without their preferred name, first name first. (e.g., John Quentin Doe). Read only.
	DisplayName             *string  `json:"displayName,omitempty"`                 // The employee's name displayed in a format configured by the user. Read only.
	Gender                  *string  `json:"gender,omitempty"`                      // The employee's gender (Male or Female).
	HireDate                *string  `json:"hireDate,omitempty"`                    // The date the employee was hired.
	OriginalHireDate        *string  `json:"originalHireDate,omitempty"`            // The date the employee was originally hired. Available starting with version 1.1.
	HomeEmail               *string  `json:"homeEmail,omitempty"`                   // The employee's home mmail address.
	HomePhone               *string  `json:"homePhone,omitempty"`                   // The employee's home phone number.
	ID                      *int     `json:"id,string,omitempty"`                   // The employee ID automatically assigned by BambooHR. Read only.
	JobTitle                *string  `json:"jobTitle,omitempty"`                    // The CURRENT value of the employee's job title, updating this field will create a new row in position history.
	LastChanged             *string  `json:"lastChanged,omitempty"`                 // The date and time that the employee record was last changed.
	LastName                *string  `json:"lastName,omitempty"`                    // The employee's last name.
	Location                *string  `json:"location,omitempty"`                    // The employee's CURRENT location.
	MaritalStatus           *string  `json:"maritalStatus,omitempty"`               // The employee's marital status (Single, Married, or Domestic Partnership).
	MiddleName              *string  `json:"middleName,omitempty"`                  // The employee's middle name.
	MobilePhone             *string  `json:"mobilePhone,omitempty"`                 // The employee's mobile phone number.
	PayChangeReason         *string  `json:"payChangeReason,omitempty"`             // The reason for the employee's last pay rate change.
	PayGroup                *string  `json:"payGroup,omitempty"`                    // The custom pay group that the employee belongs to.
	PayGroupID              *int     `json:"payGroupId,string,omitempty"`           // The ID value corresponding to the pay group that an employee belongs to.
	PayRate                 *float32 `json:"payRate,omitempty"`                     // The employee's CURRENT pay rate (e.g., $8.25).
	PayRateEffectiveDate    *string  `json:"payRateEffectiveDate,omitempty"`        // The day the most recent change was made.
	PayType                 *string  `json:"payType,omitempty"`                     // The employee's CURRENT pay type. ie: "hourly","salary","commission","exception hourly","monthly","weekly","piece rate","contract","daily","pro rata".
	PayPer                  *string  `json:"payPer,omitempty"`                      // The employee's CURRENT pay per. ie: "Hour", "Day", "Week", "Month", "Quarter", "Year".
	PaidPer                 *string  `json:"paidPer,omitempty"`                     // The employee's CURRENT pay per. ie: "Hour", "Day", "Week", "Month", "Quarter", "Year".
	PaySchedule             *string  `json:"paySchedule,omitempty"`                 // The employee's CURRENT pay schedule.
	PayScheduleID           *int     `json:"payScheduleId,string,omitempty"`        // The ID value corresponding to the pay schedule that an employee belongs to.
	PayFrequency            *string  `json:"payFrequency,omitempty"`                // The employee's CURRENT pay frequency. ie: "Weekly", "Every other week", "Twice a month", "Monthly", "Quarterly", "Twice a year", or "Yearly"
	IncludeInPayroll        *bool    `json:"includeInPayroll,string,omitempty"`     // Should employee be included in payroll (Yes or No)
	TimeTrackingEnabled     *bool    `json:"timeTrackingEnabled,string,omitempty"`  // Should time tracking be enabled for the employee (Yes or No)
	PreferredName           *string  `json:"preferredName,omitempty"`               // The employee's preferred name.
	Ssn                     *string  `json:"ssn,omitempty"`                         // The employee's Social Security number.
	Sin                     *string  `json:"sin,omitempty"`                         // The employee's Canadian Social Insurance Number.
	State                   *string  `json:"state,omitempty"`                       // The employee's state/province.
	StateCode               *string  `json:"stateCode,omitempty"`                   // The 2 character abbreviation for the employee's state (US only). Read only.
	Status                  *string  `json:"status,omitempty"`                      // The employee's employee status (Active or Inactive).
	Supervisor              *string  `json:"supervisor,omitempty"`                  // The employee’s CURRENT supervisor. Read only.
	SupervisorID            *string  `json:"supervisorId,omitempty"`                // The 'employeeNumber' of the employee's CURRENT supervisor. Read only.
	SupervisorEId           *int     `json:"supervisorEId,string,omitempty"`        // The ID of the employee's CURRENT supervisor. Read only.
	TerminationDate         *string  `json:"terminationDate,omitempty"`             // The date the employee was terminated. Read-only starting with version 1.1; update using the employmentStatus table.
	WorkEmail               *string  `json:"workEmail,omitempty"`                   // The employee's work email address.
	WorkPhone               *string  `json:"workPhone,omitempty"`                   // The employee's work phone number, without extension.
	WorkPhonePlusExtension  *string  `json:"workPhonePlusExtension,omitempty"`      // The employee's work phone and extension. Read only.
	WorkPhoneExtension      *string  `json:"workPhoneExtension,omitempty"`          // The employee's work phone extension (if any).
	Zipcode                 *string  `json:"zipcode,omitempty"`                     // The employee's ZIP code.
	IsPhotoUploaded         *string  `json:"isPhotoUploaded,omitempty"`             // Whether a photo has been uploaded for the employee. Read only.
	AcaStatus               *string  `json:"acaStatus,omitempty"`                   // The employee's ACA (Affordable Care Act) Full-Time status. Options are yes, no, non-assessment
	StandardHoursPerWeek    *int     `json:"standardHoursPerWeek,string,omitempty"` // The number of hours the employee works in a standard week.
	BonusDate               *string  `json:"bonusDate,omitempty"`                   // The date of the last bonus.
	BonusAmount             *float32 `json:"bonusAmount,string,omitempty"`          // The amount of the most recent bonus.
	BonusReason             *string  `json:"bonusReason,omitempty"`                 // The reason for the most recent bonus.
	BonusComment            *string  `json:"bonusComment,omitempty"`                // Comment about the most recent bonus.
	CommissionDate          *string  `json:"commissionDate,omitempty"`              // The date of the last commission.
	CommissionAmount        *float32 `json:"commissionAmount,string,omitempty"`     // The amount of the most recent commission.
	CommissionComment       *string  `json:"commissionComment,omitempty"`           // Comment about the most recent commission.
	Nin                     *string  `json:"nin,omitempty"`                         // The employee's nin number
	NationalID              *string  `json:"nationalID,omitempty"`                  // The employee's national ID number
	Nationality             *string  `json:"nationality,omitempty"`                 // The employee's nationality
	Notes                   *string  `json:"customNotes1,omitempty"`                // The employee's notes
}

// NewEmployeeFromJSON parses the given JSON (as byte sequence) and returns a new Employee.
func NewEmployeeUpdateFromJSON(data []byte) (*EmployeeUpdate, error) {
	var employee EmployeeUpdate
	err := json.Unmarshal(data, &employee)
	return &employee, err
}
