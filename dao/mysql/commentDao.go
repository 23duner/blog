package dao

import (
    "hw_blog0/models"
    

)

func AddComment(comment *models.Comment) error {
    // SQL 插入语句
    stmt, err := db.Prepare("INSERT INTO comments(user_id, pid, root_id, content, time, fid, module) VALUES (?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    // 执行SQL语句
    result, err := stmt.Exec(comment.UserID, comment.PID, comment.RootID, comment.Content, comment.Time, comment.FID, comment.Module)
    if err != nil {
        return err
    }

    // 获取并设置新插入评论的ID
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    comment.ID = uint(id)

    return nil
}
