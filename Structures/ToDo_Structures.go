package Structures

// Task - Структура хранящая Имя задачи ее Идентификационный номер
// И Флаг Указывающий статус задачи, выполнена или не выполнена
type Task struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	Done bool   `json:"Done"`
}
