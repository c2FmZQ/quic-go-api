//go:generate ./update-api.sh

// Package api contains auto-generated interfaces and wrappers for the [quic] data structures.
package api

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/c2FmZQ/quic-go-api"
	"github.com/c2FmZQ/quic-go-api/internal/protocol"
	"github.com/c2FmZQ/quic-go-api/internal/qerr"
)

// ### AUTO GENERATED CODE BELOW

// Transport is an auto-generated interface for [quic.Transport]
type Transport interface {
	Close() error
	Dial(context.Context, net.Addr, *tls.Config, *quic.Config) (Conn, error)
	DialEarly(context.Context, net.Addr, *tls.Config, *quic.Config) (Conn, error)
	Listen(*tls.Config, *quic.Config) (Listener, error)
	ListenEarly(*tls.Config, *quic.Config) (EarlyListener, error)
	ReadNonQUICPacket(context.Context, []uint8) (int, net.Addr, error)
	WriteTo([]uint8, net.Addr) (int, error)
}

// Listener is an auto-generated interface for [quic.Listener]
type Listener interface {
	Accept(context.Context) (Conn, error)
	Addr() net.Addr
	Close() error
}

// EarlyListener is an auto-generated interface for [quic.EarlyListener]
type EarlyListener interface {
	Accept(context.Context) (Conn, error)
	Addr() net.Addr
	Close() error
}

// Conn is an auto-generated interface for [quic.Conn]
type Conn interface {
	AcceptStream(context.Context) (Stream, error)
	AcceptUniStream(context.Context) (ReceiveStream, error)
	AddPath(Transport) (Path, error)
	CloseWithError(qerr.ApplicationErrorCode, string) error
	ConnectionState() quic.ConnectionState
	Context() context.Context
	HandshakeComplete() <-chan struct{}
	LocalAddr() net.Addr
	NextConnection(context.Context) (Conn, error)
	OpenStream() (Stream, error)
	OpenStreamSync(context.Context) (Stream, error)
	OpenUniStream() (SendStream, error)
	OpenUniStreamSync(context.Context) (SendStream, error)
	ReceiveDatagram(context.Context) ([]uint8, error)
	RemoteAddr() net.Addr
	SendDatagram([]uint8) error
}

// SendStream is an auto-generated interface for [quic.SendStream]
type SendStream interface {
	CancelWrite(qerr.StreamErrorCode)
	Close() error
	Context() context.Context
	SetWriteDeadline(time.Time) error
	StreamID() protocol.StreamID
	Write([]uint8) (int, error)
}

// ReceiveStream is an auto-generated interface for [quic.ReceiveStream]
type ReceiveStream interface {
	CancelRead(qerr.StreamErrorCode)
	Read([]uint8) (int, error)
	SetReadDeadline(time.Time) error
	StreamID() protocol.StreamID
}

// Stream is an auto-generated interface for [quic.Stream]
type Stream interface {
	CancelRead(qerr.StreamErrorCode)
	CancelWrite(qerr.StreamErrorCode)
	Close() error
	Context() context.Context
	Read([]uint8) (int, error)
	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	StreamID() protocol.StreamID
	Write([]uint8) (int, error)
}

// Path is an auto-generated interface for [quic.Path]
type Path interface {
	Close() error
	Probe(context.Context) error
	Switch() error
}

var _ Transport = (*TransportWrapper)(nil)

// TransportWrapper is an auto-generated wrapper for [quic.Transport]
type TransportWrapper struct {
	Base *quic.Transport
}

func (w *TransportWrapper) Close() error {
	return w.Base.Close()
}

func (w *TransportWrapper) Dial(a1 context.Context, a2 net.Addr, a3 *tls.Config, a4 *quic.Config) (r1 Conn, r2 error) {
	var t1 *quic.Conn
	t1, r2 = w.Base.Dial(a1, a2, a3, a4)
	if t1 != nil {
		r1 = &ConnWrapper{Base: t1}
	}
	return
}

func (w *TransportWrapper) DialEarly(a1 context.Context, a2 net.Addr, a3 *tls.Config, a4 *quic.Config) (r1 Conn, r2 error) {
	var t1 *quic.Conn
	t1, r2 = w.Base.DialEarly(a1, a2, a3, a4)
	if t1 != nil {
		r1 = &ConnWrapper{Base: t1}
	}
	return
}

func (w *TransportWrapper) Listen(a1 *tls.Config, a2 *quic.Config) (r1 Listener, r2 error) {
	var t1 *quic.Listener
	t1, r2 = w.Base.Listen(a1, a2)
	if t1 != nil {
		r1 = &ListenerWrapper{Base: t1}
	}
	return
}

