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

// EmojiGroup represents TL type `emojiGroup#7a9abda9`.
// Represents an emoji category¹.
//
// Links:
//  1. https://core.telegram.org/api/custom-emoji#emoji-categories
//
// See https://core.telegram.org/constructor/emojiGroup for reference.
type EmojiGroup struct {
	// Category name, i.e. "Animals", "Flags", "Faces" and so on...
	Title string
	// A single custom emoji used as preview for the category.
	IconEmojiID int64
	// A list of UTF-8 emojis, matching the category.
	Emoticons []string
}

// EmojiGroupTypeID is TL type id of EmojiGroup.
const EmojiGroupTypeID = 0x7a9abda9

// construct implements constructor of EmojiGroupClass.
func (e EmojiGroup) construct() EmojiGroupClass { return &e }

// Ensuring interfaces in compile-time for EmojiGroup.
var (
	_ bin.Encoder     = &EmojiGroup{}
	_ bin.Decoder     = &EmojiGroup{}
	_ bin.BareEncoder = &EmojiGroup{}
	_ bin.BareDecoder = &EmojiGroup{}

	_ EmojiGroupClass = &EmojiGroup{}
)

func (e *EmojiGroup) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Title == "") {
		return false
	}
	if !(e.IconEmojiID == 0) {
		return false
	}
	if !(e.Emoticons == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiGroup) String() string {
	if e == nil {
		return "EmojiGroup(nil)"
	}
	type Alias EmojiGroup
	return fmt.Sprintf("EmojiGroup%+v", Alias(*e))
}

// FillFrom fills EmojiGroup from given interface.
func (e *EmojiGroup) FillFrom(from interface {
	GetTitle() (value string)
	GetIconEmojiID() (value int64)
	GetEmoticons() (value []string)
}) {
	e.Title = from.GetTitle()
	e.IconEmojiID = from.GetIconEmojiID()
	e.Emoticons = from.GetEmoticons()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiGroup) TypeID() uint32 {
	return EmojiGroupTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiGroup) TypeName() string {
	return "emojiGroup"
}

// TypeInfo returns info about TL type.
func (e *EmojiGroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiGroup",
		ID:   EmojiGroupTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "IconEmojiID",
			SchemaName: "icon_emoji_id",
		},
		{
			Name:       "Emoticons",
			SchemaName: "emoticons",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiGroup) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiGroup#7a9abda9 as nil")
	}
	b.PutID(EmojiGroupTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiGroup) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiGroup#7a9abda9 as nil")
	}
	b.PutString(e.Title)
	b.PutLong(e.IconEmojiID)
	b.PutVectorHeader(len(e.Emoticons))
	for _, v := range e.Emoticons {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiGroup) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiGroup#7a9abda9 to nil")
	}
	if err := b.ConsumeID(EmojiGroupTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiGroup#7a9abda9: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiGroup) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiGroup#7a9abda9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiGroup#7a9abda9: field title: %w", err)
		}
		e.Title = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode emojiGroup#7a9abda9: field icon_emoji_id: %w", err)
		}
		e.IconEmojiID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode emojiGroup#7a9abda9: field emoticons: %w", err)
		}

		if headerLen > 0 {
			e.Emoticons = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiGroup#7a9abda9: field emoticons: %w", err)
			}
			e.Emoticons = append(e.Emoticons, value)
		}
	}
	return nil
}

// GetTitle returns value of Title field.
func (e *EmojiGroup) GetTitle() (value string) {
	if e == nil {
		return
	}
	return e.Title
}

// GetIconEmojiID returns value of IconEmojiID field.
func (e *EmojiGroup) GetIconEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.IconEmojiID
}

// GetEmoticons returns value of Emoticons field.
func (e *EmojiGroup) GetEmoticons() (value []string) {
	if e == nil {
		return
	}
	return e.Emoticons
}

// EmojiGroupGreeting represents TL type `emojiGroupGreeting#80d26cc7`.
//
// See https://core.telegram.org/constructor/emojiGroupGreeting for reference.
type EmojiGroupGreeting struct {
	// Title field of EmojiGroupGreeting.
	Title string
	// IconEmojiID field of EmojiGroupGreeting.
	IconEmojiID int64
	// Emoticons field of EmojiGroupGreeting.
	Emoticons []string
}

// EmojiGroupGreetingTypeID is TL type id of EmojiGroupGreeting.
const EmojiGroupGreetingTypeID = 0x80d26cc7

