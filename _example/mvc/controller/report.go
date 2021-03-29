/**
 * **************************************************************************
 *                                                                          *
 * Copyright (c) 2019 Baidu.com, Inc. All Rights Reserved                   *
 *                                                                          *
 * **************************************************************************
 * Created by Vscode.
 * @author AImager(yucong04@baidu.com)
 * Created Date: 2019-07-09
 * Created Time: 23:58
 */

package report

import (
	"context"

	"github.com/AImager/wool"
	"github.com/AImager/wool/_example/mvc/model"
)

type ReportTaskCtr struct {
	Type string `json:"type" form:"type" binding:"required"`
}

func init() {
	_ = wool.DI().Bind(new(model.ModelI), func(name string) model.ModelI {
		switch name {
		case "TYPE1":
			return &model.ModelOne{}
		case "TYPE2":
			return &model.ModelTwo{}
		default:
			return *new(model.ModelI)
		}
	})
}

func (ctr ReportTaskCtr) Main(ctx *context.Context) (res interface{}, err error) {
	var mi model.ModelI
	_ = wool.DI().Make(&mi, ctr.Type)
	return mi.Print(), nil
}
