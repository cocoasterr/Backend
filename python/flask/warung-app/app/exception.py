class Bad_Request(Exception):
    def __init__(self, message = "Data not found!"):
        self.message = f"error: {message}"
        super().__init__(self.message)


class Internal_Server_error(Exception):
    def __init__(self, message):
        self.message = f"error: {message}"
        super().__init__(self.message)


def handle_internal_server_error(e: str):
    err = {"error": e}
    return err, 500


def handle_bad_request(e: str):
    err = {"error": e}
    return err, 400
