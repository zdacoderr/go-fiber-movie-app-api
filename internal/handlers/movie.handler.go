package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zdacoder/go-fiber-movie-app-api/pkg/utils"
)

// ListMovies godoc
// @Summary      List movies
// @Description  get list of all movies
// @Tags         movies
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.SuccessResponse "Movies fetched successfully"
// @Failure      204	{object}  utils.ErrorResponse "Movies data is empty"
// @Failure      500  {object}  utils.ErrorResponse "Failed to fetch movies"
// @Router       /api/movies [get]
func ListMovies(ctx *fiber.Ctx) error {
	return utils.OKResponse(ctx, "Movies fetched successfully", nil)
}

// GetMovie godoc
// @Summary      Get a movie
// @Description  get movie by ID
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Movie ID"
// @Success      200  {object}  utils.SuccessResponse "Movie fetched successfully"
// @Failure      404  {object}  utils.ErrorResponse "Movie not found"
// @Failure      500  {object}  utils.ErrorResponse "Failed to fetch movie"
// @Router       /api/movies/{id} [get]
func GetMovie(ctx *fiber.Ctx) error {
	return utils.OKResponse(ctx, "Movie fetched successfully", nil)
}

// CreateMovie godoc
// @Summary      Create a movie
// @Description  create a new movie
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        movie  body      models.Movie  true  "Movie data"
// @Success      201  {object}  utils.SuccessResponse "Movie created successfully"
// @Failure      400  {object}  utils.ErrorResponse "Invalid request body or validation failed"
// @Failure      500  {object}  utils.ErrorResponse "Failed to create movie"
// @Router       /api/movies [post]
func CreateMovie(ctx *fiber.Ctx) error {
	return utils.CreatedResponse(ctx, "Movie Created successfully", nil)
}

// UpdateMovie godoc
// @Summary      Update a movie
// @Description  update an existing movie by ID
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        id     path      string       true  "Movie ID"
// @Param        movie  body      models.Movie  true  "Updated movie data"
// @Success      200  {object}  utils.SuccessResponse "Movie updated successfully"
// @Failure      400  {object}  utils.ErrorResponse "Invalid request body or validation failed"
// @Failure      404  {object}  utils.ErrorResponse "Movie not found"
// @Failure      500  {object}  utils.ErrorResponse "Failed to update movie"
// @Router       /api/movies/{id} [put]
func UpdateMovie(ctx *fiber.Ctx) error {
	return utils.OKResponse(ctx, "Movie Updated successfully", nil)
}

// DeleteMovie godoc
// @Summary      Delete a movie
// @Description  delete an existing movie by ID
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Movie ID"
// @Success      200  {object}  utils.SuccessResponse "Movie deleted successfully"
// @Failure      404  {object}  utils.ErrorResponse "Movie not found"
// @Failure      500  {object}  utils.ErrorResponse "Failed to delete movie"
// @Router       /api/movies/{id} [delete]
func DeleteMovie(ctx *fiber.Ctx) error {
	return utils.OKResponse(ctx, "Movie Deleted successfully", nil)
}
