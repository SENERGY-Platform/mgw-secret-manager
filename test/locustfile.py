from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    #@task
    #def get_secrets(self):
    #    self.client.get("/secrets")

    @task
    def set_encryption_key(self):
        self.client.post("/key", "eShVmYq3t6w9z$C&E)H@McQfTjWnZr4u")        
