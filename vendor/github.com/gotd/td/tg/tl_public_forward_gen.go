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

// PublicForwardMessage represents TL type `publicForwardMessage#1f2bf4a`.
// Contains info about a forward of a story¹ as a message.
//
// Links:
//  1. https://core.telegram.org/api/stories
//
// See https://core.telegram.org/constructor/publicForwardMessage for reference.
type PublicForwardMessage struct {
	// Info about the message with the reposted story.
	Message MessageClass
}

// PublicForwardMessageTypeID is TL type id of PublicForwardMessage.
const PublicForwardMessageTypeID = 0x1f2bf4a

// construct implements constructor of PublicForwardClass.
func (p PublicForwardMessage) construct() PublicForwardClass { return &p }

// Ensuring interfaces in compile-time for PublicForwardMessage.
var (
	_ bin.Encoder     = &PublicForwardMessage{}
	_ bin.Decoder     = &PublicForwardMessage{}
	_ bin.BareEncoder = &PublicForwardMessage{}
	_ bin.BareDecoder = &PublicForwardMessage{}

	_ PublicForwardClass = &PublicForwardMessage{}
)

func (p *PublicForwardMessage) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Message == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PublicForwardMessage) String() string {
	if p == nil {
		return "PublicForwardMessage(nil)"
	}
	type Alias PublicForwardMessage
	return fmt.Sprintf("PublicForwardMessage%+v", Alias(*p))
}

// FillFrom fills PublicForwardMessage from given interface.
func (p *PublicForwardMessage) FillFrom(from interface {
	GetMessage() (value MessageClass)
}) {
	p.Message = from.GetMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PublicForwardMessage) TypeID() uint32 {
	return PublicForwardMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*PublicForwardMessage) TypeName() string {
	return "publicForwardMessage"
}

// TypeInfo returns info about TL type.
func (p *PublicForwardMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "publicForwardMessage",
		ID:   PublicForwardMessageTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PublicForwardMessage) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode publicForwardMessage#1f2bf4a as nil")
	}
	b.PutID(PublicForwardMessageTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PublicForwardMessage) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode publicForwardMessage#1f2bf4a as nil")
	}
	if p.Message == nil {
		return fmt.Errorf("unable to encode publicForwardMessage#1f2bf4a: field message is nil")
	}
	if err := p.Message.Encode(b); err != nil {
		return fmt.Errorf("unable to encode publicForwardMessage#1f2bf4a: field message: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PublicForwardMessage) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode publicForwardMessage#1f2bf4a to nil")
	}
	if err := b.ConsumeID(PublicForwardMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode publicForwardMessage#1f2bf4a: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PublicForwardMessage) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode publicForwardMessage#1f2bf4a to nil")
	}
	{
		value, err := DecodeMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode publicForwardMessage#1f2bf4a: field message: %w", err)
		}
		p.Message = value
	}
	return nil
}

// GetMessage returns value of Message field.
func (p *PublicForwardMessage) GetMessage() (value MessageClass) {
	if p == nil {
		return
	}
	return p.Message
}

// PublicForwardStory represents TL type `publicForwardStory#edf3add0`.
// Contains info about a forward of a story¹ as a repost by a public channel.
//
// Links:
//  1. https://core.telegram.org/api/stories
//
// See https://core.telegram.org/constructor/publicForwardStory for reference.
type PublicForwardStory struct {
	// The channel that reposted the story.
	Peer PeerClass
	// The reposted story (may be different from the original story).
	Story StoryItemClass
}

// PublicForwardStoryTypeID is TL type id of PublicForwardStory.
const PublicForwardStoryTypeID = 0xedf3add0

// construct implements constructor of PublicForwardClass.
func (p PublicForwardStory) construct() PublicForwardClass { return &p }

// Ensuring interfaces in compile-time for PublicForwardStory.
var (
	_ bin.Encoder     = &PublicForwardStory{}
	_ bin.Decoder     = &PublicForwardStory{}
	_ bin.BareEncoder = &PublicForwardStory{}
	_ bin.BareDecoder = &PublicForwardStory{}

	_ PublicForwardClass = &PublicForwardStory{}
)

