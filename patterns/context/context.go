package main

import (
	"context"
	"fmt"
)

// func main() {
// 	var wg sync.WaitGroup
// 	done := make(chan interface{})
// 	defer close(done)

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		if err := printGreeting(done); err != nil {
// 			fmt.Printf("%v", err)
// 			return
// 		}
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		if err := printFarewell(done); err != nil {
// 			fmt.Printf("%v", err)
// 			return
// 		}
// 	}()

// 	wg.Wait()
// }

// func printGreeting(done <-chan interface{}) error {
// 	greeting, err := genGreeting(done)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%s world!\n", greeting)
// 	return nil
// }

// func printFarewell(done <-chan interface{}) error {
// 	farewell, err := genFarewell(done)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%s world!\n", farewell)
// 	return nil
// }

// func genGreeting(done <-chan interface{}) (string, error) {
// 	switch locale, err := locale(done); {
// 	case err != nil:
// 		return "", err
// 	case locale == "EN/US":
// 		return "hello", nil
// 	}
// 	return "", fmt.Errorf("unsupported local")
// }

// func genFarewell(done <-chan interface{}) (string, error) {
// 	switch locale, err := locale(done); {
// 	case err != nil:
// 		return "", err
// 	case locale == "EN/US":
// 		return "goodbye", nil
// 	}
// 	return "", fmt.Errorf("unsupported local")
// }

// func locale(done <-chan interface{}) (string, error) {
// 	select {
// 	case <-done:
// 		return "", fmt.Errorf("canceled")
// 	case <-time.After(1 * time.Minute):
// 	}
// 	return "EN/US", nil
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		if err := printGreeting(ctx); err != nil {
// 			fmt.Printf("cannot print greeting: %v\n", err)
// 			cancel()
// 		}
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		if err := printFarewell(ctx); err != nil {
// 			fmt.Printf("cannot print farewell: %v\n", err)
// 		}
// 	}()

// 	wg.Wait()
// }

// func printGreeting(ctx context.Context) error {
// 	greeting, err := genGreeting(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%s world!\n", greeting)
// 	return nil
// }

// func printFarewell(ctx context.Context) error {
// 	farewell, err := genFarewell(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%s world!\n", farewell)
// 	return nil
// }

// func genGreeting(ctx context.Context) (string, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	defer cancel()

// 	switch locale, err := locale(ctx); {
// 	case err != nil:
// 		return "", err
// 	case locale == "EN/US":
// 		return "hello", nil
// 	}
// 	return "", fmt.Errorf("unsupported local")
// }

// func genFarewell(ctx context.Context) (string, error) {
// 	switch locale, err := locale(ctx); {
// 	case err != nil:
// 		return "", err
// 	case locale == "EN/US":
// 		return "goodbye", nil
// 	}
// 	return "", fmt.Errorf("unsupported local")
// }

// func locale(ctx context.Context) (string, error) {
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-time.After(1 * time.Minute):
// 	}
// 	return "EN/US", nil

// func main() {
// 	var wg sync.WaitGroup
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		if err := printGreeting(ctx); err != nil {
// 			fmt.Printf("cannot print greeting: %v\n", err)
// 			cancel()
// 		}
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		if err := printFarewell(ctx); err != nil {
// 			fmt.Printf("cannot print farewell: %v\n", err)
// 		}
// 	}()

// 	wg.Wait()
// }

// func printGreeting(ctx context.Context) error {
// 	greeting, err := genGreeting(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%s world!\n", greeting)
// 	return nil
// }

// func printFarewell(ctx context.Context) error {
// 	farewell, err := genFarewell(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%s world!\n", farewell)
// 	return nil
// }

// func genGreeting(ctx context.Context) (string, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	defer cancel()

// 	switch locale, err := locale(ctx); {
// 	case err != nil:
// 		return "", err
// 	case locale == "EN/US":
// 		return "hello", nil
// 	}
// 	return "", fmt.Errorf("unsupported local")
// }

// func genFarewell(ctx context.Context) (string, error) {
// 	switch locale, err := locale(ctx); {
// 	case err != nil:
// 		return "", err
// 	case locale == "EN/US":
// 		return "goodbye", nil
// 	}
// 	return "", fmt.Errorf("unsupported local")
// }

// func locale(ctx context.Context) (string, error) {
// 	if deadline, ok := ctx.Deadline(); ok {
// 		if deadline.Sub(time.Now().Add(1*time.Minute)) <= 0 {
// 			return "", context.DeadlineExceeded
// 		}
// 	}

// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-time.After(1 * time.Minute):
// 	}
// 	return "EN/US", nil
// }

// func main() {
// 	ProcessRequest("jane", "abc123")

// 	type foo int
// 	type bar int

// 	m := make(map[interface{}]int)
// 	m[foo(1)] = 1
// 	m[bar(1)] = 2

// 	fmt.Printf("%v", m)
// }

// func ProcessRequest(userID, authToken string) {
// 	ctx := context.WithValue(context.Background(), "userID", userID)
// 	ctx = context.WithValue(ctx, "authToken", authToken)
// 	HandleResponse(ctx)
// }

// func HandleResponse(ctx context.Context) {
// 	fmt.Printf(
// 		"handling response for %v (%v)",
// 		ctx.Value("userID"),
// 		ctx.Value("authToken"),
// 	)
// }

func main() {
	ProcessRequest("jane", "abc123")
}

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func UserID(c context.Context) string {
	return c.Value(ctxUserID).(string)
}

func AuthToken(c context.Context) string {
	return c.Value(ctxAuthToken).(string)
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response for %v (auth: %v)",
		UserID(ctx),
		AuthToken(ctx),
	)
}
