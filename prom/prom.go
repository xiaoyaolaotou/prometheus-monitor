package prom

import (
	"context"
	"fmt"
	"monitor/global"
	"time"

	"github.com/prometheus/common/model"
)

func QueryVectorValue(serviceName, qpl string) (model.Vector, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	value, _, err := global.APIs[serviceName].Query(ctx, qpl, time.Now())
	if err != nil {
		return nil, fmt.Errorf("query error")
	}
	if value.Type() != model.ValVector {
		return nil, fmt.Errorf("value is not vector")
	}
	v, ok := value.(model.Vector)

	if !ok {
		return nil, fmt.Errorf("value is not vector")
	}
	return v, nil
}
