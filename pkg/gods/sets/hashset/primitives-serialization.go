package hashset

//go:generate genny   -pkg $GOPACKAGE        -in=../../../../internal/genny/gods/sets/hashset/serialization.go -out=gen-$GOFILE gen "KeyType=string,int,int32,int64,uint,uint32,uint64,float32,float64"
