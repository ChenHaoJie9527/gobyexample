---
name: golang-suite
description: 统一的 Go / Golang 总入口技能，中文优先。Use when the task spans multiple Go concerns or when Codex should first classify the problem, then route to the right installed Go skills for style, backend design, testing, concurrency, performance, security, observability, dependency management, and troubleshooting.
---

# Go Skill Suite

用这个技能作为统一入口，先判断问题类型，再加载最合适的 Go 子技能；不要一次性加载全部材料。

默认用中文输出，除非用户明确要求英文；代码、命令、标识符保持原样。

## Default Workflow

1. 先识别任务属于哪一类：编码规范、后端开发、数据库、并发、测试、性能、可观测性、安全、依赖治理、问题排查。
2. 先读取最少必要的技能；通常 1 到 3 个即可。
3. 如果任务涉及实际写 Go 代码，默认同时参考 `../golang-style/SKILL.md` 和 `../effective-go/SKILL.md`。
4. 如果任务是 Go 后端实现、重构或 review，优先再读 `../go-backend-development-cn/SKILL.md`。
5. 如果任务跨多个主题，按“主技能 + 补充技能”组合，而不是平铺所有技能。
6. 输出中说明选择了哪些技能，以及为什么这样组合。

## Chinese-First Rules

- 用户使用中文提问时，默认用中文思考、组织和回答。
- 用户没有明确指定 skill 时，也可以把这个 skill 当作 Go 任务总入口。
- 用户只说“帮我看看这段 Go 代码”：
  默认按代码审查或改进请求处理，并组合 `golang-style` + `effective-go`。
- 用户只说“帮我做个 Go 服务”：
  默认按后端实现处理，并组合 `go-backend-development-cn` + `golang-style`。
- 用户只说“这个 Go 程序有问题”：
  默认按排障处理，并优先考虑 `golang-troubleshooting`，必要时补充并发、数据库或性能专项 skill。

## Routing Rules

- 写、改、重构任意 Go 代码：
  先读 `../golang-style/SKILL.md`，必要时再读 `../effective-go/SKILL.md`。
- 设计或实现 Go 后端、HTTP、gRPC、middleware、worker、repository：
  先读 `../go-backend-development-cn/SKILL.md`。
- 任务与测试、testify、benchmark、CI 相关：
  读 [references/skill-map.md](references/skill-map.md) 中的 Testing 与 Performance 小节，并按需打开对应 skill。
- 任务与并发、取消、goroutine 泄漏、race、channel、锁相关：
  读 [references/skill-map.md](references/skill-map.md) 中的 Concurrency 与 Stability 小节。
- 任务与数据库、事务、连接池、超时、依赖注入、项目结构相关：
  读 [references/skill-map.md](references/skill-map.md) 中的 Backend 和 Architecture 小节。
- 任务与性能、可观测性、安全、依赖升级相关：
  读 [references/skill-map.md](references/skill-map.md) 中的 Production 小节。
- 如果用户只说“帮我处理 Go 问题”而没有限定方向：
  默认组合 `golang-style` + `effective-go`；若明显是服务端问题，再加 `go-backend-development-cn`。

## Combination Presets

- 通用 Go 编码：
  `golang-style` + `effective-go`
- Go 后端功能开发：
  `go-backend-development-cn` + `golang-style` + `golang-testing`
- Go PR / Code Review：
  `golang-style` + `effective-go` + 任务相关专项 skill
- 并发问题排查：
  `golang-concurrency` + `golang-context` + `golang-troubleshooting`
- 数据库与 repository：
  `go-backend-development-cn` + `golang-database` + `golang-error-handling`
- 性能优化：
  `golang-benchmark` + `golang-performance` + `golang-observability`
- 生产化加固：
  `go-backend-development-cn` + `golang-security` + `golang-observability`
- 测试与质量体系：
  `golang-testing` + `golang-stretchr-testify` + `golang-linter` + `golang-continuous-integration`

## Prompt Shapes

这些提示词可以直接复用，用来稳定触发这个 skill。

- `用 $golang-suite 帮我 review 这段 Go 代码，并指出 bug、风险和缺失测试。`
- `用 $golang-suite 帮我重构这段 Go 代码，保持行为不变，但提升可读性和可维护性。`
- `用 $golang-suite 帮我实现一个 Go HTTP API，带 handler、service、repository 分层。`
- `用 $golang-suite 帮我排查这个 Go 服务为什么会 goroutine 泄漏。`
- `用 $golang-suite 帮我给这个 Go 项目补测试，包括 table-driven tests。`
- `用 $golang-suite 帮我分析这个 Go 程序为什么慢，并给出 benchmark 和优化方案。`
- `用 $golang-suite 帮我检查这个 Go 服务离生产可用还差什么。`
- `Use $golang-suite to classify this Go task first, then apply the right installed skills and implement the solution.`

## Output Expectations

- 先给结论，再给关键依据。
- 如果是 review，先列问题，再给总结。
- 如果是实现任务，直接落代码，不要只停在方案描述。
- 如果是排障任务，先缩小问题范围，再给修复和验证方法。
- 如果涉及多个 skill，简要说明组合原因。

## On-Demand Map

完整 skill 清单与分类见 [references/skill-map.md](references/skill-map.md)。

只有在需要做更细路由时再读取这份 map，不要默认整份载入。
