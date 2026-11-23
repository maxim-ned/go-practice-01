package persistence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"example.com/go-practice-01/internal/tasks"
)

func Load(path string) (tasks.State, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return tasks.State{NextID: 1}, nil
	} else if err != nil {
		return tasks.State{NextID: 1}, fmt.Errorf("Ошибка чтения: %w", err)
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var state tasks.State
	err = decoder.Decode(&state)
	if err != nil {
		return tasks.State{NextID: 1}, fmt.Errorf("Ошибка декодирования: %w", err)
	}
	var new_state tasks.State
	err = decoder.Decode(&new_state)
	if err != io.EOF {
		return tasks.State{NextID: 1}, fmt.Errorf("Файл битый: %w", err)
	}
	err = tasks.NewMemoryStore().Restore(state)
	if err != nil {
		return tasks.State{NextID: 1}, err
	}

	return state, nil
}

func Save(path string, state tasks.State) error {
	err := tasks.NewMemoryStore().Restore(state)
	if err != nil {
		return fmt.Errorf("Неверные входные данные: %w", err)
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("Ошибка записи: %w", err)
	}
	data = append(data, '\n')
	temp, err := os.CreateTemp(filepath.Dir(path), ".tasks-*.tmp")
	if err != nil {
		return fmt.Errorf("Не удалось открыть файл: %w", err)
	}
	defer os.Remove(temp.Name())
	defer temp.Close()
	_, err = temp.Write(data)
	if err != nil {
		return fmt.Errorf("Ошибка записи в файл: %w", err)
	}
	err = temp.Sync()
	if err != nil {
		return fmt.Errorf("Не удалось синхронизировать данные: %w", err)
	}
	err = temp.Close()
	if err != nil {
		return fmt.Errorf("Не удалось закрыть файл: %w", err)
	}
	err = os.Rename(temp.Name(), path)
	if err != nil {
		return fmt.Errorf("Не удалось переименовать временный файл: %w", err)
	}
	return nil
}
