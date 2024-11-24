#!/bin/bash

# 设置项目目录
Pro_Dir=`pwd` # 注意不能有空格
echo "Project directory: $Pro_Dir"

LOG_DIR="`pwd`/log"
# 创建日志目录（如果不存在）
mkdir -p $LOG_DIR

# 定义服务名称列表
SERVICES=(
    "comment"
    "user"
    # 可以在这里继续添加其他服务
)

# 创建日志文件（rpc 和 api）
for SERVICE in "${SERVICES[@]}"; do
    for TYPE in "rpc" "api"; do
        LOG_FILE="$LOG_DIR/$SERVICE.$TYPE.log"
        if [ ! -f "$LOG_FILE" ]; then
            touch "$LOG_FILE"
            echo "Created log file: $LOG_FILE"
        fi
    done
done

# 停止之前的服务
echo "Killing existing services..."
# 先停止所有相关进程，基于进程名（你可以根据需要修改进程名）
pkill -f "go run service/comment/cmd/rpc/comment.go"
pkill -f "go run service/comment/cmd/api/comment.go"

pkill -f "go run service/user/api/user.go"
pkill -f "go run service/user/rpc/user.go"

############# 启动comment 服务 #############
# 启动 comment rpc 服务
echo "Starting comment rpc service..."
cd service/comment/cmd/rpc || exit

go run comment.go > "$LOG_DIR/comment.rpc.log" 2>&1 &
COMMENT_RPC_PID=$!
echo "Comment RPC service started with PID $COMMENT_RPC_PID"

# 启动 comment api 服务
echo "Starting comment api service..."
cd ../api || exit


go run comment.go > "$LOG_DIR/comment.api.log" 2>&1 &
COMMENT_API_PID=$!
echo "Comment API service started with PID $COMMENT_API_PID"

############# 启动用户 服务 #############
echo "Starting user rpc service..."
echo `pwd`
cd "$Pro_Dir/service/user/rpc" || exit

go run user.go > "$LOG_DIR/user.rpc.log" 2>&1 &
USE_RPC_PID=$!
echo "user RPC service started with PID $USE_RPC_PID"

# 启动 user api 服务
echo "Starting user api service..."
cd ../api || exit

go run user.go > "$LOG_DIR/user.api.log" 2>&1 &
USE_API_PID=$!
echo "user API service started with PID $USE_API_PID"

# 等待服务启动
wait $COMMENT_RPC_PID
wait $COMMENT_API_PID
wait $USE_RPC_PID
wait $USE_API_PID  

echo "All services are running."
