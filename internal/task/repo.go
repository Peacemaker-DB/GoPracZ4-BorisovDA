package task

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
)

var ErrNotFound = errors.New("task not found")

type Repo struct {
	mu       sync.RWMutex
	seq      int64
	items    map[int64]*Task
	filename string
}

func NewRepo(filename string) *Repo {
	repo := &Repo{
		items:    make(map[int64]*Task),
		filename: filename,
	}

	if err := repo.Load(); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	return repo
}

func (r *Repo) Save() error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.filename == "" {
		return nil
	}

	data, err := json.MarshalIndent(r.items, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filename, data, 0644)
}

func (r *Repo) Load() error {
	if r.filename == "" {
		return nil
	}

	data, err := os.ReadFile(r.filename)
	if err != nil {
		return err
	}

	var items map[int64]*Task
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.items = items

	r.seq = 0
	for id := range r.items {
		if id > r.seq {
			r.seq = id
		}
	}

	return nil
}

func (r *Repo) List() []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		out = append(out, t)
	}
	return out
}

func (r *Repo) Get(id int64) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	return t, nil
}

func (r *Repo) Create(title string) *Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	now := time.Now()
	t := &Task{ID: r.seq, Title: title, CreatedAt: now, UpdatedAt: now, Done: false}
	r.items[t.ID] = t
	go r.Save()
	return t
}

func (r *Repo) Update(id int64, title string, done bool) (*Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}
	t.Title = title
	t.Done = done
	t.UpdatedAt = time.Now()
	go r.Save()
	return t, nil
}

func (r *Repo) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}
	delete(r.items, id)
	go r.Save()
	return nil
}
