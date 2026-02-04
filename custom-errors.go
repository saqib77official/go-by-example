package main

import (
	"errors"
	"fmt"
	"net"
	"runtime"
	"time"
)

// 1. Basic custom error type
type AppError struct {
	Code    int
	Message string
	Details string
}

func (ae *AppError) Error() string {
	return fmt.Sprintf("App Error %d: %s - %s", ae.Code, ae.Message, ae.Details)
}

func (ae *AppError) Is(target error) bool {
	if targetApp, ok := target.(*AppError); ok {
		return ae.Code == targetApp.Code
	}
	return false
}

// 2. Error with timestamp
type TimestampedError struct {
	Timestamp time.Time
	Message   string
	Cause     error
}

func (te *TimestampedError) Error() string {
	return fmt.Sprintf("[%s] %s", te.Timestamp.Format(time.RFC3339), te.Message)
}

func (te *TimestampedError) Unwrap() error {
	return te.Cause
}

// 3. Validation error with field details
type ValidationError struct {
	Field   string
	Value   interface{}
	Rule    string
	Message string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s': %s", ve.Field, ve.Message)
}

func (ve *ValidationError) Details() string {
	return fmt.Sprintf("Field: %s, Value: %v, Rule: %s, Message: %s", 
		ve.Field, ve.Value, ve.Rule, ve.Message)
}

// 4. Network error with retry information
type NetworkError struct {
	Operation string
	URL       string
	StatusCode int
	Retryable bool
	RetryCount int
	Cause     error
}

func (ne *NetworkError) Error() string {
	return fmt.Sprintf("network error during %s to %s: status %d", 
		ne.Operation, ne.URL, ne.StatusCode)
}

func (ne *NetworkError) Unwrap() error {
	return ne.Cause
}

func (ne *NetworkError) ShouldRetry() bool {
	return ne.Retryable && ne.RetryCount < 3
}

// 5. Business logic error
type BusinessError struct {
	BusinessRule string
	Context      map[string]interface{}
	UserMessage  string
}

func (be *BusinessError) Error() string {
	return fmt.Sprintf("business rule violation: %s", be.BusinessRule)
}

func (be *BusinessError) GetUserMessage() string {
	if be.UserMessage != "" {
		return be.UserMessage
	}
	return "Operation cannot be completed due to business constraints"
}

// 6. Error with stack trace
type StackTraceError struct {
	Message    string
	Cause      error
	StackTrace []uintptr
}

func (ste *StackTraceError) Error() string {
	return fmt.Sprintf("%s: %v", ste.Message, ste.Cause)
}

func (ste *StackTraceError) Unwrap() error {
	return ste.Cause
}

func (ste *StackTraceError) StackTraceString() string {
	var result string
	for _, pc := range ste.StackTrace {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			file, line := fn.FileLine(pc)
			result += fmt.Sprintf("  %s:%d %s\n", file, line, fn.Name())
		}
	}
	return result
}

// 7. Error with severity levels
type Severity int

const (
	SeverityInfo Severity = iota
	SeverityWarning
	SeverityError
	SeverityCritical
)

type SeverityError struct {
	Severity Severity
	Message  string
	Cause    error
	Context  map[string]interface{}
}

func (se *SeverityError) Error() string {
	return fmt.Sprintf("[%s] %s", se.SeverityString(), se.Message)
}

func (se *SeverityError) Unwrap() error {
	return se.Cause
}

