class existing_data(Exception):
    def __init__(self, message="Data already!") -> None:
        super().__init__(message)


class internal_server_error(Exception):
    def __init__(self, message) -> None:
        self.message = message
        super().__init__(self.message)


class not_found_error(Exception):
    def __init__(self, message="Data not found!") -> None:
        super().__init__(message)
