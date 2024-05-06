# Full Authentication System

This project provides a full authentication system designed for seamless integration into any project. This system leverages the robustness of Go, MongoDB, and AWS S3 BUCKET, making it highly scalable and secure.

## Tech Stack 
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![MongoDB](https://img.shields.io/badge/mongodb-%2347A248.svg?style=for-the-badge&logo=mongodb&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![AWS](https://img.shields.io/badge/aws-%23232F3E.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)

## Features

- Provides robust authentication capabilities including registration, login, and user management.
- Uses Go with MongoDB for efficient and scalable back-end services.
- Containerized with Docker for easy deployment and scalability.

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop) for container management.
- [Go](https://golang.org/dl/) environment set up on your machine.
- [MongoDB](https://www.mongodb.com/try/download/community) installed locally or use MongoDB Atlas for a cloud-based solution.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/XT3RM1NATOR/full_auth.git
    cd full_auth
    ```

2. Set up environment variables:

    - Edit the `.env` file in the root directory. Fill out all the necessary credentials.


3. Configure MongoDB and MINIO settings as per your requirements.
4. Start the project:
    ```bash
    go build && ./cmd
    ```


## License

This project is licensed under the [MIT License](LICENSE).
