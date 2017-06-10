# matchers

Have you ever been in this situation, scratching your head trying to compare some complex data structures?

```
Expected
    <main_test.User>: {
        Name: "david duchovny",
        Address: "santa monica",
        Age: 56,
        Badge: {ID: 666, Color: "shiny"},
    }
to equal
    <main_test.User>: {
        Name: "david duchovny",
        Address: "santa monica",
        Age: 56,
        Badge: {ID: 666, Color: "shiny"},
    }
```

Well, you've come to the right place. The place that can answer ALL your questions 
(so long as they have to do with figuring out why values aren't equal). Because now you can get *actually helpful* errors, like this one:


```
Expected
    <main_test.User>: {
        Name: "david duchovny",
        Address: "santa monica",
        Age: 56,
        Badge: {ID: 666, Color: "shiny"},
    }
to deep equal
    <main_test.User>: {
        Name: "david duchovny",
        Address: "santa monica",
        Age: 56,
        Badge: {ID: 666, Color: "shiny"},
    }
mismatch at .Badge: type mismatch main_test.FBIBadge vs *interface {}; obtained main_test.FBI{ID:666, Color:"shiny"}; expected (*interface {})(0xc420130728)
```

D'oh! I hate it when I use a `FBIBadge` when I mean to use an `*interface{}`! (This is actually one of the more cryptic explanations we can get, but basically one instance has a pointer reference for its Badge slot, and the other has a value reference. This happened to me recently, and it sucked.)

That's right, this library provides only a single [Gomega](https://onsi.github.io/gomega/) matcher, but it's going to be
the last matcher you ever need.

To get started:

```
go get -u github.com/benmoss/matchers
```

and just swap your `Expect(...).To(Equal(...))` with a lil
`Expect(...).To(matchers.DeepEqual(...))`, and you're off to the races!

It adapts the DeepEqual function from [juju/testing](https://github.com/juju/testing) so has all the same
caveats as [their implementation](https://godoc.org/github.com/juju/testing/checkers#DeepEqual).
