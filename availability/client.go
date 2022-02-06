package availability

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	endpoint          = "https://www.ovh.com/engine/api/dedicated/server/availabilities"
	defaultDatacenter = "default"
)

type Client struct {
	Country string
}

func (kimsufi *Client) GetAvailability(region string, hardware string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s?country=%s", endpoint, kimsufi.Country))
	if err != nil {
		return "", fmt.Errorf("Can't get availabilities: %w", err)
	}
	defer resp.Body.Close()

	var instances Instances
	err = json.NewDecoder(resp.Body).Decode(&instances)
	if err != nil {
		return "", fmt.Errorf("Can't decode body: %w", err)
	}

	for _, instance := range instances {
		if instance.Region == region && instance.Hardware == hardware {
			for _, datacenter := range instance.Datacenters {
				if datacenter.Datacenter == defaultDatacenter {
					return datacenter.Availability, nil
				}
			}
		}
	}

	return "", fmt.Errorf("Can't find %s in %s at %s", hardware, region, kimsufi.Country)
}
