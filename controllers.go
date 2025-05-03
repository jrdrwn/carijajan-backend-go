package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPedagang(ctx *fiber.Ctx) error {
	args := map[string]any{
		"skip":      ctx.Query("skip", "0"),
		"limit":     ctx.Query("limit", "6"),
		"q":         ctx.Query("q", ""),
		"jenis":     ctx.Query("jenis", ""),
		"tag":       ctx.Query("tag", ""),
		"kota":      ctx.Query("kota", ""),
		"latitude":  ctx.Query("latitude", ""),
		"longitude": ctx.Query("longitude", ""),
	}

	if args["latitude"] != "" && args["longitude"] != "" {
		pedagang, _ := FindPedagangAndLocation(DB, args)
		return ctx.JSON(pedagang)
	} else {
		pedagang, _ := FindPedagang(DB, args)
		return ctx.JSON(pedagang)
	}
}

func GetPedagangById(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	pedagangDetail, _ := FindPedagangById(DB, id)
	return ctx.JSON(pedagangDetail)
}

func GetJenisDagangan(ctx *fiber.Ctx) error {
	args := map[string]any{
		"skip":  ctx.Query("skip", "0"),
		"limit": ctx.Query("limit", "6"),
		"q":     ctx.Query("q", ""),
	}
	jenisDagangan, _ := FindJenisDagangan(DB, args)

	serialized := make([]string, len(jenisDagangan))
	for i, jenis := range jenisDagangan {
		serialized[i] = jenis.Value
	}

	return ctx.JSON(serialized)
}
