package kaleido

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

func ValidateGetResponse(res *resty.Response, err error, resourceName string) {
	if res.StatusCode() != 200 {
		msg := fmt.Sprintf("Could not retrieve %s. Status code: %d.", resourceName, res.StatusCode())
		exit(msg)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%+v\n", res)
	}
}

func ValidateCreationResponse(res *resty.Response, err error, resourceName string) {
	if res.StatusCode() != 201 {
		msg := fmt.Sprintf("Could not create %s. Status code: %d.", resourceName, res.StatusCode())
		exit(msg)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%+v\n", res)
	}
}

func ValidateDeletionResponse(res *resty.Response, err error, resourceName string) {
	if res.StatusCode() != 202 && res.StatusCode() != 204 {
		msg := fmt.Sprintf("%s deletion failed. Status code: %d\n", resourceName, res.StatusCode())
		exit(msg)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%+v\n", res)
	}
}
