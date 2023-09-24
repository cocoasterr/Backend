import os
import psycopg2


class Config:
    def __init__(self):
        self.conn = None

    def connection(self):
        try:
            db = os.getenv("PG_DB")
            user = os.getenv("PG_USER")
            password = os.getenv("PG_PASSWORD")
            host = os.getenv("PG_HOST")
            port = os.getenv("PG_PORT")

            self.conn = psycopg2.connect(
                database=db, user=user, password=password, host=host, port=port
            )
            print("Connected to Postgres!")
            return self.conn
        except Exception as e:
            print(f"Unable to connect Postgres! \n {e}")
