// This code has been copied from github.com/jhump/protoreflect/dynamic/text.go
// and amended to format specifically for this app

package app

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"reflect"
	"sort"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/codec"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

type indentBuffer struct {
	bytes.Buffer
	indentCount int
	comma       bool
}

func (b *indentBuffer) start() error {
	if b.indentCount >= 0 {
		b.indentCount++
		return b.newLine(false)
	}
	return nil
}

func (b *indentBuffer) sep() error {
	if b.indentCount >= 0 {
		_, err := b.WriteString(": ")
		return err
	}
	return b.WriteByte(':')
}

func (b *indentBuffer) end() error {
	if b.indentCount >= 0 {
		b.indentCount--
		return b.newLine(false)
	}
	return nil
}

func (b *indentBuffer) maybeNext(first *bool) error {
	if *first {
		*first = false
		return nil
	}
	return b.next()
}

func (b *indentBuffer) next() error {
	if b.indentCount >= 0 {
		return b.newLine(b.comma)
	} else if b.comma {
		return b.WriteByte(',')
	} else {
		return b.WriteByte(' ')
	}
}

func (b *indentBuffer) newLine(comma bool) error {
	if comma {
		err := b.WriteByte(',')
		if err != nil {
			return err
		}
	}

	_, err := b.WriteString("<br />")
	if err != nil {
		return err
	}

	for i := 0; i < b.indentCount; i++ {
		_, err := b.WriteString("&nbsp;&nbsp;")
		if err != nil {
			return err
		}
	}
	return nil
}

type sortable []interface{}

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Less(i, j int) bool {
	vi := s[i]
	vj := s[j]
	switch reflect.TypeOf(vi).Kind() {
	case reflect.Int32:
		return vi.(int32) < vj.(int32)
	case reflect.Int64:
		return vi.(int64) < vj.(int64)
	case reflect.Uint32:
		return vi.(uint32) < vj.(uint32)
	case reflect.Uint64:
		return vi.(uint64) < vj.(uint64)
	case reflect.String:
		return vi.(string) < vj.(string)
	case reflect.Bool:
		return !vi.(bool) && vj.(bool)
	default:
		panic(fmt.Sprintf("cannot compare keys of type %v", reflect.TypeOf(vi)))
	}
}

