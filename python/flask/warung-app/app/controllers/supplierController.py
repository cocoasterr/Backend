from flask import Blueprint, request, jsonify, request
from flask_jwt_extended import jwt_required, get_jwt_identity
from exception import *
from services.supplierService import SupplierService


supplier_bp = Blueprint("supplier_bp", __name__)
supplier_service = SupplierService()


@supplier_bp.route("/supplier", methods=["POST"])
@jwt_required()
def create_product():
    try:
        product_data = request.get_json()
        user = get_jwt_identity()
        for product in product_data:
            product["username"] = user["username"]
        supplier_service.bulk_create_service(product_data)
        return jsonify({"message": "success!"}), 201
    except Exception as e:
        return handle_internal_server_error(str(e))


@supplier_bp.route("/supplier/<id>", methods=["PUT"])
@jwt_required()
def update_product(id):
    try:
        product_data = request.get_json()
        user = get_jwt_identity()
        product_data["username"] = user["username"]
        supplier_service.update_service(product_data, id)
        return jsonify({"message": "success!"}), 200
    except Bad_Request as e:
        return handle_bad_request(str(e))
    except Exception as e:
        return handle_internal_server_error(str(e))
        



@supplier_bp.route("/supplier/<id>", methods=["DELETE"])
def delete_product(id):
    try:
        supplier_service.delete_service(id)
        return jsonify({"message": "success!"}), 200
    except Bad_Request as e:
        return handle_bad_request(str(e))
    except Exception as e:
        return handle_internal_server_error(str(e))


@supplier_bp.route("/supplier/<id>", methods=["GET"])
@jwt_required()
def find_by_id_product(id):
    try:
        res = supplier_service.find_by_id_service(id)
        return jsonify({"message": "success!", "data": [res]}), 200
    except Bad_Request as e:
        return handle_bad_request(str(e))
    except Exception as e:
        return handle_internal_server_error(str(e))


@supplier_bp.route("/supplier", methods=["GET"])
@jwt_required()
def index_product():
    try:
        page = int(request.args.get("page", 1))
        limit = int(request.args.get("limit", 10))
        res, total = supplier_service.index_service(page, limit)
        return jsonify({"message": "success!", "data": [res], "total": total}), 200
    except Exception as e:
        return handle_internal_server_error(str(e))
