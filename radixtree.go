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

// WalkPath is used to walk the tree, but only visiting nodes
// from the root down to a given leaf. Where WalkPrefix walks
// all the entries *under* the given prefix, this walks the
// entries *above* the given prefix.
func (t *Tree) WalkPath(path string, fn WalkFn) {
	//max match first callback
	cbs := []func() bool{}
	cb := func(k string, v interface{}) (rt bool) {
		f := func() bool {
			return fn(k, v)
		}
		cbs = append([]func() bool{f}, cbs...)
		return
	}
	t._tree.WalkPath(path, cb)

	for _, f := range cbs {

		if f() {
			return
		}
	}
}
