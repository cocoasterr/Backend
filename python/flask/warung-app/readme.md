# Warung App Backend
Backend side warung app
## BUILD Docker Image
- Build image
    ```bash
        sudo docker build -t warung:1.0.0 .
- Make sure a docker image and container
    ```bash
        sudo docker ps
        sudo docker image
- Run Docker Image
    ```bash
    docker run -p 5000:5000 <IMAGE ID>
## Import Postmant Collection
- Open your Postman
- Copy file postman_collection.txt
- Click import
- Insert access
- Run API Postman