def personEntity(person: dict) -> dict:
    """maping value from database
    Args:
        person (dict): entity schema
    Returns:
        dict: get field database
    """
    return {
        "id": person.id,
        "name": person.name,
        "birth": person.birth,
        "sex": person.sex,
        "profile": person.profile,
        "phone_number": person.phone_number,
    }
def getMePersonEntity(person: dict) -> dict:
    """maping value from database
    Args:
        person (dict): entity schema
    Returns:
        dict: get field database
    """
    return {
        "person_id": person.id,
        "name": person.name,
        "birth": person.birth,
        "sex": person.sex,
        "profile": person.profile,
        "phone_number": person.phone_number,
    }


def personListEntity(person_list: list) -> list:
    """response when data need to mapping is morethan one
    Args:
        person_list (list): list entities schema
    Returns:
        list: get list field from database
    """
    return [personEntity(person) for person in person_list]
