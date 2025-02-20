# lbi -- A Go Weight Round Robin Load Balancer

A simple yet powerful load balancer implemented in Go. This project demonstrates how to build a load balancer from scratch, featuring weighted round-robin load balancing, health checks, and reverse proxying.

## Features
- **Weighted Round Robin:** Distributes requests based on backend server weights.
- **Health Checks:** Periodically checks the health of backend servers and removes unhealthy ones.
- **Reverse Proxy:** Forwards incoming requests to healthy backend servers.
- **Concurrency:** Uses goroutines for non-blocking health checks.

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/iliasreg/lbi.git
   cd lbi
   ```
   
Build and run the load balancer:

```bash
go run main.go
```

## Usage
Start Backend Servers
Start some backend servers using Python’s built-in HTTP server:

  ```bash
  # Terminal 1
  python3 -m http.server 8081
  
  # Terminal 2
  python3 -m http.server 8082
  
  # Terminal 3
  python3 -m http.server 8083
  ```

## Start the Load Balancer
Run the load balancer:

  ```bash
  go run main.go
  ```

## Send Requests
Send requests to the load balancer using curl or a browser:

  ```bash
  curl http://localhost:8080
  ```

---

## Observe Behavior
The load balancer will forward requests to the backend servers in a weighted round-robin fashion.
If a backend server becomes unhealthy, the load balancer will stop forwarding requests to it.

---

## Contributing
Contributions are welcome! If you’d like to contribute, please follow these steps:
Fork the repository.
Create a new branch for your feature or bugfix.
Commit your changes.
Push your branch and open a pull request.

---

## License
This project is licensed under the MIT License. See the LICENSE file for details.

---

## Acknowledgments
Inspired by Kasvith’s blog post.
Built with ❤️ using Go.

---

### **How to Use**
1. Copy the entire content above.
2. Create a file named `README.md` in the root of your project.
3. Paste the copied content into the `README.md` file.
4. Customize the sections (e.g., replace `iliasreg` with your GitHub username).

Let me know if you need further assistance! 🚀
