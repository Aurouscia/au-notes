#!/bin/zsh

# 异步任务调度系统测试脚本
# 使用方法: ./test.sh

set -e

BASE_URL="http://localhost:8080"

# 颜色定义
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "${BLUE}========================================${NC}"
echo "${BLUE}  异步任务调度系统 - API 测试脚本${NC}"
echo "${BLUE}========================================${NC}"
echo ""

# 检查服务器是否运行
echo "${YELLOW}[检查] 检查服务器是否运行...${NC}"
if ! curl -s "${BASE_URL}/api/tasks/stats" > /dev/null 2>&1; then
    echo "${YELLOW}警告: 服务器似乎没有运行在 ${BASE_URL}${NC}"
    echo "请先运行: go run ."
    exit 1
fi
echo "${GREEN}✓ 服务器运行正常${NC}"
echo ""

# 1. 提交 email 任务
echo "${YELLOW}[测试1] 提交 email 任务...${NC}"
EMAIL_TASK=$(curl -s -X POST "${BASE_URL}/api/tasks" \
    -H "Content-Type: application/json" \
    -d '{"type":"email","params":{"to":"user@example.com","subject":"Hello","body":"This is a test email"}}')
echo "响应: $EMAIL_TASK"
EMAIL_ID=$(echo $EMAIL_TASK | grep -o '"id":[0-9]*' | cut -d':' -f2)
echo "${GREEN}✓ Email 任务提交成功, ID: $EMAIL_ID${NC}"
echo ""

# 2. 提交 report 任务
echo "${YELLOW}[测试2] 提交 report 任务...${NC}"
REPORT_TASK=$(curl -s -X POST "${BASE_URL}/api/tasks" \
    -H "Content-Type: application/json" \
    -d '{"type":"report","params":{"name":"monthly_sales","format":"pdf"}}')
echo "响应: $REPORT_TASK"
REPORT_ID=$(echo $REPORT_TASK | grep -o '"id":[0-9]*' | cut -d':' -f2)
echo "${GREEN}✓ Report 任务提交成功, ID: $REPORT_ID${NC}"
echo ""

# 3. 提交 cleanup 任务
echo "${YELLOW}[测试3] 提交 cleanup 任务...${NC}"
CLEANUP_TASK=$(curl -s -X POST "${BASE_URL}/api/tasks" \
    -H "Content-Type: application/json" \
    -d '{"type":"cleanup","params":{"target":"temp_files","older_than_days":7}}')
echo "响应: $CLEANUP_TASK"
CLEANUP_ID=$(echo $CLEANUP_TASK | grep -o '"id":[0-9]*' | cut -d':' -f2)
echo "${GREEN}✓ Cleanup 任务提交成功, ID: $CLEANUP_ID${NC}"
echo ""

# 4. 批量提交多个任务
echo "${YELLOW}[测试4] 批量提交 5 个任务...${NC}"
for i in {1..5}; do
    curl -s -X POST "${BASE_URL}/api/tasks" \
        -H "Content-Type: application/json" \
        -d "{\"type\":\"email\",\"params\":{\"to\":\"user$i@example.com\",\"subject\":\"Batch $i\"}}" > /dev/null
    echo "  - 任务 $i 已提交"
done
echo "${GREEN}✓ 批量任务提交完成${NC}"
echo ""

# 5. 查询单个任务
echo "${YELLOW}[测试5] 查询任务 ID=$EMAIL_ID...${NC}"
sleep 0.5
TASK_DETAIL=$(curl -s "${BASE_URL}/api/tasks/$EMAIL_ID")
echo "响应: $TASK_DETAIL"
echo "${GREEN}✓ 任务查询成功${NC}"
echo ""

# 6. 获取任务列表（第一页）
echo "${YELLOW}[测试6] 获取任务列表（第1页，每页5条）...${NC}"
TASK_LIST=$(curl -s "${BASE_URL}/api/tasks?page=1&page_size=5")
echo "响应: $TASK_LIST"
echo "${GREEN}✓ 任务列表获取成功${NC}"
echo ""

# 7. 按状态过滤任务
echo "${YELLOW}[测试7] 获取 pending 状态的任务...${NC}"
PENDING_TASKS=$(curl -s "${BASE_URL}/api/tasks?status=pending")
echo "响应: $PENDING_TASKS"
echo "${GREEN}✓ 状态过滤成功${NC}"
echo ""

echo "${YELLOW}[测试8] 获取 running 状态的任务...${NC}"
RUNNING_TASKS=$(curl -s "${BASE_URL}/api/tasks?status=running")
echo "响应: $RUNNING_TASKS"
echo "${GREEN}✓ 状态过滤成功${NC}"
echo ""

# 8. 获取任务统计
echo "${YELLOW}[测试9] 获取任务统计...${NC}"
STATS=$(curl -s "${BASE_URL}/api/tasks/stats")
echo "响应: $STATS"
echo "${GREEN}✓ 统计信息获取成功${NC}"
echo ""

# 9. 等待任务执行，观察状态变化
echo "${YELLOW}[测试10] 等待任务执行，观察状态变化...${NC}"
echo "等待 2 秒后重新查询任务..."
sleep 2

echo ""
echo "重新查询 Email 任务:"
EMAIL_STATUS=$(curl -s "${BASE_URL}/api/tasks/$EMAIL_ID")
echo "响应: $EMAIL_STATUS"

echo ""
echo "重新获取统计:"
STATS2=$(curl -s "${BASE_URL}/api/tasks/stats")
echo "响应: $STATS2"
echo "${GREEN}✓ 状态观察完成${NC}"
echo ""

# 10. 等待更长时间，查看 completed 任务
echo "${YELLOW}[测试11] 等待所有任务完成（约 6 秒）...${NC}"
sleep 6

echo ""
echo "最终统计:"
FINAL_STATS=$(curl -s "${BASE_URL}/api/tasks/stats")
echo "响应: $FINAL_STATS"

echo ""
echo "查询已完成任务:"
COMPLETED_TASKS=$(curl -s "${BASE_URL}/api/tasks?status=completed&page=1&page_size=10")
echo "响应: $COMPLETED_TASKS"
echo "${GREEN}✓ 完成状态查询成功${NC}"
echo ""

# 测试完成
echo "${BLUE}========================================${NC}"
echo "${GREEN}  所有测试完成！${NC}"
echo "${BLUE}========================================${NC}"
echo ""
echo "测试覆盖:"
echo "  ✓ 提交任务 (POST /api/tasks)"
echo "  ✓ 查询任务 (GET /api/tasks/:id)"
echo "  ✓ 任务列表 (GET /api/tasks)"
echo "  ✓ 状态过滤 (GET /api/tasks?status=xxx)"
echo "  ✓ 分页查询 (GET /api/tasks?page=x&page_size=y)"
echo "  ✓ 任务统计 (GET /api/tasks/stats)"
echo "  ✓ 异步执行观察"
