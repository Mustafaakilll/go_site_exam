package main

import (
	"github.com/mustafaakilll/go-site-exam/db"
	"github.com/mustafaakilll/go-site-exam/internal/answer"
	"github.com/mustafaakilll/go-site-exam/internal/choice"
	"github.com/mustafaakilll/go-site-exam/internal/code"
	codeAnswer "github.com/mustafaakilll/go-site-exam/internal/code-answer"
	codeSubmission "github.com/mustafaakilll/go-site-exam/internal/code-submission"
	"github.com/mustafaakilll/go-site-exam/internal/lesson"
	"github.com/mustafaakilll/go-site-exam/internal/question"
	"github.com/mustafaakilll/go-site-exam/internal/quiz"
	"github.com/mustafaakilll/go-site-exam/internal/types"
	"github.com/mustafaakilll/go-site-exam/internal/user"
	userAnswer "github.com/mustafaakilll/go-site-exam/internal/user-answer"
	userQuiz "github.com/mustafaakilll/go-site-exam/internal/user-quiz"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	questionRepository := question.NewQuestionRepository(database)
	questionService := question.NewQuestionService(questionRepository)
	questionHandler := question.NewQuestionHandler(questionService)

	quizRepository := quiz.NewQuizRepository(database)
	quizService := quiz.NewQuizService(quizRepository)
	quizHandler := quiz.NewQuizHandler(quizService)

	userTypeRepository := types.NewUserTypeRepository(database)
	userTypeService := types.NewUserTypeService(userTypeRepository)
	userTypeHandler := types.NewUserTypeHandler(userTypeService)

	userAnswerRepository := userAnswer.NewUserAnswerRepository(database)
	userAnswerService := userAnswer.NewUserAnswerService(userAnswerRepository)
	userAnswerHandler := userAnswer.NewUserAnswerHandler(userAnswerService)

	userQuizRepository := userQuiz.NewUserQuizRepository(database)
	userQuizService := userQuiz.NewUserQuizService(userQuizRepository)
	userQuizHandler := userQuiz.NewUserQuizHandler(userQuizService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "*",
		AllowCredentials: true,
	}))
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
	userApi.Get("/student", userHandler.GetStudents)

	lessonAPI := api.Group("/lessons")
	lessonAPI.Get("/", lessonHandler.GetLessons)
	lessonAPI.Post("/", lessonHandler.CreateLessons)
	lessonAPI.Delete("/:id", lessonHandler.DeleteLessons)
	lessonAPI.Put("/", lessonHandler.UpdateLessons)

	answerAPI := api.Group("/answers")
	answerAPI.Get("/", answerHandler.GetAnswers)
	answerAPI.Post("/", answerHandler.CreateAnswer)
	answerAPI.Put("/", answerHandler.UpdateAnswer)
	answerAPI.Delete("/", answerHandler.DeleteAnswer)

	choiceAPI := api.Group("/choices")
	choiceAPI.Get("/", choiceHandler.GetChoices)
	choiceAPI.Post("/", choiceHandler.CreateChoices)
	choiceAPI.Put("/", choiceHandler.UpdateChoices)
	choiceAPI.Delete("/:id", choiceHandler.UpdateChoices)

	codeAnswerAPI := api.Group("/code-answers")
	codeAnswerAPI.Get("/", codeAnswerHandler.GetCodeAnswers)
	codeAnswerAPI.Post("/", codeAnswerHandler.CreateCodeAnswer)
	codeAnswerAPI.Delete("/:id", codeAnswerHandler.DeleteCodeAnswer)
	codeAnswerAPI.Put("/", codeAnswerHandler.UpdateCodeAnswer)

	codeSubmissionAPI := api.Group("/code-submissions")
	codeSubmissionAPI.Get("/", codeSubmissionHandler.GetCodeSubmissions)
	codeSubmissionAPI.Post("/", codeSubmissionHandler.CreateCodeSubmission)
	codeSubmissionAPI.Delete("/:id", codeSubmissionHandler.DeleteCodeSubmission)
	codeSubmissionAPI.Put("/", codeSubmissionHandler.UpdateCodeSubmission)

	codeAPI := api.Group("/codes")
	codeAPI.Get("/", codeHandler.GetCodes)
	codeAPI.Post("/", codeHandler.CreateCode)
	codeAPI.Put("/", codeHandler.UpdateCode)
	codeAPI.Delete("/:id", codeHandler.DeleteCode)

	questionAPI := api.Group("/questions")
	questionAPI.Get("/", questionHandler.GetQuestions)
	questionAPI.Post("/", questionHandler.CreateQuestions)
	questionAPI.Put("/", questionHandler.UpdateQuestion)
	questionAPI.Delete("/:id", questionHandler.DeleteQuestion)

	quizAPI := api.Group("/quizzes")
	quizAPI.Get("/", quizHandler.GetQuizzes)
	quizAPI.Post("/", quizHandler.CreateQuizzes)
	quizAPI.Put("/", quizHandler.UpdateQuiz)
	quizAPI.Delete("/:id", quizHandler.DeleteQuiz)

	userTypeAPI := api.Group("/user-types")
	userTypeAPI.Get("/", userTypeHandler.GetUserTypes)
	userTypeAPI.Post("/", userTypeHandler.CreateUserType)

	userAnswerAPI := api.Group("/user-answers")
	userAnswerAPI.Get("/", userAnswerHandler.GetUserAnswers)
	userAnswerAPI.Post("/", userAnswerHandler.CreateUserAnswers)
	userAnswerAPI.Delete("/:id", userAnswerHandler.DeleteUserAnswer)
	userAnswerAPI.Put("/", userAnswerHandler.UpdateUserAnswer)

	userQuizAPI := api.Group("/user-quizzes")
	userQuizAPI.Get("/", userQuizHandler.GetUserQuizzes)
	userQuizAPI.Post("/", userQuizHandler.CreateUserQuiz)
	userQuizAPI.Put("/", userQuizHandler.UpdateUserQuiz)
	userQuizAPI.Delete("/:id", userQuizHandler.DeleteUserQuiz)

	app.Listen(":8081")
}
