FROM drone/ca-certs

ADD src/voter /
CMD ["/voter"]
