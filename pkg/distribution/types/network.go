//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package types

const NetworkTypeVxLAN = "vxlan"

// swagger:model types_network_spec
type NetworkSpec struct {
	// Node network type
	Type string `json:"type"`
	// Node Network subnet info
	Range string `json:"range"`
	// Node Network interface
	IFace NetworkInterface `json:"iface"`
	// Node Public IP
	Addr string `json:"addr"`
}

// swagger:model types_network_interface
type NetworkInterface struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	HAddr string `json:"HAddr"`
}

func (n *NetworkSpec) Equal(nt *NetworkSpec) bool {

	switch false {
	case n.Type == nt.Type:
		return false
	case n.Range == nt.Range:
		return false
	case n.IFace.Index == nt.IFace.Index:
		return false
	case n.IFace.Name == nt.IFace.Name:
		return false
	case n.IFace.Addr == nt.IFace.Addr:
		return false
	case n.IFace.HAddr == nt.IFace.HAddr:
		return false
	case n.Addr == nt.Addr:
		return false
	}
	return true
}
