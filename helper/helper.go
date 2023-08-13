package helper

import (
	"github.com/go-playground/validator/v10"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	validationError, ok := err.(validator.ValidationErrors)
	if ok {
		for _, e := range validationError {
			errors = append(errors, e.Error())
		}
	}

	return errors
}

func FormatCommas(str string) string {
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}

func GenerateOrderID(campaignID uint) string {
	now := time.Now().Format("2006-01-02 15:04:05")
	nowSplit := strings.Split(now, " ")
	date := nowSplit[0]
	dateSplit := strings.Split(date, "-")
	campaignFormat := 1000 + campaignID
	camp := []byte(strconv.FormatInt(int64(campaignFormat), 10))
	year := []byte(dateSplit[0])

	rand.Seed(time.Now().UnixNano())
	unique := rand.Intn(900) + 100
	seq := strconv.FormatInt(int64(unique), 10)

	return "TRX-" + string(camp[1:]) + string(year[2:]) + dateSplit[1] + seq
}
