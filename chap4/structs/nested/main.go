// Nested structs
type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Person struct { // Anonymous nested struct
		FirstName string
		LastName  string
	}
	Address    Address // Named nested struct
	Department string
	Projects   []string
}

type Company struct {
	Name     string
	Location Address // Reusing Address struct
	CEO      Person  // Reusing Person struct
	Founded  time.Time
}

func main() {
	// Initialize nested structs
	emp := Employee{
		Person: struct {
			FirstName string
			LastName  string
		}{
			FirstName: "John",
			LastName:  "Doe",
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "San Francisco",
			Country: "USA",
		},
		Department: "Engineering",
		Projects:   []string{"Project A", "Project B"},
	}

	// Anonymous struct literal (defined and initialized inline)
	config := struct {
		Debug   bool
		Cache   map[string]interface{}
		Timeout time.Duration
	}{
		Debug:   true,
		Cache:   make(map[string]interface{}),
		Timeout: 30 * time.Second,
	}

	// Nested anonymous structs
	response := struct {
		Status  int
		Data    interface{}
		Headers struct {
			ContentType string
			UserAgent   string
		}
	}{
		Status: 200,
		Data:   "Success",
		Headers: struct {
			ContentType string
			UserAgent   string
		}{
			ContentType: "application/json",
			UserAgent:   "Go-Client",
		},
	}

	// Struct with embedded types (composition)
	type Logger struct {
		Level string
	}

	type Service struct {
		Logger  // Embedded struct
		name    string
		enabled bool
	}

	svc := Service{
		Logger: Logger{
			Level: "DEBUG",
		},
		name:    "MyService",
		enabled: true,
	}
}