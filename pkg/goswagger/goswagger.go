package goswagger

import (
	"github.com/zhangbao138208/goctls/pkg/goctl"
	"github.com/zhangbao138208/goctls/pkg/golang"
	"github.com/zhangbao138208/goctls/util/env"
)

const (
	Name = "swagger"
	url  = "github.com/go-swagger/go-swagger/cmd/swagger@latest"
)

func Install(cacheDir string) (string, error) {
	return goctl.Install(cacheDir, Name, func(dest string) (string, error) {
		err := golang.Install(url)
		return dest, err
	})
}

func Exists() bool {
	_, err := env.LookUpSwagger()
	return err == nil
}
