package main

import (
	"fmt"
	"time"
)

type TeamMember interface {
	PrintName()
	PrintDetails()
}

type Employee struct {
	FirstName, LastName string
	Dob                 time.Time
	JobTitle, Location  string
}

type Developer struct {
	Employee
	Skills []string
}

// Overrides the PrintDetails
func (d Developer) PrintDetails() {
	// Call Employee PrintDetails
	d.Employee.PrintDetails()
	fmt.Println("Technical Skills:")
	for _, v := range d.Skills {
		fmt.Println(v)
	}
}
func (e Employee) PrintName() {
	fmt.Printf("\n%s %s\n", e.FirstName, e.LastName)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Date of Birth: %s, Job: %s, Location: %s\n", e.Dob.String(), e.JobTitle,
		e.Location)
}

type Manager struct {
	Employee  //type embedding for composition
	Projects  []string
	Locations []string
}

type Team struct {
	Name, Description string
	TeamMembers       []TeamMember
}

func (t Team) PrintTeamDetails() {
	fmt.Printf("Team: %s - %s\n", t.Name, t.Description)
	fmt.Println("Details of the team members:")
	for _, v := range t.TeamMembers {
		v.PrintName()
		v.PrintDetails()
	}
}

// Overrides the PrintDetails
func (m Manager) PrintDetails() {
	// Call Employee PrintDetails
	m.Employee.PrintDetails()
	fmt.Println("Projects:")
	for _, v := range m.Projects {
		fmt.Println(v)
	}
	fmt.Println("Managing teams for the locations:")
	for _, v := range m.Locations {
		fmt.Println(v)
	}
}

func main() {
	emp := Employee{
		FirstName: "Bob",
		LastName:  "Marle",
		Dob:       time.Time{},
		JobTitle:  "Dev",
		Location:  "DKDKDKD",
	}

	d := Developer{
		Employee: emp,
		Skills:   []string{"Go", "Docker", "Kuberntes"},
	}
	d.PrintName()
	d.PrintDetails()
}
