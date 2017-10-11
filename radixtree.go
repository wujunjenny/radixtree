// radixtree project radixtree.go
package radixtree

import "github.com/armon/go-radix"

type Tree struct {
	_tree *radix.Tree
}

// WalkFn is used when walking the tree. Takes a
// key and value, returning if iteration should
// be terminated.
type WalkFn func(s string, v interface{}) bool

// New returns an empty Tree
func New() *Tree {
	tree := &Tree{}
	tree._tree = radix.New()
	return tree
}

// Insert is used to add a newentry or update
// an existing entry. Returns if updated.
func (t *Tree) Insert(key string, v interface{}) (interface{}, bool) {
	return t._tree.Insert(key, v)
}

// Delete is used to delete a key, returning the previous
// value and if it was deleted
func (t *Tree) Delete(s string) (interface{}, bool) {
	return t._tree.Delete(s)
}

// DeletePrefix is used to delete the subtree under a prefix
// Returns how many nodes were deleted
// Use this to delete large subtrees efficiently
func (t *Tree) DeletePrefix(s string) int {
	return t._tree.DeletePrefix(s)
}

// Get is used to lookup a specific key, returning
// the value and if it was found
func (t *Tree) Get(s string) (interface{}, bool) {
	return t._tree.Get(s)
}

// Len is used to return the number of elements in the tree
func (t *Tree) Len() int {
	return t._tree.Len()
}

// LongestPrefix is like Get, but instead of an
// exact match, it will return the longest prefix match.
func (t *Tree) LongestPrefix(s string) (string, interface{}, bool) {
	return t._tree.LongestPrefix(s)
}

// Minimum is used to return the minimum value in the tree
func (t *Tree) Minimum() (string, interface{}, bool) {
	return t._tree.Minimum()
}

// Maximum is used to return the maximum value in the tree
func (t *Tree) Maximum() (string, interface{}, bool) {
	return t._tree.Maximum()
}

// Walk is used to walk the tree
func (t *Tree) Walk(fn WalkFn) {
	t._tree.Walk(func(s string, v interface{}) bool { return fn(s, v) })
}

// WalkPath is used to walk the tree, but only visiting nodes
// from the root down to a given leaf. Where WalkPrefix walks
// all the entries *under* the given prefix, this walks the
// entries *above* the given prefix.
func (t *Tree) WalkPath(path string, fn WalkFn) {
	//max match first callback
	cbs := make([]func() bool, 0, 20)
	cb := func(k string, v interface{}) (rt bool) {
		f := func() bool {
			return fn(k, v)
		}
		cbs = append(cbs, f)
		return
	}
	t._tree.WalkPath(path, cb)

	for i := len(cbs) - 1; i >= 0; i-- {

		if cbs[i]() {
			return
		}
	}
}

// WalkPrefix is used to walk the tree under a prefix
func (t *Tree) WalkPrefix(prefix string, fn WalkFn) {
	cbs := make([]func() bool, 0, 20)
	cb := func(k string, v interface{}) (rt bool) {
		f := func() bool {
			return fn(k, v)
		}
		cbs = append(cbs, f)
		return
	}

	t._tree.WalkPrefix(prefix, cb)

	for i := len(cbs) - 1; i >= 0; i-- {

		if cbs[i]() {
			return
		}
	}
}
