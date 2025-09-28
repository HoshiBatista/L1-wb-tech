package main

import "fmt"

type StringSet struct {
    items map[string]bool
}

func NewStringSet() *StringSet {
    return &StringSet{
        items: make(map[string]bool),
    }
}

func (s *StringSet) Add(items ...string) {
    for _, item := range items {
        s.items[item] = true
    }
}

func (s *StringSet) Items() []string {
    result := make([]string, 0, len(s.items))

    for item := range s.items {
        result = append(result, item)
    }

    return result
}

func (s *StringSet) String() string {
    return fmt.Sprintf("%v", s.Items())
}

func main() {
    sequence := []string{"cat", "cat", "dog", "cat", "tree"}
    
    set := NewStringSet()
    set.Add(sequence...)
    
    fmt.Printf("Set: %s\n", set)
}