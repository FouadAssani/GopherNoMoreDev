package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewApplication(t *testing.T) {
	app, err := NewApplication()

	assert.NoError(t, err)
	assert.NotNil(t, app)
	assert.NotNil(t, app.log)
	assert.NotNil(t, app.application)
	assert.NotNil(t, app.routines)
}

func TestStart(t *testing.T) {
	app, err := NewApplication()
	assert.NoError(t, err)

	errCh := app.Start()
	require.NoError(t, <-errCh)
	app.Stop()
}

func TestStop(t *testing.T) {
	app, err := NewApplication()
	require.NoError(t, err)

	errCh := app.Start()
	require.NoError(t, <-errCh)
	app.Stop()
}
