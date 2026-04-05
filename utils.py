import os
import json
import hashlib
import uuid
from typing import Dict

class PaymentProcessor:
    def __init__(self, payment_gateway: str, secret_key: str):
        self.payment_gateway = payment_gateway
        self.secret_key = secret_key
        self.payment_gateway_url = f"https://{payment_gateway}/api/v1"

    def create_payment(self, amount: float, currency: str, recipient: str, payment_method: str, payment_type: str, payment_details: Dict[str, str] = {}):
        data = {
            "amount": amount,
            "currency": currency,
            "recipient": recipient,
            "payment_method": payment_method,
            "payment_type": payment_type,
            "payment_details": payment_details
        }
        response = self._post_request(self.payment_gateway_url, json=data)
        return response

    def _post_request(self, url: str, data: Dict[str, str]) -> Dict[str, str]:
        headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {self.secret_key}"
        }
        response = requests.post(url, headers=headers, json=data)
        response.raise_for_status()
        return response.json()

    def get_payment_status(self, payment_id: str) -> Dict[str, str]:
        response = self._get_request(f"{self.payment_gateway_url}/payments/{payment_id}")
        return response

    def _get_request(self, url: str) -> Dict[str, str]:
        headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {self.secret_key}"
        }
        response = requests.get(url, headers=headers)
        response.raise_for_status()
        return response.json()