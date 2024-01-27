
# Auth-App

Welcome to Auth-App, a serverless authentication app using AWS Lambda and API Gateway.

## Quick Start

### Clone the Repository :
```
git clone https://github.com/yourusername/auth-app.git

cd auth-app
```
## Installation
```
go mod download
```
Deploy to AWS Lambda: 
```
cd deploy
cdk deploy
```

## Deploy using the AWS CLI or your preferred method

Explore the App:

Open the provided AWS API Gateway URL to start using the authentication app.

### Technologies
   
  - Go
   
- AWS Lambda
   
 - AWS API Gateway
   
  ### Features
   - User Registration

   - User Verification
  
   - User Login

### Secure Authentication Flow

**1. POST /register :** Register a new user.

**2. POST /login :** Authenticate and log in a user.

**3. POST /otp :** Generate and validate OTP (One-Time Password).
