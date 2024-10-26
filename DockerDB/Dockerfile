FROM container-registry.oracle.com/database/free:latest-lite

COPY scripts/*.sql /opt/oracle/scripts/startup

#docker build -t reservardb:v1 .
#docker run --name=reservasdb -p 1521:1521 -e ORACLE_PDB=RESERVASDB -e ORACLE_PWD=123 reservardb:v1
#docker ps -a
#docker ps
#docker start name_container
#docker stop name_container
#docker rm