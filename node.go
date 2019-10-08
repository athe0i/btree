package btree

import (
  "fmt"
)

type Node struct {
  parent *Node
  keys []*Key
  children []*Node
  order int}

func NewNode(keys []*Key, children []*Node, parent *Node) Node {
  return Node {
    keys: keys,
    children: children,
    parent: parent}
}

func (node *Node) InsertKeyAndChild(key *Key, child *Node) {
  targetIndex := node.FindTargetIndex(key.Key)

  node.InsertKeyAtIndex(targetIndex, key)
  node.InsertChildAtIndex(targetIndex, child)
}

func (node *Node) InsertKeyAtIndex(index int, key *Key) {
  newKeys := make([]*Key, 0, len(node.keys) + 1)
  newKeys = append(newKeys, node.keys[:index]...)
  newKeys = append(newKeys, key)
  newKeys = append(newKeys, node.keys[index:]...)
  node.keys = newKeys
}

func (node *Node) InsertChildAtIndex(index int, child *Node) {
  newChildren := make([]*Node, 0, len(node.children))
  newChildren = append(newChildren, node.children[:index]...)
  newChildren = append(newChildren, child)
  newChildren = append(newChildren, node.children[index:]...)
  node.children = newChildren
}

func (node *Node) FindTargetIndex(index int) (targetIndex int) {
  targetIndex = 0

  for i := 0; i < len(node.keys); i++ {
    if index > node.keys[i].Key {
      targetIndex = i+1
    }

    endInd := len(node.keys) - (i+1)
    if index > node.keys[endInd].Key {
       targetIndex = endInd + 1
       break
    }
  }

  return
}

func (node *Node) SearchKey(needle int) (key *Key) {
  for i:=0;i<len(node.keys);i++ {
    endInd := len(node.keys) - (i + 1)
    if needle == node.keys[endInd].Key {
      return node.keys[endInd]
    }
    if needle > node.keys[endInd].Key && len(node.children) > 0 {
      return node.children[endInd+1].SearchKey(needle)
    }

    // return if found
    if node.keys[i].Key == needle {
      return node.keys[i]
    }

    // jump over key if we didn't reach the end
    if needle > node.keys[i].Key && i < len(node.keys) - 1  {
      continue
    }

    if len(node.children) > 0 {
      return node.children[i].SearchKey(needle)
    }
  }

  return nil
}

func (node *Node) GetKeysRecursive() (keys []*Key) {
  keys = make([]*Key, 0)

  if len(node.children) > 0 {
    keys = node.children[0].keys[:]
  }

  for i, key := range node.keys {
    keys = append(keys, key)
    if len(node.children) > 0 {
      keys = append(keys,  node.children[i+1].GetKeysRecursive()...)
    }
  }

  return
}

func (node *Node) PrintKeys() {
  fmt.Println()
  for _, key := range node.keys {
    fmt.Print(key.Key, " ")
  }
  fmt.Println()
}

func (node *Node) PrintRec(level int) {
  fmt.Println("Node! Depth", level)
  node.PrintKeys()

  for i, child := range node.children {
    if i > len(node.keys) - 1 {
        fmt.Println(">", node.keys[i-1])
    }

    if i < len(node.keys) {
        fmt.Println("<", node.keys[i])
    }

    child.PrintRec(level + 1)
  }

  return
}
