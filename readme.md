### readme
```shell
# api生成
添加别名`alias genapi='goctl api go -api *.api -dir ../ --style=goZero';` 
# 然后到xx.api目录下执行

# sql生成
# 到deploy/script目录下执行
# ./genModel.sh dabaseName tableName
# 比如: ./genModel.sh forum comment

# 将生成的model文件放到 对应的model下，并改其中包名为model
```


