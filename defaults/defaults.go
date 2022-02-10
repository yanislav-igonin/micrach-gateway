package defaults

import (
	"fmt"
	"log"
	"strconv"
)

func GetBoolean(value string, defaultValue bool) bool {
	if value == "" {
		return defaultValue
	}
	return value == "true"
}

func GetInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Panicln(fmt.Sprintf("Could not parse %s to int", value))
	}
	return intValue
}

func GetString(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
