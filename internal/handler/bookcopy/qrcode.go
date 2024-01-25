package bookcopyhandler

import (
	"encoding/json"
	"fmt"
	"lms-backend/internal/dataaccess/bookcopy"
	"lms-backend/internal/database"
	"lms-backend/internal/filestorage"
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
	"lms-backend/pkg/error/externalerrors"
	"lms-backend/pkg/error/internalerror"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

const (
	QRCodeFormatString = model.QRCodeFolder + "/book_qr_code_%d.jpeg"
)

func HandleGenerateQRCode(c *fiber.Ctx) error {
	param := c.Params("bookcopy_id")
	bookcopyID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid book id.", param))
	}

	qrcodeFileName := fmt.Sprintf(QRCodeFormatString, bookcopyID)

	exists := filestorage.FileExists(qrcodeFileName)
	if !exists {
		db := database.DB
		bookcopyModel, err := bookcopy.Read(db, bookcopyID)
		if err != nil {
			return err
		}

		bookcopyModel.Status = "-" // The status is not needed for the QR code
		rawJSON, err := json.Marshal(sharedview.ToBookCopyView(bookcopyModel))
		if err != nil {
			return internalerror.InternalServerError("Error marshaling book copy into JSON")
		}

		qrc, err := qrcode.New(string(rawJSON))
		if err != nil {
			return internalerror.InternalServerError("Error generating QR code")
		}

		w, err := standard.New("file_storage/" + qrcodeFileName)
		if err != nil {
			return internalerror.InternalServerError("Error creatig standard writer")
		}

		if err = qrc.Save(w); err != nil {
			return internalerror.InternalServerError("Error saving QR code")
		}
	}

	return c.SendFile("./file_storage/" + qrcodeFileName)
}
