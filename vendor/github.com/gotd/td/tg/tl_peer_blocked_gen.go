// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// PeerBlocked represents TL type `peerBlocked#e8fd8014`.
// Information about a blocked peer
//
// See https://core.telegram.org/constructor/peerBlocked for reference.
type PeerBlocked struct {
	// Peer ID
	PeerID PeerClass
	// When was the peer blocked
	Date int
}

// PeerBlockedTypeID is TL type id of PeerBlocked.
const PeerBlockedTypeID = 0xe8fd8014

// Ensuring interfaces in compile-time for PeerBlocked.
var (
	_ bin.Encoder     = &PeerBlocked{}
	_ bin.Decoder     = &PeerBlocked{}
	_ bin.BareEncoder = &PeerBlocked{}
	_ bin.BareDecoder = &PeerBlocked{}
)

func (p *PeerBlocked) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PeerID == nil) {
		return false
	}
	if !(p.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerBlocked) String() string {
	if p == nil {
		return "PeerBlocked(nil)"
	}
	type Alias PeerBlocked
	return fmt.Sprintf("PeerBlocked%+v", Alias(*p))
}

// FillFrom fills PeerBlocked from given interface.
func (p *PeerBlocked) FillFrom(from interface {
	GetPeerID() (value PeerClass)
	GetDate() (value int)
}) {
	p.PeerID = from.GetPeerID()
	p.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PeerBlocked) TypeID() uint32 {
	return PeerBlockedTypeID
}

// TypeName returns name of type in TL schema.
func (*PeerBlocked) TypeName() string {
	return "peerBlocked"
}

// TypeInfo returns info about TL type.
func (p *PeerBlocked) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "peerBlocked",
		ID:   PeerBlockedTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PeerID",
			SchemaName: "peer_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PeerBlocked) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerBlocked#e8fd8014 as nil")
	}
	b.PutID(PeerBlockedTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PeerBlocked) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerBlocked#e8fd8014 as nil")
	}
	if p.PeerID == nil {
		return fmt.Errorf("unable to encode peerBlocked#e8fd8014: field peer_id is nil")
	}
	if err := p.PeerID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode peerBlocked#e8fd8014: field peer_id: %w", err)
	}
	b.PutInt(p.Date)
	return nil
}

// Decode implements bin.Decoder.
func (p *PeerBlocked) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerBlocked#e8fd8014 to nil")
	}
	if err := b.ConsumeID(PeerBlockedTypeID); err != nil {
		return fmt.Errorf("unable to decode peerBlocked#e8fd8014: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PeerBlocked) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerBlocked#e8fd8014 to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode peerBlocked#e8fd8014: field peer_id: %w", err)
		}
		p.PeerID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerBlocked#e8fd8014: field date: %w", err)
		}
		p.Date = value
	}
	return nil
}

// GetPeerID returns value of PeerID field.
func (p *PeerBlocked) GetPeerID() (value PeerClass) {
	if p == nil {
		return
	}
	return p.PeerID
}

// GetDate returns value of Date field.
func (p *PeerBlocked) GetDate() (value int) {
	if p == nil {
		return
	}
	return p.Date
}