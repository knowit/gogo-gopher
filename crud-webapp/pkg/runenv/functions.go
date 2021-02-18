package runenv

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Printf("Could not find a value for key '%v'. Returning defaultValue '%v'\n", key, defaultValue)
		return defaultValue
	}

	return value
}

func GetEnvAsInt(name string, defaultValue int) int {
	valueStr := GetEnv(name, "")
	value, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Printf("Could not convert the value '%v' from key '%v' to int. Returning defaultValue '%v'\n",
			valueStr,
			name,
			defaultValue,
		)
		return defaultValue
	}


	return value
}

func GetEnvAsBool(name string, defaultValue bool) bool {
	valueString := GetEnv(name, "")

	if valueString == "" {
		log.Printf("Could not find a value for key '%v'. Returning defaultValue '%v'\n", name, defaultValue)
		return defaultValue
	}

	value, err := strconv.ParseBool(valueString)

	if err != nil {
		log.Printf("Could not convert the value '%v' from key '%v' to bool. Returning defaultValue '%v'\n", value, )
		return defaultValue
	}

	return value
}

func GetEnvAsSlice(name string, defaultValues []string, sep string) []string {
	valueString := GetEnv(name, "")

	if valueString == "" {
		log.Printf("Could not find the value for key '%v'. Returning defaultValues '%v'\n", name, defaultValues)
		return defaultValues
	}

	values := strings.Split(valueString, sep)

	return values
}