// construct implements constructor of EmojiGroupClass.
func (e EmojiGroupGreeting) construct() EmojiGroupClass { return &e }

// Ensuring interfaces in compile-time for EmojiGroupGreeting.
var (
	_ bin.Encoder     = &EmojiGroupGreeting{}
	_ bin.Decoder     = &EmojiGroupGreeting{}
	_ bin.BareEncoder = &EmojiGroupGreeting{}
	_ bin.BareDecoder = &EmojiGroupGreeting{}

	_ EmojiGroupClass = &EmojiGroupGreeting{}
)

func (e *EmojiGroupGreeting) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Title == "") {
		return false
	}
	if !(e.IconEmojiID == 0) {
		return false
	}
	if !(e.Emoticons == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiGroupGreeting) String() string {
	if e == nil {
		return "EmojiGroupGreeting(nil)"
	}
	type Alias EmojiGroupGreeting
	return fmt.Sprintf("EmojiGroupGreeting%+v", Alias(*e))
}

// FillFrom fills EmojiGroupGreeting from given interface.
func (e *EmojiGroupGreeting) FillFrom(from interface {
	GetTitle() (value string)
	GetIconEmojiID() (value int64)
	GetEmoticons() (value []string)
}) {
	e.Title = from.GetTitle()
	e.IconEmojiID = from.GetIconEmojiID()
	e.Emoticons = from.GetEmoticons()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiGroupGreeting) TypeID() uint32 {
	return EmojiGroupGreetingTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiGroupGreeting) TypeName() string {
	return "emojiGroupGreeting"
}

// TypeInfo returns info about TL type.
func (e *EmojiGroupGreeting) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiGroupGreeting",
		ID:   EmojiGroupGreetingTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "IconEmojiID",
			SchemaName: "icon_emoji_id",
		},
		{
			Name:       "Emoticons",
			SchemaName: "emoticons",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiGroupGreeting) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiGroupGreeting#80d26cc7 as nil")
	}
	b.PutID(EmojiGroupGreetingTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiGroupGreeting) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiGroupGreeting#80d26cc7 as nil")
	}
	b.PutString(e.Title)
	b.PutLong(e.IconEmojiID)
	b.PutVectorHeader(len(e.Emoticons))
	for _, v := range e.Emoticons {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiGroupGreeting) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiGroupGreeting#80d26cc7 to nil")
	}
	if err := b.ConsumeID(EmojiGroupGreetingTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiGroupGreeting#80d26cc7: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiGroupGreeting) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiGroupGreeting#80d26cc7 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiGroupGreeting#80d26cc7: field title: %w", err)
		}
		e.Title = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode emojiGroupGreeting#80d26cc7: field icon_emoji_id: %w", err)
		}
		e.IconEmojiID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode emojiGroupGreeting#80d26cc7: field emoticons: %w", err)
		}

		if headerLen > 0 {
			e.Emoticons = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emojiGroupGreeting#80d26cc7: field emoticons: %w", err)
			}
			e.Emoticons = append(e.Emoticons, value)
		}
	}
	return nil
}

// GetTitle returns value of Title field.
func (e *EmojiGroupGreeting) GetTitle() (value string) {
	if e == nil {
		return
	}
	return e.Title
}

// GetIconEmojiID returns value of IconEmojiID field.
func (e *EmojiGroupGreeting) GetIconEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.IconEmojiID
}

// GetEmoticons returns value of Emoticons field.
func (e *EmojiGroupGreeting) GetEmoticons() (value []string) {
	if e == nil {
		return
	}
	return e.Emoticons
}

// EmojiGroupPremium represents TL type `emojiGroupPremium#93bcf34`.
//
// See https://core.telegram.org/constructor/emojiGroupPremium for reference.
type EmojiGroupPremium struct {
	// Title field of EmojiGroupPremium.
	Title string
	// IconEmojiID field of EmojiGroupPremium.
	IconEmojiID int64
}

// EmojiGroupPremiumTypeID is TL type id of EmojiGroupPremium.
const EmojiGroupPremiumTypeID = 0x93bcf34

// construct implements constructor of EmojiGroupClass.
func (e EmojiGroupPremium) construct() EmojiGroupClass { return &e }

// Ensuring interfaces in compile-time for EmojiGroupPremium.
var (
	_ bin.Encoder     = &EmojiGroupPremium{}
	_ bin.Decoder     = &EmojiGroupPremium{}
	_ bin.BareEncoder = &EmojiGroupPremium{}
	_ bin.BareDecoder = &EmojiGroupPremium{}

	_ EmojiGroupClass = &EmojiGroupPremium{}
)

