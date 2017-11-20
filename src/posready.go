package main

import (
	"os"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/infra-integrations-sdk/metric"
	"github.com/newrelic/infra-integrations-sdk/sdk"
	// "fmt"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Filelocation string `default:"C:\\crdata\\mutex\\PosReady.flg" help:"File location to monitor"`
	Filename     string `default:"PosReady.flg" help:"File name being monitored"`
	// `default:"C:\\crdata\\mutex\\PosReady.flg" help:"File name to monitor"`
}

const (
	integrationName    = "com.newrelic.posready"
	integrationVersion = "0.1.0"
)

var args argumentList

func populateInventory(inventory sdk.Inventory) error {
	// Insert here the logic of your integration to get the inventory data
	// Ex: inventory.SetItem("softwareVersion", "value", "1.0.1")
	// --
	return nil
}

func populateMetrics(ms *metric.MetricSet) error {
	// Insert here the logic of your integration to get the metrics data

	if _, err := os.Stat(args.Filelocation); os.IsNotExist(err) {
		// fmt.Printf("File does NOT exist\n")

		log.Debug(args.Filename + " NOT Found")
		log.Debug("File Location: " + args.Filelocation)

		ms.SetMetric(args.Filename, "FALSE", metric.ATTRIBUTE)
	} else {
		log.Debug(args.Filename + " Found")
		log.Debug("File Location: " + args.Filelocation)
		ms.SetMetric(args.Filename, "TRUE", metric.ATTRIBUTE)
	}
	return nil
}

func main() {
	log.SetupLogging(args.Verbose)

	integration, err := sdk.NewIntegration(integrationName, integrationVersion, &args)
	fatalIfErr(err)

	if args.All || args.Inventory {
		fatalIfErr(populateInventory(integration.Inventory))
	}

	if args.All || args.Metrics {
		// fatalIfNotDefined(args.fileName, "Missing fileName parameter")
		ms := integration.NewMetricSet("PosHealth")
		fatalIfErr(populateMetrics(ms))
	}
	fatalIfErr(integration.Publish())
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
