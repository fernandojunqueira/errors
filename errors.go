package errors

import (
	"fmt"
	"net/http"
)

// ErrorRFC9457 representa o formato padrão de erros definido pela RFC 9457
// ("Problem Details for HTTP APIs"). Esse formato fornece uma maneira
// consistente e extensível de retornar informações de erro em APIs HTTP.
type ErrorRFC9457 struct {
	//   - Type: URI que identifica o tipo de problema. É o identificador
	//     canônico do erro. Se não for fornecido, assume-se "about:blank".
	//     Recomenda-se usar URIs absolutas e estáveis (ex.:
	//     "https://api.exemplo.com/problems/out-of-credit").
	Type string `json:"type"`

	//   - Status: Código de status HTTP gerado para esse problema.
	//     Geralmente corresponde ao código HTTP retornado na resposta
	//     (ex.: 400, 404, 403, 500).
	Status int `json:"status"`

	//   - Title: Resumo curto e legível por humanos do tipo de problema.
	//     Normalmente é uma descrição genérica (ex.: "Out of credit").
	//     Esse valor deve ser o mesmo para todas as ocorrências de um mesmo tipo.
	Title string `json:"title"`

	//   - Detail: Explicação legível por humanos específica para a
	//     ocorrência do problema. Deve fornecer contexto adicional
	//     além do campo Title (ex.: "Seu saldo atual é 30, mas a operação custa 50.").
	Detail string `json:"detail"`

	//   - Instance: URI que identifica a ocorrência específica do problema.
	//     Serve como referência única para rastreamento do erro
	//     (ex.: "/account/12345/transactions/abc"). Pode ser útil para logs e debugging.
	Instance string `json:"instance"`
}

func (e *ErrorRFC9457) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("Error: %s, Title: %s, Detail: %s, Instance: %s", e.Type, e.Title, e.Detail, e.Instance)
}

func InternalServerError(err error, title string) *ErrorRFC9457 {
	return &ErrorRFC9457{
		Status: http.StatusInternalServerError,
		Detail: err.Error(),
		Title:  title,
	}
}

func NotFoundError(err error, title string) *ErrorRFC9457 {
	return &ErrorRFC9457{
		Status: http.StatusNotFound,
		Detail: err.Error(),
		Title:  title,
	}
}

func BadRequest(err error, title string) *ErrorRFC9457 {
	return &ErrorRFC9457{
		Status: http.StatusBadRequest,
		Detail: err.Error(),
		Title:  title,
	}
}

func BadGateway(err error, title string) *ErrorRFC9457 {
	return &ErrorRFC9457{
		Status: http.StatusBadGateway,
		Detail: err.Error(),
		Title:  title,
	}
}

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
