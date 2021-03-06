package main

import (
	"os"
	"service/generator"

	"github.com/gogearbox/gearbox"
)

func main() {
	// Get the value of an Environment Variable
	port := getEnvDefault("PORT", "3000")

	// Setup
	gr := gearbox.New()
	g := generator.New()

	// Handlers
	gr.Get("/", func(ctx gearbox.Context) {
		ctx.SendString(g.Generate())
	})

	// Start service
	gr.Start(":" + port)
}

// GetEnvDefault set the environmental variable by default
func getEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
