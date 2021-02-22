/*
Copyright 2020 PayPal

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/paypal/load-watcher/pkg/watcher"
	"github.com/paypal/load-watcher/pkg/watcher/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	client, err := api.NewLibraryClient(watcher.EnvMetricProviderOpts)
	if err != nil {
		log.Fatalf("unable to create client: %v", err)
	}
	metrics, err := client.GetLatestWatcherMetrics()
	if err != nil {
		log.Errorf("unable to get watcher metrics: %v", err)
	}
	log.Infof("received metrics: %v", metrics)

	// Keep the watcher server up
	select {}
}