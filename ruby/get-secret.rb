# Use this code snippet in your app.
# If you need more information about configurations or implementing the sample code, visit the AWS docs:
# https://aws.amazon.com/developer/language/ruby/

require_relative 'lib/aws-sdk-secretsmanager'

def get_secret
  client = Aws::SecretsManager::Client.new(region: 'us-east-1')

  begin
    get_secret_value_response = client.get_secret_value(secret_id: 'soon-secret-name')
  rescue StandardError => e
    # For a list of exceptions thrown, see
    # https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
    raise e
  end

  secret = get_secret_value_response.secret_string

  # Your code goes here.

  # Expected output: {"soon-secret-key":"soon-secret-value"}
  printf "Secret: %s\n", secret
end
