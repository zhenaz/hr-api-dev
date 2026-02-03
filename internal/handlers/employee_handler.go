package handlers

import (
	"net/http"
	"strconv"

	"codeid.hr-api/internal/domain/model"
	"codeid.hr-api/internal/services"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeService services.EmployeeService
}

func NewEmployeeHandler(employeeService services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: employeeService,
	}
}
func (h *EmployeeHandler) GetEmployees(c *gin.Context) {
	employees, err := h.employeeService.GetAllEmployees(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to fetch employees",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  employees,
	})
}
func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	// get id param from endpoint router yg bertipe string
	//stcconv.ParUint : konvert string ke uint
	//parseUint(idstr,10,32) -> 10 base ten number,32 tipe bit
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid employee ID",
		})
		return
	}
	employee, err := h.employeeService.GetEmployeeByID(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "Employee not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid input",
			"details": err.Error(),
		})
		return
	}
	if err := h.employeeService.CreateEmployee(c.Request.Context(), &employee); err !=
		nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to create employee",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Employee created successfully",
		"data":    employee,
	})
}
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid country ID",
		})
		return
	}
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid input",
			"details": err.Error(),
		})
		return
	}
	employee.EmployeeID = int32(id)
	if err := h.employeeService.UpdateEmployee(c.Request.Context(), &employee); err !=
		nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to update employee",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Employee updated successfully",
		"data":    employee,
	})
}
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid employee ID",
		})
		return
	}
	if err := h.employeeService.DeleteEmployee(c.Request.Context(), int32(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to delete employee",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Employee deleted successfully",
	})
}
