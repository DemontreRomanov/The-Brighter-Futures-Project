package main

import "fmt"

// EducationProgram is a program that provides education and training to marginalized individuals and communities 
type EducationProgram struct {
	courses []Course
	resources []Resource
}

//Course is a single course type within the program
type Course struct {
	name string
	duration int 
	description string
}

//Resource provides additional learning materials to participants
type Resource struct {
	name string
	url string 
	description string
}

//NewEducationProgram creates and instance of EducationProgram
func NewEducationProgram(courses []Course, resources []Resource) *EducationProgram {
	return &EducationProgram{
		courses:courses,
		resources:resources,
	}
}

//addCourse adds a new course to the program
func (e *EducationProgram) addCourse(course Course) {
	e.courses = append(e.courses, course)
}

// removeCourse removes a course from the program
func (e *EducationProgram) removeCourse(courseName string) {
	for i, course := range e.courses {
		if course.name == courseName {
			e.courses = append(e.courses[0:i], e.courses[i+1 : len(e.courses)-1]...)
			break
		}
	}
}

// addResource adds a new resource to the program
func (e *EducationProgram) addResource(resource Resource) {
	e.resources = append(e.resources, resource)
}

// removeResource removes a resource from the program
func (e *EducationProgram) removeResource(resourceName string) {
	for i, resource := range e.resources {
		if resource.name == resourceName {
			e.resources = append(e.resources[0:i], e.resources[i+1:len(e.resources)-1]...)
			break
		}
	}
}

// printCourses prints all the courses in the program
func (e *EducationProgram) printCourses() {
	for _, course := range e.courses {
		fmt.Println(course.name)
		fmt.Println(course.duration)
		fmt.Println(course.description)
	}
}

// printResources prints all the resources in the program
func (e *EducationProgram) printResources() {
	for _, resource := range e.resources {
		fmt.Println(resource.name)
		fmt.Println(resource.url)
		fmt.Println(resource.description)
	}
}

func main () {
	// Initialize courses and resources
	courses := []Course{
		{
			name: "Course 1",
			duration: 100,
			description: "This is Course 1",
		},
		{
			name:"Course 2",
			duration: 200,
			description: "This is Course 2",
		},
	}
	resources := []Resource{
		{
			name:"Resource 1",
			url: "www.resource1.com",
			description: "This is Resource 1",
		},
		{
			name:"Resource 2",
			url: "www.resource2.com",
			description: "This is Resource 2",
		},
	}

	// Create a new EducationProgram
	newProgram := NewEducationProgram(courses, resources)

	// Add a new course
	newProgram.addCourse(Course{
		name: "Course 3",
		duration: 300,
		description: "This is Course 3",
	})

	// Remove a course
	newProgram.removeCourse("Course 2")

	// Add a new resource
	newProgram.addResource(Resource{
		name:"Resource 3",
		url: "www.resource3.com",
		description: "This is Resource 3",
	})

	// Remove a resource
	newProgram.removeResource("Resource 2")

	// Print all courses
	newProgram.printCourses()

	// Print all resources
	newProgram.printResources()
}