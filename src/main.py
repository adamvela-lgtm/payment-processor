import logging
from payment_processor.config import Config
from payment_processor.services import PaymentProcessorService
from payment_processor.db import Database

def main():
    # Initialize the logger
    logging.basicConfig(level=logging.INFO)
    logger = logging.getLogger()

    # Load the configuration
    config = Config()

    # Create a database connection
    db = Database(config.database_url)

    # Create a payment processor service
    service = PaymentProcessorService(db, config)

    # Start the payment processor
    service.start()

if __name__ == "__main__":
    main()