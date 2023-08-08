package main

import (
	"fmt"
	"strings"
	"net"
	"time"

	"github.com/google/go-cmp/cmp"
)


type (
	Gateway struct {
		SSID      string
		IPAddress net.IP
		NetMask   net.IPMask
		Clients   []Client
	}
	Client struct {
		Hostname  string
		IPAddress net.IP
		LastSeen  time.Time
	}
)

// DiffReporter is a simple custom reporter that only records differences
// detected during comparison.
type DiffReporter struct {
	path  cmp.Path
	diffs []string
}

func (r *DiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *DiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		vx, vy := r.path.Last().Values()
		r.diffs = append(r.diffs, fmt.Sprintf("%#v:\n\t-: %+v\n\t+: %+v\n", r.path, vx, vy))
	}
}

func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *DiffReporter) String() string {
	return strings.Join(r.diffs, "\n")
}

func main() {
	x, y := MakeGatewayInfo()

	var r DiffReporter
	cmp.Equal(x, y, cmp.Reporter(&r))
	fmt.Print(r.String())

}

func MakeGatewayInfo() (x, y Gateway) {
	x = Gateway{
		SSID:      "CoffeeShopWiFi",
		IPAddress: net.IPv4(192, 168, 0, 1),
		NetMask:   net.IPv4Mask(255, 255, 0, 0),
		Clients: []Client{{
			Hostname:  "ristretto",
			IPAddress: net.IPv4(192, 168, 0, 116),
		}, {
			Hostname:  "aribica",
			IPAddress: net.IPv4(192, 168, 0, 104),
			LastSeen:  time.Date(2009, time.November, 10, 23, 6, 32, 0, time.UTC),
		}, {
			Hostname:  "macchiato",
			IPAddress: net.IPv4(192, 168, 0, 153),
			LastSeen:  time.Date(2009, time.November, 10, 23, 39, 43, 0, time.UTC),
		}, {
			Hostname:  "espresso",
			IPAddress: net.IPv4(192, 168, 0, 121),
		}, {
			Hostname:  "latte",
			IPAddress: net.IPv4(192, 168, 0, 219),
			LastSeen:  time.Date(2009, time.November, 10, 23, 0, 23, 0, time.UTC),
		}, {
			Hostname:  "americano",
			IPAddress: net.IPv4(192, 168, 0, 188),
			LastSeen:  time.Date(2009, time.November, 10, 23, 3, 5, 0, time.UTC),
		}},
	}
	y = Gateway{
		SSID:      "CoffeeShopWiFi",
		IPAddress: net.IPv4(192, 168, 0, 2),
		NetMask:   net.IPv4Mask(255, 255, 0, 0),
		Clients: []Client{{
			Hostname:  "ristretto",
			IPAddress: net.IPv4(192, 168, 0, 116),
		}, {
			Hostname:  "aribica",
			IPAddress: net.IPv4(192, 168, 0, 104),
			LastSeen:  time.Date(2009, time.November, 10, 23, 6, 32, 0, time.UTC),
		}, {
			Hostname:  "macchiato",
			IPAddress: net.IPv4(192, 168, 0, 153),
			LastSeen:  time.Date(2009, time.November, 10, 23, 39, 43, 0, time.UTC),
		}, {
			Hostname:  "espresso",
			IPAddress: net.IPv4(192, 168, 0, 121),
		}, {
			Hostname:  "latte",
			IPAddress: net.IPv4(192, 168, 0, 221),
			LastSeen:  time.Date(2009, time.November, 10, 23, 0, 23, 0, time.UTC),
		}},
	}
	return x, y
}

/*
{cmp_test.Gateway}.IPAddress:
	-: 192.168.0.1
	+: 192.168.0.2

{cmp_test.Gateway}.Clients[4].IPAddress:
	-: 192.168.0.219
	+: 192.168.0.221

{cmp_test.Gateway}.Clients[5->?]:
	-: {Hostname:americano IPAddress:192.168.0.188 LastSeen:2009-11-10 23:03:05 +0000 UTC}
	+: <invalid reflect.Value>
*/
 
