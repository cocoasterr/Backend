from models.order import Order
import time


class OrderRepo:
    def __init__(self) -> None:
        self.model = Order

    def create(self, payload: dict):
        try:
            self.model = self.model(**payload)
            self.model.save()
        except Exception as e:
            raise e

    def update(self, pk: str, payload: dict):
        try:
            data = self.model.get(pk)
        except Exception as e:
            return None
        if data:
            try:
                data = self.model(**payload)
                data.save()
            except Exception as e:
                raise e
        return data

    def index(self):
        def _format(pk: str):
            data = self.model.get(pk)
            return {
                "pk": data.pk,
                "product_id": data.product_id,
                "price": data.price,
                "qty": data.qty,
                "price": data.price,
                "total": data.total,
                "fee": data.fee,
                "status": data.status,
            }

        try:
            res = [_format(pk) for pk in self.model.all_pks()]
            return res
        except Exception as e:
            raise e

    def get_by_pk(self, pk: str):
        return self.model.get(pk)

    def delete_by_pk(self, pk: str):
        self.model.delete(pk)

    def order(self, r: dict):
        def complete_order(data):
            time.sleep(10)
            data.status = "completed"
            data.save()

        try:
            data = self.model(**r)
            data.save()
            complete_order(data)
        except Exception as e:
            raise e
