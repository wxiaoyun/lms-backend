package googlebook

import (
	"encoding/json"
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/view/bookview"
	"lms-backend/internal/view/googlebookview"
	"lms-backend/pkg/error/externalerrors"
	"lms-backend/util/sliceutil"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HandleQuery(c *fiber.Ctx) error {
	queries := c.Queries()
	queryArr := []string{}

	inTitle := queries["title"]
	if len(inTitle) > 0 {
		inTitle = strings.ReplaceAll(inTitle, " ", "+")
		queryArr = append(queryArr, "intitle:"+inTitle)
	}

	inAuthor := queries["author"]
	if len(inAuthor) > 0 {
		inAuthor = strings.ReplaceAll(inAuthor, " ", "+")
		queryArr = append(queryArr, "inauthor:"+inAuthor)
	}

	inPublisher := queries["publisher"]
	if len(inPublisher) > 0 {
		inPublisher = strings.ReplaceAll(inPublisher, " ", "+")
		queryArr = append(queryArr, "inpublisher:"+inPublisher)
	}

	isbn := queries["isbn"]
	if len(isbn) > 0 {
		queryArr = append(queryArr, "isbn:"+isbn)
	}

	query := strings.Join(queryArr, "+")

	if len(query) == 0 {
		return externalerrors.BadRequest("No query is provided.")
	}

	url := fmt.Sprintf("%s/volumes?q=%s&key=%s", googleAPIBaseURL, query, GetGoogleAPIKey())

	agent := fiber.Get(url)
	statusCode, body, errs := agent.Bytes()
	if errs != nil {
		errStrs := sliceutil.Map(errs, func(err error) string {
			return err.Error()
		})
		return externalerrors.UnprocessableEntity(
			fmt.Sprintf("Failed to query Google Books API: %s", strings.Join(errStrs, ", ")),
		)
	}

	if statusCode != fiber.StatusOK {
		return externalerrors.UnprocessableEntity(
			fmt.Sprintf("Failed to query Google Books API: %s", string(body)),
		)
	}

	var response googlebookview.ResponseView
	if err := json.Unmarshal(body, &response); err != nil {
		return externalerrors.UnprocessableEntity(
			fmt.Sprintf("Failed to unmarshal Google Books API response: %s", err.Error()),
		)
	}

	return c.JSON(api.Response{
		Data: bookview.GoogleResponseToView(&response),
		Messages: api.Messages(
			api.SilentMessage("Successfully queried Google Books API."),
		),
	})
}
