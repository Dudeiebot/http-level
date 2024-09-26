package dude_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ankorstore/yokai/log/logtest"
	"github.com/ankorstore/yokai/trace/tracetest"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"

	"github.com/Dudeiebot/http-level/internal"
	"github.com/Dudeiebot/http-level/internal/model"
	"github.com/Dudeiebot/http-level/internal/repository"
)

func TestListGophersHandlerSuccess(t *testing.T) {
	// extraction
	var httpServer *echo.Echo
	var logBuffer logtest.TestLogBuffer
	var traceExporter tracetest.TestTraceExporter
	var metricsRegistry *prometheus.Registry
	var repo *repository.GopherRepository

	// run test
	internal.RunTest(
		t,
		fx.Populate(&httpServer, &logBuffer, &traceExporter, &metricsRegistry, &repo),
	)

	// populate database
	err := repo.Create(context.Background(), &model.Dude{
		Name: "bob",
		Job:  "builder",
	})
	assert.NoError(t, err)

	err = repo.Create(context.Background(), &model.Dude{
		Name: "alice",
		Job:  "doctor",
	})
	assert.NoError(t, err)

	// [GET] /gophers response assertion
	req := httptest.NewRequest(http.MethodGet, "/allpeople", nil)
	rec := httptest.NewRecorder()
	httpServer.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var gophers []*model.Dude
	err = json.Unmarshal(rec.Body.Bytes(), &gophers)
	assert.NoError(t, err)

	assert.Len(t, gophers, 2)
	assert.Equal(t, gophers[0].Name, "bob")
	assert.Equal(t, gophers[0].Job, "builder")
	assert.Equal(t, gophers[1].Name, "alice")
	assert.Equal(t, gophers[1].Job, "doctor")

	// logs assertion
	logtest.AssertHasLogRecord(t, logBuffer, map[string]interface{}{
		"level":   "info",
		"message": "called ListDudeHandler",
	})

	logtest.AssertHasLogRecord(t, logBuffer, map[string]interface{}{
		"level":   "info",
		"message": "called GopherService.List()",
	})

	// trace assertion
	tracetest.AssertHasTraceSpan(t, traceExporter, "ListDudeHandler span")

	// metrics assertion
	expectedMetric := `
        # HELP gophers_list_total The number of times gophers were listed
        # TYPE gophers_list_total counter
        gophers_list_total 1
    `

	err = testutil.GatherAndCompare(
		metricsRegistry,
		strings.NewReader(expectedMetric),
		"gophers_list_total",
	)
	assert.NoError(t, err)
}
