package main

import (
	"src/github.com/mustafaakilll/go-site-exam/db"
	"src/github.com/mustafaakilll/go-site-exam/internal/answer"
	"src/github.com/mustafaakilll/go-site-exam/internal/lesson"
	"src/github.com/mustafaakilll/go-site-exam/internal/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	database := db.DB
	userRepository := user.NewUserRepository(database)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	lessonRepository := lesson.NewLessonRepository(database)
	lessonService := lesson.NewLessonService(lessonRepository)
	lessonHandler := lesson.NewLessonHandler(lessonService)

	answerRepository := answer.NewAnswerRepository(database)
	answerService := answer.NewAnswerService(answerRepository)
	answerHandler := answer.NewAnswerHandler(answerService)

	choiceRepository := choice.NewChoiceRepository(database)
	choiceService := choice.NewChoiceService(choiceRepository)
	choiceHandler := choice.NewChoiceHandler(choiceService)

	codeAnswerRepository := codeAnswer.NewCodeAnswerRepository(database)
	codeAnswerService := codeAnswer.NewCodeAnswerService(codeAnswerRepository)
	codeAnswerHandler := codeAnswer.NewCodeAnswerHandler(codeAnswerService)

	codeSubmissionRepository := codeSubmission.NewCodeSubmissionRepository(database)
	codeSubmissionService := codeSubmission.NewCodeSubmissionService(codeSubmissionRepository)
	codeSubmissionHandler := codeSubmission.NewCodeSubmissionHandler(codeSubmissionService)

	codeRepository := code.NewCodeRepository(database)
	codeService := code.NewCodeService(codeRepository)
	codeHandler := code.NewCodeHandler(codeService)

	app := fiber.New()
	api := app.Group("/api/v1")

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	userApi := api.Group("/users")
	userApi.Get("/", userHandler.GetUsers)
	userApi.Get("/:id", userHandler.GetUserByID)
	userApi.Post("/", userHandler.CreateUser)
	userApi.Put("/:id", userHandler.UpdateUser)
	userApi.Delete("/:id", userHandler.DeleteUser)
	userApi.Get("/email/:email", userHandler.GetUserByEmail)
	userApi.Get("/name/:username", userHandler.GetUserByUsername)
	userApi.Get("/teacher/:id", userHandler.SetTeacher)

	lessonAPI := api.Group("/lessons")
	lessonAPI.Get("/", lessonHandler.GetLessons)
	lessonAPI.Post("/", lessonHandler.CreateLessons)

	answerAPI := api.Group("/answers")
	answerAPI.Get("/", answerHandler.GetAnswers)
	answerAPI.Post("/", answerHandler.CreateAnswer)

	choiceAPI := api.Group("/choices")
	choiceAPI.Get("/", choiceHandler.GetChoices)
	choiceAPI.Post("/", choiceHandler.CreateChoices)

	codeAnswerAPI := api.Group("/code-answers")
	codeAnswerAPI.Get("/", codeAnswerHandler.GetCodeAnswers)
	codeAnswerAPI.Post("/", codeAnswerHandler.CreateCodeAnswer)

	codeSubmissionAPI := api.Group("/code-submissions")
	codeSubmissionAPI.Get("/", codeSubmissionHandler.GetCodeSubmissions)
	codeSubmissionAPI.Post("/", codeSubmissionHandler.CreateCodeSubmission)

	codeAPI := api.Group("/codes")
	codeAPI.Get("/", codeHandler.GetCodes)
	codeAPI.Post("/", codeHandler.CreateCodes)

	app.Listen(":8081")
}
