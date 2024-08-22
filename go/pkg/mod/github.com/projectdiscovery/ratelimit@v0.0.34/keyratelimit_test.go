package ratelimit_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/projectdiscovery/ratelimit"
	"github.com/stretchr/testify/require"
)

func TestMultiLimiter(t *testing.T) {
	limiter, err := ratelimit.NewMultiLimiter(context.Background(), &ratelimit.Options{
		Key:         "default",
		IsUnlimited: false,
		MaxCount:    100,
		Duration:    time.Duration(3) * time.Second,
	})
	require.Nil(t, err)
	wg := &sync.WaitGroup{}
	expectedTime := (time.Duration(6) * time.Second).Round(time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defaultStart := time.Now()
		for i := 0; i < 201; i++ {
			errx := limiter.Take("default")
			require.Nil(t, errx, "failed to take")
		}
		timeTaken := time.Since(defaultStart).Round(time.Millisecond)
		require.GreaterOrEqualf(t, timeTaken.Nanoseconds(), expectedTime.Nanoseconds(), "more token sent than requested in given timeframe")
	}()

	err = limiter.Add(&ratelimit.Options{
		Key:         "one",
		IsUnlimited: false,
		MaxCount:    100,
		Duration:    time.Duration(3) * time.Second,
	})
	require.Nil(t, err)

	wg.Add(1)
	go func() {
		defer wg.Done()
		oneStart := time.Now()
		for i := 0; i < 201; i++ {
			errx := limiter.Take("one")
			require.Nil(t, errx)
		}
		timeTaken := time.Since(oneStart).Round(time.Millisecond)
		require.GreaterOrEqualf(t, timeTaken.Nanoseconds(), expectedTime.Nanoseconds(), "more token sent than requested in given timeframe")
	}()
	wg.Wait()
}

func TestAdaptiveLimit(t *testing.T) {
	limiter := ratelimit.New(context.TODO(), 1, time.Second)
	require.NotNil(t, limiter)
	start := time.Now()
	expectedDuration := (time.Duration(3) * time.Second).Round(time.Millisecond)

	go func() {
		time.Sleep(2 * time.Second)

		limiter.SetLimit(100)
	}()

	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			limiter.Take()
		}()
	}
	wg.Wait()

	require.WithinDuration(t, start.Add(expectedDuration), time.Now(), time.Second)
}
