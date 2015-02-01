//

package zib

import (
	//"fmt"
	"github.com/philetus/l33tspek/tag"
)

type XMark interface {
	tag.Duflr
	ZibHandl() tag.Handl
	MarkHandl() tag.Handl
}
type XZib struct {
	ZbHndl tag.Handl
}
func (self XZib) ZibHandl() tag.Handl {
	return self.ZbHndl
}
func (self XZib) MarkHandl() tag.Handl {
	return ""
}
func (self XZib) Dufl() tag.Dufl {
	return tag.Dufl{self.ZbHndl}
}
func (self XZib) Tukkd() bool {
	return self.ZbHndl != ""
}
type XPaan struct {
	ZbHndl tag.Handl
	MrkHndl tag.Handl
}
func (self XPaan) ZibHandl() tag.Handl {
	return self.ZbHndl
}
func (self XPaan) MarkHandl() tag.Handl {
	return self.MrkHndl
}
func (self XPaan) Dufl() tag.Dufl {
	return tag.Dufl{self.ZbHndl, self.MrkHndl}
}
type XFlit struct {
	ZbHndl tag.Handl
	MrkHndl tag.Handl
}
func (self XFlit) ZibHandl() tag.Handl {
	return self.ZbHndl
}
func (self XFlit) MarkHandl() tag.Handl {
	return self.MrkHndl
}
func (self XFlit) Dufl() tag.Dufl {
	return tag.Dufl{self.ZbHndl, self.MrkHndl}
}
type XYok struct {
	ZbHndl tag.Handl
	MrkHndl tag.Handl
}
func (self XYok) ZibHandl() tag.Handl {
	return self.ZbHndl
}
func (self XYok) MarkHandl() tag.Handl {
	return self.MrkHndl
}
func (self XYok) Dufl() tag.Dufl {
	return tag.Dufl{self.ZbHndl, self.MrkHndl}
}
type XWarp struct {
	ZbHndl tag.Handl
	MrkHndl tag.Handl
}
func (self XWarp) ZibHandl() tag.Handl {
	return self.ZbHndl
}
func (self XWarp) MarkHandl() tag.Handl {
	return self.MrkHndl
}
func (self XWarp) Dufl() tag.Dufl {
	return tag.Dufl{self.ZbHndl, self.MrkHndl}
}
func (self XWarp) Tukkd() bool {
	return (self.ZbHndl != "") || (self.MrkHndl != "")
}
type XDelta struct {
	ZbHndl tag.Handl
	MrkHndl tag.Handl
}
func (self XDelta) ZibHandl() tag.Handl {
	return self.ZbHndl
}
func (self XDelta) MarkHandl() tag.Handl {
	return self.MrkHndl
}
func (self XDelta) Dufl() tag.Dufl {
	return tag.Dufl{self.ZbHndl, self.MrkHndl}
}

