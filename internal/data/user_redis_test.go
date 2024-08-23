package data

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"my-template-with-go/bootstrap"
	"my-template-with-go/container"
	"my-template-with-go/logger"
	"os"
	"path/filepath"
	"testing"
)

func TestUserRedisRepo_GetAllUser(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)

	child := filepath.Dir(wd)
	parent := filepath.Dir(child)

	cf, err := bootstrap.InitConfig(fmt.Sprintf("%s/configs", parent))
	assert.Nil(t, err)

	zap, err := logger.InitLogger(cf)
	assert.Nil(t, err)

	result, err := container.NewContainer(cf, zap)
	assert.Nil(t, err)

	repo := NewRedisRepo(result.RedisProvider())

	res, err := repo.GetAllUser()
	assert.Nil(t, err)
	fmt.Println(res)
}
