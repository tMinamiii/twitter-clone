package appcontext

import (
	"context"
	"fmt"
	"log"
)

type ctxKey int

const CtxKeyUserID ctxKey = iota + 1

func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, CtxKeyUserID, userID)
}

func GetUserID(ctx context.Context) (int64, error) {
	prop, ok := ctx.Value(CtxKeyUserID).(int64)
	if !ok {
		err := fmt.Errorf("user id type assertion was failed")
		log.Println(err)
		return 0, err
	}

	return prop, nil
}
