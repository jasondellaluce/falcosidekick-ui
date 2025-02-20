package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

const (
	extractNumber = "^[0-9]+"
	extractUnity  = "[a-z-A-Z]+$"
)

var regExtractNumber, regExtractUnity *regexp.Regexp

func init() {
	regExtractNumber, _ = regexp.Compile(extractNumber)
	regExtractUnity, _ = regexp.Compile(extractUnity)
}

func CheckErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func WriteLog(level, message string) {
	var prefix string
	switch level {
	case "error", "fatal":
		prefix = "[ERROR]:"
	case "info":
		prefix = "[INFO] :"
	}
	if level == "fatal" {
		log.Fatalf(fmt.Sprintf("%v %v\n", prefix, message))
	}
	log.Printf("%v %v", prefix, message)
}

func ConvertToSeconds(s string) int {
	n, err := strconv.Atoi(regExtractNumber.FindString(s))
	if err != nil {
		return 0
	}
	u := regExtractUnity.FindString(s)
	switch u {
	case "s", "second", "seconds":
		return n
	case "m", "min", "minute", "minutes":
		return n * 60
	case "h", "hour", "hours":
		return n * 60 * 60
	case "d", "day", "days":
		return n * 24 * 60 * 60
	case "w", "week", "weeks":
		return n * 7 * 24 * 60 * 60
	case "M", "month", "months":
		return n * 30 * 24 * 60 * 60
	case "y", "year", "years":
		return n * 365 * 24 * 60 * 60
	}
	return 0
}

func RemoveDuplicate(input []string) []string {
	allKeys := make(map[string]bool)
	singleKeys := []string{}
	for _, i := range input {
		if _, value := allKeys[i]; !value {
			allKeys[i] = true
			singleKeys = append(singleKeys, i)
		}
	}
	sort.Strings(singleKeys)
	return singleKeys
}

func GetStringFlagOrEnvParam(flagString string, envVar string, defaultValue string, usage string) *string {
	envvar, present := os.LookupEnv(envVar)
	if present {
		defaultValue = envvar
	}
	return flag.String(flagString, defaultValue, usage)
}

func GetIntFlagOrEnvParam(flagString string, envVar string, defaultValue int, usage string) *int {
	envvar, present := os.LookupEnv(envVar)
	if present {
		val, err := strconv.Atoi(envvar)
		if err == nil {
			defaultValue = val
		}
	}
	return flag.Int(flagString, defaultValue, usage)
}
