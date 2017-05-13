// GoArrayList project GoArrayList.go; see comments in doc.go
package goArrayList

import (
//"fmt" // only needed if debugging with fmt.Println etc
)

//Obj interface matches any type; named to reflect similarity to Java 'object' in ArrayList etc.
type Obj interface{}

// ArrayList struct is intended as a fairly complete Go language substitute for the Java feature ArrayList.
// Java documents describe ArrayList as: "Resizable-array implementation of the List interface".  ArrayList 
// struct in package GoArrayList contains an array of type Obj, can contain any type of elements (this is 
// comparable to Java spec, which states that ArrayList can only contain objects, not primatives); for 
// conversion to an array of Obj (elements will usually be of some basic Go type, but can also be Obj, 
// array slice, or struct), see ToArray method.
// Note: nil is allowed in an element, much as null is allowed in Java ArrayList.
// Skipped: field modCount not implemented; could be added if there is interest.
type ArrayList struct {
    Ary []Obj
}

// ArrayListPreset func creates and returns a new ArrayList and preloads it with contents of preset, 
// preserving their type (per entry, not per array/slice; see ToArray method for how to get an array of
// Obj back out of an ArrayList).  Comparable to Java constructor "ArrayList(Collection<? extends E> c)".
func ArrayListPreset(preset []Obj) (AL *ArrayList) { // func, not method, so no 'base'
    AL = &ArrayList{preset}
    return
}

// ArrayListNew func creates and returns a new empty ArrayList with specified capacity.  Emulates Java
// constructor "ArrayList(int initialCapacity)"; to emulate Java constructor "ArrayList()", call
// ArrayListNew with param 10.
func ArrayListNew(acap int) (AL *ArrayList) {
	// create with size 0 and capacity acap
	AL = &ArrayList{make([]Obj, 0, acap)} 
    return
}

// Clear method emulates Java ArrayList.clear; does a re-slice to change size to 0, leaves capacity 
// unchanged.  Does not clear values, but this is usually harmless; if want to actually destroy old 
// values (eg for security), use New instead (for security, overkill is reasonable).
func (base *ArrayList) Clear() {
    base.Ary = base.Ary[0:0] // go-newbie note: '[0:0]' looks like a slice of 1, but is actually empty.
    return
}

// Append method emulates the append version of Java ArrayList.add; uses Go built-in 'append', which
// handles capacity increase if needed; always returns 'true', since that is the behaviour in Java.
func (base *ArrayList) Append(val Obj) bool {
    base.Ary = append(base.Ary, val)
    return true // ("as per the general contract of Collection.add")
}

// Insert method emulates the insert version of Java ArrayList.add; uses Go built-in 'append', which
// handles capacity increase if needed; inserts before entry already in position pos.
func (base *ArrayList) Insert(pos int, val Obj) {
//;    ((old version, minus closing brace))
//;    oldlen := len(base.Ary)
//;    base.Ary = append(base.Ary, 0) // expand first with dummy zero value
//;    if pos < oldlen {              // shift following entries before put new one
//;        copy(base.Ary[pos+1:], base.Ary[pos:])
//;    }
//;    base.Ary[pos] = val
//;    return
//;func (base *ArrayList) Insert(pos int, val Obj) {
    // example from reddit slice_support_and_containervector
    //a = append(a[:i], append([]T{x}, a[i:]...)...)
    base.Ary = append(base.Ary[:pos], append([]Obj{val}, base.Ary[pos:]...)...)
    return
}

// Get method emulates Java ArrayList.get, returns element at position pos (return signature is interface 
// Obj, but can use directly as underlying type, this has been tested).
func (base *ArrayList) Get(pos int) (val Obj) {
    val = base.Ary[pos]
    return
}

// Set method emulates Java ArrayList.set, overwrites element at position pos (underlying type determined 
// by param val; consistency not enforced); returns value removed, can ignore if not needed.
func (base *ArrayList) Set(pos int, val Obj) Obj {
    oldval := base.Ary[pos]
    base.Ary[pos] = val
    return oldval
}

