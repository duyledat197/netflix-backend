package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/duyledat197/netfix-backend/utils/requestinfo"
	"github.com/duyledat197/netfix-backend/utils/token"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/metadata"
)

func authenticate(authenticator token.Authenticator) func(context.Context, *http.Request) metadata.MD {
	return func(ctx context.Context, r *http.Request) metadata.MD {
		md := metadata.MD{}
		authorization := r.Header.Get(requestinfo.Authorization)

		if authorization != "" {
			bearerToken := strings.Split(authorization, requestinfo.Bearer+" ")
			if len(bearerToken) < 2 {
				return md
			}
			token := bearerToken[1]
			payload, err := authenticator.Verify(token)
			if err == nil {
				if m, err := payload.ParseToJSONString(); err != nil {
					md.Append(requestinfo.InfoKey, string(m))
				}
			}
		}
		return md
	}
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, authorization")
		if r.Method != "OPTIONS" {
			h.ServeHTTP(w, r)
		}
	})
}

// CustomMatcher ...
func CustomMatcher(key string) (string, bool) {
	return key, true
}

// ForwardResponse ...
func ForwardResponse(context.Context, http.ResponseWriter, proto.Message) error {
	return nil
}

// AppendRequestMetadata append cookies and headers to incoming context
func AppendRequestMetadata(ctx context.Context, req *http.Request) metadata.MD {
	md := metadata.MD{}

	// Append cookies
	cookies := req.Cookies()
	for _, cookie := range cookies {
		md.Append(cookie.Name, cookie.Value)
	}

	return md
}

// FormatHTTPResponse format http response from proto messages
func FormatHTTPResponse(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	// fmt.Println("format", ctx, w, resp)

	// md, _ := runtime.ServerMetadataFromContext(ctx)
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     k,
	// 	HttpOnly: false,
	// 	Value:    ,
	// 	// Domain:   ".cse30.io",
	// 	Path: "/",
	// 	// Expires:  time.Now().Add(2 * time.Hour),
	// 	// Secure:   true,
	// })
	return nil
}
