### other
1.metdata传递过程中默认会忽略大小，统一当作小写处理。

如何解决？
```go
//定一个切片
metaKeys := []string{"UserId", "RoleId", "GroupId"}

// 处理大小写转换
// 处理大小写转换
func GetMetaDataKey(oldKey string) string {
	//定一个切片
	metaKeys := []string{"UserId", "RoleId", "GroupId"}
	index := slices.IndexFunc(metaKeys, func(str string) bool {
		return strings.ToLower(str) == oldKey
	})

	if index != -1 {
		return metaKeys[index]
	} else {
		return oldKey
	}
}
```

2. Rpc传入model中的ctx是可以有值的，可以认为context在一个进程中是有效的。

