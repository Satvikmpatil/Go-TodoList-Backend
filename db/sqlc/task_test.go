package db

import (
	"context"
	"testing"

	"github.com/Satvikmpatil/tolist/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomTask(t *testing.T) Task {
	user1 := CreateRandomUser(t)
	task1 := util.RandomTask()
	arg := CreateTaskParams{
		UserID: user1.UserID,
		Task:   task1,
	}
	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)
	require.Equal(t, user1.UserID, task.UserID)
	require.Equal(t, task1, task.Task)
	require.False(t, task.Status)
	require.NotZero(t, task.TaskID)
	require.NotZero(t, task.CreatedAt)
	return task
}

func TestCreateTask(t *testing.T) {
	createRandomTask(t)
}

func TestGetTask(t *testing.T) {
	task1 := createRandomTask(t)
	task2, err := testQueries.GetTask(context.Background(), task1.TaskID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)
	require.Equal(t, task1.TaskID, task2.TaskID)
	require.Equal(t, task1.Task, task2.Task)
	require.Equal(t, task1.UserID, task2.UserID)
}

func TestListTasks(t *testing.T) {
	for i := 0; i < 3; i++ {
		createRandomTask(t)
	}
	tasks, err := testQueries.ListTasks(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, tasks)
	for _, task := range tasks {
		require.NotEmpty(t, task)
	}
}

func TestListTasksByUser(t *testing.T) {
	user := CreateRandomUser(t)
	for i := 0; i < 3; i++ {
		arg := CreateTaskParams{
			UserID: user.UserID,
			Task:   util.RandomTask(),
		}
		_, err := testQueries.CreateTask(context.Background(), arg)
		require.NoError(t, err)
	}
	tasks, err := testQueries.ListTasksByUser(context.Background(), user.UserID)
	require.NoError(t, err)
	require.Len(t, tasks, 3)
	for _, task := range tasks {
		require.Equal(t, user.UserID, task.UserID)
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	task1 := createRandomTask(t)
	arg := UpdateTaskStatusParams{
		TaskID: task1.TaskID,
		Status: true,
	}
	task2, err := testQueries.UpdateTaskStatus(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)
	require.Equal(t, task1.TaskID, task2.TaskID)
	require.True(t, task2.Status)
}

func TestUpdateTask(t *testing.T) {
	task1 := createRandomTask(t)
	newTaskText := util.RandomTask()
	arg := UpdateTaskParams{
		TaskID: task1.TaskID,
		Task:   newTaskText,
		Status: true,
	}
	task2, err := testQueries.UpdateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)
	require.Equal(t, task1.TaskID, task2.TaskID)
	require.Equal(t, newTaskText, task2.Task)
	require.True(t, task2.Status)
}

func TestDeleteTask(t *testing.T) {
	task1 := createRandomTask(t)
	err := testQueries.DeleteTask(context.Background(), task1.TaskID)
	require.NoError(t, err)

	task2, err := testQueries.GetTask(context.Background(), task1.TaskID)
	require.Error(t, err)
	require.Empty(t, task2)
}
