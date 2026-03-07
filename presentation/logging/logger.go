package logging

import "log"

func LogComponent(component, details string) {
	log.Printf("[%s] %s", component, details)
}

func LogError(component, operation string, err error) {
	log.Printf("[%s] operation=%s error=%v", component, operation, err)
}
