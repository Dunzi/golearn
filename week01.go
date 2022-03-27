package main

import "errors"

//作业：1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
//是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

//不应该抛给上层，没有记录应该给空值
//括弧：现在我们项目的dao层逻辑就是给的 error，导致很多业务中正常可以为空的逻辑中，还要针对错误码为空做处理

type ErrNoRows struct {
	Code int //code != 0 返回错误 例如 code = 1 记录为空
	Back string
}

func (e *ErrNoRows) Error() string {
	//错误来了
	sqlNoRows := &ErrNoRows{
		Code: 1,
		Back: "record not found",
	}

	if errors.Is(e, sqlNoRows) {
		return "record not found" //正常无数据逻辑
	}
	if e.Code != 0 {
		return e.Back
	}

	return "list"
}
