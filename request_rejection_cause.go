// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type RequestRejectionCause uint8

const (
	RequestRejectionCauseSourceArchived RequestRejectionCause = iota + 1
	RequestRejectionCauseNoWebhook
	RequestRejectionCauseVerificationFailed
	RequestRejectionCauseUnsupportedHttpMethod
	RequestRejectionCauseUnsupportedContentType
	RequestRejectionCauseUnparsableJson
	RequestRejectionCausePayloadTooLarge
	RequestRejectionCauseIngestionFatal
	RequestRejectionCauseUnknown
)

func (r RequestRejectionCause) String() string {
	switch r {
	default:
		return strconv.Itoa(int(r))
	case RequestRejectionCauseSourceArchived:
		return "SOURCE_ARCHIVED"
	case RequestRejectionCauseNoWebhook:
		return "NO_WEBHOOK"
	case RequestRejectionCauseVerificationFailed:
		return "VERIFICATION_FAILED"
	case RequestRejectionCauseUnsupportedHttpMethod:
		return "UNSUPPORTED_HTTP_METHOD"
	case RequestRejectionCauseUnsupportedContentType:
		return "UNSUPPORTED_CONTENT_TYPE"
	case RequestRejectionCauseUnparsableJson:
		return "UNPARSABLE_JSON"
	case RequestRejectionCausePayloadTooLarge:
		return "PAYLOAD_TOO_LARGE"
	case RequestRejectionCauseIngestionFatal:
		return "INGESTION_FATAL"
	case RequestRejectionCauseUnknown:
		return "UNKNOWN"
	}
}

func (r RequestRejectionCause) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", r.String())), nil
}

func (r *RequestRejectionCause) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "SOURCE_ARCHIVED":
		value := RequestRejectionCauseSourceArchived
		*r = value
	case "NO_WEBHOOK":
		value := RequestRejectionCauseNoWebhook
		*r = value
	case "VERIFICATION_FAILED":
		value := RequestRejectionCauseVerificationFailed
		*r = value
	case "UNSUPPORTED_HTTP_METHOD":
		value := RequestRejectionCauseUnsupportedHttpMethod
		*r = value
	case "UNSUPPORTED_CONTENT_TYPE":
		value := RequestRejectionCauseUnsupportedContentType
		*r = value
	case "UNPARSABLE_JSON":
		value := RequestRejectionCauseUnparsableJson
		*r = value
	case "PAYLOAD_TOO_LARGE":
		value := RequestRejectionCausePayloadTooLarge
		*r = value
	case "INGESTION_FATAL":
		value := RequestRejectionCauseIngestionFatal
		*r = value
	case "UNKNOWN":
		value := RequestRejectionCauseUnknown
		*r = value
	}
	return nil
}