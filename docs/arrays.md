# Arrays

`gonfique` assignes the necessary slice type to arrays. It works best when all items of an array possess the same schema or at least compatible schemas. Type assignment occurs as below when items have not same but compatible schemas:

```yaml
# input
- action: foo
  foo-details: ""
- action: bar
  bar-details: ""
```

```go
// output
[]struct {
  Action     string
  FooDetails string
  BarDetails string
}
```

> [!IMPORTANT]
> Slice type gets defined as `[]any` if shared keys have different type values. Like `detail` has given `int` and `string` values below:
>
> ```yaml
> - action: ""
>   detail: 0
> - action: ""
>   detail: ""
> ```
