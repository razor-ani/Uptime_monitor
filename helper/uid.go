package helper

import uuid "github.com/satori/go.uuid"

//GetUID ..
func GetUID() string {
	return (uuid.NewV4()).String()
}
