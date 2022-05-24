package main

import (
	"context"
	"fmt"
	"io"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"google.golang.org/api/iterator"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

func main() {
	var w io.Writer
	listMonitoredResources(w, "metis-sandbox")
}

// listMonitoredResources lists all the resources available to be monitored.
func listMonitoredResources(w io.Writer, projectID string) error {
	ctx := context.Background()
	c, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	req := &monitoringpb.ListMonitoredResourceDescriptorsRequest{
		Name: "projects/" + projectID,
	}
	iter := c.ListMonitoredResourceDescriptors(ctx, req)

	for {
		resp, err := iter.Next()
		if err != nil {
			return fmt.Errorf("Could not list time series: %v", err)
		}
		if err == iterator.Done {
			break
		}
		fmt.Fprintf(w, "%v\n", resp)
	}
	fmt.Fprintln(w, "Done")
	return nil
}
