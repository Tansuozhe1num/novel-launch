package user

import (
    "github.com/gin-gonic/gin"
    "novel-launch/novel/Biu"
)

func Profile(ctx *gin.Context) {
    uidVal, _ := ctx.Get("uid")
    unameVal, _ := ctx.Get("uname")
    Biu.Success(ctx, gin.H{
        "uid":      uidVal,
        "userName": unameVal,
    })
}

