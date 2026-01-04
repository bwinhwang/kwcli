# KWCLI 与 Klocwork Web API 对照文档

> 生成日期: 2026-01-04
> API 来源: http://cnninvmkw02.joynext.com:8080/review/api

## 概述

| 指标 | 数值 |
|------|------|
| API 总 action 数 | 51 |
| CLI 已实现 | 26 |
| CLI 未实现 | 25 |
| 覆盖率 | 51% |

## API 调用方式

所有 API 通过 POST 请求发送到 `/review/api`，必需参数：
- `user`* - Klocwork 用户名
- `ltoken` - kwauth 登录令牌
- `action`* - action 名称

## 详细对照表

### 1. 项目管理 (Projects)

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `projects` | `projects` | ✅ 已实现 | 获取项目和 stream 列表 |
| `create_project` | `create_project` | ✅ 已实现 | 创建项目 |
| `update_project` | `update_project` | ✅ 已实现 | 更新项目 |
| `delete_project` | `delete_project` | ✅ 已实现 | 删除项目 |
| `import_project` | - | ❌ 未实现 | 从其他服务器导入项目 |
| `import_status` | - | ❌ 未实现 | 查看导入状态 |
| `import_server_configuration` | - | ❌ 未实现 | 导入服务器配置 |
| `project_configuration` | - | ❌ 未实现 | 生成项目配置报告 |

### 2. 构建管理 (Builds)

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `builds` | `builds` | ✅ 已实现 | 获取构建列表 |
| `update_build` | `update_build` | ✅ 已实现 | 更新构建信息 |
| `delete_build` | `delete_build` | ✅ 已实现 | 删除构建 |
| `build_log_download` | - | ❌ 未实现 | 下载构建日志 (zip) |
| `ci_builds` | - | ❌ 未实现 | 获取 CI 构建列表 |
| `update_ci_build` | `update_ci_build` | ✅ 已实现 | 更新 CI 构建 |
| `delete_ci_build` | `delete_ci_build` | ✅ 已实现 | 删除 CI 构建 |

### 3. 模块管理 (Modules)

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `modules` | `modules` | ✅ 已实现 | 获取模块列表 |
| `create_module` | `create_module` | ✅ 已实现 | 创建模块 |
| `update_module` | `update_module` | ✅ 已实现 | 更新模块 |
| `delete_module` | `delete_module` | ✅ 已实现 | 删除模块 |

### 4. 视图管理 (Views)

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `views` | `views` | ✅ 已实现 | 获取视图列表 |
| `create_view` | `create_view` | ✅ 已实现 | 创建视图 |
| `update_view` | `update_view` | ✅ 已实现 | 更新视图 |
| `delete_view` | `delete_view` | ✅ 已实现 | 删除视图 |

### 5. 问题检索与管理 (Issues)

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `search` | `search` | ✅ 已实现 | 搜索问题 |
| `issue_details` | `issue_details` | ✅ 已实现 | 获取问题详情 |
| `ci_issue_details` | - | ❌ 未实现 | 获取 CI 问题详情 |
| `update_status` | - | ❌ 未实现 | **更新问题状态/Owner/Comment** |

### 6. 合规性报告 (Compliance)

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `compliance_report` | `compliance_report` | ✅ 已实现 | 生成合规性报告 |
| `compliance_download` | `compliance_download` | ✅ 已实现 | 下载合规性报告 |
| `compliance_delete` | `compliance_delete` | ✅ 已实现 | 删除合规性报告 |

### 7. 缺陷类型与分类 (Defects & Taxonomies)

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `defect_types` | `defect_types` | ✅ 已实现 | 获取缺陷类型列表 |
| `taxonomies` | `taxonomies` | ✅ 已实现 | 获取分类术语 |
| `update_defect_type` | - | ❌ 未实现 | 启用/禁用缺陷类型 |

