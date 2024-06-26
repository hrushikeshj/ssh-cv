package cv

// https://github.com/charmbracelet/bubbletea/blob/master/examples/glamour/main.go

type Education struct {
	uni    string
	degree string
	start  string
	end    string
	cgpa   string
}

type Experience struct {
	role    string
	company string
	start   string
	end     string
	points  []string
}

type Project struct {
	name   string
	tech   string
	points []string
}

type Skill struct {
	typ    string
	skills string
}

type link struct {
	name string
	url  string
}

type CV struct {
	contact     string
	education   []Education
	experiences []Experience
	projects    []Project
	skills      []Skill
	links       []link
}

var MyCV = CV{
	contact: "hrushi2002j@gmail.com • +91 6363337912",
	education: []Education{
		{
			uni:    "National Institute of Technology Karnataka",
			degree: "Bachelor of Technology in CSE, Minor in ECE",
			start:  "December 2020",
			end:    "April 2024",
			cgpa:   "8.54",
		},
	},
	experiences: []Experience{
		{
			role:    "Web Developer",
			company: "IRIS, NITK",
			start:   "March 2021",
			end:     "April 2024",
			points: []string{
				"IRIS is the Management Information System(MIS) developed for automating all administrative and academic activities at NITK.",
				//"Developed Time Table module for students and professors to keep track of classes, exams, assignments, and other course activities",
				"Added Post Doctoral Fellowship application to the existing Applications module",
				"Implemented API endpoints for the Time Table and Forms module, and wrote API documentation for the App Team.",
				"Implemented new features for various modules like Admissions, Branch Change, Courses and No-Dues which are used by over 7000 students.",
			},
		},
		{
			role:    "Web Developer",
			company: "Centre for Innovation, IPR and Industrial Consultancy(CIC), NITK",
			start:   "January 2023",
			end:     "June 2023",
			points: []string{
				"Developed a portal to track consultancy and testing work at NITK and digitized the approval process.",
				"Faculty can create proposals and track their approval status.",
				"Generate reports like the projects under each center, department, and faculty and their total value.Developed interactive data visualizations.",
			},
		},
	},
	projects: []Project{
		{
			name: "Club recruitment",
			tech: "Ruby On Rails, HTML, Bootstrap",
			points: []string{
				"Developed a full-stack web application using Ruby On Rails to simplify the club recruitment process.",
				"Implemented a Role-based authorization system. Each user can be a convener, applicant, council member, or admin.",
				"Designed an admin dashboard to view statistics and export the allotment list in either PDF or Comma-separated values(CSV) format.",
				"Deployed through Heroku, and uses a PostgreSQL database.",
			},
		},
		{
			name: "Blog Aggregator",
			tech: "Python, Flask, beautifulsoup4",
			points: []string{
				"Developed a web application using Flask to view the latest blog posts from various technical and non-technical clubs at NITK.",
				"Implemented feature to fetch posts by the date published.",
				"Shows the latest posts from each club on a carousel",
			},
		},
		{
			name: "MySQL ORM",
			tech: "Python",
			points: []string{
				"Developed a simple Object Relational Mapper(ORM) for MySql.",
				"Using this package records can be created, fetched, updated, and destroyed.",
				"Wrote unit tests using the pytest package.",
				"Implemented continuous delivery using GitHub Actions to build and publish the package to PyPI repository upon a new release.",
			},
		},
	},
	skills: []Skill{
		{
			typ:    "Languages  ",
			skills: "Python, C/C++, Ruby, Golang, JavaScript, SQL",
		},
		{
			typ:    "Frameworks ",
			skills: "Django, Ruby on Rails, Flask, React",
		},
		{
			typ:    "Dev Tools  ",
			skills: "Git, GitHub Actions, GCP, Docker, Terraform, Kubernetes",
		},
	},
	links: []link{
		{
			name: "Personal Website",
			url:  "http://hrushi.dev",
		},
		{
			name: "GitHub",
			url:  "https://github.com/hrushikeshj",
		},
		{
			name: "LinkedIn",
			url:  "https://www.linkedin.com/in/hrushikeshj/",
		},
	},
}
