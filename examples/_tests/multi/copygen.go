// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.

// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/switchupcb/copygen/examples/_tests/multi/complex"
	"github.com/switchupcb/copygen/examples/_tests/multi/external"
)

// Placeholder represents a basic type..
type Placeholder bool

// Collection represents a type that holds collection field types.
type Collection struct {
	Arr [16]byte
	S   []string
	M   map[string]bool
	C   chan int
}

// Array copies a [16]byte to a [16]byte.
func Array(tb [16]byte, fb [16]byte) {
	// [16]byte fields
	tb = fb
}

// ArrayComplex copies a [16]map[byte]string to a *complex.Collection.
func ArrayComplex(tC *complex.Collection, fm [16]map[byte]string) {
	// *complex.Collection fields
	tC.Arr = fm
}

// ArrayExternal copies a [16]external.Placeholder to a [16]external.Placeholder.
func ArrayExternal(tg [16]external.Placeholder, fg [16]external.Placeholder) {
	// [16]external.Placeholder fields
	tg = fg
}

// ArrayExternalComplex copies a [16]map[*external.Collection]string to a *complex.ComplexCollection.
func ArrayExternalComplex(tC *complex.ComplexCollection, fm [16]map[*external.Collection]string) {
	// *complex.ComplexCollection fields
	tC.Arr = fm
}

// ArraySimple copies a [16]byte to a *Collection.
func ArraySimple(tC *Collection, fb [16]byte) {
	// *Collection fields
	tC.Arr = fb
}

// Basic copies a bool to a bool.
func Basic(tb bool, fb bool) {
	// bool fields
	tb = fb
}

// BasicExternal copies a *external.Placeholder to a external.Placeholder.
func BasicExternal(tP external.Placeholder, fP *external.Placeholder) {
	// external.Placeholder fields
	tP = *fP
}

// BasicExternalMulti copies a *external.Placeholder to a external.Placeholder, *external.Placeholder.
func BasicExternalMulti(tP external.Placeholder, tP1 *external.Placeholder, fP *external.Placeholder) {
	// external.Placeholder fields
	tP = *fP

	// *external.Placeholder fields
	tP1 = fP
}

// BasicPointer copies a Placeholder to a *Placeholder.
func BasicPointer(tP *Placeholder, fP Placeholder) {
	// *Placeholder fields
	tP = &fP
}

// BasicPointerMulti copies a *Placeholder to a *Placeholder, *Placeholder, string.
func BasicPointerMulti(tP *Placeholder, tP1 *Placeholder, ts string, fP *Placeholder) {
	// *Placeholder fields
	tP = fP

	// *Placeholder fields

	// string fields
}

// BasicSimple copies a Placeholder to a Placeholder.
func BasicSimple(tP Placeholder, fP Placeholder) {
	// Placeholder fields
	tP = fP
}

// NoMatchArraySimple copies a [16]byte to a Collection.
func NoMatchArraySimple(tC Collection, fb [16]byte) {
	// Collection fields
}

// NoMatchBasic copies a Placeholder to a Placeholder.
func NoMatchBasic(tP Placeholder, fP Placeholder) {
	// Placeholder fields
}

// NoMatchBasicExternal copies a *Placeholder to a external.Placeholder, *external.Placeholder, bool.
func NoMatchBasicExternal(tP external.Placeholder, tP1 *external.Placeholder, tb bool, fP *Placeholder) {
	// external.Placeholder fields

	// *external.Placeholder fields

	// bool fields
}

// NoMatchSliceSimple copies a []string to a Collection.
func NoMatchSliceSimple(tC Collection, fs []string) {
	// Collection fields
}

// Slice copies a []string to a []string.
func Slice(ts []string, fs []string) {
	// []string fields
	ts = fs
}

// SliceComplex copies a []map[string][16]int to a *complex.Collection.
func SliceComplex(tC *complex.Collection, fm []map[string][16]int) {
	// *complex.Collection fields
	tC.S = fm
}

// SliceExternal copies a []external.Placeholder to a []external.Placeholder.
func SliceExternal(tg []external.Placeholder, fg []external.Placeholder) {
	// []external.Placeholder fields
	tg = fg
}

// SliceExternalComplex copies a []map[string]func(*external.Collection)string to a *complex.ComplexCollection.
func SliceExternalComplex(tC *complex.ComplexCollection, fm []map[string]func(*external.Collection) string) {
	// *complex.ComplexCollection fields
	tC.S = fm
}

// SliceSimple copies a []string to a *Collection.
func SliceSimple(tC *Collection, fs []string) {
	// *Collection fields
	tC.S = fs
}