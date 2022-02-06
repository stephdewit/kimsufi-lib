package availability

import (
	"fmt"
)

type Client struct {
	Country string
}

func (kimsufi *Client) GetAvailability(region string, hardware string) string {
	return fmt.Sprintf("Looking at %s for %s in %s", kimsufi.Country, hardware, region)
}