func (w *TransportWrapper) ListenEarly(a1 *tls.Config, a2 *quic.Config) (r1 EarlyListener, r2 error) {
	var t1 *quic.EarlyListener
	t1, r2 = w.Base.ListenEarly(a1, a2)
	if t1 != nil {
		r1 = &EarlyListenerWrapper{Base: t1}
	}
	return
}

func (w *TransportWrapper) ReadNonQUICPacket(a1 context.Context, a2 []uint8) (int, net.Addr, error) {
	return w.Base.ReadNonQUICPacket(a1, a2)
}

func (w *TransportWrapper) WriteTo(a1 []uint8, a2 net.Addr) (int, error) {
	return w.Base.WriteTo(a1, a2)
}

var _ Listener = (*ListenerWrapper)(nil)

// ListenerWrapper is an auto-generated wrapper for [quic.Listener]
type ListenerWrapper struct {
	Base *quic.Listener
}

func (w *ListenerWrapper) Accept(a1 context.Context) (r1 Conn, r2 error) {
	var t1 *quic.Conn
	t1, r2 = w.Base.Accept(a1)
	if t1 != nil {
		r1 = &ConnWrapper{Base: t1}
	}
	return
}

func (w *ListenerWrapper) Addr() net.Addr {
	return w.Base.Addr()
}

func (w *ListenerWrapper) Close() error {
	return w.Base.Close()
}

var _ EarlyListener = (*EarlyListenerWrapper)(nil)

// EarlyListenerWrapper is an auto-generated wrapper for [quic.EarlyListener]
type EarlyListenerWrapper struct {
	Base *quic.EarlyListener
}

func (w *EarlyListenerWrapper) Accept(a1 context.Context) (r1 Conn, r2 error) {
	var t1 *quic.Conn
	t1, r2 = w.Base.Accept(a1)
	if t1 != nil {
		r1 = &ConnWrapper{Base: t1}
	}
	return
}

func (w *EarlyListenerWrapper) Addr() net.Addr {
	return w.Base.Addr()
}

func (w *EarlyListenerWrapper) Close() error {
	return w.Base.Close()
}

var _ Conn = (*ConnWrapper)(nil)

// ConnWrapper is an auto-generated wrapper for [quic.Conn]
type ConnWrapper struct {
	Base *quic.Conn
}

func (w *ConnWrapper) AcceptStream(a1 context.Context) (r1 Stream, r2 error) {
	var t1 *quic.Stream
	t1, r2 = w.Base.AcceptStream(a1)
	if t1 != nil {
		r1 = &StreamWrapper{Base: t1}
	}
	return
}

func (w *ConnWrapper) AcceptUniStream(a1 context.Context) (r1 ReceiveStream, r2 error) {
	var t1 *quic.ReceiveStream
	t1, r2 = w.Base.AcceptUniStream(a1)
	if t1 != nil {
		r1 = &ReceiveStreamWrapper{Base: t1}
	}
	return
}

func (w *ConnWrapper) AddPath(a1 Transport) (r1 Path, r2 error) {
	var t1 *quic.Path
	t1, r2 = w.Base.AddPath(a1.(*TransportWrapper).Base)
	if t1 != nil {
		r1 = &PathWrapper{Base: t1}
	}
	return
}

func (w *ConnWrapper) CloseWithError(a1 qerr.ApplicationErrorCode, a2 string) error {
	return w.Base.CloseWithError(a1, a2)
}

func (w *ConnWrapper) ConnectionState() quic.ConnectionState {
	return w.Base.ConnectionState()
}

func (w *ConnWrapper) Context() context.Context {
	return w.Base.Context()
}

func (w *ConnWrapper) HandshakeComplete() <-chan struct{} {
	return w.Base.HandshakeComplete()
}

func (w *ConnWrapper) LocalAddr() net.Addr {
	return w.Base.LocalAddr()
}

func (w *ConnWrapper) NextConnection(a1 context.Context) (r1 Conn, r2 error) {
	var t1 *quic.Conn
	t1, r2 = w.Base.NextConnection(a1)
	if t1 != nil {
		r1 = &ConnWrapper{Base: t1}
	}
	return
}

func (w *ConnWrapper) OpenStream() (r1 Stream, r2 error) {
	var t1 *quic.Stream
	t1, r2 = w.Base.OpenStream()
	if t1 != nil {
		r1 = &StreamWrapper{Base: t1}
	}
	return
}

func (w *ConnWrapper) OpenStreamSync(a1 context.Context) (r1 Stream, r2 error) {
	var t1 *quic.Stream
	t1, r2 = w.Base.OpenStreamSync(a1)
	if t1 != nil {
		r1 = &StreamWrapper{Base: t1}
	}
	return
}

