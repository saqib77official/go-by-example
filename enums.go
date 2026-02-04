package main

import "fmt"

func main() {
	fmt.Println("=== Enums Examples ===")

	// 1. Basic enum using iota
	fmt.Println("\n1. Basic enum with iota:")
	type Day int
	
	const (
		Sunday Day = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	
	fmt.Printf("Sunday: %d\n", Sunday)
	fmt.Printf("Monday: %d\n", Monday)
	fmt.Printf("Tuesday: %d\n", Tuesday)
	
	// Add String method for better printing
	func (d Day) String() string {
		switch d {
		case Sunday:
			return "Sunday"
		case Monday:
			return "Monday"
		case Tuesday:
			return "Tuesday"
		case Wednesday:
			return "Wednesday"
		case Thursday:
			return "Thursday"
		case Friday:
			return "Friday"
		case Saturday:
			return "Saturday"
		default:
			return "Unknown"
		}
	}
	
	var today Day = Wednesday
	fmt.Printf("Today is %s\n", today.String())

	// 2. Enum with custom values
	fmt.Println("\n2. Enum with custom values:")
	type Status int
	
	const (
		StatusActive Status = iota + 1
		StatusInactive
		StatusPending
		StatusSuspended
	)
	
	func (s Status) String() string {
		switch s {
		case StatusActive:
			return "Active"
		case StatusInactive:
			return "Inactive"
		case StatusPending:
			return "Pending"
		case StatusSuspended:
			return "Suspended"
		default:
			return "Unknown"
		}
	}
	
	func (s Status) IsActive() bool {
		return s == StatusActive
	}
	
	userStatus := StatusActive
	fmt.Printf("User status: %s\n", userStatus.String())
	fmt.Printf("Is active: %t\n", userStatus.IsActive())

	// 3. Enum with string values
	fmt.Println("\n3. Enum with string values:")
	type Color string
	
	const (
		ColorRed   Color = "red"
		ColorGreen Color = "green"
		ColorBlue  Color = "blue"
		ColorYellow Color = "yellow"
	)
	
	func (c Color) RGB() (r, g, b int) {
		switch c {
		case ColorRed:
			return 255, 0, 0
		case ColorGreen:
			return 0, 255, 0
		case ColorBlue:
			return 0, 0, 255
		case ColorYellow:
			return 255, 255, 0
		default:
			return 0, 0, 0
		}
	}
	
	favoriteColor := ColorBlue
	fmt.Printf("Favorite color: %s\n", favoriteColor)
	r, g, b := favoriteColor.RGB()
	fmt.Printf("RGB values: %d, %d, %d\n", r, g, b)

	// 4. Enum with bitmask values
	fmt.Println("\n4. Enum with bitmask values:")
	type Permission int
	
	const (
		PermissionRead   Permission = 1 << iota // 1
		PermissionWrite                           // 2
		PermissionExecute                         // 4
		PermissionDelete                          // 8
		PermissionAdmin    Permission = PermissionRead | PermissionWrite | PermissionDelete
	)
	
	func (p Permission) String() string {
		var permissions []string
		
		if p&PermissionRead != 0 {
			permissions = append(permissions, "Read")
		}
		if p&PermissionWrite != 0 {
			permissions = append(permissions, "Write")
		}
		if p&PermissionExecute != 0 {
			permissions = append(permissions, "Execute")
		}
		if p&PermissionDelete != 0 {
			permissions = append(permissions, "Delete")
		}
		
		if len(permissions) == 0 {
			return "None"
		}
		
		result := permissions[0]
		for i := 1; i < len(permissions); i++ {
			result += "|" + permissions[i]
		}
		return result
	}
	
	func (p Permission) Has(permission Permission) bool {
		return p&permission != 0
	}
	
	userPermissions := PermissionRead | PermissionWrite
	fmt.Printf("User permissions: %s\n", userPermissions.String())
	fmt.Printf("Can read: %t\n", userPermissions.Has(PermissionRead))
	fmt.Printf("Can execute: %t\n", userPermissions.Has(PermissionExecute))
	
	// Add permission
	userPermissions |= PermissionExecute
	fmt.Printf("After adding execute: %s\n", userPermissions.String())
	
	// Remove permission
	userPermissions &^= PermissionWrite
	fmt.Printf("After removing write: %s\n", userPermissions.String())

	// 5. Enum with validation
	fmt.Println("\n5. Enum with validation:")
	type Priority int
	
	const (
		PriorityLow Priority = iota
		PriorityMedium
		PriorityHigh
		PriorityCritical
	)
	
	var validPriorities = map[Priority]bool{
		PriorityLow:      true,
		PriorityMedium:   true,
		PriorityHigh:     true,
		PriorityCritical: true,
	}
	
	func (p Priority) IsValid() bool {
		return validPriorities[p]
	}
	
	func (p Priority) String() string {
		switch p {
		case PriorityLow:
			return "Low"
		case PriorityMedium:
			return "Medium"
		case PriorityHigh:
			return "High"
		case PriorityCritical:
			return "Critical"
		default:
			return "Invalid"
		}
	}
	
	func ParsePriority(s string) (Priority, error) {
		switch s {
		case "Low":
			return PriorityLow, nil
		case "Medium":
			return PriorityMedium, nil
		case "High":
			return PriorityHigh, nil
		case "Critical":
			return PriorityCritical, nil
		default:
			return PriorityLow, fmt.Errorf("invalid priority: %s", s)
		}
	}
	
	priority := PriorityHigh
	fmt.Printf("Priority: %s\n", priority.String())
	fmt.Printf("Is valid: %t\n", priority.IsValid())
	
	invalidPriority := Priority(99)
	fmt.Printf("Invalid priority: %s\n", invalidPriority.String())
	fmt.Printf("Is valid: %t\n", invalidPriority.IsValid())
	
	if parsed, err := ParsePriority("Medium"); err == nil {
		fmt.Printf("Parsed priority: %s\n", parsed.String())
	}

	// 6. Enum with associated data
	fmt.Println("\n6. Enum with associated data:")
	type LogLevel int
	
	const (
		LogDebug LogLevel = iota
		LogInfo
		LogWarning
		LogError
		LogFatal
	)
	
	var logLevelData = map[LogLevel]struct {
		name  string
		color string
		level int
	}{
		LogDebug:   {"DEBUG", "gray", 0},
		LogInfo:    {"INFO", "blue", 1},
		LogWarning: {"WARNING", "yellow", 2},
		LogError:   {"ERROR", "red", 3},
		LogFatal:   {"FATAL", "magenta", 4},
	}
	
	func (l LogLevel) Name() string {
		return logLevelData[l].name
	}
	
	func (l LogLevel) Color() string {
		return logLevelData[l].color
	}
	
	func (l LogLevel) Level() int {
		return logLevelData[l].level
	}
	
	func (l LogLevel) IsHigherThan(other LogLevel) bool {
		return l.Level() > other.Level()
	}
	
	currentLevel := LogError
	fmt.Printf("Log level: %s\n", currentLevel.Name())
	fmt.Printf("Color: %s\n", currentLevel.Color())
	fmt.Printf("Level: %d\n", currentLevel.Level())
	fmt.Printf("Is higher than Warning: %t\n", currentLevel.IsHigherThan(LogWarning))

	// 7. Enum with iteration
	fmt.Println("\n7. Enum iteration:")
	type Month int
	
	const (
		January Month = iota + 1
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)
	
	var months = []Month{
		January, February, March, April, May, June,
		July, August, September, October, November, December,
	}
	
	var monthNames = map[Month]string{
		January:   "January",
		February:  "February",
		March:     "March",
		April:     "April",
		May:       "May",
		June:      "June",
		July:      "July",
		August:    "August",
		September: "September",
		October:   "October",
		November:  "November",
		December:  "December",
	}
	
	fmt.Println("All months:")
	for _, month := range months {
		fmt.Printf("%d: %s\n", month, monthNames[month])
	}

	// 8. Enum with JSON marshaling
	fmt.Println("\n8. Enum with JSON marshaling:")
	type UserRole string
	
	const (
		RoleGuest     UserRole = "guest"
		RoleUser      UserRole = "user"
		RoleModerator UserRole = "moderator"
		RoleAdmin     UserRole = "admin"
	)
	
	func (r UserRole) IsValid() bool {
		switch r {
		case RoleGuest, RoleUser, RoleModerator, RoleAdmin:
			return true
		default:
			return false
		}
	}
	
	func (r UserRole) Permissions() []string {
		switch r {
		case RoleGuest:
			return []string{"read"}
		case RoleUser:
			return []string{"read", "write"}
		case RoleModerator:
			return []string{"read", "write", "moderate"}
		case RoleAdmin:
			return []string{"read", "write", "moderate", "admin"}
		default:
			return []string{}
		}
	}
	
	userRole := RoleModerator
	fmt.Printf("User role: %s\n", userRole)
	fmt.Printf("Is valid: %t\n", userRole.IsValid())
	fmt.Printf("Permissions: %v\n", userRole.Permissions())

	// 9. Enum with state machine
	fmt.Println("\n9. Enum with state machine:")
	type OrderStatus int
	
	const (
		StatusPending OrderStatus = iota
		StatusConfirmed
		StatusProcessing
		StatusShipped
		StatusDelivered
		StatusCancelled
	)
	
	var validTransitions = map[OrderStatus][]OrderStatus{
		StatusPending:    {StatusConfirmed, StatusCancelled},
		StatusConfirmed:  {StatusProcessing, StatusCancelled},
		StatusProcessing: {StatusShipped, StatusCancelled},
		StatusShipped:    {StatusDelivered},
		StatusDelivered:  {}, // Terminal state
		StatusCancelled:  {}, // Terminal state
	}
	
	func (os OrderStatus) String() string {
		switch os {
		case StatusPending:
			return "Pending"
		case StatusConfirmed:
			return "Confirmed"
		case StatusProcessing:
			return "Processing"
		case StatusShipped:
			return "Shipped"
		case StatusDelivered:
			return "Delivered"
		case StatusCancelled:
			return "Cancelled"
		default:
			return "Unknown"
		}
	}
	
	func (os OrderStatus) CanTransitionTo(newStatus OrderStatus) bool {
		allowedTransitions, exists := validTransitions[os]
		if !exists {
			return false
		}
		
		for _, allowed := range allowedTransitions {
			if allowed == newStatus {
				return true
			}
		}
		return false
	}
	
	currentStatus := StatusProcessing
	fmt.Printf("Current status: %s\n", currentStatus.String())
	
	nextStatus := StatusShipped
	fmt.Printf("Can transition to %s: %t\n", nextStatus.String(), currentStatus.CanTransitionTo(nextStatus))
	
	invalidStatus := StatusPending
	fmt.Printf("Can transition to %s: %t\n", invalidStatus.String(), currentStatus.CanTransitionTo(invalidStatus))

	// 10. Enum with database mapping
	fmt.Println("\n10. Enum with database mapping:")
	type DatabaseType string
	
	const (
		DatabaseMySQL    DatabaseType = "mysql"
		DatabasePostgres DatabaseType = "postgres"
		DatabaseSQLite   DatabaseType = "sqlite"
		DatabaseMongoDB  DatabaseType = "mongodb"
	)
	
	var databasePorts = map[DatabaseType]int{
		DatabaseMySQL:    3306,
		DatabasePostgres: 5432,
		DatabaseSQLite:   0, // File-based
		DatabaseMongoDB:  27017,
	}
	
	var databaseDrivers = map[DatabaseType]string{
		DatabaseMySQL:    "github.com/go-sql-driver/mysql",
		DatabasePostgres: "github.com/lib/pq",
		DatabaseSQLite:   "github.com/mattn/go-sqlite3",
		DatabaseMongoDB:  "go.mongodb.org/mongo-driver",
	}
	
	func (dt DatabaseType) Port() int {
		return databasePorts[dt]
	}
	
	func (dt DatabaseType) Driver() string {
		return databaseDrivers[dt]
	}
	
	func (dt DatabaseType) IsRelational() bool {
		return dt == DatabaseMySQL || dt == DatabasePostgres || dt == DatabaseSQLite
	}
	
	dbType := DatabasePostgres
	fmt.Printf("Database type: %s\n", dbType)
	fmt.Printf("Default port: %d\n", dbType.Port())
	fmt.Printf("Driver: %s\n", dbType.Driver())
	fmt.Printf("Is relational: %t\n", dbType.IsRelational())
}