func (se *SeverityError) SeverityString() string {
	switch se.Severity {
	case SeverityInfo:
		return "INFO"
	case SeverityWarning:
		return "WARNING"
	case SeverityError:
		return "ERROR"
	case SeverityCritical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

// 8. Error with user-friendly messages
type UserError struct {
	TechnicalMessage string
	UserMessage     string
	ErrorCode       string
	Cause           error
}

func (ue *UserError) Error() string {
	return ue.TechnicalMessage
}

func (ue *UserError) GetUserMessage() string {
	if ue.UserMessage != "" {
		return ue.UserMessage
	}
	return "An unexpected error occurred. Please try again."
}

func (ue *UserError) GetErrorCode() string {
	return ue.ErrorCode
}

// 9. Error with recovery suggestions
type RecoverableError struct {
	Message      string
	Suggestions []string
	Cause        error
}

func (re *RecoverableError) Error() string {
	return re.Message
}

func (re *RecoverableError) Unwrap() error {
	return re.Cause
}

func (re *RecoverableError) GetSuggestions() []string {
	return re.Suggestions
}

// 10. Error with metadata
type MetadataError struct {
	BaseError error
	Metadata  map[string]interface{}
}

func (me *MetadataError) Error() string {
	return me.BaseError.Error()
}

func (me *MetadataError) Unwrap() error {
	return me.BaseError
}

func (me *MetadataError) GetMetadata(key string) (interface{}, bool) {
	val, exists := me.Metadata[key]
	return val, exists
}

// 11. Error with retry policy
type RetryPolicy struct {
	MaxRetries    int
	InitialDelay  time.Duration
	MaxDelay      time.Duration
	BackoffFactor float64
}

type RetryableError struct {
	Operation   string
	RetryPolicy RetryPolicy
	Attempt     int
	LastError   error
}

func (re *RetryableError) Error() string {
	return fmt.Sprintf("operation '%s' failed on attempt %d: %v", 
		re.Operation, re.Attempt, re.LastError)
}

func (re *RetryableError) ShouldRetry() bool {
	return re.Attempt < re.RetryPolicy.MaxRetries
}

func (re *RetryableError) NextDelay() time.Duration {
	delay := time.Duration(float64(re.RetryPolicy.InitialDelay) * 
		pow(re.RetryPolicy.BackoffFactor, float64(re.Attempt)))
	if delay > re.RetryPolicy.MaxDelay {
		return re.RetryPolicy.MaxDelay
	}
	return delay
}

func pow(base, exp float64) float64 {
	result := 1.0
	for i := 0; i < int(exp); i++ {
		result *= base
	}
	return result
}

// 12. Error with context information
type ContextError struct {
	Message string
	Context map[string]interface{}
	Cause   error
}

func (ce *ContextError) Error() string {
	return ce.Message
}

func (ce *ContextError) Unwrap() error {
	return ce.Cause
}

func (ce *ContextError) AddContext(key string, value interface{}) {
	if ce.Context == nil {
		ce.Context = make(map[string]interface{})
	}
	ce.Context[key] = value
}

func (ce *ContextError) GetContext(key string) (interface{}, bool) {
	if ce.Context == nil {
		return nil, false
	}
	val, exists := ce.Context[key]
	return val, exists
}

// 13. Error aggregator
type ErrorAggregator struct {
	errors []error
}

func (ea *ErrorAggregator) Add(err error) {
	if err != nil {
		ea.errors = append(ea.errors, err)
	}
}

func (ea *ErrorAggregator) HasErrors() bool {
	return len(ea.errors) > 0
}

func (ea *ErrorAggregator) Error() string {
	if len(ea.errors) == 0 {
		return "no errors"
	}
	
	result := fmt.Sprintf("%d error(s) occurred:\n", len(ea.errors))
	for i, err := range ea.errors {
		result += fmt.Sprintf("  %d: %v\n", i+1, err)
	}
	return result
}

func (ea *ErrorAggregator) GetErrors() []error {
	return ea.errors
}

// Example functions that create custom errors
func validateUserInput(name, email string) error {
	var aggregator ErrorAggregator
	
	if name == "" {
		aggregator.Add(&ValidationError{
			Field:   "name",
			Value:   name,
			Rule:    "required",
			Message: "name cannot be empty",
		})
	}
	
	if len(name) < 2 {
		aggregator.Add(&ValidationError{
			Field:   "name",
			Value:   name,
			Rule:    "min_length",
			Message: "name must be at least 2 characters",
		})
	}
	
	if email == "" {
		aggregator.Add(&ValidationError{
			Field:   "email",
			Value:   email,
			Rule:    "required",
			Message: "email cannot be empty",
		})
	}
	
	if !contains(email, "@") {
		aggregator.Add(&ValidationError{
			Field:   "email",
			Value:   email,
			Rule:    "format",
			Message: "email must contain @ symbol",
		})
	}
	
	if aggregator.HasErrors() {
		return &aggregator
	}
	
	return nil
}

func processPayment(amount float64, cardNumber string) error {
	if amount <= 0 {
		return &BusinessError{
			BusinessRule: "positive_amount",
			Context: map[string]interface{}{
				"amount": amount,
			},
			UserMessage: "Payment amount must be positive",
		}
	}
	
	if len(cardNumber) != 16 {
		return &UserError{
			TechnicalMessage: "invalid card number length",
			UserMessage:     "Please enter a valid 16-digit card number",
			ErrorCode:       "INVALID_CARD",
		}
	}
	
	// Simulate network error
	return &NetworkError{
		Operation:   "payment_charge",
		URL:         "https://api.payment.com/charge",
		StatusCode:  503,
		Retryable:   true,
		RetryCount:  0,
		Cause:       errors.New("service unavailable"),
	}
}

func connectToDatabase() error {
	// Simulate connection error
	baseErr := errors.New("connection timeout")
	
	return &ContextError{
		Message: "failed to connect to database",
		Context: map[string]interface{}{
			"host":     "localhost",
			"port":     5432,
			"database": "myapp",
			"timeout":  "30s",
		},
		Cause: baseErr,
	}
}

func main() {
	fmt.Println("=== Custom Errors Examples ===")

	// 1. Basic custom error
	fmt.Println("\n1. Basic custom error:")
	err := &AppError{
		Code:    1001,
		Message: "User not found",
		Details: "User ID 123 does not exist in the system",
	}
	fmt.Printf("Error: %v\n", err)

	// 2. Timestamped error
	fmt.Println("\n2. Timestamped error:")
	baseErr := errors.New("file not found")
	timestampedErr := &TimestampedError{
		Timestamp: time.Now(),
		Message:   "Failed to load configuration",
		Cause:     baseErr,
	}
	fmt.Printf("Error: %v\n", timestampedErr)
	fmt.Printf("Unwrapped: %v\n", errors.Unwrap(timestampedErr))

	// 3. Validation error
	fmt.Println("\n3. Validation error:")
	validationErr := &ValidationError{
		Field:   "email",
		Value:   "invalid-email",
		Rule:    "format",
		Message: "email must be a valid email address",
	}
	fmt.Printf("Error: %v\n", validationErr)
	fmt.Printf("Details: %s\n", validationErr.Details())

	// 4. Network error
	fmt.Println("\n4. Network error:")
	netErr := &NetworkError{
		Operation:  "GET",
		URL:        "https://api.example.com/users",
		StatusCode: 404,
		Retryable:  false,
		RetryCount: 0,
		Cause:      errors.New("resource not found"),
	}
	fmt.Printf("Error: %v\n", netErr)
	fmt.Printf("Should retry: %t\n", netErr.ShouldRetry())

	// 5. Business error
	fmt.Println("\n5. Business error:")
	bizErr := &BusinessError{
		BusinessRule: "insufficient_balance",
		Context: map[string]interface{}{
			"required": 100.0,
			"available": 50.0,
		},
		UserMessage: "Insufficient balance for this transaction",
	}
	fmt.Printf("Error: %v\n", bizErr)
	fmt.Printf("User message: %s\n", bizErr.GetUserMessage())

	// 6. Error with stack trace
	fmt.Println("\n6. Error with stack trace:")
	stackErr := &StackTraceError{
		Message: "critical system error",
		Cause:   errors.New("memory allocation failed"),
	}
	stackErr.StackTrace = make([]uintptr, 10)
	runtime.Callers(1, stackErr.StackTrace)
	
	fmt.Printf("Error: %v\n", stackErr)
	fmt.Printf("Stack trace:\n%s", stackErr.StackTraceString())

	// 7. Severity error
	fmt.Println("\n7. Severity error:")
	sevErr := &SeverityError{
		Severity: SeverityError,
		Message:  "Database connection lost",
		Cause:    errors.New("connection timeout"),
		Context: map[string]interface{}{
			"database": "primary",
			"host":     "db.example.com",
		},
	}
	fmt.Printf("Error: %v\n", sevErr)
	fmt.Printf("Severity: %s\n", sevErr.SeverityString())

	// 8. User error
	fmt.Println("\n8. User error:")
	userErr := &UserError{
		TechnicalMessage: "password hash verification failed",
		UserMessage:     "Invalid username or password",
		ErrorCode:       "AUTH_FAILED",
		Cause:           errors.New("hash mismatch"),
	}
	fmt.Printf("Technical: %v\n", userErr.Error())
	fmt.Printf("User: %s\n", userErr.GetUserMessage())
	fmt.Printf("Code: %s\n", userErr.GetErrorCode())

	// 9. Recoverable error
	fmt.Println("\n9. Recoverable error:")
	recErr := &RecoverableError{
		Message: "File upload failed",
		Suggestions: []string{
			"Check your internet connection",
			"Verify the file format is supported",
			"Try uploading a smaller file",
		},
		Cause: errors.New("network timeout"),
	}
	fmt.Printf("Error: %v\n", recErr)
	fmt.Printf("Suggestions:\n")
	for i, suggestion := range recErr.GetSuggestions() {
		fmt.Printf("  %d. %s\n", i+1, suggestion)
	}

	// 10. Metadata error
	fmt.Println("\n10. Metadata error:")
	metaErr := &MetadataError{
		BaseError: errors.New("operation failed"),
		Metadata: map[string]interface{}{
			"user_id":    12345,
			"operation":  "create_order",
			"timestamp":  time.Now(),
			"request_id": "req-123",
		},
	}
	fmt.Printf("Error: %v\n", metaErr)
	if userID, exists := metaErr.GetMetadata("user_id"); exists {
		fmt.Printf("User ID: %v\n", userID)
	}

	// 11. Retryable error
	fmt.Println("\n11. Retryable error:")
	retryErr := &RetryableError{
		Operation: "send_email",
		RetryPolicy: RetryPolicy{
			MaxRetries:    3,
			InitialDelay:  time.Second,
			MaxDelay:      30 * time.Second,
			BackoffFactor: 2.0,
		},
		Attempt:   1,
		LastError: errors.New("SMTP server busy"),
	}
	fmt.Printf("Error: %v\n", retryErr)
	fmt.Printf("Should retry: %t\n", retryErr.ShouldRetry())
	fmt.Printf("Next delay: %v\n", retryErr.NextDelay())

	// 12. Context error
	fmt.Println("\n12. Context error:")
	ctxErr := &ContextError{
		Message: "API request failed",
		Cause:   errors.New("rate limit exceeded"),
	}
	ctxErr.AddContext("endpoint", "/api/users")
	ctxErr.AddContext("method", "GET")
	ctxErr.AddContext("user_id", 12345)
	
	fmt.Printf("Error: %v\n", ctxErr)
	if endpoint, exists := ctxErr.GetContext("endpoint"); exists {
		fmt.Printf("Endpoint: %v\n", endpoint)
	}

	// 13. Error aggregator
	fmt.Println("\n13. Error aggregator:")
	aggregator := &ErrorAggregator{}
	aggregator.Add(errors.New("first error"))
	aggregator.Add(errors.New("second error"))
	aggregator.Add(errors.New("third error"))
	
	if aggregator.HasErrors() {
		fmt.Printf("Aggregated error:\n%s", aggregator.Error())
	}

	// 14. Real-world examples
	fmt.Println("\n14. Real-world validation:")
	err = validateUserInput("", "invalid-email")
	if err != nil {
		fmt.Printf("Validation failed:\n%s", err.Error())
	}

	fmt.Println("\n15. Real-world payment processing:")
	err = processPayment(-100, "1234")
	if err != nil {
		var bizErr *BusinessError
		var userErr *UserError
		var netErr *NetworkError
		
		switch {
		case errors.As(err, &bizErr):
			fmt.Printf("Business error: %s\n", bizErr.GetUserMessage())
		case errors.As(err, &userErr):
			fmt.Printf("User error: %s (Code: %s)\n", 
				userErr.GetUserMessage(), userErr.GetErrorCode())
		case errors.As(err, &netErr):
			fmt.Printf("Network error: %v\n", netErr)
			if netErr.ShouldRetry() {
				fmt.Printf("This error is retryable\n")
			}
		default:
			fmt.Printf("Other error: %v\n", err)
		}
	}

	fmt.Println("\n16. Real-world database connection:")
	err = connectToDatabase()
	if err != nil {
		var ctxErr *ContextError
		if errors.As(err, &ctxErr) {
			fmt.Printf("Context error: %v\n", ctxErr.Error())
			if host, exists := ctxErr.GetContext("host"); exists {
				fmt.Printf("Failed host: %v\n", host)
			}
		}
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && 
			(s[:len(substr)] == substr || 
			 s[len(s)-len(substr):] == substr ||
			 findSubstring(s, substr))))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
