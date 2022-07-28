package controllers

import (
	"cesargdd/rest-api/servers"
	"cesargdd/rest-api/tools"
	"cesargdd/rest-api/userspb"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/status"
)

var userSvc = servers.UsersSvc()

func CreateUser(ctx *fiber.Ctx) error {
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
	return ctx.Status(fiber.StatusOK).JSON(res.User)
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	res, err := userSvc.GetUser(context.Background(), &userspb.GetUserRequest{Id: id})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			ctx.Status(fiber.StatusBadRequest).SendString(e.Message())
		} else {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(res.User)
}

func GetUserByUsername(ctx context.Context, username string) (*userspb.User, error) {
	res, err := userSvc.GetUserByUsername(ctx, &userspb.GetUserByUsernameRequest{
		Username: username,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res.User, nil
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	res, err := userSvc.DeleteUser(context.Background(), &userspb.DeleteUserRequest{Id: id})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			ctx.Status(fiber.StatusBadRequest).SendString(e.Message())
		} else {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(res.User)
}
