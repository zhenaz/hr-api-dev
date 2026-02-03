package routes

import (
	"codeid.hr-api/internal/handlers"
	"codeid.hr-api/internal/repositories"
	"codeid.hr-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	//5. Initialize repositories
	regionRepo := repositories.NewRegionRepository(db)
	countryRepo := repositories.NewCountryRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	//6. Initialize services
	regionService := services.NewRegionService(regionRepo)
	countryService := services.NewCountryService(countryRepo)
	employeeService := services.NewEmployeeService(employeeRepo)
	//7. Initialize handlers
	regionHandler := handlers.NewRegionHandler(regionService)
	countryHandler := handlers.NewCountryHandler(countryService)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	//3.1 call handler
	// router.GET("/", handlers.WelcomeHandler)
	//9.call basepath
	basePath := viper.GetString("SERVER.BASE_PATH")
	//8. grouping subroute with prefix /api
	api := router.Group(basePath)
	{
		// region routes endpoints
		regions := api.Group("/regions")
		{
			regions.GET("", regionHandler.GetRegions)
			regions.GET("/:id", regionHandler.GetRegion)
			regions.GET("/countries", regionHandler.GetRegionsWithCountry)
			regions.GET("/:id/countries", regionHandler.GetRegionByIdWithCountry)
			regions.POST("", regionHandler.CreateRegion)
			regions.PUT("/:id", regionHandler.UpdateRegion)
			regions.DELETE("/:id", regionHandler.DeleteRegion)
		}
		countries := api.Group("/countries")
		{
			countries.GET("", countryHandler.GetCountries)
			countries.GET("/:id", countryHandler.GetCountry)
			countries.POST("", countryHandler.CreateCountry)
			countries.PUT("/:id", countryHandler.UpdateCountry)
			countries.DELETE("/:id", countryHandler.DeleteCountry)
		}
		employees := api.Group("/employees")
		{
			employees.GET("", employeeHandler.GetEmployees)
			employees.GET("/:id", employeeHandler.GetEmployee)
			employees.POST("", employeeHandler.CreateEmployee)
			employees.PUT("/:id", employeeHandler.UpdateEmployee)
			employees.DELETE("/:id", employeeHandler.DeleteEmployee)
		}
	}
}