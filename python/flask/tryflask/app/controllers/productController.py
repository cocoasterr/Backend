from flask import Blueprint, request, jsonify, request
from services.productService import ProductService
from exceptions import *
from flask_jwt_extended import jwt_required, get_jwt_identity

product_bp = Blueprint("product_bp", __name__)
product_service = ProductService()


@product_bp.route("/products", methods=["POST"])  # type: ignore
@jwt_required()
def create_product():
    try:
        product_data = request.get_json()
        user = get_jwt_identity()
        for product in product_data:
            product["username"] = user["username"]
        product_service.bulk_create_service(product_data)
        return jsonify({"message": "success!"}), 201
    except Exception as e:
        handle_internal_server_error_exception(e)


@product_bp.route("/products/<id>", methods=["PUT"])  # type: ignore
@jwt_required()
def update_product(id):
    try:
        product_data = request.get_json()
        user = get_jwt_identity()
        product_data["username"] = user["username"]
        product_service.update_service(product_data, id)
        return jsonify({"message": "success!"}), 201
    except InternalServerErrorException as e:
        handle_internal_server_error_exception(e)
    except NotFoundException as e:
        handle_not_found_exception(e)


@product_bp.route("/products/<id>", methods=["DELETE"])  # type: ignore
def delete_product(id):
    try:
        product_service.delete_service(id)
        return jsonify({"message": "success!"}), 201
    except NotFoundException as e:
        handle_not_found_exception(e)
    except InternalServerErrorException as e:
        handle_internal_server_error_exception(e)


@product_bp.route("/products/<id>", methods=["GET"])  # type: ignore
@jwt_required()
def find_by_id_product(id):
    try:
        res = product_service.find_by_id_service(id)
        return jsonify({"message": "success!", "data": [res]}), 200
    except NotFoundException as e:
        handle_not_found_exception(e)
    except InternalServerErrorException as e:
        return handle_internal_server_error_exception(str(e))


@product_bp.route("/products", methods=["GET"])  # type: ignore
@jwt_required()
def index_product():
    try:
        page = int(request.args.get("page", 1))
        limit = int(request.args.get("limit", 10))
        res, total = product_service.index_service(page, limit)
        return jsonify({"message": "success!", "data": [res], "total": total}), 200
    except InternalServerErrorException as e:
        handle_internal_server_error_exception(e)
