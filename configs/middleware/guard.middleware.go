package middleware

import (
	"core/configs"
	connection "core/connections"
	"core/models"
	"core/utils"
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func Authenticate(env *configs.Env) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		tokenString := strings.TrimPrefix(authorization, "Bearer ")
		method := c.Method()
		path := c.Path()

		if (path == utils.LoginPath || path == utils.UserPath) && method == utils.PostMethod {
			return c.Next()
		}

		if tokenString != "" {
			claim, _ := utils.GetUserFromToken(tokenString, env.JwtSecret)

			user := new(models.User)
			result := connection.DB.Find(&user, claim["sub"])

			response := result.First(user)
			if response.Error != nil {
				if errors.Is(response.Error, gorm.ErrRecordNotFound) {
					return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
				}
			}

			type Policy struct {
				Methods pq.StringArray `gorm:"type:text[]"`
				Path    string
			}

			var policies []Policy

			query := connection.DB.Table("policies").Select("policies.ls_method as methods", "policies.tx_path as path")
			query = query.Joins("inner join groups on groups.id = policies.id_group")
			query = query.Joins("inner join user_groups on user_groups.id_group = groups.id")
			query.Where("user_groups.id_user = ?", user.ID).Scan(&policies)

			if len(policies) == 0 {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			}

			for _, policy := range policies {
				// o if tem tres condicoes com '||' por que ele itera em cima da lista policies,
				// se um dos paths for diferente e ele for o ultimo aparecera unauthorized
				// e usar o break para quebrar o loop abruptamente e seria uma má pratica
				if path == policy.Path || policy.Path == utils.AllPaths || utils.InMethod(method, policy.Methods) {
					return c.Next()
				}
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})

		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
}
