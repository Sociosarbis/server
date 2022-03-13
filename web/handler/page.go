// Copyright (c) 2022 Trim21 <trim21.me@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/bangumi/server/internal/errgo"
)

const defaultPageLimit = 30
const defaultMaxPageLimit = 100

type pageQuery struct {
	Limit  int
	Offset int
}

func (q pageQuery) check(count int64) error {
	if q.Limit > int(count) {
		return fiber.NewError(fiber.StatusBadRequest,
			"limit should less equal than "+strconv.FormatInt(count, 10))
	}

	if int64(q.Offset) >= count {
		return fiber.NewError(fiber.StatusBadRequest, "offset is greater than count")
	}

	return nil
}

func getPageQuery(c *fiber.Ctx, defaultLimit int, maxLimit int) (pageQuery, error) {
	q := pageQuery{Limit: defaultLimit}
	var err error

	raw := c.Query("limit")
	if raw != "" {
		q.Limit, err = strconv.Atoi(raw)
		if err != nil {
			return q, errgo.Wrap(err, "ParseUint")
		}

		if q.Limit > maxLimit {
			return q, fiber.NewError(fiber.StatusBadRequest,
				"limit should less equal than "+strconv.Itoa(maxLimit))
		}
	}

	raw = c.Query("offset")
	if raw != "" {
		q.Offset, err = strconv.Atoi(raw)
		if err != nil {
			return q, errgo.Wrap(err, "ParseUint")
		}
	}

	return q, nil
}
