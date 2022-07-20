# server-full:4.181
#FROM        payara/server-full:4.181

#ARG         PAYARA_HOME=$PAYARA_PATH
#ARG         DOMAIN_NAME=domain1
#ARG         JDK_VER=8
#ENV         JAVA_HOME         /usr/lib/jvm/java-8-openjdk-amd64
#ENV         PASSWORD_FILE     $PAYARA_HOME/passwordFile
#RUN         echo "AS_ADMIN_PASSWORD=${ADMIN_PASSWORD}" >> ${PASSWORD_FILE}


# server-full:5.2020.5
#FROM        payara/server-full:5.2020.5

#ARG         PAYARA_HOME=$PAYARA_DIR
#ARG         DOMAIN_NAME=production
#ARG         JDK_VER=8
#ENV         JAVA_HOME         /usr/lib/jvm/zulu8-ca-amd64


# server-full:5.2020.5-jdk11
FROM        payara/server-full:5.2020.5-jdk11

ARG         PAYARA_HOME=$PAYARA_DIR
ARG         DOMAIN_NAME=production
ARG         JDK_VER=11
ENV         JAVA_HOME         /usr/lib/jvm/zulu-11-amd64



ARG         JDBC_CON_POOL_REF=AnpiRef
ARG         JDBC_CON_POOL_UPD=AnpiUpd
ARG         JDBC_RESOURCE_REF=jdbc_AnpiRef
ARG         JDBC_RESOURCE_UPD=jdbc_AnpiUpd
ARG         JDBC_USER=rescuenow
ARG         JDBC_PWD=rescuenow119
ARG         JDBC_SERVER_REF
ARG         JDBC_SERVER_UPD
ARG         JDBC_PORT=3306
ARG         JDBC_DBNAME=anpidb

ENV         PATH              $PATH:$JAVA_HOME/bin:$PAYARA_HOME/glassfish/bin
ENV         TZ=Asia/Tokyo

USER        root

