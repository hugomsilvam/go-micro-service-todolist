FROM alpine
ADD todo.list.micro-srv /todo.list.micro-srv
ENTRYPOINT [ "/todo.list.micro-srv" ]
