package auth

// keep this file in auth package and import it in main.go to use the Login function
// make it capitalized to make it public and accessible from other packages
func Login(username, password string) string {
	return "Login successful for user: " + username + " with password: " + password
}