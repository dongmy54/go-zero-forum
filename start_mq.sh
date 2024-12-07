#!/bin/bash

# 创建日志目录
mkdir -p log

# 检查 kafka 是否运行
check_kafka() {
    nc -z localhost 9092 > /dev/null 2>&1
    return $?
}

# 启动 kafka
start_kafka() {
    echo "正在启动 Kafka..."
    docker-compose up -d >> log/kafka.log 2>&1
    
    # 等待 Kafka 完全启动
    for i in {1..30}; do
        if check_kafka; then
            echo "Kafka 已成功启动"
            return 0
        fi
        echo "等待 Kafka 启动... $i/30"
        sleep 2
    done
    echo "Kafka 启动超时"
    return 1
}

# 检查并启动 Kafka
if ! check_kafka; then
    start_kafka
    if [ $? -ne 0 ]; then
        echo "Kafka 启动失败，退出程序"
        exit 1
    fi
else
    echo "Kafka 已在运行中"
fi

# 检查 mq.go 是否已运行
MQ_PID=$(pgrep -f "mq.go|mq$")
if [ ! -z "$MQ_PID" ]; then
    echo "发现 MQ 服务已运行，PID: $MQ_PID"
    echo "正在停止旧的 MQ 服务..."
    kill -15 $MQ_PID
    sleep 2
    
    # 确保进程已经终止
    if ps -p $MQ_PID > /dev/null; then
        echo "强制终止 MQ 服务..."
        kill -9 $MQ_PID
    fi
fi

# 清空 mq.log 文件（如果存在）
if [ -f "log/mq.log" ]; then
    echo "清空 MQ 日志文件..."
    : > log/mq.log
fi

# 启动 mq.go
echo "正在启动 MQ 服务..."
go run service/mq/mq.go >> log/mq.log 2>&1 &

echo "MQ 服务已启动，日志文件：log/mq.log" 