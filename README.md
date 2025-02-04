# heap

`heap` is like stdlib heap but more convenient for some cases.
It supports only slice for backing storage but doesn't require to implement interfaces.

## Usage

Usage is pretty straightforward.

```
    h := heap.Make[*Element](func(a []*Element, i, j int) bool {
            return a[i].X < a[j].X
    })

    h.Push(&Element{X: 5})
    h.Push(&Element{X: 2})
    h.Push(&Element{X: 10})

    _ = h.Pop() // &Element{X: 2}

    h.Data[0].X = 20 // 5 -> 20
    h.Fix(0)

    h.Del(0) // 10

    _ = h.Len() // 1
```
