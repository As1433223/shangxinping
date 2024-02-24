package es

import (
	"context"
	"server/global"
)

func EsIkIndex(index string) error {
	_, err := global.EsClient.CreateIndex(index).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
