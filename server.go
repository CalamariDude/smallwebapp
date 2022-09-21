package main

import ("github.com/gin-gonic/gin"
	"os/exec"
	"fmt"
)

const ShellToUse = "bash"

func main() {
	// Initialize gin router with default settings
	r := gin.Default()

	// Obtain local interfaces addresses/masks
	cmd := exec.Command(ShellToUse, "-c", "ip --brief address show | awk '{print $3}'")
	stdout, err := cmd.Output()

	// Print error to stdout if interfaces 
	// cannot be obtained and exit without
	// starting server
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Handle /addresses route by returning the 
	// stdout of iputils command
	r.GET("/addresses", func(c *gin.Context) {
		c.String(200, string(stdout))
	})

	// Bind server to listen on port 3000
	r.Run(":3000")
}
