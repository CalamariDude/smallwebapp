# SmallWebApp

SmallWebApp is a small web application meant to display information about the local network interfaces of a linux machine. The program initializes a webapp that serves up a line-separated list of ip addresses and masks about your local interfaces to http://localhost:3000

## Requirements

Ubuntu 22.04 LTS or a linux machine with `bash`, `iputils`, and `awk` installed

Go 1.19

## Usage

Clone the repository and cd into it

`git clone https://github.com/CalamariDude/smallwebapp.git && cd smallwebapp `

Install Go dependencies

`go install`

Run the program

`go run .`

Navigate to http://localhost:3000

## Tests

To run test(s) do

`go test`
