# GOSON

## goson is a json marshal/unmarshal tool with function

## Useage

- you can use it like json.Marshal

```go
    goson.Marshal()
```

- and json.Unmarshal

```go
    goson.Unmarshal
```

- you can use set/get value on marshal/unmarshal
  
```go
type Example struct {
    version              string       `json:"version"`
}

func (t *Example) SetVersion(version string) {
    t.version = version
}

func (t Example) Version() string {
    return t.version
}
```

- and you can define the custom function to set/get value on marshal/unmarshal
  
```go
type Example struct {
    version              string       `json:"version" json-getter:"Version" json-setter:"SetMyVersion"`
}


func (t *Example) SetMyVersion(version string) {
    t.version = version
}

func (t Example) Version() string {
    return t.version
}
```
