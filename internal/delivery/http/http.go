package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/duyledat197/netfix-backend/pb"
	"github.com/duyledat197/netfix-backend/utils/storage"
	"github.com/duyledat197/netfix-backend/utils/token"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// NewHTTPHandler ...
func NewHTTPHandler(
	ctx context.Context,
	authenticator token.Authenticator,
	authpb pb.AuthServiceServer,
	postpb pb.PostServiceServer,
	categorypb pb.CategoryServiceServer,
	bucket *storage.Bucket,
) (http.Handler, error) {
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(CustomMatcher),
		runtime.WithMetadata(authenticate(authenticator)),
		// runtime.WithForwardResponseOption(FormatHTTPResponse),
	)

	if err := pb.RegisterAuthServiceHandlerServer(ctx, mux, authpb); err != nil {
		return nil, err
	}

	if err := pb.RegisterPostServiceHandlerServer(ctx, mux, postpb); err != nil {
		return nil, err
	}
	if err := pb.RegisterCategoryServiceHandlerServer(ctx, mux, categorypb); err != nil {
		return nil, err
	}
	muxh := http.NewServeMux()
	muxh.Handle("/", allowCORS(mux))
	muxh.Handle("/v1/files", allowCORS(UploadFile(bucket)))

	return muxh, nil
}

func UploadFile(bucket *storage.Bucket) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		file, header, err := r.FormFile("file")
		if err != nil {
			log.Println("get file err", err)
			responseWithJson(w, http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}
		defer file.Close()
		fileName := header.Filename

		link, err := bucket.UploadImage(ctx, file, fileName)
		if err != nil {
			responseWithJson(w, http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}
		responseWithJson(w, http.StatusOK, map[string]interface{}{
			"link": link,
		})
	})
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}