### 8. 指标与报告 (Metrics & Reports)

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `metrics` | `metrics` | ✅ 已实现 | 获取指标数据 |
| `report` | - | ❌ 未实现 | 生成构建摘要报告 |
| `fchurns` | - | ❌ 未实现 | 生成文件变更报告 |

### 9. 用户管理 (Users) ⚠️ 完全未实现

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `users` | - | ❌ 未实现 | 列出用户 |
| `create_user` | - | ❌ 未实现 | 创建用户 |
| `delete_user` | - | ❌ 未实现 | 删除用户 |

### 10. 用户组管理 (Groups) ⚠️ 完全未实现

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `groups` | - | ❌ 未实现 | 列出用户组 |
| `create_group` | - | ❌ 未实现 | 创建用户组 |
| `update_group` | - | ❌ 未实现 | 更新用户组 |
| `delete_group` | - | ❌ 未实现 | 删除用户组 |

### 11. 角色管理 (Roles) ⚠️ 完全未实现

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `roles` | - | ❌ 未实现 | 列出角色 |
| `create_role` | - | ❌ 未实现 | 创建角色 |
| `delete_role` | - | ❌ 未实现 | 删除角色 |
| `role_assignments` | - | ❌ 未实现 | 列出角色分配 |
| `update_role_assignment` | - | ❌ 未实现 | 更新角色分配 |
| `update_role_permissions` | - | ❌ 未实现 | 更新角色权限 |

### 12. 系统管理 (System)

| API Action | CLI 命令 | 状态 | 说明 |
|------------|----------|------|------|
| `version` | `version` | ✅ 已实现 | 获取服务器版本 |
| `license_count` | - | ❌ 未实现 | 获取许可证数量 |
| `task_status` | - | ❌ 未实现 | 列出运行中的任务 |

---

## 未实现命令优先级建议

### 高优先级 (常用功能)

| 命令 | 理由 |
|------|------|
| `update_status` | 批量更新问题状态，日常使用频率最高 |
| `ci_builds` | CI/CD 集成必备 |
| `ci_issue_details` | CI 问题详情查看 |
| `report` | 生成报告，汇报用 |

### 中优先级 (管理功能)

| 命令 | 理由 |
|------|------|
| `users` | 用户列表查询 |
| `groups` | 用户组管理 |
| `roles` / `role_assignments` | 权限管理 |
| `update_defect_type` | 启用/禁用检查器 |
| `build_log_download` | 下载构建日志排查问题 |

### 低优先级 (高级功能)

| 命令 | 理由 |
|------|------|
| `import_project` | 服务器迁移时使用 |
| `import_server_configuration` | 服务器迁移时使用 |
| `fchurns` | 特殊报告 |
| `license_count` | 管理员偶尔使用 |
| `task_status` | 管理员偶尔使用 |

---

## 架构设计评价

### 优点

1. **命令与 API 一一对应** - CLI 命令名直接使用 API action 名
2. **参数透传** - cobra flags 直接映射到 API 参数
3. **认证兼容** - 读取 `~/.klocwork/ltoken`，与官方工具兼容
4. **扩展性好** - 每个命令独立文件，添加新命令简单

### 待改进

1. **功能覆盖不全** - 用户/组/角色管理完全缺失
2. **参数校验** - 可加强必填参数校验
3. **帮助信息** - 可将 API 文档描述作为 flag help

---

## 附录：API 参数速查

### 常用必填参数

| 参数 | 说明 | 使用场景 |
|------|------|---------|
| `project` | 项目名 | 大部分命令 |
| `name` | 资源名称 | create/update/delete |
| `id` | 问题 ID | issue_details |
| `query` | 搜索条件 | search, create_view |
| `taxonomy` | 分类名 | compliance_report |

### 常用可选参数

| 参数 | 说明 |
|------|------|
| `view` | 视图名称过滤 |
| `build` | 指定构建版本 |
| `limit` | 结果数量限制 |
| `tags` | 标签 (逗号分隔) |
