### _____Building and Deploying a Go Application with Docker Swarm on Azure______

View the application by opening the following URL in any web browser: 'http://4.248.12.230:8081/'

### 1. Build and Test the Go Application Locally  
### Download the ZIP file from GitHub, extract it, and open the extracted folder in VS Code

1. Open `main.go` in VSCode:

2. Run the Go Application:
   
   go run main.go
   
3. Open a Browser:
   - Navigate to http://localhost:8080 

_____________________________________________________________________________________


### 2. Creating the Docker Image

1. Build the Docker Image:
   
   docker build -t saimeghana1/go-hostname-app:latest .
   
2. Run the Docker Container Locally:
   
   docker run -d -p 8081:8080 saimeghana1/go-hostname-app:latest
   
3. Push the Docker Image to Docker Hub:
   
   docker push saimeghana1/go-hostname-app:latest
_____________________________________________________________________________________

### 3. Setting Up Azure Virtual Machine

1. Log in to the Azure Portal:

2. Create a Debian-Based VM:
   - Use the Azure Portal GUI to create the VM with the specified details.

3. Configure Security Groups:
   - Open the “Networking” tab in the Azure Portal and add inbound rules for:
     - Port 22: For SSH.
     - Port 80: For HTTP .
_____________________________________________________________________________________

### 4. Deploying Application with Docker Swarm

1. Initialize Docker Swarm:
   
   sudo docker swarm init
   
2. Deploy the Service:
   
   sudo docker service create --name myservice --publish 8081:8080 --replicas 4 saimeghana1/go-hostname-app:latest

3. Check the Status of the Deployed Service:
   
   sudo docker service ls
   
4. Update the Service to Use the Latest Image Version:
   
   sudo docker service update --image saimeghana1/go-hostname-app:latest myservice

5. (Optional) Recreate the Service with the Same Settings:
   
   sudo docker service create --name myservice --publish 8081:8080 --replicas 4 saimeghana1/go-hostname-app:latest

6. Test the Application:
   - Open a web browser and navigate to `http://4.248.12.230:8081/`.
_____________________________________________________________________________________

### Contributions

- Shivaji Burgula: Contributed to developing the Go application and assisted in the PowerPoint presentation.
- Jyothi Prasanna Kambam: Worked on Dockerfile creation, pushed the Docker image to Docker Hub, and helped with commands.
- Pranay Sai Chava: Set up the Azure VM, deployed Docker Swarm, and contributed to the Word document.
- Yamini Kunapareddy: Managed security group configuration, performed service testing, and created the PowerPoint presentation.
- Sai Meghana Sangawar: Worked on the Go application’s idea and development, worked on main.go.
- Sushma Kambam: Documented the process and prepared the Word document.
_____________________________________________________________________________________

### Screenshots

![4 Replicas of images](https://github.com/user-attachments/assets/3cfc044d-d7e1-413c-b7bb-09ddcc06f94d)

![Azure Virtual Machine Overview](https://github.com/user-attachments/assets/31ee5610-e221-463d-96f8-fbbed0c61ea9)

![Dockerhub Container image](https://github.com/user-attachments/assets/b2aebb1d-e706-4373-a386-99f115cdc60a)

![Replicas of image](https://github.com/user-attachments/assets/6e92bbc1-99a2-4dfd-8a99-cf41964ebbf9)
