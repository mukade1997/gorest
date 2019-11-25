package constructors

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/barber"
	"github.com/kepkin/gorest/internal/generator/translator"
)

func TestMakePathParamsConstructor(t *testing.T) {
	def := translator.TypeDef{
		Name: "IncomeRequestPath",
		Fields: []translator.Field{
			{Name: "UserID", GoType: "int64", Parameter: "user_id", Type: translator.IntegerField},
			{Name: "Role", GoType: "string", Parameter: "role", Type: translator.StringField},
		},
	}

	b := &strings.Builder{}
	if !assert.NoError(t, MakePathParamsConstructor(b, def)) {
		return
	}
	result := strings.NewReader("package api\n" + b.String())

	prettyResult := &strings.Builder{}
	if !assert.NoError(t, barber.PrettifySource(result, prettyResult)) {
		return
	}

	assert.Equal(t, `package api

func MakeIncomeRequestPath(c *gin.Context) (result IncomeRequestPath, errors []FieldError) {
	var err error

	userIdStr, _ := c.Params.Get("user_id")
	result.UserID, err = strconv.ParseInt(userIdStr, 10, 0)
	if err != nil {
		errors = append(errors, NewFieldError(InPath, "user_id", "can't parse as integer", err))
	}

	result.Role, _ = c.Params.Get("role")
	return
}
`, prettyResult.String())
}