func (p *PublicForwardStory) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Peer == nil) {
		return false
	}
	if !(p.Story == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PublicForwardStory) String() string {
	if p == nil {
		return "PublicForwardStory(nil)"
	}
	type Alias PublicForwardStory
	return fmt.Sprintf("PublicForwardStory%+v", Alias(*p))
}

// FillFrom fills PublicForwardStory from given interface.
func (p *PublicForwardStory) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetStory() (value StoryItemClass)
}) {
	p.Peer = from.GetPeer()
	p.Story = from.GetStory()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PublicForwardStory) TypeID() uint32 {
	return PublicForwardStoryTypeID
}

// TypeName returns name of type in TL schema.
func (*PublicForwardStory) TypeName() string {
	return "publicForwardStory"
}

// TypeInfo returns info about TL type.
func (p *PublicForwardStory) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "publicForwardStory",
		ID:   PublicForwardStoryTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Story",
			SchemaName: "story",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PublicForwardStory) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode publicForwardStory#edf3add0 as nil")
	}
	b.PutID(PublicForwardStoryTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PublicForwardStory) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode publicForwardStory#edf3add0 as nil")
	}
	if p.Peer == nil {
		return fmt.Errorf("unable to encode publicForwardStory#edf3add0: field peer is nil")
	}
	if err := p.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode publicForwardStory#edf3add0: field peer: %w", err)
	}
	if p.Story == nil {
		return fmt.Errorf("unable to encode publicForwardStory#edf3add0: field story is nil")
	}
	if err := p.Story.Encode(b); err != nil {
		return fmt.Errorf("unable to encode publicForwardStory#edf3add0: field story: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PublicForwardStory) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode publicForwardStory#edf3add0 to nil")
	}
	if err := b.ConsumeID(PublicForwardStoryTypeID); err != nil {
		return fmt.Errorf("unable to decode publicForwardStory#edf3add0: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PublicForwardStory) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode publicForwardStory#edf3add0 to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode publicForwardStory#edf3add0: field peer: %w", err)
		}
		p.Peer = value
	}
	{
		value, err := DecodeStoryItem(b)
		if err != nil {
			return fmt.Errorf("unable to decode publicForwardStory#edf3add0: field story: %w", err)
		}
		p.Story = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (p *PublicForwardStory) GetPeer() (value PeerClass) {
	if p == nil {
		return
	}
	return p.Peer
}

// GetStory returns value of Story field.
func (p *PublicForwardStory) GetStory() (value StoryItemClass) {
	if p == nil {
		return
	}
	return p.Story
}

// PublicForwardClassName is schema name of PublicForwardClass.
const PublicForwardClassName = "PublicForward"

// PublicForwardClass represents PublicForward generic type.
//
// See https://core.telegram.org/type/PublicForward for reference.
//
// Example:
//
//	g, err := tg.DecodePublicForward(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.PublicForwardMessage: // publicForwardMessage#1f2bf4a
//	case *tg.PublicForwardStory: // publicForwardStory#edf3add0
//	default: panic(v)
//	}
type PublicForwardClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PublicForwardClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodePublicForward implements binary de-serialization for PublicForwardClass.
func DecodePublicForward(buf *bin.Buffer) (PublicForwardClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PublicForwardMessageTypeID:
		// Decoding publicForwardMessage#1f2bf4a.
		v := PublicForwardMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PublicForwardClass: %w", err)
		}
		return &v, nil
	case PublicForwardStoryTypeID:
		// Decoding publicForwardStory#edf3add0.
		v := PublicForwardStory{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PublicForwardClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PublicForwardClass: %w", bin.NewUnexpectedID(id))
	}
}

// PublicForward boxes the PublicForwardClass providing a helper.
type PublicForwardBox struct {
	PublicForward PublicForwardClass
}

// Decode implements bin.Decoder for PublicForwardBox.
func (b *PublicForwardBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PublicForwardBox to nil")
	}
	v, err := DecodePublicForward(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PublicForward = v
	return nil
}

// Encode implements bin.Encode for PublicForwardBox.
func (b *PublicForwardBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PublicForward == nil {
		return fmt.Errorf("unable to encode PublicForwardClass as nil")
	}
	return b.PublicForward.Encode(buf)
}