// Remove method emulates indexed version of Java ArrayList.remove; returns value removed, can ignore if 
// value not needed.
func (base *ArrayList) Remove(pos int) Obj {
//;    ((old version, minus closing brace))
//;    oldlen := len(base.Ary)
//;    val := base.Ary[pos]
//;    copy(base.Ary[pos:], base.Ary[pos+1:])
//;    // needs to return reduced len, so 're-slice' base.Ary
//;    base.Ary = base.Ary[:oldlen-1]
//;    return val
//;func (base *ArrayList) Remove(pos int) Obj {
    // example from reddit slice_support_and_containervector
    //a = append(a[:i], a[i+1:]...)
    val := base.Ary[pos]
    base.Ary = append(base.Ary[:pos], base.Ary[pos+1:]...)
    return val
}

// EnsureCapacity method emulates Java ArrayList.ensureCapacity; if needed creates new array, copies 
// existing elements to new array, makes new array the base.Ary, then reslices to old length.
func (base *ArrayList) EnsureCapacity(minCapacity int) {
//;    ((old version, minus closing brace))
//;    oldlen := len(base.Ary)
//;    if minCapacity > cap(base.Ary) { // need to increase size by amt delta, then reslice to oldlen
//;        newSlice := make([]Obj, minCapacity)
//;        copy(newSlice, base.Ary)
//;        base.Ary = newSlice
//;        base.Ary = base.Ary[:oldlen]
//;    }
//;func (base *ArrayList) EnsureCapacity(minCapacity int) {
    // example from reddit slice_support_and_containervector
    //a = append(a, make([]T, j)...)
    oldlen := len(base.Ary)
    if minCapacity > cap(base.Ary) { // need to increase size by amt delta, then reslice to oldlen
        base.Ary = append(base.Ary, make([]Obj, minCapacity - cap(base.Ary))...)
        base.Ary = base.Ary[:oldlen]
    }
}

// IsEmpty method emulates Java ArrayList.isEmpty, returns true if this list contains no elements.
func (base *ArrayList) IsEmpty() bool {
    return (len(base.Ary) <= 0)
}

// Size method emulates Java ArrayList.size, returns size of slice.  Also implemented Cap to return slice 
// capacity and SizeCap to return both at once; not in Java, but makes sense for Go.
func (base *ArrayList) Size() int {
    return len(base.Ary)
}

// Cap method returns slice capacity; not in Java, but makes sense for Go.
func (base *ArrayList) Cap() int {
    return cap(base.Ary)
}

// SizeCap method returns both slice size and slice capacity; not in Java, but makes sense for Go.
func (base *ArrayList) SizeCap() (int, int) { // convenience method
    return len(base.Ary), cap(base.Ary)
}

// TrimToSize method emulates Java ArrayList.trimToSize, trims slice cap to slice len; actually saves old
// values to a tmp array, makes a new one as base.Ary, and copies old values to new -- suggest to use 
// sparingly, as TrimToSize copies array twice and will take time on a really large ArrayList, but has 
// little value on a small one.
func (base *ArrayList) TrimToSize() {
    oldlen := len(base.Ary)
    tmpa := make([]Obj, oldlen)
    copy(tmpa, base.Ary)
    base.Ary = make([]Obj, oldlen, oldlen) // two copies; open to suggestions if a better way
    copy(base.Ary, tmpa)
    return
}

