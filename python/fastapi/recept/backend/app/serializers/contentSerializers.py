def contentEntity(content: dict) -> dict:
    """maping value from database
    Args:
        content (dict): entity schema
    Returns:
        dict: get field database
    """
    return {
        "id": content.id,
        "user_id": content.user_id,
        "username": content.username,
        "title": content.title,
        "like": content.like,
        "comment": content.comment,
        "status": content.status,
        "ingridient": content.ingridient,
        "status": content.status,
        "intructions": content.intructions,
        "created_at": content.created_at,
        "updated_at": content.updated_at,
    }


def contentListEntity(content_list: list) -> list:
    """response when data need to mapping is morethan one
    Args:
        content_list (list): list entities schema
    Returns:
        list: get list field from database
    """
    return [contentEntity(content) for content in content_list]
