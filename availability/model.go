package availability

type Instances []Instance

type Instance struct {
	Region      string       `json:region`
	Hardware    string       `json:hardware`
	Datacenters []Datacenter `json:datacenters`
}

type Datacenter struct {
	Datacenter   string `json:datacenter`
	Availability string `json:availability`
}
