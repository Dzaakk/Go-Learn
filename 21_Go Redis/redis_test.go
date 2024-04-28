package golearn

import (
	"context"
	"testing"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

var ctx = context.Background()

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	err := client.Close()
	assert.Nil(t, err)
}

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "agent1", 3*time.Second)

	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "agent1", result)

	time.Sleep(5 * time.Second)

	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)

}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "first name")
	client.RPush(ctx, "names", "middle name")
	client.RPush(ctx, "names", "last name")

	assert.Equal(t, "first name", client.LPop(ctx, "names").Val())
	assert.Equal(t, "middle name", client.LPop(ctx, "names").Val())
	assert.Equal(t, "last name", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}
