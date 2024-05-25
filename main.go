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
		AllowOrigins: "*",
		AllowMethods: "*",
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

	questionAPI := api.Group("/questions")
	questionAPI.Get("/", questionHandler.GetQuestions)
	questionAPI.Post("/", questionHandler.CreateQuestions)

	quizAPI := api.Group("/quizzes")
	quizAPI.Get("/", quizHandler.GetQuizzes)
	quizAPI.Post("/", quizHandler.CreateQuizzes)

	userTypeAPI := api.Group("/user-types")
	userTypeAPI.Get("/", userTypeHandler.GetUserTypes)
	userTypeAPI.Post("/", userTypeHandler.CreateUserType)

	userAnswerAPI := api.Group("/user-answers")
	userAnswerAPI.Get("/", userAnswerHandler.GetUserAnswers)
	userAnswerAPI.Post("/", userAnswerHandler.CreateUserAnswers)

	userQuizAPI := api.Group("/user-quizzes")
	userQuizAPI.Get("/", userQuizHandler.GetUserQuizzes)
	userQuizAPI.Post("/", userQuizHandler.CreateUserQuizzes)

	app.Listen(":8081")
}
