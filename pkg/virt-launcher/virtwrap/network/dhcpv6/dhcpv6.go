/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2020 Red Hat, Inc.
 *
 */

package dhcpv6

import (
	"fmt"
	"net"
	"time"

	"github.com/insomniacslk/dhcp/dhcpv6"
	"github.com/insomniacslk/dhcp/dhcpv6/server6"
	"github.com/insomniacslk/dhcp/iana"

	"kubevirt.io/client-go/log"
)

const (
	infiniteLease = 999 * 24 * time.Hour
)

type DHCPv6Handler struct {
	clientIP    net.IP
	serverIface string
}

func SingleClientDHCPv6Server(clientIP net.IP, serverIface string) error {
	log.Log.Info("Starting SingleClientDHCPv6Server")

	handler := &DHCPv6Handler{
		clientIP:    clientIP,
		serverIface: serverIface,
	}

	s, err := server6.NewServer(serverIface, nil, handler.ServeDHCPv6)
	if err != nil {
		return fmt.Errorf("couldn't create DHCPv6 server: %v", err)
	}

	err = s.Serve()
	if err != nil {
		return fmt.Errorf("failed to run DHCPv6 server: %v", err)
	}

	return nil
}

func (h *DHCPv6Handler) ServeDHCPv6(conn net.PacketConn, peer net.Addr, m dhcpv6.DHCPv6) {
	log.Log.V(4).Info("DHCPv6 serving a new request")

	// TODO if we extend the server to support bridge binding, we need to filter out non-vm requests

	var response *dhcpv6.Message

	optIAAddress := dhcpv6.OptIAAddress{IPv6Addr: h.clientIP, PreferredLifetime: infiniteLease, ValidLifetime: infiniteLease}

	iface, err := net.InterfaceByName(h.serverIface)
	if err != nil {
		log.Log.V(4).Info("DHCPv6 - couldn't get the server interface")
		return
	}
	duid := dhcpv6.Duid{Type: dhcpv6.DUID_LL, HwType: iana.HWTypeEthernet, LinkLayerAddr: iface.HardwareAddr}

	dhcpv6Msg := m.(*dhcpv6.Message)
	switch dhcpv6Msg.Type() {
	case dhcpv6.MessageTypeSolicit:
		log.Log.V(4).Info("DHCPv6 - the request has message type Solicit")
		response, err = dhcpv6.NewAdvertiseFromSolicit(dhcpv6Msg, dhcpv6.WithIANA(optIAAddress), dhcpv6.WithServerID(duid))
	default:
		log.Log.V(4).Info("DHCPv6 - non Solicit request recieved")
		response, err = dhcpv6.NewReplyFromMessage(dhcpv6Msg, dhcpv6.WithIANA(optIAAddress), dhcpv6.WithServerID(duid))
	}

	if err != nil {
		log.Log.V(4).Errorf("DHCPv6 failed sending a response to the client: %v", err)
		return
	}

	if _, err := conn.WriteTo(response.ToBytes(), peer); err != nil {
		log.Log.V(4).Errorf("DHCPv6 cannot reply to client: %v", err)
	}
}
