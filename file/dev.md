## 关键点说明
1. 关于zset的说明语法

当我们使用 z := redis.Z{Score: 1, Member: task.UserID} 这个语法时，我们正在创建一个 redis.Z 类型的结构体对象，并将其赋值给变量 z。
redis.Z 是在 "github.com/go-redis/redis" 包中定义的一个结构体类型。它表示有序集合（sorted set）中的一个元素，包含两个字段：Score 和 Member。

Score 字段表示有序集合元素的分数，它可以是一个浮点数，用于排序和比较元素的顺序。
Member 字段表示有序集合元素的成员，可以是任何类型的值，用于唯一标识该元素。
在这个例子中，我们使用了结构体字面量的语法来创建一个 redis.Z 类型的对象。通过指定字段名和对应的值，我们可以按需设置 Score 和 Member 字段的值。

Score: 1 表示将 Score 字段的值设置为 1。
Member: task.UserID 表示将 Member 字段的值设置为 task.UserID，其中 task.UserID 是一个变量，用于表示成员的唯一标识。
通过这种方式，我们可以创建一个包含特定分数和成员的 redis.Z 对象，并将其用于有序集合操作，比如将其添加到有序集合中。