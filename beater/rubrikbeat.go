package beater

import (
	"fmt"
	"time"
	"log"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"

	"github.com/railroadmanuk/rubrikbeat/config"
)

// Rubrikbeat configuration.
type Rubrikbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of rubrikbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Rubrikbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts rubrikbeat.
func (bt *Rubrikbeat) Run(b *beat.Beat) error {
	logp.Info("rubrikbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	rubrik, err := rubrikcdm.ConnectEnv()
	if err != nil {
		log.Fatal(err)
	}

	clusterVersion, err := rubrik.ClusterVersion()
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"version": clusterVersion,
				"id": "1234",
				"clusterName": "test",
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
		counter++
	}
}

// Stop stops rubrikbeat.
func (bt *Rubrikbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