// InsertAll method replaces offset version of Java's addAll; called by AppendAll so must handle offset to 
// just beyond end.  Also, must panic if input ary nil, and return false if input ary empty.
func (base *ArrayList) InsertAll(pos int, ary []Obj) bool {
//;    ((old version, minus closing brace))
//;    if ary == nil {
//;        panic("ArrayList.InsertAll input array nil")
//;    }
//;    // check capacity sufficient before shift, expand if needed
//;    addlen := len(ary)
//;    if addlen < 1 {
//;        return false
//;    } // fail silently, empty insert could be a design choice
//;    oldlen := len(base.Ary)
//;    newmin := oldlen + addlen
//;    if cap(base.Ary) < newmin {
//;        base.EnsureCapacity(newmin)
//;    }
//;    // always need append loop
//;    for ix := 0; ix < addlen; ix++ {
//;        base.Ary = append(base.Ary, 0) // expand first with dummy zero value
//;    }
//;    // now can do shift
//;    if pos < oldlen { // shift following entries before put new one
//;        copy(base.Ary[pos+addlen:], base.Ary[pos:])
//;    }
//;    // and a loop to copy ary contents
//;    for ix, elem := range ary {
//;        base.Ary[pos+ix] = elem
//;    } // for
//;    return true
//;func (base *ArrayList) InsertAll(pos int, ary []Obj) bool {
    if ary == nil {
        panic("ArrayList.InsertAll input array nil") // required by spec
    }
    // example from reddit slice_support_and_containervector
    //a = append(a[:i], append(b, a[i:]...)...)    //// eliminates 2 loops
    base.Ary = append(base.Ary[:pos], append(ary, base.Ary[pos:]...)...)
    return true // required by spec
}

// AppendAll method replaces 1-param version of Java's addAll; calls InsertAll with pos just beyond end of 
// existing ArrayList.  Also, must panic if input ary nil, and return false if input ary empty.
func (base *ArrayList) AppendAll(ary []Obj) bool {
    if ary == nil {
        panic("ArrayList.AppendAll input array nil")
    }
    return base.InsertAll(len(base.Ary), ary)
}

// ToArray method emulates param version of Java's ArrayList.ToArray(arrayT) for interface Obj (both input 
// and output are []Obj, but elements will usually be of some basic Go type); returns updated input array 
// if large enough, or a new array of same type if base.Ary is larger than input array, with all elements 
// copied (in order found) from base.Ary.  Access to nth element is a simple 'elem := result[n]', and each 
// element can be used directly as its underlying type.  Type consistency of elements not enforced, but 
// see doc.go for type-safe suggestions.
// Note: ignored Java spec feature of nul (Go nil) after last element, as Go slice makes that unreachable.
func (base *ArrayList) ToArray(ary []Obj) []Obj {
//;    ((old version, minus closing brace))
//;    alen := len(base.Ary)
//;    if alen > cap(ary) {
//;        ary = make([]Obj, alen)
//;    }
//;    ary = ary[:alen] // reslice in case len < cap or len > alen
//;    for ix, elem := range base.Ary {
//;        ary[ix] = elem.(Obj)
//;    }
//;    return ary
//;func (base *ArrayList) ToArray(ary []Obj) []Obj {
    alen := len(base.Ary)
    if alen > cap(ary) {
        ary = make([]Obj, alen)
    }
    ary = ary[:alen] // reslice in case len < cap or len > alen
    // example from reddit slice_support_and_containervector
    //b = make([]T, len(a)); copy(b, a)    //// eliminates a loop
    copy(ary, base.Ary)
    return ary
}

// ToArrayNew method emulates non-param version of Java's ArrayList.ToArray(arrayT) for interface Obj; 
// creates and returns a new array of interface Obj with all elements copied (in order found) from 
// base.Ary.  Access to nth element is a simple 'elem := result[n]', and each element can be used directly 
// as its underlying type.  Type consistency not enforced, but see doc.go for type-safe suggestions.
// Note: ignored Java spec feature of nul (Go nil) after last element, as Go slice makes that unreachable.
func (base *ArrayList) ToArrayNew() []Obj {
    return base.ToArray(make([]Obj, len(base.Ary))) // dimension optional but might as well
}

