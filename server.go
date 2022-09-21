package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"os/exec"
	"strings"
	"log"
)

const ShellToUse = "bash"

func main() {
	log.SetPrefix("server: ")

	addresses, err := GetLocalAddresses()

	// If we can't get local addresses and masks
	// log error and exit program fatally
	if err != nil {
		log.Fatal(err)
	}
	
	
	// Initialize gin router with default settings
	r := gin.Default()
	
	api := r.Group("/api")

	api.GET("/addresses", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"addresses": addresses,
		})
	})

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Bind server to listen on port 3000
	r.Run(":3000")
}


// GetLocalAddresses runs a bash command to obtain
// a splice of local addresses/masks
func GetLocalAddresses() ([]string, error) {
	// Obtain local interfaces addresses/masks
	cmd := exec.Command(ShellToUse, "-c", "ip --brief address show | awk '{print $3}' ")
	stdout, err := cmd.Output()

	if err != nil {
		log.Print("Could not obtain ip addresses, please make sure"+
				" your machine has iputils and awk installed")
		return nil, err
	}
	

	// Addresses are delimited by newlines from stdout
	// remove whitespace and return addresses 
	return strings.Split(strings.TrimSpace(string(stdout)), "\n"), nil
}
