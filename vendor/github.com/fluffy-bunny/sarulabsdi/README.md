# sarulabsdi  

This is a baseline drop of [v2.4.2 of github.com/sarulabsdi](https://github.com/sarulabs/di/releases/tag/v2.4.2)  

## New Features

The following are non breaking changes to sarulabs/di.  

## SafeGetByType  

```go
// SafeGetByType retrieves the last object added from the Container.
// The object has to belong to this scope or a more generic one.
// If the object does not already exist, it is created and saved in the Container.
// If the object can not be created, it returns an error.
SafeGetByType(rt reflect.Type) (interface{}, error)
```

## GetByType  

```go
// GetByType is similar to SafeGetByType but it does not return the error.
// Instead it panics.
GetByType(rt reflect.Type) interface{}
```

## SafeGetManyByType  

```go
// SafeGetManyByType retrieves an array of objects from the Container.
// The objects have to belong to this scope or a more generic one.
// If the objects do not already exist, it is created and saved in the Container.
// If the objects can not be created, it returns an error.
SafeGetManyByType(rt reflect.Type) ([]interface{}, error)
```

## GetManyByType  

```go
// GetManyByType is similar to SafeGetManyByType but it does not return the error.
// Instead it panics.
GetManyByType(rt reflect.Type) []interface{}
```

Just like dotnetcore, I would like to register a bunch of services that all implement the same interface.  I would like to get back an array of all registered objects, and to do that I need to ask for them by type

For efficiency, type validation is done during registration.  After that it should just be lookups.  

When you do a ```GetByType```, don't keep calling reflect to get a type of what is a compliled type.  do it once and put it in a map.  

### Registration

type [DEF](https://github.com/fluffy-bunny/sarulabsdi/blob/8a200c4fa3aefa0a28ddc66739aac1631f2a95aa/definition.go#L19) struct  

```go
// Def contains information to build and close an object inside a Container.
type Def struct {
  Build            func(ctn Container) (interface{}, error)
  Close            func(obj interface{}) error
  Name             string //[ignored] if Type is used this is overriden and hidden.
  Scope            string
  Tags             []Tag
  Type             reflect.Type //[optional] only if you want to claim that this object also implements these types.
  ImplementedTypes TypeSet
  Unshared         bool
}
```

Two new fields where added to the Def struct.  

```go
  Type             reflect.Type
  ImplementedTypes TypeSet
```

```Type```:              is the type of the object being registered.  
```ImplementedTypes```:  is the types that this object either is or implements  

Only the ```Type``` option is needed to register and it is automatically added as an implemented type.   The ```ImplementedTypes``` option is primarily for claiming that the added type supports a given set of interfaces.  If this option is added, the original ```Type``` is automatically added to the ```ImplementedTypes``` set.  

The following [code](https://github.com/fluffy-bunny/sarulabsdi/blob/909f303f513ce84953164cc78b311a57ae959544/builder.go#L90) will return an error if the added type **DOES NOT** implemented the ```ImplementedTypes``` that were claimed.  

```go
type ISomething interface {
  GetName() string
  SetName(name string)
}

type Service struct {
  name            string
}
// SetName ...
func (s *Service) SetName(in string) {
  s.name = in
}

// SetName ...
func (s *Service) GetName() string {
  return s.name
}

func AddTransientService(builder *di.Builder) {
  log.Info().Msg("IoC: AddTransientService")

  types := di.NewTypeSet()
  inter := di.GetInterfaceReflectType((*exampleServices.ISomething)(nil))
  types.Add(inter)
  
  builder.Add(di.Def{
    Name:             "[overidden and hidded if added by Type, can be empty]",
    Scope:            di.App,
    ImplementedTypes: types, //[optional] only if you want to claim that this object also implements these types.
    Type:             reflect.TypeOf(&Service{}),
    Unshared: true,
    Build: func(ctn di.Container) (interface{}, error) {
      service := &Service{}
      return service, nil
    },
  })
}
```

### Retrieval  

```go
// please put this into a lookup map
inter := di.GetInterfaceReflectType((*exampleServices.ISomething)(nil))

dd := ctn.GetManyByType(inter)
for _, d := range dd {
  ds := d.(exampleServices.ISomething)
  ds.SetName("rabbit")
  log.Info().Msg(ds.GetName())
}
```

## NON-Interface Types  

When you register a non-interface type, like a struct, we use the following;

```go
Type: reflect.TypeOf(&mockObject{}),
```

but to retrieve the object, a .Elem() is needed for the lookup.

```go
rt := reflect.TypeOf(&mockObject2{}).Elem()
```

```go
func TestTypedObject_Unshared_OneAdded_FailedRetrieve(t *testing.T) {
  b, _ := NewBuilder()

  // add mockObject
  b.Add(Def{
    Type: reflect.TypeOf(&mockObject{}),
    Build: func(ctn Container) (interface{}, error) {
      return &mockObject{}, nil
    },
      Unshared: true,
  })

  var app = b.Build()

  // try to get mockObject2
  rt := reflect.TypeOf(&mockObject2{}).Elem()
  _, err := app.SafeGetByType(rt)
  require.NotNil(t, err)

  // try to get IGetterSetter
  rt = GetInterfaceReflectType((*IGetterSetter)(nil))
  _, err = app.SafeGetByType(rt)
  require.NotNil(t, err)
}
```

## Default BuildByType

Its more expensive but you don't need a BUILD func to create the objects if the object and its dependencies are all registred by type.  

```go
type mockObject3 struct {
  GetterSetter  IGetterSetter   `inject:""`
  GetterSetters []IGetterSetter `inject:""`
}
```

This assumes that there exists a registration for ```IGetterSetter```.  

```go
b.Add(Def{
    Type:     reflect.TypeOf(&mockObject3{}),
    Unshared: true,
  })
```  

## Maps and Slices 

### Map

```go
func TestTypedObjects_map_type(t *testing.T) {
    type MyMapType map[string]string
    rtPtr := reflect.TypeOf(&MyMapType{})
    rtElem := rtPtr.Elem()

    b, _ := NewBuilder()
    b.Add(Def{
        Type:     rtPtr,
        Unshared: true,
        Build: func(ctn Container) (interface{}, error) {
            result := MyMapType{
                "test": "dog",
            }
            return &result, nil
        },
    })
    var app = b.Build()
    obj := app.GetByType(rtElem).(*MyMapType)
    assert.NotNil(t, obj)
    dObj := *obj

    assert.Equal(t, "dog", dObj["test"])
}
```

### Slice

```go
func TestTypedObjects_slice_type(t *testing.T) {
    type MySliceType []string
    rtPtr := reflect.TypeOf(&MySliceType{})
    rtElem := rtPtr.Elem()

    b, _ := NewBuilder()
    b.Add(Def{
        Type:     rtPtr,
        Unshared: true,
        Build: func(ctn Container) (interface{}, error) {
            result := MySliceType{
                "dog",
            }
            return &result, nil
        },
    })
    var app = b.Build()
    obj := app.GetByType(rtElem).(*MySliceType)
    assert.NotNil(t, obj)
    dObj := *obj

    assert.Equal(t, "dog", dObj[0])
}
```

## Ctor or Not to Ctor

If your typed object contains the following method it will get called during the build AFTER all services have been injected  

```go
type mockObjectWithCtor struct {
        CtorCalled bool
    }

func (m *mockObjectWithCtor) Ctor() {
    m.CtorCalled = true
}
```
