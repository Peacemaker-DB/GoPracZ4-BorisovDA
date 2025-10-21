package task

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.list)          // GET /tasks
	r.Post("/", h.create)       // POST /tasks
	r.Get("/{id}", h.get)       // GET /tasks/{id}
	r.Put("/{id}", h.update)    // PUT /tasks/{id}
	r.Delete("/{id}", h.delete) // DELETE /tasks/{id}
	return r
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var doneFilter *bool
	if d := query.Get("done"); d != "" {
		if done, err := strconv.ParseBool(d); err == nil {
			doneFilter = &done
		}
	}

	allTasks := h.repo.List()
	var filteredTasks []*Task

	for _, task := range allTasks {
		if doneFilter == nil || task.Done == *doneFilter {
			filteredTasks = append(filteredTasks, task)
		}
	}

	page := 1
	if p := query.Get("page"); p != "" {
		if newPage, err := strconv.Atoi(p); err == nil && newPage > 0 {
			page = newPage
		}
	}

	limit := 10
	if l := query.Get("limit"); l != "" {
		if newLimit, err := strconv.Atoi(l); err == nil && newLimit > 0 {
			limit = newLimit
		}
	}

	start := (page - 1) * limit
	end := start + limit
	if start > len(filteredTasks) {
		start = len(filteredTasks)
	}
	if end > len(filteredTasks) {
		end = len(filteredTasks)
	}

	result := filteredTasks[start:end]

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"tasks": result,
		"total": len(filteredTasks),
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, bad := parseID(w, r)
	if bad {
		return
	}
	t, err := h.repo.Get(id)
	if err != nil {
		httpError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, t)
}

type createReq struct {
	Title string `json:"title"`
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
		httpError(w, http.StatusBadRequest, "invalid json: require non-empty title")
		return
	}
	title := strings.TrimSpace(req.Title)
	if len(title) < 3 || len(title) > 100 {
		httpError(w, http.StatusBadRequest, "Длина должна быть >3 и <100")
		return
	}
	t := h.repo.Create(req.Title)
	writeJSON(w, http.StatusCreated, t)
}

type updateReq struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, bad := parseID(w, r)
	if bad {
		return
	}
	var req updateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
		httpError(w, http.StatusBadRequest, "invalid json: require non-empty title")
		return
	}
	title := strings.TrimSpace(req.Title)
	if len(title) < 3 || len(title) > 100 {
		httpError(w, http.StatusBadRequest, "Длина должна быть >3 и <100")
		return
	}
	t, err := h.repo.Update(id, req.Title, req.Done)
	if err != nil {
		httpError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, bad := parseID(w, r)
	if bad {
		return
	}
	if err := h.repo.Delete(id); err != nil {
		httpError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// helpers

func parseID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	raw := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		httpError(w, http.StatusBadRequest, "invalid id")
		return 0, true
	}
	return id, false
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func httpError(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}
