package btree

import (
  "testing"
  "math/rand"
  "fmt"
)

func TestBTree(t *testing.T) {

  testSet, btree := initBench(5, 100);

  //btree.PrintTree()

  for _,i := range testSet {
    //r := rand.Intn(n)
    //fmt.Println("Add", i)
    find := btree.Find(i)

    if find == nil {
      t.Error("We didn't find value ", i)

      return
    }

    if find.Value.(string) != fmt.Sprintf("test%d", i) {
      t.Errorf("We've found something wrong! Expected %v got %v", "test5", find.Value.(string))

      return
    }
  }
}

func BenchmarkBTreeAdd(b *testing.B) {
  btree := NewBTree(20)
  testSet := make([]int, b.N)

  for i := 0; i < b.N; i++ {
    testSet[i] = rand.Intn(b.N*10)
  }

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    btree.Add(testSet[i], "test")
  }

}

func BenchmarkBTreeSearch(b *testing.B) {
  b.N = 5000000
  testSet, btree := initBench(50, b.N)

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    btree.Find(testSet[i])
  }
}

func initBench(order, n int) (testSet []int, btree BTree) {
  btree = NewBTree(order)
  testSet = make([]int, n)

  for i := 0; i < n; i++ {
    testSet[i] = rand.Intn(n*10)
    btree.Add(testSet[i], fmt.Sprintf("test%d", testSet[i]))
  }

  return
}
