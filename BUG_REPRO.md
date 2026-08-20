# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

运营人员提交三张整改照片后继续复用客户端数组上传，已经保存的第一张照片会被后续请求替换。请修复保存结果与调用方数据相互影响的问题，确保已入库证据不会再变化。照片列表相关测试文件不能修改，也不得降低内容断言的严格程度。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-22
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-22.git
- parent SHA：49dffa8e13ef1a81b1aaea9cb470613acf897810

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-22.git bug-repro
cd bug-repro
git checkout --detach 49dffa8e13ef1a81b1aaea9cb470613acf897810
go test ./internal/httpapi -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.01s)
    task_behavior_test.go:16: body={"first":"overwritten!"}
FAIL
FAIL	chargeguard/internal/httpapi	0.069s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:16: body={"first":"overwritten!"}
FAIL
FAIL	chargeguard/internal/httpapi	0.003s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，在题面描述的触发条件下应得到预期业务结果且不再出现原始症状；定向验证命令修复前必须失败、应用修复后必须通过，相关回归和仓库全量测试必须通过；不得新增、删除或修改测试文件，不得跳过测试、降低断言或绕过目标逻辑。
