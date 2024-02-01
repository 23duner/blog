
package models


type Comment struct {
    ID      uint   `json:"id"`         // 评论的唯一标识符
    PID     uint   `json:"pid"`        // 父级评论的ID，表示该评论的父级评论ID
    RootID  uint   `json:"rootId"`     // 根评论的ID，表示整个评论线的根节点ID
    Content string `json:"content"`    // 评论内容
    FID     uint   `json:"fid"`        // 所属的博客或文章的ID
    Module  string `json:"module"`     // 所属的模块或类别
}