func (e *EmojiGroupPremium) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Title == "") {
		return false
	}
	if !(e.IconEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiGroupPremium) String() string {
	if e == nil {
		return "EmojiGroupPremium(nil)"
	}
	type Alias EmojiGroupPremium
	return fmt.Sprintf("EmojiGroupPremium%+v", Alias(*e))
}

// FillFrom fills EmojiGroupPremium from given interface.
func (e *EmojiGroupPremium) FillFrom(from interface {
	GetTitle() (value string)
	GetIconEmojiID() (value int64)
}) {
	e.Title = from.GetTitle()
	e.IconEmojiID = from.GetIconEmojiID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiGroupPremium) TypeID() uint32 {
	return EmojiGroupPremiumTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiGroupPremium) TypeName() string {
	return "emojiGroupPremium"
}

// TypeInfo returns info about TL type.
func (e *EmojiGroupPremium) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiGroupPremium",
		ID:   EmojiGroupPremiumTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "IconEmojiID",
			SchemaName: "icon_emoji_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiGroupPremium) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiGroupPremium#93bcf34 as nil")
	}
	b.PutID(EmojiGroupPremiumTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiGroupPremium) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiGroupPremium#93bcf34 as nil")
	}
	b.PutString(e.Title)
	b.PutLong(e.IconEmojiID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiGroupPremium) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiGroupPremium#93bcf34 to nil")
	}
	if err := b.ConsumeID(EmojiGroupPremiumTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiGroupPremium#93bcf34: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiGroupPremium) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiGroupPremium#93bcf34 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiGroupPremium#93bcf34: field title: %w", err)
		}
		e.Title = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode emojiGroupPremium#93bcf34: field icon_emoji_id: %w", err)
		}
		e.IconEmojiID = value
	}
	return nil
}

// GetTitle returns value of Title field.
func (e *EmojiGroupPremium) GetTitle() (value string) {
	if e == nil {
		return
	}
	return e.Title
}

// GetIconEmojiID returns value of IconEmojiID field.
func (e *EmojiGroupPremium) GetIconEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.IconEmojiID
}

// EmojiGroupClassName is schema name of EmojiGroupClass.
const EmojiGroupClassName = "EmojiGroup"

// EmojiGroupClass represents EmojiGroup generic type.
//
// See https://core.telegram.org/type/EmojiGroup for reference.
//
// Example:
//
//	g, err := tg.DecodeEmojiGroup(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.EmojiGroup: // emojiGroup#7a9abda9
//	case *tg.EmojiGroupGreeting: // emojiGroupGreeting#80d26cc7
//	case *tg.EmojiGroupPremium: // emojiGroupPremium#93bcf34
//	default: panic(v)
//	}
type EmojiGroupClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() EmojiGroupClass

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

	// Category name, i.e. "Animals", "Flags", "Faces" and so on...
	GetTitle() (value string)

	// A single custom emoji used as preview for the category.
	GetIconEmojiID() (value int64)
}

// DecodeEmojiGroup implements binary de-serialization for EmojiGroupClass.
func DecodeEmojiGroup(buf *bin.Buffer) (EmojiGroupClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EmojiGroupTypeID:
		// Decoding emojiGroup#7a9abda9.
		v := EmojiGroup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiGroupClass: %w", err)
		}
		return &v, nil
	case EmojiGroupGreetingTypeID:
		// Decoding emojiGroupGreeting#80d26cc7.
		v := EmojiGroupGreeting{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiGroupClass: %w", err)
		}
		return &v, nil
	case EmojiGroupPremiumTypeID:
		// Decoding emojiGroupPremium#93bcf34.
		v := EmojiGroupPremium{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmojiGroupClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EmojiGroupClass: %w", bin.NewUnexpectedID(id))
	}
}

// EmojiGroup boxes the EmojiGroupClass providing a helper.
type EmojiGroupBox struct {
	EmojiGroup EmojiGroupClass
}

// Decode implements bin.Decoder for EmojiGroupBox.
func (b *EmojiGroupBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EmojiGroupBox to nil")
	}
	v, err := DecodeEmojiGroup(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EmojiGroup = v
	return nil
}

// Encode implements bin.Encode for EmojiGroupBox.
func (b *EmojiGroupBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EmojiGroup == nil {
		return fmt.Errorf("unable to encode EmojiGroupClass as nil")
	}
	return b.EmojiGroup.Encode(buf)
}