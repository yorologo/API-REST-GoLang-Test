package main

import (
	"os"

	"github.com/gogearbox/gearbox"
)

func main() {
	// Get the value of an Environment Variable
	port := getEnvDefault("PORT", "3001")

	// Setup gearbox
	g := gearbox.New()

	// Define your handlers
	g.Get("/", func(ctx gearbox.Context) {
		ctx.SendString("Hello World!")
	})

	// Start service
	g.Start(":" + port)
}

// GetEnvDefault set the environmental variable by default
func getEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
