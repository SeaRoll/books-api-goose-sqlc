package handler

import (
	"books-api/db"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func GetConditions(c echo.Context) error {
	returnConditions := []BucketConditionDTO{}
	err := db.WithTX(c.Request().Context(), func(ctx context.Context, q *db.Queries, tx pgx.Tx) error {
		conditions, err := q.GetConditions(ctx, tx)
		if err != nil {
			return err
		}
		for _, condition := range conditions {
			returnConditions = append(returnConditions, BucketConditionDTO{
				Day:     condition.Bucket.Time.UTC(),
				AvgTemp: condition.AvgTemp,
			})
		}
		return err
	})
	if err != nil {
		return returnServerError(c, err)
	}

	return c.JSON(200, returnConditions)
}

func InsertCondition(c echo.Context) error {
	insertDTO := InsertConditionDTO{}
	err := c.Bind(&insertDTO)
	if err != nil {
		return returnUserError(c, err)
	}
	err = validateStruct(insertDTO)
	if err != nil {
		return returnUserError(c, err)
	}

	err = db.WithTX(c.Request().Context(), func(ctx context.Context, q *db.Queries, tx pgx.Tx) error {
		err := q.InsertCondition(ctx, tx, db.InsertConditionParams{
			Time: pgtype.Timestamptz{
				Time:  insertDTO.OccurredAt.UTC(),
				Valid: true,
			},
			Location:    insertDTO.Location,
			Device:      insertDTO.Device,
			Temperature: insertDTO.Temperature,
			Humidity:    insertDTO.Humidity,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return returnServerError(c, err)
	}

	return returnSuccess(c, 201)
}
