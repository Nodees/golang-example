package main

import (
	postgres "core/connections"
)

func main() {
	postgres.InitDB()
}