func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func marshalTextFormatted(m *dynamic.Message) ([]byte, error) {
	var b indentBuffer
	if err := marshalText(m, &b); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func marshalText(m *dynamic.Message, b *indentBuffer) error {
	first := true
	// first the known fields
	for _, fd := range m.GetKnownFields() {
		v := m.GetField(fd)
		if fd.IsMap() {
			md := fd.GetMessageType()
			kfd := md.FindFieldByNumber(1)
			vfd := md.FindFieldByNumber(2)
			mp := v.(map[interface{}]interface{})
			keys := make([]interface{}, 0, len(mp))
			for k := range mp {
				keys = append(keys, k)
			}
			sort.Sort(sortable(keys))
			for _, mk := range keys {
				mv := mp[mk]
				err := b.maybeNext(&first)
				if err != nil {
					return err
				}
				err = marshalKnownFieldMapEntryText(b, fd, kfd, mk, vfd, mv)
				if err != nil {
					return err
				}
			}
		} else if fd.IsRepeated() {
			sl := v.([]interface{})
			for _, slv := range sl {
				err := b.maybeNext(&first)
				if err != nil {
					return err
				}
				err = marshalKnownFieldText(b, fd, slv)
				if err != nil {
					return err
				}
			}
		} else {
			err := b.maybeNext(&first)
			if err != nil {
				return err
			}
			err = marshalKnownFieldText(b, fd, v)
			if err != nil {
				return err
			}
		}
	}
	// then the unknown fields
	for _, tag := range m.GetUnknownFields() {
		ufs := m.GetUnknownField(tag)
		for _, uf := range ufs {
			err := b.maybeNext(&first)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintf(b, "%d", tag)
			if err != nil {
				return err
			}
			if uf.Encoding == proto.WireStartGroup {
				err = b.WriteByte('{')
				if err != nil {
					return err
				}
				err = b.start()
				if err != nil {
					return err
				}
				in := codec.NewBuffer(uf.Contents)
				err = marshalUnknownGroupText(b, in, true)
				if err != nil {
					return err
				}
				err = b.end()
				if err != nil {
					return err
				}
				err = b.WriteByte('}')
				if err != nil {
					return err
				}
			} else {
				err = b.sep()
				if err != nil {
					return err
				}
				if uf.Encoding == proto.WireBytes {
					err = writeString(b, string(uf.Contents))
					if err != nil {
						return err
					}
				} else {
					_, err = b.WriteString(strconv.FormatUint(uf.Value, 10))
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func marshalKnownFieldMapEntryText(b *indentBuffer, fd *desc.FieldDescriptor, kfd *desc.FieldDescriptor, mk interface{}, vfd *desc.FieldDescriptor, mv interface{}) error {
	var name string
	if fd.IsExtension() {
		name = fmt.Sprintf("[%s]", fd.GetFullyQualifiedName())
	} else {
		name = fd.GetName()
	}
	_, err := b.WriteString(name)
	if err != nil {
		return err
	}
	err = b.sep()
	if err != nil {
		return err
	}

	err = b.WriteByte('<')
	if err != nil {
		return err
	}
	err = b.start()
	if err != nil {
		return err
	}

	err = marshalKnownFieldText(b, kfd, mk)
	if err != nil {
		return err
	}
	err = b.next()
	if err != nil {
		return err
	}
	if !isNil(mv) {
		err = marshalKnownFieldText(b, vfd, mv)
		if err != nil {
			return err
		}
	}

	err = b.end()
	if err != nil {
		return err
	}
	return b.WriteByte('>')
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Ptr && rv.IsNil()
}

func marshalKnownFieldText(b *indentBuffer, fd *desc.FieldDescriptor, v interface{}) error {
	b.WriteString("<span class='name'>")
	group := fd.GetType() == descriptor.FieldDescriptorProto_TYPE_GROUP
	if group {
		var name string
		if fd.IsExtension() {
			name = fmt.Sprintf("[%s]", fd.GetMessageType().GetFullyQualifiedName())
		} else {
			name = fd.GetMessageType().GetName()
		}
		_, err := b.WriteString(name)
		if err != nil {
			return err
		}
	} else {
		var name string
		if fd.IsExtension() {
			name = fmt.Sprintf("[%s]", fd.GetFullyQualifiedName())
		} else {
			name = fd.GetName()
		}
		_, err := b.WriteString(name)
		if err != nil {
			return err
		}
		err = b.sep()
		if err != nil {
			return err
		}
	}
	b.WriteString("</span>")
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int32, reflect.Int64:
		b.WriteString("<span class='num'>")
		ed := fd.GetEnumType()
		if ed != nil {
			n := int32(rv.Int())
			vd := ed.FindValueByNumber(n)
			if vd == nil {
				_, err := b.WriteString(strconv.FormatInt(rv.Int(), 10))
				return err
			}
			_, err := b.WriteString(vd.GetName())
			b.WriteString("</span>")
			return err
		}
		_, err := b.WriteString(strconv.FormatInt(rv.Int(), 10))
		b.WriteString("</span>")
		return err
	case reflect.Uint32, reflect.Uint64:
		b.WriteString("<span class='num'>")
		_, err := b.WriteString(strconv.FormatUint(rv.Uint(), 10))
		b.WriteString("</span>")
		return err
	case reflect.Float32, reflect.Float64:
		f := rv.Float()
		b.WriteString("<span class='num'>")
		var str string
		if math.IsNaN(f) {
			str = "nan"
		} else if math.IsInf(f, 1) {
			str = "inf"
		} else if math.IsInf(f, -1) {
			str = "-inf"
		} else {
			var bits int
			if rv.Kind() == reflect.Float32 {
				bits = 32
			} else {
				bits = 64
			}
			str = strconv.FormatFloat(rv.Float(), 'g', -1, bits)
		}
		_, err := b.WriteString(str)
		b.WriteString("</span>")
		return err
	case reflect.Bool:
		b.WriteString("<span class='bool'>")
		_, err := b.WriteString(strconv.FormatBool(rv.Bool()))
		b.WriteString("</span>")
		return err
	case reflect.Slice:
		return writeString(b, string(rv.Bytes()))
	case reflect.String:
		return writeString(b, rv.String())
	default:
		var err error
		if group {
			err = b.WriteByte('{')
		} else {
			_, err = b.WriteString("<span class='bkt'>&lt;</span>")
		}
		if err != nil {
			return err
		}
		err = b.start()
		if err != nil {
			return err
		}
		// must be a message
		if dm, ok := v.(*dynamic.Message); ok {
			err = marshalText(dm, b)
			if err != nil {
				return err
			}
		} else {
			err = proto.CompactText(b, v.(proto.Message))
			if err != nil {
				return err
			}
		}
		err = b.end()
		if err != nil {
			return err
		}
		if group {
			return b.WriteByte('}')
		}
		_, err := b.WriteString("<span class='bkt'>&gt;</span>")
		return err
	}
}

// writeString writes a string in the protocol buffer text format.
// It is similar to strconv.Quote except we don't use Go escape sequences,
// we treat the string as a byte sequence, and we use octal escapes.
// These differences are to maintain interoperability with the other
// languages' implementations of the text format.
func writeString(b *indentBuffer, s string) error {
	// use WriteByte here to get any needed indent
	if _, err := b.WriteString(`<span class='str'>"`); err != nil {
		return err
	}
	// Loop over the bytes, not the runes.
	for i := 0; i < len(s); i++ {
		var err error
		// Divergence from C++: we don't escape apostrophes.
		// There's no need to escape them, and the C++ parser
		// copes with a naked apostrophe.
		switch c := s[i]; c {
		case '\n':
			_, err = b.WriteString("\\n")
		case '\r':
			_, err = b.WriteString("\\r")
		case '\t':
			_, err = b.WriteString("\\t")
		case '"':
			_, err = b.WriteString("\\\"")
		case '\\':
			_, err = b.WriteString("\\\\")
		default:
			if c >= 0x20 && c < 0x7f {
				err = b.WriteByte(c)
			} else {
				_, err = fmt.Fprintf(b, "\\%03o", c)
			}
		}
		if err != nil {
			return err
		}
	}
	_, err := b.WriteString(`"</span>`)
	return err
}

func marshalUnknownGroupText(b *indentBuffer, in *codec.Buffer, topLevel bool) error {
	first := true
	for {
		if in.EOF() {
			if topLevel {
				return nil
			}
			// this is a nested message: we are expecting an end-group tag, not EOF!
			return io.ErrUnexpectedEOF
		}
		tag, wireType, err := in.DecodeTagAndWireType()
		if err != nil {
			return err
		}
		if wireType == proto.WireEndGroup {
			return nil
		}
		err = b.maybeNext(&first)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(b, "%d", tag)
		if err != nil {
			return err
		}
		if wireType == proto.WireStartGroup {
			err = b.WriteByte('{')
			if err != nil {
				return err
			}
			err = b.start()
			if err != nil {
				return err
			}
			err = marshalUnknownGroupText(b, in, false)
			if err != nil {
				return err
			}
			err = b.end()
			if err != nil {
				return err
			}
			err = b.WriteByte('}')
			if err != nil {
				return err
			}
			continue
		} else {
			err = b.sep()
			if err != nil {
				return err
			}
			if wireType == proto.WireBytes {
				contents, err := in.DecodeRawBytes(false)
				if err != nil {
					return err
				}
				err = writeString(b, string(contents))
				if err != nil {
					return err
				}
			} else {
				var v uint64
				switch wireType {
				case proto.WireVarint:
					v, err = in.DecodeVarint()
				case proto.WireFixed32:
					v, err = in.DecodeFixed32()
				case proto.WireFixed64:
					v, err = in.DecodeFixed64()
				default:
					return proto.ErrInternalBadWireType
				}
				if err != nil {
					return err
				}
				_, err = b.WriteString(strconv.FormatUint(v, 10))
				if err != nil {
					return err
				}
			}
		}
	}
}
