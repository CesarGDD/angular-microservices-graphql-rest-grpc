package controllers

import (
	"cesargdd/rest-api/models"
	"cesargdd/rest-api/tools"
	"cesargdd/rest-api/userspb"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/status"
)

func SignIn(ctx *fiber.Ctx) error {
	body := new(userspb.User)
	err := ctx.BodyParser(body)
	if err != nil {
		return err
	}
	body.Password = tools.HashPassword(body.Password)

	res, err := userSvc.CreateUser(context.Background(), &userspb.CreateUserRequest{
		Username: body.Username,
		Password: body.Password,
	})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			ctx.Status(fiber.StatusBadRequest).SendString(e.Message())
		} else {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}
	token, err := JwtGenerate(context.Background(), res.User.Id)
	if err != nil {
		return err
	}
	authRes := &models.AuthResponse{
		Token: token,
		User:  res.User,
	}
	return ctx.Status(fiber.StatusOK).JSON(authRes)
}

// Login is the resolver for the login field.
func Login(ctx *fiber.Ctx) error {
	body := new(userspb.User)
	err := ctx.BodyParser(body)
	if err != nil {
		return err
	}
	res, err := GetUserByUsername(context.Background(), body.Username)
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			ctx.Status(fiber.StatusBadRequest).SendString(e.Message())
		} else {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}
	if err := tools.ComparePassword(res.Password, body.Password); err != nil {
		return err
	}
	token, err := JwtGenerate(context.Background(), res.Id)
	if err != nil {
		return err
	}
	user := &userspb.User{
		Id:       res.Id,
		Username: res.Username,
	}
	authRes := &models.AuthResponse{
		Token: token,
		User:  user,
	}
	return ctx.Status(fiber.StatusOK).JSON(authRes)
}
