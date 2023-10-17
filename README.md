# Chatty - A Live Chat Application

![Chatty Logo](https://github.com/liel-almog/chatty/blob/main/public/chatty-logo.png =500x500)

## Table of Contents

- [Chatty - A Live Chat Application](#chatty---a-live-chat-application)
  - [Table of Contents](#table-of-contents)
  - [Background](#background)
  - [Features](#features)
    - [Easy User Identification 👤](#easy-user-identification-)
    - [Variety of Chat Rooms 🗨️](#variety-of-chat-rooms-️)
    - [Real-Time Conversation 💬](#real-time-conversation-)
  - [Technologies 🛠️](#technologies-️)
    - [Backend Powered by Go 🚀](#backend-powered-by-go-)
    - [Frontend in TypeScript ⚙️](#frontend-in-typescript-️)
    - [PostgreSQL for Robust Data Management 🗃️](#postgresql-for-robust-data-management-️)
    - [Docker for Containerization 📦](#docker-for-containerization-)
  - [Usage](#usage)
    - [Get Started with Docker 🐳](#get-started-with-docker-)
  - [License](#license)

## Background

Chatty is a live chat application that enables real-time communication between users. Built with modern web technologies like Go for the backend and TypeScript for the frontend, Chatty offers a seamless chatting experience. Users can select a username, join an existing chat room, and start sending messages immediately. The application uses WebSocket protocol to provide live messaging capabilities.

## Features

### Easy User Identification 👤

Upon entering Chatty, you're prompted to select a username. This name identifies all your messages across different chat rooms, making it easy to keep track of conversations.

    Note: Your chosen username is your virtual identity in the Chatty universe, allowing for an intuitive and user-centric chat experience.

![Demo Chat](https://github.com/liel-almog/chatty/blob/main/public/chatty-intro.png =800x400)

### Variety of Chat Rooms 🗨️

Chatty offers a selection of predefined chat rooms, each catering to different interests and discussions. Simply choose a room and join the conversation.

    Insight: The chat rooms are designed to foster community and segmented discussions, enhancing user engagement and conversation quality.

### Real-Time Conversation 💬

Thanks to the power of WebSocket technology, all messages are updated in real-time. This ensures that you're always up to date with the latest messages.

    Tech Highlight: WebSocket provides full-duplex communication channels over a single, long-lived connection, making your chats instantaneous and lively.

![Demo Chat](https://github.com/liel-almog/chatty/blob/main/public/chatty-demo-chat.png =400x400)

## Technologies 🛠️

### Backend Powered by Go 🚀

The backend is built using Go, a programming language prized for its efficiency and speed, ensuring a smooth and responsive user experience.

    Did You Know?: Go was designed at Google and is optimized for simple, reliable, and efficient software development.

- **Key Dependencies**:
  - [Gorilla WebSocket](https://github.com/gorilla/websocket) - A websocket package that enables live streaming to other clients.
  - [Gin Web Framework](https://github.com/gin-gonic/gin) - A popular web servers package. Used to create HTTP routes.
  - [PGX](https://github.com/jackc/pgx) - PostgreSQL pure driver. Enables SQL escaping to prevent SQL injection.

### Frontend in TypeScript ⚙️

TypeScript serves as the backbone of the frontend, offering strong type checking and modern ES6 features.

    Quick Fact: TypeScript is a superset of JavaScript, providing optional static typing and interface-based programming.

- **Key Dependencies**:
  - [React](https://react.dev/) - The **most** popular library for web and native user interfaces.
  - [React Query](https://tanstack.com/query/latest/docs/react/overview) - A very popular React data fetching package.
  - [Sass](https://sass-lang.com/) - Used to create beautiful styles inside a web application.
  - [React WebSocket](https://www.npmjs.com/package/react-use-websocket) - A library that makes it simple to use WebSockets inside a react application.

### PostgreSQL for Robust Data Management 🗃️

The application leverages PostgreSQL, a powerful, open-source object-relational database system, to manage and store chat data. With its proven architecture, strong community support, and ease of use, PostgreSQL offers reliability and data integrity, making it an ideal choice for Chatty's database needs.

    Interesting Fact: PostgreSQL supports advanced data types and is ACID-compliant, ensuring data consistency and reliability.

### Docker for Containerization 📦

Chatty is containerized using Docker, making it extremely easy to deploy and scale.

    Insight: Docker containers encapsulate all the dependencies, ensuring that the app runs the same regardless of where it's deployed.

## Usage

### Get Started with Docker 🐳

Getting Chatty up and running is as simple as executing a single command, thanks to Docker Compose.

1. Clone the repository
   ```
   git clone https://github.com/liel-almog/chatty.git
   ```
2. Navigate to the project directory
   ```
   cd chatty
   ```
3. Start the application
   ```
   docker compose up
   ```

Visit `http://localhost:8080` to access the Chatty web interface.

## License

This project is licensed under the MIT License - see the [MIT LICENSE](https://github.com/liel-almog/chatty/blob/main/LICENSE) file for details.