func (w *ConnWrapper) OpenUniStream() (r1 SendStream, r2 error) {
	var t1 *quic.SendStream
	t1, r2 = w.Base.OpenUniStream()
	if t1 != nil {
		r1 = &SendStreamWrapper{Base: t1}
	}
	return
}

func (w *ConnWrapper) OpenUniStreamSync(a1 context.Context) (r1 SendStream, r2 error) {
	var t1 *quic.SendStream
	t1, r2 = w.Base.OpenUniStreamSync(a1)
	if t1 != nil {
		r1 = &SendStreamWrapper{Base: t1}
	}
	return
}

func (w *ConnWrapper) ReceiveDatagram(a1 context.Context) ([]uint8, error) {
	return w.Base.ReceiveDatagram(a1)
}

func (w *ConnWrapper) RemoteAddr() net.Addr {
	return w.Base.RemoteAddr()
}

func (w *ConnWrapper) SendDatagram(a1 []uint8) error {
	return w.Base.SendDatagram(a1)
}

var _ SendStream = (*SendStreamWrapper)(nil)

// SendStreamWrapper is an auto-generated wrapper for [quic.SendStream]
type SendStreamWrapper struct {
	Base *quic.SendStream
}

func (w *SendStreamWrapper) CancelWrite(a1 qerr.StreamErrorCode) {
	w.Base.CancelWrite(a1)
}

func (w *SendStreamWrapper) Close() error {
	return w.Base.Close()
}

func (w *SendStreamWrapper) Context() context.Context {
	return w.Base.Context()
}

func (w *SendStreamWrapper) SetWriteDeadline(a1 time.Time) error {
	return w.Base.SetWriteDeadline(a1)
}

func (w *SendStreamWrapper) StreamID() protocol.StreamID {
	return w.Base.StreamID()
}

func (w *SendStreamWrapper) Write(a1 []uint8) (int, error) {
	return w.Base.Write(a1)
}

var _ ReceiveStream = (*ReceiveStreamWrapper)(nil)

// ReceiveStreamWrapper is an auto-generated wrapper for [quic.ReceiveStream]
type ReceiveStreamWrapper struct {
	Base *quic.ReceiveStream
}

func (w *ReceiveStreamWrapper) CancelRead(a1 qerr.StreamErrorCode) {
	w.Base.CancelRead(a1)
}

func (w *ReceiveStreamWrapper) Read(a1 []uint8) (int, error) {
	return w.Base.Read(a1)
}

func (w *ReceiveStreamWrapper) SetReadDeadline(a1 time.Time) error {
	return w.Base.SetReadDeadline(a1)
}

func (w *ReceiveStreamWrapper) StreamID() protocol.StreamID {
	return w.Base.StreamID()
}

var _ Stream = (*StreamWrapper)(nil)

// StreamWrapper is an auto-generated wrapper for [quic.Stream]
type StreamWrapper struct {
	Base *quic.Stream
}

func (w *StreamWrapper) CancelRead(a1 qerr.StreamErrorCode) {
	w.Base.CancelRead(a1)
}

func (w *StreamWrapper) CancelWrite(a1 qerr.StreamErrorCode) {
	w.Base.CancelWrite(a1)
}

func (w *StreamWrapper) Close() error {
	return w.Base.Close()
}

func (w *StreamWrapper) Context() context.Context {
	return w.Base.Context()
}

func (w *StreamWrapper) Read(a1 []uint8) (int, error) {
	return w.Base.Read(a1)
}

func (w *StreamWrapper) SetDeadline(a1 time.Time) error {
	return w.Base.SetDeadline(a1)
}

func (w *StreamWrapper) SetReadDeadline(a1 time.Time) error {
	return w.Base.SetReadDeadline(a1)
}

func (w *StreamWrapper) SetWriteDeadline(a1 time.Time) error {
	return w.Base.SetWriteDeadline(a1)
}

func (w *StreamWrapper) StreamID() protocol.StreamID {
	return w.Base.StreamID()
}

func (w *StreamWrapper) Write(a1 []uint8) (int, error) {
	return w.Base.Write(a1)
}

var _ Path = (*PathWrapper)(nil)

// PathWrapper is an auto-generated wrapper for [quic.Path]
type PathWrapper struct {
	Base *quic.Path
}

func (w *PathWrapper) Close() error {
	return w.Base.Close()
}

func (w *PathWrapper) Probe(a1 context.Context) error {
	return w.Base.Probe(a1)
}

func (w *PathWrapper) Switch() error {
	return w.Base.Switch()
}
