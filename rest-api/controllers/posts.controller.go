package controllers

import (
	"cesargdd/rest-api/postspb"
	"cesargdd/rest-api/servers"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/status"
)

var svc = servers.PostSvc()

func CreatePost(ctx *fiber.Ctx) error {
	fmt.Println("Create Post rest")
	body := new(postspb.Post)
	err := ctx.BodyParser(body)
	if err != nil {
		return err
	}
	res, err := svc.CreatePost(context.Background(), &postspb.CreatePostRequest{
		User: body.User,
		Content: &postspb.Content{
			Tittle: body.Content.Tittle,
			Body:   body.Content.Body,
			Image:  body.Content.Image,
		},
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
	return ctx.Status(fiber.StatusOK).JSON(res.Post)
}

func UpdatePost(ctx *fiber.Ctx) error {
	fmt.Println("Update Post rest")
	body := new(postspb.Post)
	err := ctx.BodyParser(body)
	if err != nil {
		return err
	}
	res, err := svc.UpdatePost(context.Background(), &postspb.UpdatePostRequest{
		Id: body.Id,
		Content: &postspb.Content{
			Tittle: body.Content.Tittle,
			Body:   body.Content.Body,
			Image:  body.Content.Image,
		},
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
	return ctx.Status(fiber.StatusOK).JSON(res.Post)
}

func DeletePost(ctx *fiber.Ctx) error {
	fmt.Println("Delete Post rest")
	id := ctx.Params("id")
	res, err := svc.DeletePost(context.Background(), &postspb.DeletePostRequest{Id: id})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			ctx.Status(fiber.StatusBadRequest).SendString(e.Message())
		} else {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(res.Post)
}

func Posts(ctx *fiber.Ctx) error {
	fmt.Println("Posts rest")
	res, err := svc.GetPosts(context.Background(), &postspb.GetPostsRequest{})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			ctx.Status(fiber.StatusBadRequest).SendString(e.Message())
		} else {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(res.Post)
}

func Post(ctx *fiber.Ctx) error {
	fmt.Println("Get Post rest")
	id := ctx.Params("id")
	res, err := svc.GetPost(context.Background(), &postspb.GetPostRequest{Id: id})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			ctx.Status(fiber.StatusBadRequest).SendString(e.Message())
		} else {
			ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(res.Post)
}