RUN         apt-get update -y && \
            apt-get install -y curl ant && \
            rm -rf /var/lib/apt/lists/*

# Create JDBC resource
RUN         ${PAYARA_HOME}/bin/asadmin --user=${ADMIN_USER} --passwordfile=${PASSWORD_FILE} start-domain ${DOMAIN_NAME} && \
            ${PAYARA_HOME}/bin/asadmin --user=${ADMIN_USER} --passwordfile=${PASSWORD_FILE} create-jdbc-connection-pool \
              --datasourceclassname com.mysql.jdbc.jdbc2.optional.MysqlXADataSource \
              --restype javax.sql.XADataSource \
              --property User=${JDBC_USER}:Password=${JDBC_PWD}:ServerName=${JDBC_SERVER_UPD}:Url="jdbc\:mysql\://${JDBC_SERVER_UPD}\:${JDBC_PORT}/${JDBC_DBNAME}":URL="jdbc\:mysql\://${JDBC_SERVER_UPD}\:${JDBC_PORT}/${JDBC_DBNAME}" ${JDBC_CON_POOL_UPD} && \
            ${PAYARA_HOME}/bin/asadmin --user=${ADMIN_USER} --passwordfile=${PASSWORD_FILE} create-jdbc-connection-pool \
              --datasourceclassname com.mysql.jdbc.jdbc2.optional.MysqlXADataSource \
              --restype javax.sql.XADataSource \
              --property User=${JDBC_USER}:Password=${JDBC_PWD}:ServerName=${JDBC_SERVER_REF}:Url="jdbc\:mysql\://${JDBC_SERVER_REF}\:${JDBC_PORT}/${JDBC_DBNAME}":URL="jdbc\:mysql\://${JDBC_SERVER_REF}\:${JDBC_PORT}/${JDBC_DBNAME}" ${JDBC_CON_POOL_REF} && \
            ${PAYARA_HOME}/bin/asadmin --user=${ADMIN_USER} --passwordfile=${PASSWORD_FILE} create-jdbc-resource \
              --connectionpoolid ${JDBC_CON_POOL_UPD} ${JDBC_RESOURCE_UPD} && \
            ${PAYARA_HOME}/bin/asadmin --user=${ADMIN_USER} --passwordfile=${PASSWORD_FILE} create-jdbc-resource \
              --connectionpoolid ${JDBC_CON_POOL_REF} ${JDBC_RESOURCE_REF} && \
            ${PAYARA_HOME}/bin/asadmin --user=${ADMIN_USER} --passwordfile=${PASSWORD_FILE} stop-domain ${DOMAIN_NAME}

# Install mysql-connector
RUN         curl -L -o mysql-connector-java.tar.gz https://downloads.mysql.com/archives/get/p/3/file/mysql-connector-java-5.1.47.tar.gz && \
            tar xvfz mysql-connector-java.tar.gz && \
            mv mysql-connector-java-5.1.47/mysql-connector-java-5.1.47.jar $PAYARA_HOME/glassfish/lib && \
            rm -rf mysql-connector-java*

# Copy files into the docker image
ADD         ./ /home
WORKDIR     /home

RUN         chmod 777 /home/redeploy.sh

# Deploy anpiclient
RUN         cp -f /home/anpiclient/lib/* $PAYARA_HOME/glassfish/lib && \
            cd /home/anpiclient/build && \
            sed -i \
             -e "s%^compile.jvm=.*%compile.jvm=${JDK_VER}%" \
             -e "s%^glassfish.home.dir=.*%glassfish.home.dir=${PAYARA_HOME}/glassfish%" \
             -e "s%^glassfish.domain1.name=.*%glassfish.domain1.name=${DOMAIN_NAME}%" \
             -e "s%^java.tools=.*%java.tools=%" \
             build.properties && \
            sed -i \
             -e "s%  </path>%    <fileset dir=\"\${glassfish.home.dir}/modules\" includes=\"*.jar\" /></path>%" \
             build.xml && \
            ant make-service && \
            cp /home/anpiclient/war/target/anpi.war $DEPLOY_DIR && \
            ant clean

# Deploy anpiadmin
RUN         cp -f /home/anpiadmin/lib/* $PAYARA_HOME/glassfish/lib && \
            cd /home/anpiadmin/build && \
            sed -i \
             -e "s%^compile.jvm=.*%compile.jvm=${JDK_VER}%" \
             -e "s%^glassfish.home.dir=.*%glassfish.home.dir=${PAYARA_HOME}/glassfish%" \
             -e "s%^glassfish.domain1.name=.*%glassfish.domain1.name=${DOMAIN_NAME}%" \
             -e "s%^java.tools=.*%java.tools=%" \
             build.properties && \
            sed -i \
             -e "s%  </path>%    <fileset dir=\"\${glassfish.home.dir}/modules\" includes=\"*.jar\" /></path>%" \
             build.xml && \
            ant make-service && \
            cp /home/anpiadmin/war/target/admin.war $DEPLOY_DIR && \
            ant clean

# Deploy famiclient
RUN         cp -f /home/famiclient/lib/* $PAYARA_HOME/glassfish/lib && \
            cd /home/famiclient/build && \
            sed -i \
             -e "s%^compile.jvm=.*%compile.jvm=${JDK_VER}%" \
             -e "s%^glassfish.home.dir=.*%glassfish.home.dir=${PAYARA_HOME}/glassfish%" \
             -e "s%^glassfish.domain1.name=.*%glassfish.domain1.name=${DOMAIN_NAME}%" \
             -e "s%^java.tools=.*%java.tools=%" \
             build.properties && \
            sed -i \
             -e "s%  </path>%    <fileset dir=\"\${glassfish.home.dir}/modules\" includes=\"*.jar\" /></path>%" \
             build.xml && \
            ant make-service && \
            cp /home/famiclient/war/target/famianpi.war $DEPLOY_DIR && \
            ant clean
