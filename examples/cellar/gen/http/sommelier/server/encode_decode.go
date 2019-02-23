// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/goa/examples/cellar/design -o
// $(GOPATH)/src/goa.design/goa/examples/cellar

package server

import (
	"context"
	"io"
	"net/http"

	goa "goa.design/goa"
	sommelier "goa.design/goa/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/goa/examples/cellar/gen/sommelier/views"
	goahttp "goa.design/goa/http"
)

// EncodePickResponse returns an encoder for responses returned by the
// sommelier pick endpoint.
func EncodePickResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(sommelierviews.StoredBottleCollection)
		enc := encoder(ctx, w)
		body := NewStoredBottleResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodePickRequest returns a decoder for requests sent to the sommelier pick
// endpoint.
func DecodePickRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body PickRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewPickCriteria(&body)

		return payload, nil
	}
}

// EncodePickError returns an encoder for errors returned by the pick sommelier
// endpoint.
func EncodePickError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "no_criteria":
			res := v.(sommelier.NoCriteria)
			enc := encoder(ctx, w)
			body := NewPickNoCriteriaResponseBody(res)
			w.Header().Set("goa-error", "no_criteria")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "no_match":
			res := v.(sommelier.NoMatch)
			enc := encoder(ctx, w)
			body := NewPickNoMatchResponseBody(res)
			w.Header().Set("goa-error", "no_match")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalSommelierviewsWineryViewToWineryResponseTiny builds a value of type
// *WineryResponseTiny from a value of type *sommelierviews.WineryView.
func marshalSommelierviewsWineryViewToWineryResponseTiny(v *sommelierviews.WineryView) *WineryResponseTiny {
	res := &WineryResponseTiny{
		Name: *v.Name,
	}

	return res
}