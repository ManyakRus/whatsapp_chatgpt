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

// ChatOnlines represents TL type `chatOnlines#f041e250`.
// Number of online users in a chat
//
// See https://core.telegram.org/constructor/chatOnlines for reference.
type ChatOnlines struct {
	// Number of online users
	Onlines int
}

// ChatOnlinesTypeID is TL type id of ChatOnlines.
const ChatOnlinesTypeID = 0xf041e250

// Ensuring interfaces in compile-time for ChatOnlines.
var (
	_ bin.Encoder     = &ChatOnlines{}
	_ bin.Decoder     = &ChatOnlines{}
	_ bin.BareEncoder = &ChatOnlines{}
	_ bin.BareDecoder = &ChatOnlines{}
)

func (c *ChatOnlines) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Onlines == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatOnlines) String() string {
	if c == nil {
		return "ChatOnlines(nil)"
	}
	type Alias ChatOnlines
	return fmt.Sprintf("ChatOnlines%+v", Alias(*c))
}

// FillFrom fills ChatOnlines from given interface.
func (c *ChatOnlines) FillFrom(from interface {
	GetOnlines() (value int)
}) {
	c.Onlines = from.GetOnlines()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatOnlines) TypeID() uint32 {
	return ChatOnlinesTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatOnlines) TypeName() string {
	return "chatOnlines"
}

// TypeInfo returns info about TL type.
func (c *ChatOnlines) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatOnlines",
		ID:   ChatOnlinesTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Onlines",
			SchemaName: "onlines",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatOnlines) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatOnlines#f041e250 as nil")
	}
	b.PutID(ChatOnlinesTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatOnlines) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatOnlines#f041e250 as nil")
	}
	b.PutInt(c.Onlines)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatOnlines) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatOnlines#f041e250 to nil")
	}
	if err := b.ConsumeID(ChatOnlinesTypeID); err != nil {
		return fmt.Errorf("unable to decode chatOnlines#f041e250: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatOnlines) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatOnlines#f041e250 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatOnlines#f041e250: field onlines: %w", err)
		}
		c.Onlines = value
	}
	return nil
}

// GetOnlines returns value of Onlines field.
func (c *ChatOnlines) GetOnlines() (value int) {
	if c == nil {
		return
	}
	return c.Onlines
}