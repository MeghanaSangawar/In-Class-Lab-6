package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// UserDetails represents the details of a user.
type UserDetails struct {
	Name        string
	Email       string
	PhoneNumber string
}

// students is a list of students with their names and IDs.
var students = []struct {
	Name string
	ID   string
}{
	{"Shivaji Burgula", "500224628"},
	{"Jyothi Prasanna Kambam", "500233773"},
	{"Pranay Sai Chava", "500236056"},
	{"Yamini Kunapareddy", "500234808"},
	{"Sai Meghana Sangawar", "500235387"},
	{"Sushma Kambam", "500233832"},
}

// emailDomains is a list of email domains.
var emailDomains = []string{"example.com", "test.com", "demo.com"}

// generateUniqueEmail generates a unique email for a given name.
func generateUniqueEmail(name string) string {
	domain := emailDomains[rand.Intn(len(emailDomains))]
	return fmt.Sprintf("%s@%s", name, domain)
}

// generateUniquePhoneNumber generates a unique phone number.
func generateUniquePhoneNumber() string {
	return fmt.Sprintf("+1(555)%03d-%04d", rand.Intn(1000), rand.Intn(10000))
}

// generateUserDetails generates user details for a given name.
func generateUserDetails(name string) UserDetails {
	return UserDetails{
		Name:        name,
		Email:       generateUniqueEmail(name),
		PhoneNumber: generateUniquePhoneNumber(),
	}
}

// handler is the HTTP request handler.
func handler(w http.ResponseWriter, r *http.Request) {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Get hostname or unique identifier for the machine
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Unable to get hostname", http.StatusInternalServerError)
		return
	}

	// Use hostname to seed the random number generator
	rand.Seed(time.Now().UnixNano() + int64(len(hostname)))

	// Randomly select a student
	studentIndex := rand.Intn(len(students))
	selectedStudent := students[studentIndex]

	details := generateUserDetails(selectedStudent.Name)

	// Print the details in bold, centered, and use Calibri font as HTML
	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Student Information</title>
			<style>
				body {
					font-family: Calibri, sans-serif;
					text-align: center;
					margin: 0;
					padding: 0;
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
				}
				.container {
					text-align: center;
				}
				strong {
					font-weight: bold;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Student Information</h1>
				<p><strong>ID:</strong> %s</p>
				<p><strong>Name:</strong> %s</p>
				<p><strong>Email:</strong> %s</p>
				<p><strong>Phone Number:</strong> %s</p>
			</div>
		</body>
		</html>`,
		selectedStudent.ID,
		selectedStudent.Name,
		details.Email,
		details.PhoneNumber,
	)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
