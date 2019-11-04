FROM        scratch

COPY ghcs /bin/ghcs

ENTRYPOINT [ "/bin/ghcs" ]