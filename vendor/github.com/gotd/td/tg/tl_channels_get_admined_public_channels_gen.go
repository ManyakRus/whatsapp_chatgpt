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

// ChannelsGetAdminedPublicChannelsRequest represents TL type `channels.getAdminedPublicChannels#f8b036af`.
// Get channels/supergroups/geogroups¹ we're admin in. Usually called when the user
// exceeds the limit² for owned public channels/supergroups/geogroups³, and the user is
// given the choice to remove one of his channels/supergroups/geogroups.
//
// Links:
//  1. https://core.telegram.org/api/channel
//  2. https://core.telegram.org/constructor/config
//  3. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getAdminedPublicChannels for reference.
type ChannelsGetAdminedPublicChannelsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Get geogroups
	ByLocation bool
	// If set and the user has reached the limit of owned public
	// channels/supergroups/geogroups¹, instead of returning the channel list one of the
	// specified errors² will be returned.Useful to check if a new public channel can indeed
	// be created, even before asking the user to enter a channel username to use in channels
	// checkUsername³/channels.updateUsername⁴.
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	//  2) https://core.telegram.org#possible-errors
	//  3) https://core.telegram.org/method/channels.checkUsername
	//  4) https://core.telegram.org/method/channels.updateUsername
	CheckLimit bool
	// ForPersonal field of ChannelsGetAdminedPublicChannelsRequest.
	ForPersonal bool
}

// ChannelsGetAdminedPublicChannelsRequestTypeID is TL type id of ChannelsGetAdminedPublicChannelsRequest.
const ChannelsGetAdminedPublicChannelsRequestTypeID = 0xf8b036af

// Ensuring interfaces in compile-time for ChannelsGetAdminedPublicChannelsRequest.
var (
	_ bin.Encoder     = &ChannelsGetAdminedPublicChannelsRequest{}
	_ bin.Decoder     = &ChannelsGetAdminedPublicChannelsRequest{}
	_ bin.BareEncoder = &ChannelsGetAdminedPublicChannelsRequest{}
	_ bin.BareDecoder = &ChannelsGetAdminedPublicChannelsRequest{}
)

func (g *ChannelsGetAdminedPublicChannelsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ByLocation == false) {
		return false
	}
	if !(g.CheckLimit == false) {
		return false
	}
	if !(g.ForPersonal == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetAdminedPublicChannelsRequest) String() string {
	if g == nil {
		return "ChannelsGetAdminedPublicChannelsRequest(nil)"
	}
	type Alias ChannelsGetAdminedPublicChannelsRequest
	return fmt.Sprintf("ChannelsGetAdminedPublicChannelsRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetAdminedPublicChannelsRequest from given interface.
func (g *ChannelsGetAdminedPublicChannelsRequest) FillFrom(from interface {
	GetByLocation() (value bool)
	GetCheckLimit() (value bool)
	GetForPersonal() (value bool)
}) {
	g.ByLocation = from.GetByLocation()
	g.CheckLimit = from.GetCheckLimit()
	g.ForPersonal = from.GetForPersonal()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetAdminedPublicChannelsRequest) TypeID() uint32 {
	return ChannelsGetAdminedPublicChannelsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetAdminedPublicChannelsRequest) TypeName() string {
	return "channels.getAdminedPublicChannels"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetAdminedPublicChannelsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getAdminedPublicChannels",
		ID:   ChannelsGetAdminedPublicChannelsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ByLocation",
			SchemaName: "by_location",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "CheckLimit",
			SchemaName: "check_limit",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "ForPersonal",
			SchemaName: "for_personal",
			Null:       !g.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *ChannelsGetAdminedPublicChannelsRequest) SetFlags() {
	if !(g.ByLocation == false) {
		g.Flags.Set(0)
	}
	if !(g.CheckLimit == false) {
		g.Flags.Set(1)
	}
	if !(g.ForPersonal == false) {
		g.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (g *ChannelsGetAdminedPublicChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getAdminedPublicChannels#f8b036af as nil")
	}
	b.PutID(ChannelsGetAdminedPublicChannelsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetAdminedPublicChannelsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getAdminedPublicChannels#f8b036af as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getAdminedPublicChannels#f8b036af: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetAdminedPublicChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getAdminedPublicChannels#f8b036af to nil")
	}
	if err := b.ConsumeID(ChannelsGetAdminedPublicChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getAdminedPublicChannels#f8b036af: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetAdminedPublicChannelsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getAdminedPublicChannels#f8b036af to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.getAdminedPublicChannels#f8b036af: field flags: %w", err)
		}
	}
	g.ByLocation = g.Flags.Has(0)
	g.CheckLimit = g.Flags.Has(1)
	g.ForPersonal = g.Flags.Has(2)
	return nil
}

// SetByLocation sets value of ByLocation conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) SetByLocation(value bool) {
	if value {
		g.Flags.Set(0)
		g.ByLocation = true
	} else {
		g.Flags.Unset(0)
		g.ByLocation = false
	}
}

// GetByLocation returns value of ByLocation conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) GetByLocation() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// SetCheckLimit sets value of CheckLimit conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) SetCheckLimit(value bool) {
	if value {
		g.Flags.Set(1)
		g.CheckLimit = true
	} else {
		g.Flags.Unset(1)
		g.CheckLimit = false
	}
}

// GetCheckLimit returns value of CheckLimit conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) GetCheckLimit() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(1)
}

// SetForPersonal sets value of ForPersonal conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) SetForPersonal(value bool) {
	if value {
		g.Flags.Set(2)
		g.ForPersonal = true
	} else {
		g.Flags.Unset(2)
		g.ForPersonal = false
	}
}

// GetForPersonal returns value of ForPersonal conditional field.
func (g *ChannelsGetAdminedPublicChannelsRequest) GetForPersonal() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(2)
}

// ChannelsGetAdminedPublicChannels invokes method channels.getAdminedPublicChannels#f8b036af returning error if any.
// Get channels/supergroups/geogroups¹ we're admin in. Usually called when the user
// exceeds the limit² for owned public channels/supergroups/geogroups³, and the user is
// given the choice to remove one of his channels/supergroups/geogroups.
//
// Links:
//  1. https://core.telegram.org/api/channel
//  2. https://core.telegram.org/constructor/config
//  3. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 CHANNELS_ADMIN_LOCATED_TOO_MUCH: The user has reached the limit of public geogroups.
//	400 CHANNELS_ADMIN_PUBLIC_TOO_MUCH: You're admin of too many public channels, make some channels private to change the username of this channel.
//
// See https://core.telegram.org/method/channels.getAdminedPublicChannels for reference.
func (c *Client) ChannelsGetAdminedPublicChannels(ctx context.Context, request *ChannelsGetAdminedPublicChannelsRequest) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}