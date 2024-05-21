# Use this code snippet in your app.
# If you need more information about configurations
# or implementing the sample code, visit the AWS docs:
# https://aws.amazon.com/developer/language/python/

# Dependencies
import boto3
from botocore.exceptions import ClientError

def get_secret():
    # This is the secret created via the AWS Console.
    # Please refer to the PNG screenshot for the details.
    secret_name = "soon-secret-name"
    region_name = "us-east-1"

    # Create a Secrets Manager client
    session = boto3.session.Session()
    client = session.client(
        service_name='secretsmanager',
        region_name=region_name
    )

    try:
        get_secret_value_response = client.get_secret_value(
            SecretId=secret_name
        )
    except ClientError as e:
        # For a list of exceptions thrown, see
        # https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
        raise e

    secret = get_secret_value_response['SecretString']

    # Your code goes here.

    # Expected output: {"soon-secret-key":"soon-secret-value"}
    print("Secret: ", secret)

if __name__ == "__main__":
    get_secret()
