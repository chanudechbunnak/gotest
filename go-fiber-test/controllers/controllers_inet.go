package controllers

import (
	"strconv"
	"strings"

	"regexp"

	"go-fiber-test/database"
	m "go-fiber-test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Factorial(c *fiber.Ctx) error {
	numberParam := c.Params("number")
	number, err := strconv.Atoi(numberParam)
	if err != nil || number < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid number, please provide a non-negative integer",
		})
	}
	result := 1
	for i := 2; i <= number; i++ {
		result *= i
	}
	response := "!" + strconv.Itoa(number) + "=" + strconv.Itoa(result)
	return c.JSON(fiber.Map{
		"result": response,
	})
}

func NickName(c *fiber.Ctx) error {
	taxID := c.Query("tax_id", "")
	if taxID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "tax_id query parameter is required",
		})
	}
	asciiValues := []int{}
	for _, char := range taxID {
		asciiValues = append(asciiValues, int(char))
	}
	return c.JSON(fiber.Map{
		"results": asciiValues,
	})
}

func Register(c *fiber.Ctx) error {
	var req m.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error",
		})
	}

	if !isValidEmail(req.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "กรอกเมล์ให้ถูกต้อง",
		})
	}

	if !isValidUsername(req.Username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ช้อักษรภาษาอังกฤษ (a-z) ตัวเลข (0-9) และเครื่องหมาย (-, _) เท่านั้น เช่น john_doe",
		})
	}

	if len(req.Password) < 6 || len(req.Password) > 20 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ความยาว 6 -20 ตัวอักษร",
		})
	}

	if req.Username == "" || req.Password == "" || req.LineID == "" || req.Tel == "" || req.BusinessType == "" || req.Website == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "กรุณากรอกข้อมูลให้ครบถ้วน",
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(req)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func isValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	return re.MatchString(username)
}

// <------------------------------------- Dogs ------------------------------------->
func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func GetRemoveDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var removedDogs []m.Dogs
	result := db.Unscoped().Where("deleted_at IS NOT NULL").Find(&removedDogs)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(removedDogs)
}

func GetDogsInRange(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	result := db.Where("dog_id > ? AND dog_id < ?", 50, 100).Find(&dogs)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(dogs)
}

func AddDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	var dataResults []m.DogsRes
	for _, v := range dogs {
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}
		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs),
	}
	return c.Status(200).JSON(r)
}

func GetDogCount(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	var dataResults []m.DogsRes
	sumRed := 0
	sumGreen := 0
	sumPink := 0
	sumNoColor := 0

	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sumRed++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sumGreen++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sumPink++
		} else {
			typeStr = "no color"
			sumNoColor++
		}
		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:       dataResults,
		Name:       "golang-test",
		Count:      len(dogs),
		SumRed:     sumRed,
		SumGreen:   sumGreen,
		SumPink:    sumPink,
		SumNoColor: sumNoColor,
	}
	return c.Status(200).JSON(r)
}

// <------------------------------------- Company ------------------------------------->
func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var company []m.Company
	db.Find(&company)
	return c.Status(200).JSON(company)
}

func AddCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&company)
	return c.Status(201).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")
	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company
	result := db.Delete(&company, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}

func GetCompanyId(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company
	result := db.First(&company, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(company)
}

// <------------------------------------- UsersProfile ------------------------------------->

func GetUsersProfiles(c *fiber.Ctx) error {
	db := database.DBConn
	var usersProfile []m.UsersProfile
	db.Find(&usersProfile)
	return c.Status(200).JSON(usersProfile)
}

func AddUsersProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var usersProfile m.UsersProfile
	if err := c.BodyParser(&usersProfile); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&usersProfile)
	return c.Status(201).JSON(usersProfile)
}

func UpdateUsersProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var usersProfile m.UsersProfile
	id := c.Params("id")
	if err := c.BodyParser(&usersProfile); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Where("id = ?", id).Updates(&usersProfile)
	return c.Status(200).JSON(usersProfile)
}

func RemoveUsersProfile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var usersProfile m.UsersProfile
	result := db.Delete(&usersProfile, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
