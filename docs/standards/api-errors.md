# API Error Standard

All errors return:

{
"error": {
"code": "STRING_STABLE_CODE",
"message": "Human readable summary",
"details": { "optional": "object" },
"request_id": "uuid"
}
}

Required headers:

- X-Request-Id: uuid (client may send; server must generate if missing)
- X-Correlation-Id: optional (propagate through services + events)

Common codes:

- VALIDATION_FAILED
- UNAUTHORIZED
- FORBIDDEN
- NOT_FOUND
- CONFLICT
- RATE_LIMITED
- INTERNAL
- DEPENDENCY_FAILED