// The findObj method is a local utility, not exported, in support of Contains, IndexOf, LastIndexOf & 
// RemoveObj.  A refactoring of common code, direction allows it to search forward or backward.
func (base *ArrayList) findObj(direction int, obj Obj) (pos int) { // name pos not used, is documentation
    end := len(base.Ary) - 1
    if direction > 0 { // search fwd from 0
        for ix, elem := range base.Ary {
            if elem == obj {
                return ix
            } // found
        }
        return -1 // not found
    } // else search backward from end
    for ix, _ := range base.Ary {
        if base.Ary[end-ix] == obj {
            return end - ix
        } // found
    }
    return -1 // not found
} // findObj

// Contains method emulates Java's ArrayList.contains; calls local utility method findObj; returns true if
// item is found, else false.
func (base *ArrayList) Contains(obj Obj) bool {
    if base.findObj(1, obj) >= 0 {
        return true
    }
    return false
}

// IndexOf method emulates Java's ArrayList.indexOf; returns position where item first found, or -1 if 
// not found.
func (base *ArrayList) IndexOf(obj Obj) int {
    return base.findObj(1, obj)
}

// LastIndexOf method emulates Java's ArrayList.lastIndexOf; returns position where item last found, or 
// -1 if not found.
func (base *ArrayList) LastIndexOf(obj Obj) int {
    return base.findObj(-1, obj)
}

// RemoveObj method emulates Java's ArrayList.remove with object param, does not return value (calls 
// Remove with position param); returns false if object not found.
func (base *ArrayList) RemoveObj(obj Obj) bool {
    pos := base.findObj(1, obj)
    if pos > 0 {
        base.Remove(pos) // discards returned value
        return true
    } // no else case, silent if not found
    return false
}

// RemoveRange method emulates Java's ArrayList.remove with 2 int params, does not return value; silent
// return if toIndex <= fromIndex; shifts elements and reslices w/o changing capacity.
func (base *ArrayList) RemoveRange(fromIndex, toIndex int) {
//;    ((old version, minus closing brace))
//;    if toIndex <= fromIndex {
//;        return
//;    } // silent if no action; required by spec
//;    newlen := len(base.Ary) + fromIndex - toIndex
//;    copy(base.Ary[fromIndex:], base.Ary[toIndex:])
//;    base.Ary = base.Ary[:newlen] // capacity not affected
//;    return
//;func (base *ArrayList) RemoveRange(fromIndex, toIndex int) {
    if toIndex <= fromIndex {
        return
    } // silent if no action; required by spec
    // example from reddit slice_support_and_containervector
    //a = append(a[:i], a[j:]...)
    base.Ary = append(base.Ary[:fromIndex], base.Ary[toIndex:]...)
    return
}

// Copy method replaces but does *not* emulate Java ArrayList.clone (which some advise not to use in Java),
// returns a new ArrayList with elements copied from existing ArrayList; unlike Java ArrayList.clone, 
// changes to values in the copy will not change values in the original (for 'shallow copy' in Go, simply 
// use ':=' or '=', so there's no need for a method Clone to do that).
func (base *ArrayList) Copy() (AL *ArrayList) {
//;    ((old version, minus closing brace))
//;    alen := len(base.Ary)
//;    tmp := make([]Obj, alen, alen) // create with size and capacity both alen
//;    for ix, elem := range base.Ary {
//;        tmp[ix] = elem
//;    }
//;    AL = &ArrayList{tmp}
//;    return
//;func (base *ArrayList) Copy() (AL *ArrayList) {
    alen := len(base.Ary)
    tmp := make([]Obj, alen, alen) // create with size and capacity both alen
    // example from reddit slice_support_and_containervector
    //b = make([]T, len(a)); copy(b, a)    //// eliminates a loop
    copy(tmp, base.Ary)
    AL = &ArrayList{tmp}
    return
}

// [start,end)
func (base *ArrayList) SubList(start,end int) *ArrayList {
    val := base.Ary[start:end]
    return &ArrayList{val}
}