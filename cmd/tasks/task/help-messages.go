package task

const (
	createLong   = "\n$toodles add 'Task title' 'Task description'"
	editLong     = "\n$toodles TASK_ID 'New task title' 'New task description'"
	priorityLong = "\n$toodles TASK_ID normal|mid|high"
)

const removeLong = `

Singular: $toodles remove TASK_ID
Plural:   $toodles remove TASK_ID_1 TASK_ID_2
    `

const doneLong = `

Singular: $toodles done TASK_ID
Plural:   $toodles done TASK_ID_1 TASK_ID_2
    `
