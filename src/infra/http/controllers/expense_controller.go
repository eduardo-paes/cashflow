package expense_controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	core "github.com/eduardo-paes/cashflow/core/entities"
	"github.com/eduardo-paes/cashflow/core/ports"
	"github.com/gorilla/mux"
)

type Service struct {
	UseCase core.ExpenseUseCases
}

// New returns contract implementation of ExpenseService
func New(usecase core.ExpenseUseCases) core.ExpenseService {
	return &Service{
		UseCase: usecase,
	}
}

// @Summary Create a new expense
// @Description Create a new expense
// @Tags expenses
// @Accept json
// @Produce json
// @Param request body ports.ExpenseInput true "Expense data"
// @Success 200 {object} core.Expense
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /expense [post]
func (s *Service) Create(response http.ResponseWriter, request *http.Request) {
	expenseRequest, err := ports.FromJSONCreateExpense(request.Body)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}

	expense, err := s.UseCase.Create(expenseRequest)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(expense)
}

// @Summary Delete an expense by ID
// @Description Delete an expense by ID
// @Tags expenses
// @Accept json
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {object} core.Expense
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /expense/{id} [delete]
func (s *Service) Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid ID"))
		return
	}

	expense, err := s.UseCase.Delete(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(expense)
}

// @Summary Get one or multiple expenses
// @Description Get one or multiple expenses based on provided parameters
// @Tags expenses
// @Accept json
// @Produce json
// @Param id query int false "Expense ID"
// @Param skip query int false "Number of items to skip"
// @Param take query int false "Number of items to take"
// @Success 200 {array} core.Expense
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /expense [get]
func (s *Service) GetOneOrMany(response http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	id, idErr := strconv.ParseInt(queryValues.Get("id"), 10, 64)
	skip, _ := strconv.Atoi(queryValues.Get("skip"))
	take, _ := strconv.Atoi(queryValues.Get("take"))

	var expenses []core.Expense
	var err error

	if idErr == nil && id > 0 {
		// If an id is provided, fetch a specific expense
		expenses, err = s.UseCase.GetOneOrMany(skip, take, id)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(err.Error()))
			return
		}
	} else {
		// If no id is provided, fetch multiple expenses
		expenses, err = s.UseCase.GetOneOrMany(skip, take)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(err.Error()))
			return
		}
	}

	json.NewEncoder(response).Encode(expenses)
}

// @Summary Update an expense by ID
// @Description Update an expense by ID
// @Tags expenses
// @Accept json
// @Produce json
// @Param id path int true "Expense ID"
// @Param request body ports.ExpenseInput true "Expense data"
// @Success 200 {object} core.Expense
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /expense/{id} [put]
func (s *Service) Update(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Invalid ID"))
		return
	}

	expenseRequest, err := ports.FromJSONCreateExpense(request.Body)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}

	expense, err := s.UseCase.Update(id, expenseRequest)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(expense)
}