package main

import (
	"os"
	"service/generator"

	"github.com/gogearbox/gearbox"
)

// Response struct for json response
type Response struct {
	Value string
}

func main() {
	// Get the value of an Environment Variable
	port := getEnvDefault("PORT", "3000")

	// Setup
	gb := gearbox.New()
	g := generator.New()

	// Handlers
	gb.Get("/", func(ctx gearbox.Context) {
		ctx.SendString(g.Generate())
	})
	gb.Get("/json", func(ctx gearbox.Context) {
		var r Response
		r.Value = g.Generate()

		ctx.Set("Access-Control-Allow-Origin", "*")
		ctx.SendJSON(&r)
	})

	// Start service
	gb.Start(":" + port)
}

// GetEnvDefault set the environmental variable by default
func getEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
