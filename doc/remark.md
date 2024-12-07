## remark

### 端口命名规则
```
rpc 以900x 命名 比如 9001
api 以800x 命名 比如 8001

相同的服务 api和rpc要对应 比如api 8002 则 rpc对应 9002
```

### 项目启动
`./start.sh`

### 缓存
在测试时，查数据我们直接在redis中删掉它来测试

### 元数据
元数据不会自动在各个rpc之间自动往下传递，比如api 调用A rpc A又去调用了B rpc api发起时的元数据不会自动往下传递，必须在rpc服务发起时候自动添加才行


### 消息队列
1. 执行`./start_mq.sh`启动消息队列
2. kafka是在`docker-compose.yml`中启动的，启动后通过`http://localhost:8080/` 查看kafkaq消息情况
