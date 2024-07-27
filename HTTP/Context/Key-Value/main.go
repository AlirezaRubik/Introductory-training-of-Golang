package main

import (
	"context"
	"fmt"
)

func main() {
	// Create a background context
	ctx := context.Background()

	// Add an API key value to the context
	ctx = context.WithValue(ctx, "ApiKey", "Alireza")

	// Create a new context with a username value
	ctx2 := context.WithValue(ctx, "username", "Im_Programmer")

	// Call function f1 with the first context
	f1(ctx)

	// Call function f2 with the second context
	f2(ctx2)
}

// f1 prints the API key from the context and calls other functions
func f1(ctx context.Context) {
	fmt.Println("The Function Name is: f1 and The Value in Context Is:", ctx.Value("ApiKey"))
	f11(ctx)
	f12(ctx)
	f13(ctx)
}

// f11 prints the API key from the context
func f11(ctx context.Context) {
	fmt.Println("The Function Name is: f11 and The Value in Context Is:", ctx.Value("ApiKey"))
}

// f12 prints the API key from the context and calls f121
func f12(ctx context.Context) {
	fmt.Println("The Function Name is: f12 and The Value in Context Is:", ctx.Value("ApiKey"))
	f121(ctx)
}

// f13 prints the API key from the context and calls f131
func f13(ctx context.Context) {
	fmt.Println("The Function Name is: f13 and The Value in Context Is:", ctx.Value("ApiKey"))
	f131(ctx)
}

// f121 prints the API key from the context
func f121(ctx context.Context) {
	fmt.Println("The Function Name is: f121 and The Value in Context Is:", ctx.Value("ApiKey"))
}

// f131 prints the API key from the context
func f131(ctx context.Context) {
	fmt.Println("The Function Name is: f131 and The Value in Context Is:", ctx.Value("ApiKey"))
}

// f2 prints the username from the context and calls other functions
func f2(ctx context.Context) {
	fmt.Println("The Function Name is: f2 and The Value in Context Is:", ctx.Value("username"))
	f21(ctx)
	f22(ctx)
	f23(ctx)
}

// f21 prints the username from the context
func f21(ctx context.Context) {
	fmt.Println("The Function Name is: f21 and The Value in Context Is:", ctx.Value("username"))
}

// f22 prints the username from the context and calls f221
func f22(ctx context.Context) {
	fmt.Println("The Function Name is: f22 and The Value in Context Is:", ctx.Value("username"))
	f221(ctx)
}

// f221 prints the username from the context
func f221(ctx context.Context) {
	fmt.Println("The Function Name is: f221 and The Value in Context Is:", ctx.Value("username"))
}

// f23 prints the username from the context and calls f231
func f23(ctx context.Context) {
	fmt.Println("The Function Name is: f23 and The Value in Context Is:", ctx.Value("username"))
	f231(ctx)
}

// f231 prints the username from the context
func f231(ctx context.Context) {
	fmt.Println("The Function Name is: f231 and The Value in Context Is:", ctx.Value("username"))
}
