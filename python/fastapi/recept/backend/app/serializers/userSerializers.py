def userEntity(user: dict) -> dict:
    """maping value from database

    Args:
        user (dict): entity schema

    Returns:
        dict: get field database
    """
    return {
        "id": user.id,
        "username": user.username,
        "password": user.password,
        "email": user.email,
    }


def getMeResponseEntity(user:dict) ->dict:
    return {
        "id": user['id'],
        "username": user['username'],
        "email": user['email'],
    }


def userResponseEntity(user: dict) -> dict:
    """this function used for response entity

    Args:
        user (dict): entity schema


    Returns:
        dict: field database
    """
    return {
        "id": user.id,
        "username": user.username,
        "email": user.email,
    }

    # name: str
    # birth: Optional[str]
    # sex: Optional[str]
    # profile: Optional[str]
    # phone_number: Optional[str]
def userCreateResponseEntity(user: dict) -> dict:
    """this function used for response entity

    Args:
        user (dict): entity schema


    Returns:
        dict: field database
    """
    return {
        "id": user.id,
        "name": user.name,
        "email": user.email,
    }


def userListEntity(users: list) -> list:
    """response when data need to mapping is morethan one

    Args:
        users (list): list entities schema

    Returns:
        list: get list field from database
    """
    return [userEntity(user) for user in users]
