# Use an official Node.js image as a parent image
FROM node:16

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json (or yarn.lock) from your project into the working directory
COPY package*.json ./

# Install dependencies in the container
RUN npm install

# Copy the remaining files from your project into the working directory
COPY . .

# Expose port 3000 to the outside once the container is running
EXPOSE 3000

# Run the command inside your image filesystem
CMD ["npm", "start"]
