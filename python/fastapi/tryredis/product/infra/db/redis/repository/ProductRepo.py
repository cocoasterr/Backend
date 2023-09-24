from models.product import Product


class ProductRepo:
    def __init__(self) -> None:
        self.model = Product

    def create(self, payload: dict):
        try:
            self.model = Product(**payload)
            self.model.save()
            return {"status": "Created"}
        except Exception as e:
            pass

    def update(self, pk: str, payload: dict):
        try:
            data = self.model.get(pk)
        except Exception as e:
            return None
        if data:
            try:
                self.model = Product(**payload)
                self.model.save()
            except Exception as e:
                pass
        return data

    def index(self):
        def _format(pk: str):
            product = self.model.get(pk)
            return {
                "id": product.pk,
                "name": product.name,
                "qty": product.qty,
                "price": product.price,
            }

        try:
            res = [_format(pk) for pk in self.model.all_pks()]
            return res
        except Exception as e:
            pass

    def get_by_pk(self, pk: str):
        return self.model.get(pk)

    def delete_by_pk(self, pk: str):
        self.model.delete(pk)
