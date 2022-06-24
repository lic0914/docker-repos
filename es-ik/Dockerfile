FROM elasticsearch:6.8.10
WORKDIR /usr/share/elasticsearch/plugins/ik
COPY elasticsearch-analysis-ik-6.8.10.tgz  .
RUN tar zxvf elasticsearch-analysis-ik-6.8.10.tgz
#RUN ./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v6.8.23/elasticsearch-analysis-ik-6.8.23.zip