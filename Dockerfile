# Use a specific Go version to ensure consistency
FROM golang:1.22.3

# Set the working directory inside the container
WORKDIR /ascii-art-web-dockerize

# Add metadata to the image
LABEL maintainer="VON"
LABEL version="1.0"
LABEL description="ascii art web dockerize"

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build 

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD [ "./ascii" ]