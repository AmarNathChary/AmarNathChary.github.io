package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Profile struct {
	Name    string
	Title   string
	Contact string
	Summary string
}

type Skill struct {
	Category string
	Items    []string
}

type Experience struct {
	Title        string
	Company      string
	Duration     string
	Description  []string
	Technologies []string
}

type Education struct {
	Degree      string
	Institution string
	Year        string
}

type Certificate struct {
	Name string
}

type Award struct {
	Name        string
	Description string
}

type Portfolio struct {
	Profile      Profile
	Skills       []Skill
	Experiences  []Experience
	Educations   []Education
	Certificates []Certificate
	Awards       []Award
}

var portfolio = Portfolio{
	Profile: Profile{
		Name:    "N. Amarnath Chary",
		Title:   "Software Engineer",
		Contact: "+91-7794812511 | amarnathnaduminti@gmail.com | https://www.linkedin.com/in/amarnath-chary-25436a278/",
		Summary: "Enthusiastic and results-driven Software Engineer with a passion for creating innovative solutions. Bringing over 2+ years of hands-on experience in developing optimized, high-performing web applications using Go, Kubernetes, Docker, and MySQL. Committed to enhancing user experiences and driving business value through robust software solutions.",
	},
	Skills: []Skill{
		{Category: "Programming Language", Items: []string{"Golang"}},
		{Category: "Database", Items: []string{"MySQL"}},
		{Category: "DevOps Tools", Items: []string{"Docker", "Kubernetes", "Git", "Linux"}},
	},
	Experiences: []Experience{
		{
			Title:    "Software Engineer",
			Company:  "NEC Corporation India PVT LTD",
			Duration: "2024 (Sept) – Present | Noida",
			Description: []string{
				"Developed and designed a web application using Go.",
				"Fetched user data and stored it in the database.",
				"Built optimized, high-performing websites using Golang.",
				"Implemented form validations, API integrations, and JWT token authentication.",
				"Handled error management and utilized goroutines and channels for concurrent processing.",
				"Developed an API for managing application workflow and database maintenance.",
				"Designed and implemented user authentication to ensure robust performance.",
				"Business Value Addition: Contributed to the development and design of APIs and modules.",
			},
			Technologies: []string{"Golang", "Linux", "MySQL", "Git"},
		},
		{
			Title:    "OSS Community Developer (Kubernetes)",
			Company:  "NEC Corporation India PVT LTD",
			Duration: "2024 (Sept) – Present",
			Description: []string{
				"Resolved issues raised by the OSS community while contributing to open-source projects.",
				"Reviewed changes in feature implementations in the Kubernetes project.",
				"Created a backup solution and documented pod backup procedures in Kubernetes.",
				"Enhanced project support by adding features to redirect users to Kubernetes.",
				"Business Value Addition: Documented and implemented new features in the Kubernetes project.",
			},
			Technologies: []string{"Golang", "Docker", "Linux"},
		},
		{
			Title:    "Software Engineer",
			Company:  "C Ahead Digital",
			Duration: "2022 (Aug) – 2024 (Mar) | Bengaluru",
			Description: []string{
				"Developed RESTful APIs using Golang for a Feedback Management System project.",
				"Designed and implemented end-to-end solutions, including database schema design.",
				"Utilized MySQL for efficient data storage and retrieval.",
				"Optimized system scalability and performance for a responsive user experience.",
				"Used Git for version control to ensure smooth collaboration.",
				"Collaborated with cross-functional teams to resolve project challenges.",
			},
			Technologies: []string{"Golang", "MySQL", "Git"},
		},
	},
	Educations: []Education{
		{
			Degree:      "Bachelor of Business Administration (BBA)",
			Institution: "Rukmini College, Osmania University",
			Year:        "Completed in 2021",
		},
	},
	Certificates: []Certificate{
		{Name: "Golang"},
		{Name: "Kubernetes"},
		{Name: "Docker"},
		{Name: "Git"},
	},
	Awards: []Award{
		{
			Name:        "Team Spot Award, NEC",
			Description: "Awarded for exceptional performance and teamwork at NEC India.",
		},
		{
			Name:        "Member Role in Kubernetes Community Project",
			Description: "Recognized for exceptional contributions and performance in the Kubernetes project.",
		},
	},
}

func main() {
	tmpl := template.New("index.html").Funcs(template.FuncMap{
		"join": strings.Join,
	})
	tmpl, err := tmpl.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, portfolio)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
