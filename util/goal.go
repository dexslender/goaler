package util

import "github.com/disgoorg/snowflake/v2"

type GoalType int

const (
	TypeProject GoalType = iota + 1
	TypeMilestone
	TypeTODO
	TypeQuest
	TypeHabit
)

type TaskStatuses string

type Task struct {
	UserID snowflake.ID
	Title,
	Description,
	Status string
	Deadline,
	Reminder *string
	Quest bool
}

type Container struct {
	UserID snowflake.ID
	Title,
	Description string
	Type ContainerType	
}
