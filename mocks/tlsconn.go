// SPDX-License-Identifier: GPL-3.0-or-later

package mocks

import "crypto/tls"

// TLSConn is a mockable TLS connection.
type TLSConn struct {
	// We embed Conn to handle the net.Conn interface.
	Conn

	// MockConnectionState is the function to call when ConnectionState is called.
	MockConnectionState func() tls.ConnectionState

	// MockHandshakeContext is the function to call when HandshakeContext is called.
	MockHandshakeContext func() error
}

// ConnectionState calls MockConnectionState.
func (c *TLSConn) ConnectionState() tls.ConnectionState {
	return c.MockConnectionState()
}

// HandshakeContext calls MockHandshakeContext.
func (c *TLSConn) HandshakeContext() error {
	return c.MockHandshakeContext()
}
