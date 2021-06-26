// 命名空间
db.namespace.ensureIndex({code: 1}, {unique: true, background: true})

// 名单类型
db.list_type.ensureIndex({namespace: 1, code: 1}, {unique: true, background: true})

// 名单项
db.list_item.ensureIndex({namespace: 1, code: 1, value: 1}, {unique: true, background: true})
