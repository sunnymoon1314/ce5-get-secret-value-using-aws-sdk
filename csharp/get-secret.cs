/*
 *	Use this code snippet in your app.
 *	If you need more information about configurations or implementing the sample code, visit the AWS docs:
 *	https://aws.amazon.com/developer/language/net/getting-started
 */

// 21.05.2024 Soon: Project file copied from:
// https://stackoverflow.com/questions/51393395/msbuild-error-msb1003-specify-a-project-or-solution-file

using Amazon;
using Amazon.SecretsManager;
using Amazon.SecretsManager.Model;

namespace Amazon
static async Task GetSecret()
{
    string secretName = "soon-secret-name";
    string region = "us-east-1";

    IAmazonSecretsManager client = new AmazonSecretsManagerClient(RegionEndpoint.GetBySystemName(region));

    GetSecretValueRequest request = new GetSecretValueRequest
    {
        SecretId = secretName,
        VersionStage = "AWSCURRENT", // VersionStage defaults to AWSCURRENT if unspecified.
    };

    GetSecretValueResponse response;

    try
    {
        response = await client.GetSecretValueAsync(request);
    }
    catch (Exception e)
    {
        // For a list of the exceptions thrown, see
        // https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
        throw e;
    }

    string secret = response.SecretString;

    // Your code goes here

    // Expected output: {"soon-secret-key":"soon-secret-value"}
    Console.WriteLine("Secret: ", secret);
}
