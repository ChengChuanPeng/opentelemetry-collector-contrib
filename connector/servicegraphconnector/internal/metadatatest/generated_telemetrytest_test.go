// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector/internal/metadata"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := SetupTelemetry()
	tb, err := metadata.NewTelemetryBuilder(
		testTel.NewTelemetrySettings(),
	)
	require.NoError(t, err)
	require.NotNil(t, tb)
	tb.ConnectorServicegraphDroppedSpans.Add(context.Background(), 1)
	tb.ConnectorServicegraphExpiredEdges.Add(context.Background(), 1)
	tb.ConnectorServicegraphTotalEdges.Add(context.Background(), 1)

	testTel.AssertMetrics(t, []metricdata.Metrics{
		{
			Name:        "otelcol_connector_servicegraph_dropped_spans",
			Description: "Number of spans dropped when trying to add edges",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_connector_servicegraph_expired_edges",
			Description: "Number of edges that expired before finding its matching span",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_connector_servicegraph_total_edges",
			Description: "Total number of unique edges",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
	}, metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreValue())
	require.NoError(t, testTel.Shutdown(context.Background()))
}