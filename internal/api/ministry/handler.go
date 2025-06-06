package ministry

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/ministry/dto"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/ministry/dto/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/request"
)

type MinistryHandler interface {
	All() http.HandlerFunc
	Apply() http.HandlerFunc
}
type handler struct {
	logger          *zap.Logger
	MinistryService ministry.MinistryService
}

func NewMinistryHandler(
	logger *zap.Logger,
	ministryService ministry.MinistryService,
) MinistryHandler {
	return &handler{
		logger:          logger,
		MinistryService: ministryService,
	}
}

const (
	InternalServerErrorMsg = "Internal server error. Please try again later."
)

func (mh *handler) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ministries, err := mh.MinistryService.All(ctx)
		if err != nil {
			mh.logger.Sugar().Infow("Failed to get all ministries", "error", err)
			render.Json(w, http.StatusInternalServerError, mappers.ToGetAllMinistriesResponse(nil, "Failed to fetch ministries"))

			return
		}

		resp := mappers.ToGetAllMinistriesResponse(ministries, "Successfully fetched ministries")
		render.Json(w, http.StatusOK, resp)
	}
}

func (mh *handler) Apply() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.MinistryApplicationRequest
		if err := request.DecodeAndValidate(r.Body, &req); err != nil {
			mh.logger.Sugar().Errorw("failed to decode and validate ministry application request", "error", err)
			render.Json(w, http.StatusBadRequest,
				mappers.ToMinistryApplicationResponse(
					"Please double check your application and try again",
				),
			)

			return
		}

		// Get the user ID from the context
		userID, err := context.GetUserIDString(ctx)
		if err != nil {
			mh.logger.Sugar().Errorw("User ID not found in session", "error", err)
			render.Json(
				w,
				http.StatusUnauthorized,
				mappers.ToMinistryApplicationResponse(
					"Unauthorized",
				),
			)

			return
		}

		err = mh.MinistryService.SendApplication(ctx, userID, req.MinistryID, req.Reason)
		if err != nil {
			mh.logger.Sugar().Errorw("Failed to send application", "error", err)

			statusCode, errMsg := mapErrorsToStatusCodeAndUserFriendlyMessages(err)

			render.Json(
				w,
				statusCode,
				mappers.ToMinistryApplicationResponse(
					errMsg,
				),
			)

			return
		}

		render.Json(
			w,
			http.StatusOK,
			mappers.ToMinistryApplicationResponse(
				"Your application was successfully sent.",
			),
		)
	}
}

func mapErrorsToStatusCodeAndUserFriendlyMessages(err error) (int, string) {
	switch {
	case errors.Is(err, ministry.ErrMinistryNotFound):
		return http.StatusNotFound, "The selected ministry could not be found."
	default:
		return http.StatusInternalServerError, InternalServerErrorMsg
	}
}
