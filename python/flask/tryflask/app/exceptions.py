from flask import jsonify, make_response


class AlreadyExistsException(Exception):
    def __init__(self, data):
        self.data = data
        self.message = f"{data} is already exist!"
        super().__init__(self.message)


class InvalidInputException(Exception):
    def __init__(self, message):
        self.message = f"Invalid input: {message}"
        super().__init__(self.message)


class DatabaseErrorException(Exception):
    def __init__(self, message):
        self.message = message
        super().__init__(self.message)


class InternalServerErrorException(Exception):
    def __init__(self, message):
        self.message = message
        super().__init__(self.message)


class NotFoundException(Exception):
    def __init__(self, message="data not found!"):
        self.message = message
        super().__init__(self.message)


def handle_invalid_input(e):
    err = jsonify({"message": "already exist", "error": str(e)})
    return make_response(err, 400)


def handle_not_found_exception(e):
    err = jsonify({"message": "not found", "error": str(e)})
    return make_response(err, 404)


def handle_internal_server_error_exception(e):
    err = {"message": "internal server error", "error": str(e)}
    return err, 500
