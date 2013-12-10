package dockerapi



/* Container type
 */
type Container struct {
	ID string
	Created string
	Path string
	Args []string
	Name string
	Image string
	NetworkSettings NetworkSettings
	Volumes map[string]string
}

type NetworkSettings struct {
	IPAddress string
	Gateway string
	Ports map[string][]Address
}

type Address struct {
	HostIp string
	HostPort string
}
