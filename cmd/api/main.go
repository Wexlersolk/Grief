package main

const version = "1.1.0"

func main() { 
	cfg:= config (
		addr: env.GetString("ADDR", "8080"),
	)
}

