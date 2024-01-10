package test

import (
	"books-api/handler"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/stretchr/testify/assert"
)

// Insert around 1000 conditions
// Every 100 conditions, increase day by 1
// Temperature and humidity is random between 0 and 1
// Location is random between "LA" and "NY"
// Device is "A"
// OccurredAt is random between 00:00:00 and 23:59:59 for that day

func generateConditionDTO(day int) *handler.InsertConditionDTO {

	occurredAt := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, day).Add(time.Duration(rand.Intn(24)) * time.Hour).Add(time.Duration(rand.Intn(60)) * time.Minute).Add(time.Duration(rand.Intn(60)) * time.Second)
	randomLocation := rand.Intn(2)
	location := "LA"
	if randomLocation == 1 {
		location = "NY"
	}

	randomTemperature := rand.Float64()
	randomHumidity := rand.Float64()

	return &handler.InsertConditionDTO{
		OccurredAt:  occurredAt,
		Location:    location,
		Device:      "A",
		Temperature: randomTemperature,
		Humidity:    randomHumidity,
	}
}

func insertCondition(dto handler.InsertConditionDTO) {
	res, err := PerformRequest(BuildJsonRequest(http.MethodPost, "/conditions", dto), handler.InsertCondition, nil)
	if err != nil || res.Code != 201 {
		panic(err)
	}
}

func (suite *HandlerTestSuite) TestMassInsertions() {
	for i := 0; i < 1000; i++ {
		insertCondition(*generateConditionDTO(i / 100))
	}
	res, err := PerformRequest(BuildJsonRequest(http.MethodGet, "/conditions", nil), handler.GetConditions, nil)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, res.Code)
	fmt.Println(res.Body.String())
}
