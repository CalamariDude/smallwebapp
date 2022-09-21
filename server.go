package main

import ("github.com/gin-gonic/gin"
	"os/exec"
	"fmt"
)

const ShellToUse = "bash"

func main() {
	r := gin.Default()
	cmd := exec.Command(ShellToUse, "-c", "ip --brief address show | awk '{print $3}'")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	//print output
	fmt.Println(string(stdout))
	
	
	r.GET("/addresses", func(c *gin.Context) {
		c.String(200, string(stdout))
	})

	r.Run(":3000")
}
