package btree

type BTree struct {
  root *Node
  order int
}

func NewBTree(order int) BTree {
  root := NewNode(make([]*Key, 0, order), make([]*Node, 0, order + 1), nil)
  return BTree{
    order: order,
    root: &root}
}

func (btree *BTree) Add(index int, value interface{}) {
  key := Key{Key: index, Value: value}
  node := btree.root

  for node != nil {
    if len(node.keys) >= btree.order {
      if len(node.keys) == btree.order {
        btree.SplitNode(node)
        node = node.parent
        continue
      }
    }

    targetIndex := node.FindTargetIndex(index)
    // we are at leaf
    if (len(node.children) == 0) {
      node.InsertKeyAtIndex(targetIndex, &key)
      node = nil
    } else {
      node = node.children[targetIndex]
    }
  }
}

func (btree *BTree) SplitNode(node *Node) {
  splitAt := btree.order / 2
  newChild := NewNode(make([]*Key, 0, btree.order), make([]*Node, 0, btree.order+1), nil)

  splitKey := node.keys[splitAt]
  newChild.keys = append(make([]*Key, 0), node.keys[:splitAt]...)
  node.keys = append(make([]*Key, 0), node.keys[splitAt+1:]...)

  if len(node.children) > 0 {
    newChild.children = append(make([]*Node, 0), node.children[:splitAt+1]...)
    node.children = append(make([]*Node, 0), node.children[splitAt+1:]...)

    for _, child := range newChild.children {
      child.parent = &newChild
    }
  }

  // if we are splitting root - make new root
  if node.parent == nil {
    //fmt.Println("New Root", splitKey)
    newRoot := NewNode(make([]*Key, 0, btree.order), make([]*Node, 0, btree.order + 1), nil)

    newRoot.children = append(newRoot.children, &newChild, node)
    newRoot.keys = append(newRoot.keys, splitKey)

    btree.root = &newRoot
    newChild.parent = &newRoot
    node.parent = &newRoot
  } else {
    newChild.parent = node.parent
    node.parent.InsertKeyAndChild(splitKey, &newChild)
  }
}

func (btree *BTree) Delete(needle int) {
  node := btree.root

  for node != nil {
    targetIndex := node.FindTargetIndex(needle)

    if node.keys[targetIndex].Key == needle {
      // delete
    }

    if len(node.children) > 0 {
      
    }
  }
}

func (btree *BTree) Find(needle int) (key *Key) {
  return btree.root.SearchKey(needle)
}

func (btree *BTree) PrintTree() {
  btree.root.PrintRec(0)
}
