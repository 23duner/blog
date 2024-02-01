package controller

import (
    "github.com/gin-gonic/gin"
    "hw_blog0/dao"
	 openai "github.com/sashabaranov/go-openai"
    "hw_blog0/models"
    "net/http"
	"context"
	"fmt"
	"os"
	"time"
	"strconv"
)



// AddComment handles the addition of a new comment
func AddComment(c *gin.Context) {
    var comment models.Comment
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 使用 OpenAI API 过滤评论内容
    filteredContent, err := filterCommentContent(comment.Content)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter comment content"})
        return
    }
    comment.Content = filteredContent

    // 设置当前时间为评论时间
    comment.Time = time.Now()

    // 将评论添加到数据库
    if err := dao.AddComment(&comment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully", "commentID": comment.ID})
}

// filterCommentContent uses OpenAI API to filter inappropriate content from comments
func filterCommentContent(content string) (string, error) {
    openaiAPIKey := os.Getenv("OPENAI_API_KEY")
    if openaiAPIKey == "" {
        return "", fmt.Errorf("OpenAI API key is not set")
    }

    client := openai.NewClient(openaiAPIKey)

    // Set proxy if needed
    os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")
    os.Setenv("HTTP_PROXY", "http://127.0.0.1:7890")

    resp, err := client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: openai.GPT3Dot5Turbo,
            Messages: []openai.ChatCompletionMessage{
                {Role: "system", Content: "你将会被提供一些博客下的评论，你的任务是将这些评论中的脏话过滤掉，并以文明和易于接受的方式重述出来"},
                {Role: "user", Content: content},
            },
            Temperature: 0.7,
            MaxTokens:   64,
            TopP:        1,
        },
    )
    if err != nil {
        return "", err
    }

    return resp.Choices[0].Message.Content, nil
}

func SelectCommentCount(c *gin.Context) {
    // 从请求中获取 fid 和 module
    fid := c.Query("fid")
    module := c.Query("module")

    // 调用 DAO 层的 SelectCommentCount 方法查询数据库以获取评论数量
    count, err := dao.SelectCommentCount(fid, module)
    if err != nil {
        
        c.JSON(http.StatusInternalServerError, gin.H{"code": "错误码", "msg": err.Error()})
        return
    }

    // 返回评论数量和成功的状态码
    c.JSON(http.StatusOK, gin.H{"code": "200", "data": count})
}

func DeleteComment(c *gin.Context) {
    // 从请求中获取评论的 ID
    commentIDStr := c.Param("id")
    commentID, err := strconv.Atoi(commentIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
        return
    }

    // 调用 DAO 层的 DeleteComment 方法在数据库中删除评论
    if err := dao.DeleteComment(uint(commentID)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
        return
    }

    // 返回成功的状态码
    c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}